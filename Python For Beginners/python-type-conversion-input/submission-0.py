def read_integer() -> int:
    i = int(input())
    if type(i) is int:
        return i 
    pass


def read_float() -> float:
    i = float(input())
    if type(i) is float:
        return i 
    pass


# do not modify below this line
print(read_integer())
print(read_integer())
print(read_integer())

print(read_float())
print(read_float())
print(read_float())
