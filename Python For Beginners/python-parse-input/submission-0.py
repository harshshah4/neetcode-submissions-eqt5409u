from typing import List

def read_integers() -> List[int]:
    i = input()
    l = i.split(",")
    int_list = [int(x) for x in l]
    return int_list
    pass

# do not modify the code below
print(read_integers())
print(read_integers())
print(read_integers())
