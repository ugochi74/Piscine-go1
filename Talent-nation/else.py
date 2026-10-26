balance = 4500
account_locked = False
if account_locked:
    print("Account locked")
elif balance <= 0:
    print("No funds")
elif balance < 5000:
    print("Low balance warning")
else:
    print("Balance is healthy")