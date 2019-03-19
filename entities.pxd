cimport cython

cdef class TestObj(object):

    cdef int num
    cdef str string

    cpdef str num_times_string(self)
    cpdef TestObj clone(self, int times)
    cpdef int fibonacchi(self)
    cpdef str md5(self) except -1
    cpdef str explode(self) except -1