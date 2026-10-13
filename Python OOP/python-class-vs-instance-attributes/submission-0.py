class BankAccount: 
    # TODO: Add class and instance attributes at their appropriate places
    total_accounts = 0
    total_balance = 0
    def __init__(self, name, balance) -> None:
        self.name=name
        self.balance=balance
        BankAccount.total_accounts += 1
        BankAccount.total_balance += balance
        pass
    
    def printb(self):
        print("{}'s balance: ${}".format(self.name, self.balance))


# TODO: Create two accounts
a = BankAccount("Alice", 1000)
b = BankAccount("Bob", 2000)
# TODO: Print the information using the mentioned format
a.printb()
b.printb()
print("Total Accounts: {}".format(BankAccount.total_accounts))
print("Total Balance: ${}".format(BankAccount.total_balance))
