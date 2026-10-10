from typing import Dict # this adds type hinting for Dict

def count_characters(word: str) -> Dict[str, int]:
    d = {}
    for w in word:
        if not (w in d):
            d[w] = word.count(w)
    return d
    pass




# don't modify below this line
print(count_characters("hello"))
print(count_characters("world"))
print(count_characters("hello world"))
print(count_characters("this is a longer sentence"))
