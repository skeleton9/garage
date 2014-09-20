global_var = 1

# it is ok to access global variable's value
def foo():
    return global_var # 1

# but when you try to set, this is actually a local variable
def foo2():
    global_var = 2 # you created a local variable

# if you want to set value to global var, you have to use `global`
def foo3():
    global global_var
    global_var = 3

assert global_var == 1

assert foo() == 1

foo2()
assert global_var == 1

foo3()
assert global_var == 3
