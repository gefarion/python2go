package main

/*
#cgo pkg-config: python2
#define Py_LIMITED_API
#include <Python.h>

static inline int pyParseTuple1(PyObject *a, PyObject **o) {return PyArg_ParseTuple(a, "O", o);}
static inline int pyParseTuple2(PyObject *a, PyObject **o1, PyObject **o2) {return PyArg_ParseTuple(a, "OO", o1, o2);}
static inline int pyParseTuple2Opt(PyObject *a, PyObject **o1, PyObject **o2) {return PyArg_ParseTuple(a, "O|O", o1, o2);}
static inline PyObject* pyInitModule(const char *n, PyMethodDef *m) { return Py_InitModule(n, m); }
static inline void pyDecref(PyObject *o) { Py_DECREF(o); }

// go protos
extern PyObject* pySum(PyObject*, PyObject*);
extern PyObject* pyNumTimesString(PyObject*, PyObject*);
extern PyObject* pyClone(PyObject*, PyObject*);
extern PyObject* pyMD5(PyObject*, PyObject*);
extern PyObject* pyFibonacci(PyObject*, PyObject*);
extern PyObject* pyCatchExplode(PyObject*, PyObject*);
extern PyObject* pyNoCatchExplode(PyObject*, PyObject*);
extern PyObject* pyCreateObject(PyObject*, PyObject*);
extern PyObject* pyCallback(PyObject*, PyObject*);
extern PyObject* pyListOfClones(PyObject*, PyObject*);

*/
import "C"
import "fmt"

var testObj *C.PyObject

var pyMethods = []C.PyMethodDef{
	{C.CString("sum"), C.PyCFunction(C.pySum), C.METH_VARARGS, nil},
	{C.CString("num_times_string"), C.PyCFunction(C.pyNumTimesString), C.METH_VARARGS, nil},
	{C.CString("sum"), C.PyCFunction(C.pySum), C.METH_VARARGS, nil},
	{C.CString("clone"), C.PyCFunction(C.pyClone), C.METH_VARARGS, nil},
	{C.CString("md5"), C.PyCFunction(C.pyMD5), C.METH_VARARGS, nil},
	{C.CString("fibonacci"), C.PyCFunction(C.pyFibonacci), C.METH_VARARGS, nil},
	{C.CString("catch_explode"), C.PyCFunction(C.pyCatchExplode), C.METH_VARARGS, nil},
	{C.CString("no_catch_explode"), C.PyCFunction(C.pyNoCatchExplode), C.METH_VARARGS, nil},
	{C.CString("create_object"), C.PyCFunction(C.pyCreateObject), C.METH_VARARGS, nil},
	{C.CString("callback"), C.PyCFunction(C.pyCallback), C.METH_VARARGS, nil},
	{C.CString("list_of_clones"), C.PyCFunction(C.pyListOfClones), C.METH_VARARGS, nil},
	{nil, nil, 0, nil},
}

func pyCall1(f func(a *C.PyObject) *C.PyObject, args *C.PyObject) *C.PyObject {
	var a *C.PyObject
	if C.pyParseTuple1(args, &a) == 0 {
		return nil
	}
	if a == nil {
		fmt.Println("pyCall1: a = nil")
		return nil
	}
	return f(a)
}

func pyCall2(f func(a, b *C.PyObject) *C.PyObject, args *C.PyObject) *C.PyObject {
	var a, b *C.PyObject
	if C.pyParseTuple2(args, &a, &b) == 0 {
		return nil
	}
	return f(a, b)
}

func pyCall2Opt(f func(a, b *C.PyObject) *C.PyObject, args *C.PyObject) *C.PyObject {
	var a, b *C.PyObject
	if C.pyParseTuple2Opt(args, &a, &b) == 0 {
		return nil
	}
	return f(a, b)
}

func pyCallInstanceMethod(o *C.PyObject, s string) *C.PyObject {
	if o == nil {
		fmt.Println("pyCallInstanceMethod: o == nil")
		return nil
	}
	m := C.PyObject_GetAttrString(o, C.CString(s))
	if m == nil {
		return nil
	}
	//C.pyDecref(o);
	a := C.PyTuple_New(0)
	r := C.PyObject_CallObject(m, a)
	C.pyDecref(m)
	C.pyDecref(a)
	return r
}

func sum(a, b *C.PyObject) *C.PyObject {
	al := C.PyInt_AsLong(a)
	bl := C.PyInt_AsLong(b)
	return C.PyInt_FromLong(al + bl)
}

//export pySum
func pySum(self, args *C.PyObject) *C.PyObject {
	return pyCall2(sum, args)
}

func numTimesString(o *C.PyObject) *C.PyObject {
	return C.PyObject_GetAttrString(o, C.CString("num_times_string"))
}

//export pyNumTimesString
func pyNumTimesString(self, args *C.PyObject) *C.PyObject {
	return pyCall1(numTimesString, args)
}

//export pyClone
func pyClone(self, args *C.PyObject) *C.PyObject {
	return pyCall2Opt(func(a, b *C.PyObject) *C.PyObject {
		return pyCallInstanceMethod(a, "clone")
	}, args)
}

//export pyMD5
func pyMD5(self, args *C.PyObject) *C.PyObject {
	return pyCall1(func(o *C.PyObject) *C.PyObject {
		return pyCallInstanceMethod(o, "md5")
	}, args)
}

//export pyFibonacci
func pyFibonacci(self, args *C.PyObject) *C.PyObject {
	return pyCall1(func(o *C.PyObject) *C.PyObject {
		return pyCallInstanceMethod(o, "fibonacci")
	}, args)
}

//export pyNoCatchExplode
func pyNoCatchExplode(self, args *C.PyObject) *C.PyObject {
	return pyCall1(func(o *C.PyObject) *C.PyObject {
		return pyCallInstanceMethod(o, "explode")
	}, args)
}

func catchExplode(o *C.PyObject) *C.PyObject {
	r := pyCallInstanceMethod(o, "explode")
	if r == nil {
		return C.PyInt_FromLong(C.long(1))
	}
	return r

}

//export pyCatchExplode
func pyCatchExplode(self, args *C.PyObject) *C.PyObject {
	return pyCall1(catchExplode, args)
}

func createObject(n, s *C.PyObject) *C.PyObject {
	t := C.PyTuple_New(2)
	C.PyTuple_SetItem(t, 0, n)
	C.PyTuple_SetItem(t, 1, s)
	r := C.PyObject_CallObject(testObj, t)
	C.pyDecref(t)
	return r
}

//export pyCreateObject
func pyCreateObject(self, args *C.PyObject) *C.PyObject {
	return pyCall2(createObject, args)
}

func callback(a, b *C.PyObject) *C.PyObject {
	t := C.PyTuple_New(1)
	if t == nil {
		return nil
	}
	C.PyTuple_SetItem(t, 0, b)
	r := C.PyObject_CallObject(a, t)
	C.pyDecref(t)
	return r
}

//export pyCallback
func pyCallback(self, args *C.PyObject) *C.PyObject {
	return pyCall2(callback, args)
}

func listOfClones(a, b *C.PyObject) *C.PyObject {
	n := C.PyInt_AsLong(b)
	l := C.PyList_New(n)
	if l == nil {
		return nil
	}
	for i := 0; i < int(n); i++ {
		c := pyCallInstanceMethod(a, "clone")
		if c == nil {
			return nil
		}
		if C.PyList_SetItem(l, C.long(i), c) != 0 {
			return nil
		}
	}
	return l
}

//export pyListOfClones
func pyListOfClones(self, args *C.PyObject) *C.PyObject {
	return pyCall2(listOfClones, args)
}

//export inita
func inita() {
	m := C.pyInitModule(C.CString("a"), &pyMethods[0])
	if m == nil {
		return
	}

	// import object
	e := C.PyImport_ImportModule(C.CString("entities"))
	if e == nil {
		return
	}
	c := C.PyObject_GetAttrString(e, C.CString("TestObj"))
	if c == nil {
		return
	}
	testObj = c
}

func main() {}
