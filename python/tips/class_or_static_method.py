# @classmethod vs. @staticmethod
# see more on http://stackoverflow.com/questions/136097/what-is-the-difference-between-staticmethod-and-classmethod-in-python

class Test:
    NAME = "jack"
    def foo(self, x):
        print self.NAME # access variable by self
        print x

    # the class is passed to it
    @classmethod
    def class_foo(cls, x):
        print cls.NAME # access variable by cls
        print (cls, x)

    # no class or self is passed
    @staticmethod
    def static_foo(x):
        print Test.NAME # access variable by class
        print x

print Test.NAME

t = Test()
t.foo(1)

# class_foo can be called by a instance or the class
# the instance is implicity passed
t.class_foo(2)
Test.class_foo(2)

# static_foo can only be called by the class
Test.static_foo(3)
