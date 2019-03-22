package main

/*
#include <stdlib.h>
typedef struct{
	int n;
	char *s;
}SrvObj;


static inline int callSrvObjCallback(void *f, SrvObj *a, SrvObj *b) {
	return ((int (*)(SrvObj*, SrvObj*))f)(a, b);
}

*/
import "C"
import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"unsafe"
)

type Srv struct {
	cb unsafe.Pointer
}

type Msg struct {
	Number int    `json:"n"`
	String string `json:"s"`
}

type Cache struct {
	req, resp   Msg
	cReq, cResp C.SrvObj
}

var cachePool sync.Pool

func (s *Srv) reqHandler(w http.ResponseWriter, r *http.Request) {

	var b []byte = nil
	var err error

	if r.Method != "POST" {
		http.Error(w, "only POST supporte", http.StatusMethodNotAllowed)
		return
	}
	cv := cachePool.Get()
	if cv == nil {
		cv = &Cache{}
	}
	c := cv.(*Cache)
	defer cachePool.Put(c)

	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&c.req)
	if err != nil {
		goto Error
	}
	c.cReq = C.SrvObj{
		n: C.int(c.req.Number),
		s: C.CString(c.req.String),
	}
	defer C.free(unsafe.Pointer(c.cReq.s))

	if C.callSrvObjCallback(s.cb, &c.cReq, &c.cResp) != 0 {
		err = errors.New("callSrvObjCallback: failed")
		goto Error
	}
	c.resp = Msg{Number: int(c.cResp.n), String: C.GoString(c.cResp.s)}
	b, err = json.Marshal(c.resp)
	if err != nil {
		goto Error
	}
	w.Write(b)
	return
Error:
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func (s *Srv) ListenAndServe(addr string) error {
	http.HandleFunc("/req", s.reqHandler)
	return http.ListenAndServe(addr, nil)
}

var once sync.Once

func doSignals() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Println(sig)
		os.Exit(0)
	}()
}

//export Serve
func Serve(addr string, callback unsafe.Pointer) {
	flag.Parse()
	once.Do(doSignals)

	s := &Srv{cb: callback}
	log.Fatal(s.ListenAndServe(addr))
}

func main() {}
