from ctypes import *

_go = cdll.LoadLibrary("b.so")
def importgo(name, ret, args):
	f = getattr(_go, name)
	f.restype = ret
	f.argtypes = args
	return f

GoInt = c_int64

class GoString(Structure):
	_fields_ = [("p", c_char_p), ("n", GoInt)]

	def __init__(self, s):
		self.p = c_char_p(s)
		self.n = GoInt(len(s))
	def __eq__(self, o):
		return self.p == o.p and self.n == o.n
	def __repr__(self):
		return '{}({!r}, {!r})'.format(
			self.__class__.__name__, self.p, self.n)

class GoSlice(Structure):
	_fields_ = [("data", POINTER(c_void_p)),
		("len", GoInt), ("cap", GoInt)]
	def __repr__(self):
		return '{}({!r}, {!r}, {!r})'.format(
			self.__class__.__name__, self.data, self.len, self.cap)

class TestObj(Structure):
	_fields_ = [("n", GoInt), ("s", GoString)]

	def __init__(self, n, s):
		self.n = n
		self.s = GoString(s)

	@property
	def num_times_string(self):
		return self.s.p * self.n

	def md5(self):
		return _go.md5(self)

	def fibonacci(self):
		return _go.fibonacci(self)

	def clone(self, times=1):
		return TestObj(self.n * times, self.s.p * times)

	def __eq__(self, o):
		return self.n == o.n and self.s == o.s

	def __repr__(self):
		return '{}({!r}, {!r})'.format(
			self.__class__.__name__, self.n, self.s)


CALLBACK = CFUNCTYPE(py_object, py_object)
importgo('callback', py_object, [CALLBACK, py_object])
def callback(a, b):
	return _go.callback(CALLBACK(a), py_object(b))

importgo('sum', GoInt, [GoInt, GoInt])
def sum(a, b):
	return _go.sum(a, b)

importgo('md5', c_char_p, [TestObj])
importgo('fibonacci', GoInt, [TestObj])
