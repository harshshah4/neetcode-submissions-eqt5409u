import heapq
from typing import List


def heap_pop(heap: List[int]) -> List[int]:
    temp_heap = heap.copy()
    return [heapq.heappop(temp_heap) for _ in heap]
    pass


# do not modify below this line
print(heap_pop([1, 2, 3]))
print(heap_pop([1, 3, 2]))
print(heap_pop([6, 7, 8, 12, 9, 10]))
