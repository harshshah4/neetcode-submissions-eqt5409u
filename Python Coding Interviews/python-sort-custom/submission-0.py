from typing import List

def words_ordering(word: str) -> int:
    return len(word)

def sort_words(words: List[str]) -> List[str]:
    words.sort(key=words_ordering, reverse=True)
    return words

def absolute(number: int) -> int:
    return abs(number)

def sort_numbers(numbers: List[int]) -> List[int]:
    numbers.sort(key=absolute)
    return numbers
    pass


# do not modify below this line
print(sort_words(["cherry", "apple", "blueberry", "banana", "watermelon", "zucchini", "kiwi", "pear"]))

print(sort_numbers([1, -5, -3, 2, 4, 11, -19, 9, -2, 5, -6, 7, -4, 2, 6]))
