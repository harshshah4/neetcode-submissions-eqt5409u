class TextProcessor:
    # Implement method overloading for format_text method        
    def format_text(self, val1: str, val2: str = None) -> str:
        if not val2:
            return val1.upper()
        return f"{val1}{val2}"
    pass



# Don't modify the code below
processor = TextProcessor()
print(processor.format_text("hello"))
print(processor.format_text("hello", "world"))
