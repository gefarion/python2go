package main

/*
typedef struct{
	int n;
	_GoString_ s;
}TestObj;

static inline void* call(void *a, void *b) {return ((void* (*)(void*))a)(b);}
*/
import "C"
import (
	crypto_md5 "crypto/md5"
	"fmt"
	"io"
	"unsafe"
)

type GoObj C.TestObj

//export sum
func sum(a, b int) int {
	return a + b
}

//export md5
func md5(o GoObj) *C.char {
	h := crypto_md5.New()
	io.WriteString(h, o.s)
	s := fmt.Sprintf("%x", h.Sum(nil))
	return C.CString(s)
}

//export fibonacci
func fibonacci(o GoObj) int {
	n := uint64(o.n - 1)
	if o.n <= 1 {
		return int(n)
	}

	var n2, n1 uint64 = 0, 1

	for i := uint64(2); i < n; i++ {
		n2, n1 = n1, n1+n2
	}

	return int(n2 + n1)
}

//export callback
func callback(a, b unsafe.Pointer) unsafe.Pointer {
	r := C.call(a, b)
	return r
}

func main() {}
