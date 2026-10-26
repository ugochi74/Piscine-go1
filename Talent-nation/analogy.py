card_valid = True
pin_correct = False
amount = 4000
balance = 15000
if not card_valid:
    print("Card rejected")
elif not pin_correct:
    print("incorrect PIN")
elif amount > balance:
    print("Insufficient balance")
else:
    print("Transaction successful")
