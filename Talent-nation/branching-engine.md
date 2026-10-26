LESSON CONTENT
The simplest branch: if
The if statement allows you to run a block of code only when a condition is true.

The syntax is:

if condition:
    statement
    statement
The colon at the end of the if line is required. The indented block below it is the code that runs if the condition is truthy.

For example:

balance = 12000
if balance > 5000:
    print("You can make a large transfer")
The output:

You can make a large transfer
If the condition is false, the indented block is skipped.

balance = 2000
if balance > 5000:
    print("You can make a large transfer")
print("End of check")
Output:

End of check
Notice that the last print is not indented, so it always runs. Only the indented block belongs to the if.

 
 
Adding an alternative: else
The else statement provides a block of code that runs when the if condition is false.

balance = 2000
if balance > 5000:
    print("Large balance")
else:
    print("Balance is not large")
Output:

Balance is not large
The else line must also end with a colon. It must be at the same indentation level as its matching if.

The else block is optional. If you omit else, nothing happens when the condition is false.

 
 
Multiple alternatives: elif
Sometimes you need more than two paths. That is what elif is for. elif stands for “else if.” It checks another condition only if the previous condition was false.

The general structure is:

if condition1:
    block1
elif condition2:
    block2
elif condition3:
    block3
else:
    block4

Python checks the conditions in order. The first one that is true runs its block. If none are true, the else block runs. The else is optional.

For example, a delivery fee based on distance:

distance = 7
if distance <= 5:
    fee = 1000
elif distance <= 10:
    fee = 1500
elif distance <= 20:
    fee = 2500
else:
    fee = 4000
print("Delivery fee:", fee)
Since distance is 7, the first condition distance <= 5 is false. The second condition distance <= 10 is true. So fee becomes 1500.

This is much cleaner than writing three separate if statements. With elif, only one block can run.

 
 
Conditions and indented blocks
Every branch contains an indented block. The block can contain multiple statements. All statements in the same block must be indented at the same level.

status = "active"
if status == "active":
    print("Account is active")
    print("You can proceed")
else:
    print("Account is not active")
Both print calls in the if block run because they are indented at the same level. If you indent one less or more, Python may raise an IndentationError or change the block’s meaning.

The indented block is the heart of the branch. The condition decides whether the block runs, and indentation defines what belongs to the block.

 
 
Compound conditions
You can combine multiple conditions in a single if using and, or, and not.

For example, a user can withdraw money only if the PIN is correct and the balance is enough:

pin_correct = True
balance = 8000
amount = 3000
if pin_correct and amount <= balance:
    print("Transaction approved")
else:
    print("Transaction declined")
Output:

Transaction approved
An or example:

payment_method = "transfer"
if payment_method == "cash" or payment_method == "transfer":
    print("Payment method accepted")
else:
    print("Unknown payment method")
A not example:

account_locked = False
if not account_locked:
    print("Account is open")
else:
    print("Account is locked")
Compound conditions make branching powerful. One if can express a complete business rule.

 
 
Selecting an execution path
The whole purpose of if, elif, and else is to select one execution path among several.

Think of a router in a logistics hub. A package arrives. The system checks:

If the destination is within Lagos, send it to the local van.
Else if the destination is within Nigeria, send it to the interstate truck.
Else, send it to the international courier.
Only one path is chosen. The package does not go to all three places. The same is true in Python. Only the first matching branch executes.

Let us model this:
```python
destination = "Abuja"
if destination == "Lagos":
    print("Send to local van")
elif destination == "Abuja":
    print("Send to interstate truck")
elif destination == "Accra":
    print("Send to international courier")
else:
    print("Route not found")
Output:

Send to interstate truck
```
Even though there is an elif for Accra later, it is not checked because Abuja already matched. This is top-to-bottom resolution, which we will cover in the next sub-lesson.

 
 
Analogy
Let us write a complete P.O.S decision:
```python
card_valid = True
pin_correct = False
amount = 4000
balance = 15000
if not card_valid:
    print("Card rejected")
elif not pin_correct:
    print("Incorrect PIN")
elif amount > balance:
    print("Insufficient balance")
else:
    print("Transaction successful")
```
This is a guard-like sequence of checks. It reads naturally:

If the card is invalid, reject.
Else if the PIN is wrong, show incorrect PIN.
Else if the amount is too high, show insufficient balance.
Else, approve.
Each branch is mutually exclusive. Python checks them in order and runs only one.

 
 
Branching without elif
You can often achieve the same result with nested if statements, but nesting can become messy. We will address that in 9.1.4.

For now, use elif for mutually exclusive conditions. It keeps the code flat and readable.

 
 
Common mistakes and misconceptions
A common mistake is to forget the colon after if, elif, or else. This is a SyntaxError.

Another mistake is to indent the else at a different level than its if. This confuses Python.

A third mistake is to write if x = 5: instead of if x == 5:. Assignment is not a condition.

A fourth mistake is to use multiple if statements when you want only one branch. If you want mutual exclusivity, use elif.

A fifth mistake is to use elif after else. else must come last.

A sixth mistake is to write else if as two words. In Python, it is elif.

A seventh mistake is to think elif is required. It is optional. You can have just if or just if and else.

An eighth mistake is to put a condition after else. else does not take a condition. It catches everything remaining.

 
 
Summary
if runs a block only when a condition is truthy.
else runs a block when the if condition is false.
elif checks another condition if previous ones are false.
Each branch has a colon and an indented block.
Compound conditions use and, or, and not.
Python selects exactly one branch in an if/elif/else chain.
Use elif for mutually exclusive conditions.
Additional Resources
Python Documentation: The if statement
Real Python: Conditional Statements in Python
W3Schools: Python If ... Else
Programiz: Python if...else Statement