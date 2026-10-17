
LESSON CONTENT

Types live on objects, not names
You already know Python is dynamically typed. Let us now understand what that means at the memory level.

In Python, the type belongs to the object, not to the variable name.

When you write:

value = 42
The name value is bound to an object of type int. The name itself has no fixed type. It is just a label.

If you later write:

value = "forty-two"
The name value is now bound to an object of type str. The name did not change type, because the name never had a type. The object changed.

This is different from statically typed languages, where a variable has a declared type and cannot be rebound to a different type.

In Python, type checking happens at runtime. When the interpreter reaches an operation, it looks at the type of the objects involved and decides whether the operation is valid.



Runtime type association
When Python executes an assignment, it determines the type of the value at that moment.

For example:

data = 100
print(type(data))
This prints <class 'int'>.

Now:

data = 99.5
print(type(data))
This prints <class 'float'>.

The type is associated with the object at runtime. There is no separate type declaration.

This is why you can write flexible functions and scripts without knowing in advance what type a value will have. The interpreter figures it out as it runs.

However, this flexibility has a cost. If you make a wrong assumption about a type, the error appears only when the bad line executes. You cannot rely on a compile-time type checker to catch the mistake before running.



Rebinding across types
Let us follow a variable through several types in one script.

record = 5000
print(type(record))   # <class 'int'>
record = 5000.75
print(type(record))   # <class 'float'>
record = "5000 naira"
print(type(record))   # <class 'str'>
record = False
print(type(record))   # <class 'bool'>
record = None
print(type(record))   # <class 'NoneType'>
The same name record points to five different types. This is rebinding across types.

This is extremely useful in exploratory programming. You can start with one representation, then switch to another as you understand the problem better.

But it can also be dangerous if you are not careful. If a function expects record to be an integer and someone rebinds it to a string, a later arithmetic operation will fail.



Mutation vs. rebinding
There is an important distinction between mutation and rebinding.

Rebinding means pointing a name to a different object. You already understand this.

Mutation means changing an object in place without creating a new object. For immutable primitives, mutation is impossible. But for mutable objects like lists and dictionaries, mutation is possible.

Let us focus on immutables for now.

When you do:

name = "Ada"
name = "Adaeze"
This is rebinding. The string "Ada" was not mutated. A new string "Adaeze" was created, and name was pointed to it.

When you do:

count = 10
count = count + 5
This is also rebinding. The integer 10 was not changed. A new integer 15 was created, and count was pointed to it.

Even when it looks like mutation, immutables do not mutate. For example:

greeting = "Good"
greeting += " morning"
This looks like an in-place change, but under the hood Python creates a new string "Good morning" and rebinds greeting. The old string "Good" is discarded.

Understanding this distinction prevents confusion about identity and id(). Rebinding changes which object a name points to. Mutation changes the object itself.



Why dynamic typing speeds up iteration
Dynamic typing is a major reason Python is popular for data work, scripting, and prototyping.

You can load data from a CSV file without knowing in advance whether a column is numeric or text. You can assign values, inspect them with type(), and convert as needed.

For example, a data analyst in Lagos might read a column called price. The values might come in as strings because the CSV file has commas or naira symbols. The analyst can use float(price.replace(",", "")) to convert. They do not need to declare a fixed type for price.

This flexibility allows fast iteration. You can experiment, see what types you actually have, and adapt.

The downside is that you must be disciplined. You must test your code because the interpreter will not prevent type errors before execution. But for many projects, the speed benefit outweighs the risk.



analogy
Think of a minibus driver. The driver’s route can change depending on the day.

Sometimes the bus carries passengers from Obalende to CMS. That is like an integer value: a fixed count.

Sometimes it carries goods from a market to a depot. That is like a string value: text labels.

The driver is not permanently assigned to only passengers or only goods. The driver is just a person who can carry whatever is assigned. The assignment changes as the day goes.

The name of the driver is like a variable. The passengers or goods are like objects. The nature of what is carried can change because the driver is not fixed to one type of load.

Dynamic typing means the name can carry any type. It is flexible, but the driver must know what they are carrying before they act on it.



Common mistakes and misconceptions
A common mistake is to think a variable has a type. In Python, a variable has no type; the object it refers to has a type.

Another mistake is to assume that because a variable was an integer before, it will always be an integer. Someone may have rebound it to a string.

A third mistake is to confuse concatenation and addition. + works differently for numbers and strings. If value is "5" and you try value + 1, you get a TypeError.

A fourth mistake is to think that int("5") modifies the string. It does not. It creates a new integer 5. The original string remains.

A fifth mistake is to assume dynamic typing means no type errors. It means type errors occur at runtime instead of compile time. You still need to handle them.

A sixth mistake is to ignore None. A variable can be None, and if you try to perform arithmetic on it, you will get an error.



Summary
In Python, types live on objects, not on variable names.
Type is determined at runtime.
A name can be rebound across different types.
Rebinding points a name to a new object.
Mutation changes an object in place.
Primitive scalar types are immutable, so “changes” are actually rebinding.
Dynamic typing speeds up iteration but requires discipline and testing.


Additional Resources
Python Documentation: Dynamic Typing
Real Python: Why is Python dynamically typed?
GeeksforGeeks: Dynamic Typing in Python
Wikipedia: Type system