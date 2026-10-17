LESSON CONTENT

A variable is a name, not a box
Many programming tutorials tell you that a variable is like a box that holds a value. That picture is simple, but it is misleading for Python.

In Python, a variable is a name that refers to an object. The object lives somewhere in memory. The name is just a label that points to it.

Think of a phone contact. You save a number under the name “Amaka.” The name “Amaka” is not the phone number itself. It is a label in your contact list. The actual number is stored in the phone’s memory. If you change the number saved under “Amaka,” you are not changing the name. You are changing what the name points to.

A Python variable works the same way. When you write:

customer = "Amaka"
Python creates a string object "Amaka" in memory. Then it binds the name customer to that object. The name is not the string. It is a reference to the string.

This is why I said the box metaphor is wrong. A box contains a thing. A Python name does not contain an object; it points to an object. The object has its own location in memory.



Names and objects
Let us separate two ideas clearly:

Name: the identifier you write in your code, such as customer, balance, or age.
Object: the actual data value in memory, such as "Amaka", 10000, or 25.
When you run:

age = 25
Python does two things:

It creates or obtains the integer object 25.
It binds the name age to that object.
The name age now refers to the integer 25.

You can have multiple names referring to the same object.

a = 10
b = a
After these two lines, both a and b refer to the same integer object 10. There are not two tens in memory. There is one 10, and two labels pointing to it.

This is like two people saving the same phone number under different names. The phone stores one number, but “Doctor” and “Pharmacy” both point to it.



Assignment and rebinding
Assignment is the act of binding a name to an object.

city = "Lagos"
This binds the name city to the string "Lagos".

Later, you can bind the same name to a different object.

city = "Abuja"
Now city points to "Abuja". The string "Lagos" still exists in memory for a moment, but if no name refers to it, Python’s garbage collector will eventually clean it up. The name city has changed its binding.

This is called rebinding. It is not the same as changing the object. You are not modifying "Lagos" into "Abuja". You are pointing the name to a new object.

Let us trace this carefully:

location = "Kano"
print(location)   # Kano
location = "Enugu"
print(location)   # Enugu
The first assignment binds location to "Kano". The second assignment rebinds location to "Enugu". The string "Kano" is not changed. It is simply no longer referenced.



Tracking a name across different objects
Because a name can be rebound, you can track it as it moves from one object to another.

Consider this script:

value = 100
print(value)      # 100
value = 99.5
print(value)      # 99.5
value = "one hundred"
print(value)      # one hundred
value = False
print(value)      # False
The name value points to four different objects at different times. It starts with an integer, then a float, then a string, then a boolean. The name itself has no fixed type. It is just a label. The objects have types.

This ability to reassign a name to different types is part of Python’s dynamic typing, which we will explore later.

For now, the key is: assignment creates a binding. Reassignment changes the binding. The name is not permanently attached to any object.



Objects can live without names
It is possible to have an object without a name.

print("I am a nameless string")
Here, Python creates the string "I am a nameless string" and passes it directly to print. No variable name is bound to it. The object exists briefly while print uses it, then it becomes garbage and is cleaned up.

This shows that objects are the real data. Names are just convenient handles.

Think of a market goods delivery. The goods exist whether or not a label is stuck on them. The label helps you find the right package, but the package is the real thing. A package can be delivered without a label, but then it is hard to track.



Why the reference model matters
Understanding references helps you avoid common confusion.

For example, if you do:

a = [1, 2, 3]
b = a
In a box model, you might think b is a copy of the list. It is not. Both a and b refer to the same list object. This distinction matters when lists are mutable. But we are not dealing with lists yet. Even with numbers and strings, the distinction helps you understand identity and id().

Another example:

x = 10
y = x
x = 20
print(y)   # 10
Because y was bound to the integer 10, not to x itself. Changing x later does not affect y. y still points to 10.

If you thought of variables as boxes, you might think y contains a copy of what was in x. But Python’s model is simpler: y points to the same object 10. Later, x points to 20, but y remains pointing at 10.



analogy
Imagine you have a bank account. The account number is like a variable name. The money in the account is like the object.

When someone sends you ₦5,000, the account balance changes. The account number does not change. The name points to the balance.

If you close the account and open a new one, the old account number no longer has a balance. The name has been rebound.

If two people share a joint account, both names point to the same money. That is like two variables pointing to the same object.

This analogy is imperfect, but it reinforces that the name is separate from the value.



Common mistakes and misconceptions
A common mistake is to say, “The variable contains a string.” It is better to say, “The variable refers to a string.” This keeps the model correct.

Another mistake is to think reassignment modifies the old object. It does not. Reassignment points the name to a new object.

A third mistake is to think that if b = a, then b is a copy. For immutable values like numbers and strings, the distinction is less visible, but with mutable objects it becomes critical. We will reach mutability later.

A fourth mistake is to believe that deleting a name deletes the object. Deleting a name removes the reference. The object may remain if other names refer to it.

A fifth mistake is to confuse the name with its string representation. name and "name" are different. The first is a variable, the second is a string literal.



Summary
A variable is a name bound to an object in memory.
Objects live in memory; names are labels that refer to them.
Assignment binds a name to an object.
Reassignment changes the binding.
Multiple names can refer to the same object.
Objects can exist without names.
Reassignment does not change the old object; it points the name elsewhere.


Additional Resources
Python Documentation: Assignment Statements
Real Python: Variables in Python
GeeksforGeeks: Python Variables
Programiz: Python Variables and Literals
Lessons in this topic
01
LESSON 6.1: MEMORY ALLOCATION, VARIABLES & PRIMITIVE TYPES
02
LESSON 6.1.1 VARIABLES AS MEMORY REFERENCES, NOT CONTAINERS
03
LESSON 6.1.2 PRIMITIVES IN MEMORY
04
LESSON 6.1.3 DYNAMIC TYPING ENGINE
05
LESSON 6.1.4 INTROSPECTION: TYPE, VALUE, AND IDENTITY
In this lesson

LESSON CONTENTA variable is a name, not a box

Names and objects

Assignment and rebinding

Tracking a name across different objects

Objects can live without names

Why the reference model matters

analogy

Common mistakes and misconceptions
