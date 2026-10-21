import heapq
from typing import List

def heap_function(arr: List[Any]) -> List[Any]:
    heapq.heapify(arr)
    return arr

def heapify_strings(strings: List[str]) -> List[str]:
    return heap_function(strings)
    pass


def heapify_integers(integers: List[int]) -> List[int]:
    return heap_function(integers)
    pass


def heap_sort(nums: List[int]) -> List[int]:
    heapq.heapify(nums)
    l = len(nums)
    return [heapq.heappop(nums) for _ in range(l)]
    pass


# do not modify below this line
print(heapify_strings(["b", "a", "e", "c", "d"]))
print(heapify_integers([3, 4, 5, 1, 2, 6]))
print(heap_sort([3, 4, 5, 1, 2, 6]))
