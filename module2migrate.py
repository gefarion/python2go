from entities import TestObj

def create_object(num, str):
    return TestObj(num, str)

def num_times_string(obj):
    return obj.num_times_string

def clone(obj, times=1):
    return obj.clone()

def list_of_clones(obj, count):
    return [obj.clone() for n in range(count)]

def fibonacchi(obj):
    return obj.fibonacchi()

def md5(obj):
    return obj.md5()

def catch_explode(obj):
    try:
        obj.explode()
    except Exception:
        return 1

def no_catch_explode(obj):
    obj.explode()

def callback(obj, callback):
    return callback(obj)
