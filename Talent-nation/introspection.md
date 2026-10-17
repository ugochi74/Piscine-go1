
Three facts about every object
Every object in Python has three important characteristics:

Value — the data it contains, such as 42, "Lagos", or False.
Type — what kind of object it is, such as int, str, or bool.
Identity — a unique identifier that distinguishes this object from all others.
You can inspect these three facts using built-in tools. This is called introspection, because you are looking inside Python’s runtime to understand what is happening.

The three tools are:

The object itself for value.
type() for type.
id() for identity.
Let us explore each.



type() reveals the type
The type() function returns the type of an object.

print(type(25))
Output:

<class 'int'>
It tells you that 25 is an integer.

More examples:

print(type(3.14))      # <class 'float'>
print(type("Hello"))   # <class 'str'>
print(type(True))      # <class 'bool'>
print(type(None))      # <class 'NoneType'>
You can assign a value to a variable and then check its type.

amount = 1500
print(type(amount))
Output:

<class 'int'>
This is useful when you are not sure what type of data you are dealing with, especially when data comes from external sources like user input or files.



id() reveals identity
The id() function returns a unique integer that represents the object’s identity. This is usually related to its address in memory, but Python does not guarantee the exact value. It only guarantees that two objects that are the same have the same id, and two different objects have different ids during their lifetimes.

a = 100
b = 100
print(id(a))
print(id(b))
Depending on Python’s caching of small integers, these two may have the same id because both names point to the same cached integer object. Do not rely on this behavior for correctness.

For a clearer example:

x = [1, 2, 3]
y = [1, 2, 3]
print(id(x))
print(id(y))
These two lists are separate objects. They have different ids even though their contents are the same. This shows that equality of value does not mean identity.

For now, you can use id() to observe whether two names refer to the same object.



Observing rebinding through type() and id()
Let us track a variable through rebinding and watch how type and id change.

value = 10
print(value, type(value), id(value))
value = 10.5
print(value, type(value), id(value))
value = "ten"
print(value, type(value), id(value))
When you run this, you will see that the type changes at each step, and the id also changes because the object is different.

The name value remains the same, but the object it points to is different. The id reveals this.

Let us also observe two names pointing to the same object.

a = 7
b = a
print(id(a))
print(id(b))
Both prints show the same id because a and b refer to the same integer object.

If we rebind a:

a = 8
print(id(a))
print(id(b))
Now a has a different id, while b keeps the old id. This is proof that rebinding a did not affect b.

This type of introspection is powerful for debugging. You can check exactly when a name starts pointing to a new object.



Value, type, and identity compared
Let us compare two objects.

m = 50
n = 50
Values: both are 50, so equal in value.
Types: both are int.
Identity: maybe the same or maybe different, depending on caching. For small integers, usually the same.
Now:

m = 1000
n = 1000
Values: both 1000.
Types: both int.
Identity: often different for larger integers because Python may not cache them. But this is an implementation detail.
The important distinction is that equality (==) checks value, while identity (is) checks whether two names refer to the same object.

a = [1, 2]
b = [1, 2]
print(a == b)   # True, same value
print(a is b)   # False, different objects
For immutable primitives, you will mostly use ==. Identity checks are more relevant for mutable objects and None. For example, use x is None rather than x == None.



Using introspection in daily workflow
Introspection helps you:

Confirm the type of a variable before performing an operation.
Check whether a function returned None.
Debug by printing type and id.
Understand whether two variables refer to the same object.
Verify that a conversion produced the expected type.
For example, if you read input from a user, you might check:

raw = input("Enter amount: ")
print(type(raw))
This will show <class 'str'>, because input() always returns a string. Knowing this, you will convert it to int or float before arithmetic.

This is exactly the kind of defensive awareness you will practice in the drills.



analogy
Think of a customs officer inspecting goods at Apapa port.

The officer wants to know:

What is the cargo? That is the value.
What category does the cargo belong to? Electronics, food, textiles? That is the type.
Which container is it in, and is it distinct from the container next to it? That is the identity.
The officer uses tools to inspect the cargo without unloading everything. type() and id() are inspection tools. They let you look at the object without changing it.

A good officer does not assume the cargo is food because the invoice says so. The officer opens and checks. Similarly, a good programmer checks the type instead of assuming.



Common mistakes and misconceptions
A common mistake is to confuse type() with id(). type() tells you the category; id() tells you the identity.

Another mistake is to think that because two values are equal, they must have the same id. Equality and identity are different.

A third mistake is to use id() to compare values. id() is for identity, not value.

A fourth mistake is to assume id() remains stable after an object is garbage-collected. Python may reuse ids later. Do not store ids as permanent records.

A fifth mistake is to use type() for comparison in a way that is too verbose. For example, type(x) == int. This works, but the built-in isinstance(x, int) is often preferred. We will not cover isinstance yet, but know that type() is mainly for inspection.

A sixth mistake is to think that printing a value is enough to understand its type. "5" and 5 may look similar but have different types. Use type() to be sure.



Summary
Every object has a value, a type, and an identity.
type() returns the type of an object.
id() returns a unique identifier for an object.
Reassignment changes the object a name points to, so id() changes.
Equality (==) checks value; identity (is) checks whether two names point to the same object.
Introspection helps you debug and verify assumptions.
Use type() and id() as inspection tools, not as a replacement for logic.


Additional Resources
Python Documentation: Built-in Functions — type()
Python Documentation: Built-in Functions — id()
Real Python: Python's id() and type()
Programiz: Python type() and id()
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

Three facts about every object

type() reveals the type

id() reveals identity

Observing rebinding through type() and id()

Value, type, and identity compared

Using introspection in daily workflow

analogy

Common mistakes and misconceptions
