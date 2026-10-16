
LESSON CONTENT
Words Python keeps for itself
A keyword is a word that has a special meaning in Python. Keywords are also called reserved words because Python reserves them for its own use. You cannot use a keyword as a variable name, function name, or ordinary identifier.

For example, if is a keyword. You cannot write:

if = 10
If you try, Python raises a SyntaxError. The interpreter sees if and expects a condition. It does not accept if as a label for a value. The word is already locked.

Here are some common Python keywords you will meet:

if
else
elif
for
while
def
class
import
return
True
False
None
and
or
not
yield
try
except
pass
This is not the full list, but these are among the most frequently used.

You do not need to memorize all Python keywords today. What matters is that you understand the concept: certain words belong to the language and cannot be reused by you.

Keywords vs. identifiers
An identifier is the name you give to things you create. Variable names, function names, and class names are identifiers.

For example, these are identifiers:

customer_name
age
total_amount
send_money
A keyword is not an identifier. It is a word with a fixed role in the language. The interpreter recognizes keywords and treats them differently from ordinary names.

Think of a courtroom. The judge has a reserved seat. No matter how important a lawyer is, the lawyer cannot sit in the judge’s chair. That chair is reserved. Keywords are the judge’s seat in Python. They are reserved for the language itself.

Identifiers are like the chairs for everyone else. You can choose where your variables sit, but you cannot take the reserved seat.



Basic identifier rules
Python has simple rules for creating identifiers.

A name must start with a letter or an underscore _. It cannot start with a number.
# valid
name = "Amaka"
_total = 5000
# invalid
2name = "Chidi"   # SyntaxError
``
After the first character, a name may contain letters, numbers, and underscores.
student_1 = "Fatima"
order_id = 7788
``
Python is case-sensitive. Name, name, and NAME are three different identifiers.
age = 25
Age = 30
AGE = 35
``
These are three separate variables.
A name cannot contain spaces or special characters. Use underscores to separate words.
# valid
customer_phone = "08031234567"
# invalid
customer phone = "08031234567"   # SyntaxError
customer-phone = "08031234567"   # SyntaxError
``
A name cannot be a Python keyword.
class = "SS2"   # SyntaxError
``
These rules are strict, but they are easy. If you follow them, Python will accept your identifiers.



Naming conventions
Conventions are not laws. The interpreter will not crash if you ignore them. But good developers follow them because they make code easier to read.

The most common convention in Python is snake_case. This means all lowercase words separated by underscores.

Examples:

customer_name
account_balance
total_order_amount
number_of_students
For constant values, Python developers often use all uppercase with underscores. A constant is a value that should not change.

MAX_RETRY = 5
DEFAULT_CURRENCY = "NGN"
MIN_AGE = 18
For class names, Python uses PascalCase, where each word starts with an uppercase letter.

class BankAccount:
    pass
For now, we are mainly naming variables. Use snake_case. It is clean and conventional.

Avoid names that are too short unless they are clear in context. x and y may be fine in small math code, but customer_balance is clearer than cb.

Also avoid names that are too long and repetitive. A name like the_total_amount_of_money_in_the_customers_bank_account is excessive. Find a balance between clarity and brevity.

A good name should tell you what the value represents. If you see days_since_last_login, you immediately know what it stores. If you see d, you know nothing.



Why reserved keywords matter
The Python interpreter reads your code and tries to make sense of it. Keywords are the signposts that tell the interpreter what structure you are building.

When you write:

if balance > 5000:
    print("Enough balance")
Python sees if and knows this is a conditional statement. It expects a condition and a block. If you try to use if as a variable, Python gets confused. It expects the start of a conditional, but instead it sees an equal sign. That confusion produces a SyntaxError.

This is why using a keyword as an identifier causes a fatal syntax error. The interpreter refuses to continue because the structure of the program is broken. It cannot guess your intention.

The compiler is strict because ambiguity in programming is dangerous. If Python allowed you to redefine if, your code would become unreadable and unpredictable. The rule protects you.



AN analogy: reserved phone numbers
Imagine the Nigerian police have a short code: 112 for emergencies. You cannot register that number as your personal line. The telecommunications regulator reserves it.

If someone tries to call 112 to buy a car, the call goes to the emergency center, not a car dealer. The number has a fixed purpose.

Python keywords are like that. They are fixed-function words. You cannot use them for your own variables because they already mean something specific to the interpreter.

Ordinary phone numbers are like identifiers. You can assign them to your contacts as you wish. But you cannot assign 112 to “My Car Dealer.”



Common mistakes and misconceptions
A common mistake is to name a variable list or dict. These are not strict keywords, but they are built-in names. In Python, list and dict are built-in functions. You can technically overwrite them, but it is a bad idea because you will lose access to the original function. We will not discuss this deeply, but know that some names are discouraged even if not reserved keywords.

Another mistake is to use a capital letter in a variable name because you think it looks nice. It is allowed, but not conventional. CustomerName is not snake_case. Use customer_name.

A third mistake is to start a variable name with a number. 1st_name is invalid. Use first_name.

A fourth mistake is to confuse True and False with true and false. Python is case-sensitive. True and False are keywords. true and false are ordinary identifiers, though you should not use them for that purpose.

A fifth mistake is to think that naming conventions are optional rules that the interpreter enforces. They are not. The interpreter only enforces identifier rules. Conventions are for humans.



Summary
Keywords are words reserved by Python for its own use.
You cannot use keywords as variable names, function names, or identifiers.
Identifiers are names you create.
Valid identifiers start with a letter or underscore, may contain letters, numbers, and underscores, and cannot be keywords.
Python is case-sensitive.
Use snake_case for variables and constants, and PascalCase for class names.
Good names communicate purpose.
Using a keyword as a variable causes a fatal syntax error because the interpreter’s structure breaks.


Additional Resources
Python Documentation: Keywords
Python Documentation: Identifiers and keywords
W3Schools: Python Keywords
PEP 8: Naming Conventions
Lessons in this topic
01
LESSON 5.1: PYTHON KEYWORDS, SYNTAX & THE STRICT COMPILER
02
LESSON 5.1.1 RESERVED KEYWORDS
03
LESSON 5.1.2 THE INDENTATION ENGINE
04
LESSON 5.1.3 SYNTAX ERRORS VS. RUNTIME ERRORS
05
LESSON 5.1.4 CODE AS DOCUMENTATION
In this lesson

LESSON CONTENT

Words Python keeps for itself

Keywords vs. identifiers

Basic identifier rules

Naming conventions

Why reserved keywords matter

AN analogy: reserved phone numbers

Common mistakes and misconceptions
Next