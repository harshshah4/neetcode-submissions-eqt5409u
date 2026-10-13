class SuperHero:
    def __init__(self, name: str, health: int, power_level: int):
        self.name = name
        self.health = health
        self.power_level = power_level
        # TODO: Add the private attributes
    
    # TODO: Add the getter and setter methods
    def get_name(self) -> str:
        return self.name
        pass
    
    def get_health(self) -> int:
        return self.health
        pass
    
    def get_power_level(self) -> int:
        return self.power_level
        pass
    

    def name(self, name: str):
        self.name = name
        pass
    
    def set_health(self, health: int):
        if health > 100:
            print("You can't set the health to more than 100")
        elif health < 0:
            print("You can't set the health to less than 0")
        else: self.health = health
        
    
    def set_power_level(self, power_level: int):
        if power_level > 10:
            print("You can't set the power level to more than 10")
            return
        elif power_level < 1:
            print("You can't set the power level to less than 1")
            return
        else: self.power_level=power_level
        pass






super_hero = SuperHero("Batman", 80, 9)

print(super_hero.get_health()) # this should print 80
super_hero.set_health(110) # this should print You can't set the health to more than 100
super_hero.set_health(-10) # this should print You can't set the health to less than 100
super_hero.set_health(70)

print(super_hero.get_power_level()) # this should print 9
super_hero.set_power_level(11) # this should print You can't set the power level to more than 10
super_hero.set_power_level(0) # this should print You can't set the power level to less than 1
super_hero.set_power_level(7)



# TODO: print the hero's attributes
print("{} has {} health and {} power level".format(super_hero.get_name(), super_hero.get_health(), super_hero.get_power_level()))
