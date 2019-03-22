from ctypes import *

_go = cdll.LoadLibrary("c.so")
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
	def __repr__(self):
		return '{}({!r}, {!r})'.format(
			self.__class__.__name__, self.p, self.n)

class SrvObj(Structure):
	_fields_ = [("n", c_int), ("s", c_char_p)]

	def __init__(self, n, s):
		self.n = n
		self.s = c_char_p(s)


SRV_CALLBACK = CFUNCTYPE(GoInt, POINTER(SrvObj), POINTER(SrvObj))
importgo('Serve', None, [GoString, SRV_CALLBACK])

_callback = None
def handle(a, b):
	try:
		resp = _callback(a[0])
		if resp == None:
			return -1
		b[0] = resp
		return 0
	except Exception:
		return -2

def serve(a, cb):
	global _callback
	_callback = cb
	return _go.Serve(GoString(a), SRV_CALLBACK(handle))
