import math

class AreaCalc:
    # TODO: Implement calculate method
    def calculate(self, val1, val2=None) -> float:
        if not val2:
            return round(math.pi * math.pow(val1, 2),2)
        return val1 * val2
    pass
    

    
# Don't modify the following code
calc = AreaCalc()
print(calc.calculate(5))    
print(calc.calculate(4, 6))
