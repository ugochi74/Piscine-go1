Welcome & Overview
You already use operators every day, even if you have never called them that.

When you calculate the change a customer should receive from ₦5,000 after buying goods worth ₦3,200, you are using arithmetic. When you compare your phone balance to a recharge card price to see if you can afford it, you are using comparison. When you decide to go to a market only if you have enough cash and the weather is good, you are using logic.

Python gives you a precise set of symbols to express these everyday calculations and decisions. Those symbols are called operators. An expression is a combination of values, variables, and operators that Python can evaluate to produce a single result.

Operators are the grammar of computation. They allow you to turn raw data into useful results. A bank app uses arithmetic to subtract withdrawals. A transport app uses comparison to check if a route is longer than another. A POS terminal uses logical operators to decide whether to approve a transaction only when the PIN is correct and the balance is sufficient.

In this lesson, you will learn the four major groups of operators:

Arithmetic operators: +, -, *, /, //, %, **
Comparison operators: ==, !=, >, <, >=, <=
Logical operators: and, or, not
Order of operations: parentheses, precedence, and associativity
By the end, you should be able to:

Evaluate arithmetic expressions correctly, including floor division and modulo.
Use comparison operators to produce boolean values.
Combine booleans with and, or, and not.
Explain short-circuit evaluation and why it matters.
Trace complex expressions step by step using Python’s precedence rules.
Translate a real-world mathematical formula like Body Mass Index into Python code safely.
Take your time. Operators are not just symbols. They are the building blocks of every computation you will ever write






The basic tools of calculation
Python can do arithmetic, and it does it very well. The arithmetic operators are the symbols you already know from mathematics, plus a few that may be new.

Here is the complete set we will cover:

Operator	Meaning	Example	Result
+	Addition	7 + 2	9
-	Subtraction	7 - 2	5
*	Multiplication	7 * 2	14
/	Division	7 / 2	3.5
//	Floor division	7 // 2	3
%	Modulo	7 % 2	1
**	Exponentiation	7 ** 2	49

These are enough to handle most everyday calculations.

Let us examine each operator, because some of them have surprising behavior.



Addition + and subtraction -
Addition and subtraction work exactly as you expect.

print(1500 + 2000)   # 3500
print(5000 - 3200)   # 1800
You can add or subtract integers and floats. Python automatically handles the type of the result.

print(10 + 2.5)   # 12.5
print(7 - 2.0)    # 5.0
When you mix an integer and a float, Python produces a float. This is natural because a float is a more general representation of a number.



Unary + and -
The plus and minus signs can also be used as unary operators. A unary operator acts on a single value.

For example:

x = 5
y = -x
print(y)   # -5
Here, -x negates the value of x.

The unary plus is less useful but valid:

z = +5
print(z)   # 5
Unary minus is common for negative amounts, such as a debt of ₦1,500:

balance = -1500
print(balance)   # -1500
This represents that the account is overdrawn or that money is owed.

Unary operators are evaluated before most binary operators. For example, -5 + 3 is -2, not -(5 + 3).



Multiplication *
Multiplication uses the asterisk *, not the letter x.

price = 2500
quantity = 4
total = price * quantity
print(total)   # 10000
You can also multiply a string by an integer, which repeats the string. But that is a later topic. For now, focus on numbers.



Division /
The forward slash / performs ordinary division. The result is always a float, even if the numbers divide evenly.

print(10 / 2)   # 5.0
print(7 / 2)    # 3.5
Notice that 10 / 2 gives 5.0, not 5. This is important. In Python 3, division always returns a float unless you use floor division.

If you want to divide and get an integer result, you must use // or convert the result.



Floor division //
Floor division divides and then rounds down to the nearest integer. It is sometimes called integer division.

print(7 // 2)   # 3
print(10 // 2)  # 5
print(9 // 2)   # 4
Why is 7 // 2 equal to 3? Because 7 / 2 is 3.5, and floor division drops the decimal part. It rounds down, not up.

For negative numbers, floor division rounds down further away from zero:

print(-7 // 2)   # -4
Because -7 / 2 is -3.5, and floor division rounds down to -4. This may feel strange, but it is consistent with the mathematical definition of floor: the greatest integer less than or equal to the true result.

Floor division is useful when you need a whole number of units. For example, if a bus can carry 18 passengers and 40 people are waiting, you need:

print(40 // 18)   # 2
So 2 full buses can be filled. The remaining passengers need another bus.



Modulo %
The modulo operator % gives the remainder of a division.

print(7 % 2)   # 1
Because 7 divided by 2 is 3 with a remainder of 1.

More examples:

print(10 % 3)   # 1
print(9 % 3)    # 0
print(14 % 5)   # 4
Modulo is extremely useful for checking divisibility. If a number modulo 2 is 0, the number is even. If it is 1, the number is odd.

number = 17
if number % 2 == 0:
    print("Even")
else:
    print("Odd")
Modulo is also useful for cycling through a sequence. For example, if you want to know what the hour will be after 15 hours on a 12-hour clock:

hour = 9
later = (hour + 15) % 12
print(later)   # 0 (12 o'clock)
This gives the hour on a 12-hour cycle. The modulo operator keeps values within a range.

In Nigerian commerce, modulo might help you figure out how many items are left after packing goods into cartons. If you have 45 eggs and each carton holds 12, you can do:

remaining = 45 % 12
print(remaining)   # 9
So 9 eggs remain after filling 3 cartons.



Exponentiation ``
The double asterisk ** raises a number to a power.

print(2 ** 3)   # 8
print(5 ** 2)   # 25
print(10 ** 3)  # 1000
This is useful for area calculations. For example, the area of a square with side length 7 meters:

side = 7
area = side ** 2
print(area)   # 49
Exponentiation also works with floats:

print(9 ** 0.5)   # 3.0 (square root)
Because raising to the power of 0.5 is the same as taking the square root.



Division behavior summary
A common source of confusion is the difference between / and //.

/ always returns a float.
// returns the floor integer.
If you need an integer result from division, use //. If you need the exact decimal result, use /.

For example, when calculating the average of three numbers:

total = 10 + 15 + 20
average = total / 3
print(average)   # 15.0
The float result is appropriate because the average may have a decimal. But if you are calculating the number of full rows in a grid, // is better.



Operator combinations
You can combine arithmetic operators in a single expression.

result = (1000 + 2000) * 0.9 - 500
print(result)   # 2200.0
We will explore the order of operations in detail in 7.1.4. For now, know that Python follows mathematical precedence.



Common mistakes and misconceptions
A common mistake is to think / returns an integer. In Python 3, it returns a float. Use // for integer division.

Another mistake is to confuse % with division. % gives the remainder, not the quotient.

A third mistake is to write x or ^ for multiplication or exponentiation. Python uses * for multiplication and ** for exponentiation. The caret ^ in Python means something else entirely.

A fourth mistake is to forget that floor division rounds down, even for negative numbers. -7 // 2 is -4, not -3.

A fifth mistake is to divide by zero. 10 / 0 raises a ZeroDivisionError. This is a runtime error, not a syntax error.

A sixth mistake is to use % with floats and expect an exact integer remainder. It works with floats too, but the result may be a float.



Summary
Arithmetic operators include +, -, *, /, //, %, and **.
Unary + and - act on a single value.
/ always returns a float; // returns the floor integer.
% returns the remainder.
** raises to a power.
Arithmetic operators can be combined.
Division by zero raises a ZeroDivisionError.


Additional Resources
Python Documentation: Numeric types and operations
W3Schools: Python Arithmetic Operators
GeeksforGeeks: Python Arithmetic Operators
Programiz: Python Operators







Asking questions about values
Comparison operators compare two values and answer a yes-or-no question. The answer is always a boolean: True or False.

You use comparison in everyday life. Is my balance enough? Is my age up to 18? Is this item the same as the one I ordered? Python lets you ask these questions in code.

Here are the comparison operators:

Operator	Meaning	Example	Result
==	Equal to	5 == 5	True
!=	Not equal to	5 != 3	True
>	Greater than	7 > 2	True
<	Less than	2 < 7	True
>=	Greater than or equal to	7 >= 7	True
<=	Less than or equal to	3 <= 7	True
Let us examine each.



Equality ==
The == operator checks whether two values are equal.

print(10 == 10)   # True
print(10 == 5)    # False
It works with strings too:

print("Lagos" == "Lagos")   # True
print("Lagos" == "lagos")   # False, because Python is case-sensitive
Equality is used to check whether an input matches an expected value. For example, a bank app might ask:

entered_pin = "1234"
stored_pin = "5678"
print(entered_pin == stored_pin)   # False
This boolean can then be used in an if statement.

Do not confuse = and ==. A single = is assignment. It binds a name to a value. A double == is comparison. It checks equality. This is one of the most common beginner mistakes.



Inequality !=
The != operator checks whether two values are not equal.

print(10 != 5)    # True
print(10 != 10)   # False
If you want to know whether a product is not the same as another, use !=.

selected = "rice"
available = "beans"
print(selected != available)   # True
This is useful for detecting differences.



Ordering comparisons: >, <, >=, <=
These operators compare the order of values.

> checks if the left value is greater than the right.

print(7 > 2)   # True
print(2 > 7)   # False
< checks if the left value is less than the right.

print(2 < 7)   # True
print(7 < 2)   # False
>= checks if the left is greater than or equal to the right.

print(7 >= 7)   # True
print(6 >= 7)   # False
<= checks if the left is less than or equal to the right.

print(6 <= 7)   # True
print(7 <= 6)   # False
These operators are used constantly in conditions. For example, a POS terminal might check:

withdrawal_amount = 5000
balance = 12000
print(withdrawal_amount <= balance)   # True
If this boolean is True, the withdrawal can proceed.

Comparing strings and other values
Comparison operators work on more than numbers. Strings can be compared lexicographically, which means dictionary order.

print("apple" < "banana")   # True, because "a" comes before "b"
This is useful for sorting names. However, be careful: uppercase letters are considered “less than” lowercase letters in Python’s Unicode ordering. So "Z" < "a" is True. This may not be what you expect if you are sorting names naturally. For now, focus on numbers.



Boolean results
Every comparison expression produces a boolean value.

result = (5 > 3)
print(result)   # True
print(type(result))   # <class 'bool'>
This boolean can be stored in a variable and used later.

is_eligible = age >= 18
if is_eligible:
    print("You can register")
else:
    print("You are too young")
The variable is_eligible holds either True or False. This makes the if statement readable.



Combining comparisons
You can combine comparison operators with logical operators, which we cover next. But you can also chain comparisons in Python:

age = 25
print(18 <= age <= 65)   # True
This is a Python special feature. It checks both 18 <= age and age <= 65. It is a clean way to express a range.

For beginners, it is fine to use logical operators instead:

print(18 <= age and age <= 65)
Both are correct.

Real-world example: fuel price check
Imagine you want to buy fuel only if the price is between ₦600 and ₦800 per liter.

price = 650
print(price >= 600 and price <= 800)   # True
Or using chained comparison:

print(600 <= price <= 800)   # True
This boolean can then control a purchase decision.



Common mistakes and misconceptions
A common mistake is to use = instead of == for comparison. = is assignment. == is equality.

Another mistake is to write != but mean something else, or to put a space between ! and =. It must be != with no space.

A third mistake is to assume comparison is only for numbers. It works for strings and other types, but the result may not be meaningful for mixed types.

A fourth mistake is to compare values of different types. For example, 5 < "5" raises a TypeError in Python 3 because the types are incomparable.

A fifth mistake is to think == compares identity. It compares value. For identity, use is. But for beginners, value comparison is almost always what you want.

A sixth mistake is to use => or =< instead of >= and <=. The correct order is >= and <=.



Summary
Comparison operators answer yes/no questions and return booleans.
== checks equality; != checks inequality.
>, <, >=, <= compare ordering.
A single = is assignment, not comparison.
Comparisons can be chained in Python.
Comparing incompatible types raises a TypeError.


Additional Resources
Python Documentation: Comparisons
W3Schools: Python Comparison Operators
GeeksforGeeks: Python Comparison Operators
Programiz: Python Comparison Operators








Combining truths
Logical operators allow you to combine boolean values. They answer questions like:

Is the balance enough and is the PIN correct?
Do we have rice or beans in stock?
Is the item not expired?
Python has three logical operators: and, or, and not.

and — both must be true
The and operator returns True only if both operands are True.

a = True
b = True
print(a and b)   # True
If one is false, the result is false.

a = True
b = False
print(a and b)   # False
You can use and with comparisons:

balance = 12000
withdrawal = 5000
pin_correct = True
can_withdraw = pin_correct and withdrawal <= balance
print(can_withdraw)   # True
Here, can_withdraw is True only if both conditions hold.

A Nigerian example: A bank transfer is approved only if the sender has enough money and the recipient account exists.

sender_balance = 8000
amount = 3000
recipient_exists = True
approved = sender_balance >= amount and recipient_exists
print(approved)   # True
If either condition fails, the transfer is not approved.

 

or — at least one must be true
The or operator returns True if at least one operand is True.

a = True
b = False
print(a or b)   # True
It returns False only when both are false.

a = False
b = False
print(a or b)   # False
or is useful for alternative conditions. For example, a delivery company may accept either cash or transfer:

payment_method = "transfer"
accepted = payment_method == "cash" or payment_method == "transfer"
print(accepted)   # True
If the payment method is neither cash nor transfer, accepted becomes False.

not — negation
The not operator flips a boolean value.

a = True
print(not a)   # False
If a is false, not a is true.

a = False
print(not a)   # True
not is useful for checking the opposite of a condition. For example, if you want to know whether a user is not logged in:

logged_in = False
if not logged_in:
    print("Please log in first")
This is cleaner than writing logged_in == False.

Truthiness
In Python, logical operators can work with non-boolean values. Python interprets certain values as truthy or falsy. For now, know that:

0, 0.0, "", None, and empty containers are considered falsey.
Most other values are considered truthy.
For example:

print(5 and "hello")   # "hello"
This may look odd. The and operator returns the first falsey value or the last value if all are truthy. This is a more advanced topic. For now, use and, or, not with explicit booleans.

When you write:

if balance:
    print("Balance exists")
Python checks whether balance is truthy. If balance is 0, the condition is false. If it is any other number, it is true. This is convenient, but be careful: a balance of 0 might still be a valid state, and using truthiness could hide that.



Short-circuit evaluation
This is one of the most important concepts in logical operators.

When Python evaluates an and expression, it evaluates the left side first. If the left side is false, the whole expression is already false, so Python does not evaluate the right side.

When Python evaluates an or expression, it evaluates the left side first. If the left side is true, the whole expression is already true, so Python does not evaluate the right side.

This behavior is called short-circuit evaluation.

Why does it matter? Because the right side may have side effects, such as a function call or a division that could fail.

For example:

balance = 0
if balance > 0 and 100 / balance > 2:
    print("High ratio")
Because balance > 0 is false, Python does not evaluate 100 / balance. If Python had evaluated the right side, it would have raised a ZeroDivisionError. Short-circuit evaluation prevents that.

Another example with or:

name = ""
display_name = name or "Guest"
print(display_name)   # "Guest"
Because name is an empty string (falsey), Python returns "Guest". If name were truthy, Python would return name without evaluating "Guest".

Short-circuit logic is especially useful for safe checks:

if price > 0 and 1000 / price > 5:
    print("Good value")
The division only happens if price is positive. This prevents a ZeroDivisionError.



Logical operators in control flow
Logical operators are often used inside if conditions.

age = 20
has_id = True
if age >= 18 and has_id:
    print("Can vote")
else:
    print("Cannot vote")
This checks both conditions. Only if both are true does the person vote.

day = "Saturday"
if day == "Saturday" or day == "Sunday":
    print("It is the weekend")
This checks if either condition is true.



Precedence of logical operators
The not operator has higher precedence than and, which has higher precedence than or.

So:

not True and False
is the same as:

(not True) and False
This evaluates to False.

For clarity, use parentheses when mixing logical operators. It makes the intention obvious.

(a and b) or (c and d)


analogy
Think of a security gate at an event in Lagos.

The guard says, “You can enter if you have a ticket and your name is on the list.” That is and. Both conditions must be true.

Another gate says, “You can enter if you have a ticket or a VIP pass.” That is or. One condition is enough.

The guard also says, “If you are not on the banned list.” That is not.

Short-circuit evaluation is like the guard checking your ticket first. If you do not have a ticket, the guard does not even check the list. The decision is already made.



Common mistakes and misconceptions
A common mistake is to write & for and, or | for or. In Python, & and | are bitwise operators, not logical operators. Use the words and and or.

Another mistake is to think that and and or always return True or False. They can return one of the operands. For beginner code, use them in conditions where the result is interpreted as a boolean.

A third mistake is to forget short-circuit evaluation and write code that depends on the right side always evaluating. Do not do that. If the right side must run, write it separately.

A fourth mistake is to write if a and b or c without parentheses and assume it groups as (a and b) or c. Actually, that is correct because and has higher precedence than or, but parentheses improve readability.

A fifth mistake is to use = instead of == inside a condition. For example, if a = 5: is a syntax error. You mean if a == 5:.

A sixth mistake is to compare a boolean to True explicitly: if x == True:. This is usually unnecessary. Write if x: instead.



Summary
Logical operators combine booleans: and, or, not.
and returns true only if both operands are true.
or returns true if at least one operand is true.
not flips the boolean.
Short-circuit evaluation skips the right side when the result is already determined.
Short-circuiting can prevent errors.
Use parentheses to clarify complex logical expressions.


Additional Resources
Python Documentation: Boolean operations
W3Schools: Python Logical Operators
GeeksforGeeks: Python Logical Operators
Real Python: Operators and Expressions in Python




 
Python follows mathematical precedence
When you write a complex expression, Python does not always evaluate it from left to right. It follows a fixed set of rules called operator precedence.

Precedence determines which operations are performed first. For example, in mathematics, multiplication happens before addition. Python follows the same rule.

Consider:

result = 2 + 3 * 4
If Python evaluated left to right, it would be (2 + 3) * 4 = 20. But Python follows mathematical precedence, so it is 2 + (3 * 4) = 14.

This is why understanding precedence is crucial. If you write 2 + 3 * 4 expecting 20, you will be surprised when Python returns 14.



Parentheses override precedence
Parentheses () are used to group sub-expressions and force a specific order of evaluation. Anything inside parentheses is evaluated first.

result = (2 + 3) * 4
print(result)   # 20
Here, 2 + 3 is evaluated first because it is inside parentheses. Then the result 5 is multiplied by 4.

Parentheses can be nested.

result = ((2 + 3) * 4) - 1
print(result)   # 19
The innermost parentheses are evaluated first. 2 + 3 gives 5. Then 5 * 4 gives 20. Then 20 - 1 gives 19.

When in doubt, use parentheses. It makes your intention clear and avoids relying on memorized precedence rules.



Operator precedence table
Here is the precedence order for the operators covered in this lesson, from highest to lowest:

Parentheses ()
Exponentiation **
Unary + and -
Multiplication *, division /, floor division //, modulo %
Addition +, subtraction -
Comparison operators ==, !=, >, <, >=, <=
Logical not
Logical and
Logical or
This is a simplified table. It covers the operators we have learned.

Let us trace through an example.

result = 2 + 3 * 4 ** 2
Step by step:

Exponentiation: 4 ** 2 = 16
Multiplication: 3 * 16 = 48
Addition: 2 + 48 = 50
So result is 50.

If you wanted (2 + 3) * (4 ** 2), you would write parentheses:

result = (2 + 3) * (4 ** 2)
Now:

Parentheses: 2 + 3 = 5
Exponentiation inside parentheses: 4 ** 2 = 16
Multiplication: 5 * 16 = 80
The parentheses changed the outcome.



Associativity
What happens when two operators have the same precedence? Python uses associativity to decide the order.

For most arithmetic operators, associativity is left-to-right. This means evaluation proceeds from left to right.

For example:

result = 20 - 5 - 3
Since subtraction is left-associative, it evaluates as:

(20 - 5) - 3 = 12
Not 20 - (5 - 3) = 18.

For exponentiation **, associativity is right-to-left. This is an important exception.

result = 2 ** 3 ** 2
Right-associative means:

2 ** (3 ** 2) = 2 ** 9 = 512
Not (2 ** 3) ** 2 = 8 ** 2 = 64.

This is a common source of errors. When in doubt, use parentheses.

For comparison and logical operators, associativity is also left-to-right, but the effect is subtle. For clarity, avoid writing long chains without parentheses.



Evaluating complex expressions step by step
Let us trace a compound expression:

x = 5
y = 10
z = 3
result = x + y * z > x ** 2
We need to determine the final boolean value.

Step 1: Exponentiation has highest precedence.

x ** 2 = 5 ** 2 = 25

Expression becomes: x + y * z > 25

Step 2: Multiplication next.

y * z = 10 * 3 = 30

Expression becomes: x + 30 > 25

Step 3: Addition.

x + 30 = 5 + 30 = 35

Expression becomes: 35 > 25

Step 4: Comparison.

35 > 25 is True

So result is True.

Python would evaluate exactly in this order.



Precedence and logical operators
Logical operators also have precedence.

not has higher precedence than and, which has higher precedence than or.

This means:

a = True
b = False
c = True
result = a or b and c
b and c is False and True = False.

Then a or False = True or False = True.

So result is True.

If you wanted (a or b) and c, you would write parentheses.



Why order of operations matters in real systems
Incorrect precedence can cause serious bugs.

Imagine a bank calculates interest as:

interest = balance * rate + 0.5
If the programmer meant balance * (rate + 0.5), the result is wrong. Every customer’s interest could be calculated incorrectly.

In a medical application, a wrong BMI calculation could misclassify a patient as underweight or overweight. We will see this in the drill.

In a delivery fee calculation, forgetting parentheses could charge customers too much or too little.

This is why parentheses are not just for style. They are safety tools.



Best practices
Use parentheses whenever the expression is not instantly clear.
Break complex expressions into smaller named variables.
Avoid relying on the right-associativity of **; use parentheses.
Test your expressions with known values.
Write code that a tired developer can read at 2 a.m.
For example, instead of:

result = a + b * c ** d - e
Write:

power = c ** d
product = b * power
total = a + product
result = total - e
This is longer, but much clearer. Each step is easy to verify.



Common mistakes and misconceptions
A common mistake is to assume left-to-right evaluation for all expressions. Python follows precedence.

Another mistake is to forget that ** is right-associative. 2 ** 3 ** 2 is 512, not 64.

A third mistake is to omit parentheses and then be surprised by the result. When in doubt, parenthesize.

A fourth mistake is to confuse == with = inside an expression. = is assignment and cannot be used in an expression that expects a value.

A fifth mistake is to chain comparisons without understanding the internal logic. Python allows 1 < x < 10, which is convenient, but it is effectively (1 < x) and (x < 10). This is safe, but other chained comparisons can be tricky.

A sixth mistake is to ignore precedence between arithmetic and comparison. Arithmetic is evaluated before comparison. So 2 + 3 > 4 is 5 > 4, which is True.



Summary
Python follows mathematical precedence.
Parentheses override precedence and should be used for clarity.
Precedence order: parentheses, exponentiation, unary, multiplication/division/modulo, addition/subtraction, comparison, not, and, or.
Most arithmetic operators are left-associative.
Exponentiation ** is right-associative.
Complex expressions should be broken into steps or parenthesized.
Precedence bugs can have real-world consequences.


Additional Resources
Python Documentation: Operator precedence
Real Python: Operators and Expressions in Python
GeeksforGeeks: Precedence and Associativity in Python
Programiz: Python Operator Precedence







Lesson Overview
Imagine you ask a human friend to make you a smoothie. You say: "Use 2 bananas and 1.5 cups of milk, and write 'Smoothie' on the cup." Your friend knows the difference between the number 2 and the word "2". They know 1.5 is a measurement. They know "Smoothie" is a label, not an amount. They can easily switch between numbers and text when needed.

Now imagine you give the same instruction to a kitchen robot. It has no common sense. It doesn't know the difference between the number 2 and the text "2". It doesn't know 1.5 can be written as "1.5" for display. If you tell it to write "2 bananas" on the label, it tries to use the number 2 as text—but numbers and text are stored differently in its memory system. The robot gets confused and crashes.

For this entire lesson, we are going to use one simple scene: a robot making a smoothie using bananas, milk, a blender, and a cup with a paper label. The same bananas, milk, and cup will appear in every example, drill, and exercise. You will learn:

What type casting is and why we need it (it is like translating between languages).
How to convert text to numbers using int() and float().
How to convert numbers to text using str().
How to handle cases where the conversion fails.
By the end of this lesson, you will be able to safely convert between types and handle conversion errors without crashing your programs.

 



Key Concepts
Concept	Plain-Language Explanation	Storyline Reference
Type Casting	Changing a value from one data type to another.	Translating the text "2" written on a paper order slip into the actual number 2 so the robot can count it out.
int()	A built-in function that converts text characters or decimals into a whole integer.	Reading the text "2" and understanding it means exactly 2 physical bananas.
float()	A built-in function that converts text characters or whole integers into a decimal value.	Reading the text "1.5" on a receipt and measuring out exactly 1.5 cups of liquid milk.
str()	A built-in function that converts numbers or boolean values into text characters.	Converting the calculated number 3.5 into the text "3.5" so the robot can print it on the cup label.
Conversion Error	An execution error that occurs when you try to convert a value into a format that does not make sense.	Asking the robot to convert the word "banana" into a number; the robot does not know how to evaluate this and halts.
 

 
Lesson Content
Imagine a kitchen counter. On it sit three bananas, a carton of milk, a blender, and a cup with a blank paper label.

The robot stands ready. You type your first instruction:

fruit_count = "2"
milk_amount = "1.5"
total_ingredients = fruit_count + milk_amount
The robot whirs to life. It reads fruit_count as the text string "2" and milk_amount as the text string "1.5". It tries to add them together. Instead of doing math, it joins the text strings together and writes "21.5" on the cup label. Then it grabs 2 bananas and 1.5 cups of milk, but the blender reads the total label "21.5" and tries to blend 21.5 total items. The blender overflows, and your smoothie is ruined.

You wanted a smoothie with 3.5 total ingredient units. You got a label that says "21.5" and a confused robot. The robot did exactly what you wrote. You gave it text instead of numbers, and it does not know how to change "2" into 2 unless you explicitly tell it to.

 

 
Step 1: Why We Need Type Casting
In programming, data types behave differently. Numbers can perform math operations, while text strings can only be joined together (concatenated). Sometimes your script receives data as text (like when a customer types an order on a touchscreen), but you must perform math calculations on it before making the smoothie.

BAD Instruction (Adding Raw Strings)
banana_count = "2"      # text string
milk_cups = "1.5"       # text string
total = banana_count + milk_cups  # Tries to add text
What goes wrong: The robot sees two text values. It performs string concatenation and outputs "21.5".

GOOD Instruction (Casting to Numbers First)
banana_count = "2"      # text string
milk_cups = "1.5"       # text string

# Convert the text variables into numeric variables
banana_number = int(banana_count)   # Becomes the integer 2
milk_number = float(milk_cups)      # Becomes the float 1.5

# Perform mathematical addition on the numbers
total = banana_number + milk_number  # 2 + 1.5 = 3.5
Why this works: We explicitly tell the robot to translate the text strings into their equivalent numeric forms. Now the robot can perform the addition correctly.

💻 Try It Yourself
Run the interactive Python interpreter in your terminal and type:

bananas = "2"
milk = "1.5"
print(bananas + milk)
Verify that the output displays 21.5 instead of 3.5.

 

 
Step 2: Converting Text to Numbers: int() and float()
To convert a text string containing digits into a number, we use int() for whole counts and float() for quantities with decimal points.

int() converts text to a whole number: "2" becomes 2.
float() converts text to a decimal number: "1.5" becomes 1.5.
BAD Instruction (Using int() on Decimals)
milk_amount = "1.5"
milk_number = int(milk_amount)  # Tries to force a decimal into an integer
What goes wrong: The robot looks at "1.5" and expects a whole number. Because there is a decimal point, the conversion fails, and the program crashes with a ValueError.

GOOD Instruction (Using float() for Decimals)
milk_amount = "1.5"
milk_number = float(milk_amount)   # Successfully converts to 1.5
Why this works: The float() function is designed to read decimal points, converting the characters safely.

💻 Try It Yourself
Test how int() handles incorrect values by typing this in your interpreter:

print(float("1.5"))
# Now try:
print(int("1.5"))
Observe how the second line immediately triggers a ValueError crash.

 

 
Step 3: Converting Numbers to Text: str()
To print a number on the paper label or combine it with a text message, you must convert the numeric value back into a string of text characters using str().

BAD Instruction (Combining Text and Raw Numbers)
banana_count = 2
message = "I have " + banana_count + " bananas"
What goes wrong: The robot tries to add a string ("I have ") to a raw number (2). It does not know how to blend these two different types and crashes with a TypeError.

GOOD Instruction (Using str() to Convert)
banana_count = 2
message = "I have " + str(banana_count) + " bananas"
# message becomes "I have 2 bananas"
Why this works: str(2) converts the numerical value 2 into the character string "2", which can be joined with the surrounding text.

💻 Try It Yourself
Type this combined string expression in your terminal:

ounces = 8.5
print("Pouring " + str(ounces) + " ounces of milk.")
Verify that it prints the complete sentence without throwing an error.

 

 
Step 4: Handling Conversion Errors
If the user inputs characters that do not represent numbers, trying to convert them will crash the program. We must verify that the characters are valid digits before we attempt to convert them.

BAD Instruction (Unchecked Conversion)
user_input = "banana"
banana_count = int(user_input)  # Tries to convert alphabetical characters
What goes wrong: The robot cannot find any numeric digits in "banana". The program immediately crashes.

GOOD Instruction (Checking with .isdigit())
user_input = "banana"

# Check if the string consists entirely of digits (0-9)
if user_input.isdigit():
    banana_count = int(user_input)
    print("Conversion successful!")
else:
    print("Warning: That is not a valid number! Defaulting to 0.")
    banana_count = 0
Why this works: The script checks if the string contains only digits before running int(). If the check fails, it handles the mistake by setting a default value instead of crashing.

💻 Try It Yourself
Test the .isdigit() check in your terminal:

value1 = "15"
value2 = "banana"
print(value1.isdigit())
print(value2.isdigit())
Confirm that the first prints True and the second prints False.

 

 
Real-World Application
Type casting is an essential part of programming any system that accepts user input, such as retail checkout kiosks.

When you scan items at a supermarket:

The barcode scanner reads the item quantity as a string: "3".
The internal scale reads the item weight as a string: "1.45".
The store database returns the price as a string: "2.99".
Before the kiosk can calculate your transaction total, the backend system must run explicit conversions:

quantity = int("3")
price = float("2.99")
total_cost = quantity * price
If the software developer forgot to cast the variables, the kiosk would try to multiply "2.99" by "3", causing a crash that halts the entire checkout lane. When the machine is ready to display your receipt, it runs str(total_cost) to convert the numeric total back into text characters for the monitor screen. Every interface on your phone or computer uses type casting to convert raw user input into numbers and back again.

 

 
Common Mistakes & Self-Check Clarifications
Common Mistake	Clarification / Fix	Corrected Example (Smoothie Storyline)
Using int() on text with decimals	int() can only parse text that represents a whole number. If the text has a decimal point, you must use float().	Incorrect: milk_cups = int("1.5") (crashes with a ValueError).

Correct: milk_cups = float("1.5")
Expecting the original variable to change automatically	Functions like int(), float(), and str() do not modify the original variable. They generate a new value that you must store.	Incorrect: int(banana_order) on one line, then using banana_order (which is still a string) for addition.

Correct: banana_count = int(banana_order)
Casting unverified text inputs	Attempting to convert alphabetical strings directly to numbers will crash your script.	Incorrect: bananas = int("two") (causes a crash).

Correct: Check if "two".isdigit(): before calling int().
 
Research Prompt
What to Search
Search the internet for: "Python ValueError invalid literal for int() with base 10"

What You Should Find / Be Able to Do Afterward
After researching this error, you should be able to:

Explain exactly what causes Python to throw this specific ValueError message.
Identify which string patterns (like spaces, letters, or punctuation characters) trigger this crash when passed to int().
Explain how to use a try/except block to catch this error and keep your program running smoothly.
Additional Resources: https://automatetheboringstuff.com,Taking Input in Python

Lessons in this topic
01
LESSON 8.1: TYPE CASTING
02
LESSON 8.1.1 IMPLICIT VS. EXPLICIT CONVERSION: WHEN PYTHON GUESSES VS. WHEN YOU COMMAND
03
LESSON 8.1.2 DATA LOSS AND TRUNCATION: THE DANGER OF CASTING FLOATS TO INTEGERS
04
LESSON 8.1.3 STRING PARSING: TRANSLATING ASCII/UTF-8 TEXT INTO COMPUTABLE NUMBERS
05
LESSON 8.1.4 HANDLING VALUEERROR: WHAT HAPPENS WHEN TYPE CASTING FAILS
In this lesson

Lesson Overview

Key Concepts

Lesson Content

Step 1: Why We Need Type Casting

Step 2: Converting Text to Numbers: int() and float()

Step 3: Converting Numbers to Text: str()

Step 4: Handling Conversion Errors

Real-World Application
Next