# split empty string with delimiter will get ONE element in the result

# this works fine
s1 = "a,b,c"
print s1.split(",") # ['a', 'b', 'c']

# when split a emtpy string with delimiter, the result is not empty
s2 = ""
print s2.split(",") # [''] the result is NOT empty
print s2.split()    # []

