from typing import List


def combine_elements(arr1: List[int], arr2: List[int]) -> List[int]:
    arr = []
    arr.extend(arr1)
    arr.extend(arr2)
    return arr
    pass



# do not modify below this line
arr1 = [1, 3, 5]
arr2 = [4, 6, 8]

print(combine_elements(arr1, arr2))
print(arr1)
print(arr2)
