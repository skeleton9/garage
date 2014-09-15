# extend unittest.testcase to support @for_examples decorator to enable multiple inputs for a test
# See http://stackoverflow.com/a/4455312/987180
#
# e.g. test_less will run twice with input (1,3) and (2,4)
# class MyTest(TestCaseExt):
#    @for_examples((1,3), (2,4))
#    def test_less(self, a, b):
#       self.assertTrue(a < b)

import unittest

__examples__ = "__examples__"

def for_examples(*examples):
    def decorator(f, examples=examples):
        setattr(f, __examples__, getattr(f, __examples__,()) + examples)
        return f
    return decorator

class TestCaseExtMetaclass(type):
    def __new__(meta, name, bases, dict):
        def tuplify(x):
            if not isinstance(x, tuple):
              return (x,)
            return x
        for methodname, method in dict.items():
            if hasattr(method, __examples__):
                dict.pop(methodname)
                examples = getattr(method, __examples__)
                delattr(method, __examples__)
                for example in (tuplify(x) for x in examples):
                    def method_for_example(self, method = method, example = example):
                        method(self, *example)
                    methodname_for_example = methodname + "(" + ", ".join(str(v) for v in example) + ")"
                    dict[methodname_for_example] = method_for_example
        return type.__new__(meta, name, bases, dict)

class TestCaseExt(unittest.TestCase):
    __metaclass__ = TestCaseExtMetaclass
    pass



