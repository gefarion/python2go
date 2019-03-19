import md5
import cython

class TestObj(object):

    def __init__(self, num, string):
        self.num = num
        self.string = string

    @property
    def num_times_string(self):
        return self.string * self.num

    @cython.locals(a = int, b = int, c = int)
    def fibonacchi(self):
        a = 0
        b = 1

        for i in xrange(self.num):
            t = a
            a += b
            b = t

        return t

    def md5(self):
        m = md5.new()
        m.update(self.string)
        return m.hexdigest()

    def clone(self, times=1):
        return TestObj(self.num * times, self.string * times)

    def explode(self):
        raise Exception("Booom!")

    def __eq__(self, other):
        return self.num == other.num and self.string  == other.string
