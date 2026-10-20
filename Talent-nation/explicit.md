
LESSON OVERVIEW
You've seen that Python can handle different types of data. You can add an integer and a float. You can multiply a string by a number. But how does Python know what to do when you mix types?

The answer is type conversion — also called type casting or coercion.

This is one of those concepts that seems simple on the surface but has deep implications. Understanding exactly how and when Python converts types will save you from countless bugs.


LESSON CONTENT

Implicit Conversion: Python Guesses for You
Sometimes Python automatically converts one type to another. This is called implicit conversion or coercion.

# Integer + Float = Float
result = 5 + 3.14
print(result)  # 8.14
print(type(result))  # <class 'float'>
Python saw that you were adding an integer and a float. It automatically converted the integer 5 to the float 5.0 before doing the addition. The result was a float.

This happens with other operations too:

# Integer / Integer = Float (always)
result = 10 / 2
print(result)  # 5.0
print(type(result))  # <class 'float'>
# Integer * Float = Float
result = 5 * 2.5
print(result)  # 12.5
print(type(result))  # <class 'float'>
# Integer + Boolean = Integer
result = 5 + True
print(result)  # 6
print(type(result))  # <class 'int'>
Wait, True is a number? Yes. In Python, True is treated as 1 and False as 0 in numeric contexts. This is a holdover from Python's C origins and can be surprisingly useful:

# Count how many True values are in a list
data = [True, False, True, True, False, True]
true_count = sum(data)  # Sums the booleans as 1s and 0s
print(true_count)  # 4
Python does this to avoid losing information. Converting an integer to a float is safe — you don't lose any data. But converting a float to an integer can lose the decimal part, so Python won't do that automatically.

Implicit conversion rules:

Integer → Float (safe, no data loss)
Integer → Complex (safe)
Float → Complex (safe)
Integer → Boolean (0 becomes False, anything else becomes True) — this happens in boolean contexts
Never Float → Integer (could lose data)
Never Complex → Float or Integer (could lose data)
Never String → Number (Python doesn't know how)
For mixed-type operations, Python always converts "up" the hierarchy:

Integer → Float → Complex
So int + float becomes float + float. int + complex becomes complex + complex.



Why Implicit Conversion Exists
Imagine if Python didn't do this automatically:

# Without implicit conversion, this would fail
result = 5 + 3.14  # TypeError: can't add int and float
You'd have to manually convert everything:

result = float(5) + 3.14  # 8.14
That would be tedious. Implicit conversion makes Python feel more natural. It "just works" in most cases.



When Implicit Conversion Can Surprise You
The implicit conversion rules are generally sensible, but there are some cases where they can catch you off guard:

# Adding a boolean to a number
result = 10 + True
print(result)  # 11
# Adding a boolean to a string
result = "Hello" + True  # TypeError!
This inconsistency can be confusing. Python will implicitly convert booleans to numbers in numeric contexts, but not in string contexts.



Explicit Conversion: You Take Control
Sometimes Python's automatic conversion isn't what you want. Sometimes you need to tell Python exactly what type you need. This is called explicit conversion.

Python gives you functions for this:

Function	What It Does
int()	Converts to an integer
float()	Converts to a float
str()	Converts to a string
bool()	Converts to a boolean
list()	Converts to a list
tuple()	Converts to a tuple
dict()	Converts to a dictionary
set()	Converts to a set
These are all covered in the official Python documentation for built-in functions: docs.python.org/3/library/functions.html



Converting to Integer: int()
# Float to int (truncates)
value = int(3.14)
print(value)  # 3
value = int(3.99)
print(value)  # 3 (not 4)
# String to int (must be a valid number)
value = int("42")
print(value)  # 42
value = int("-5")
print(value)  # -5
# Boolean to int
value = int(True)
print(value)  # 1
value = int(False)
print(value)  # 0
# This will fail
value = int("hello")  # ValueError
Warning: int() truncates toward zero. It doesn't round. So int(3.99) is 3, not 4. This catches many beginners off guard.



Converting to Float: float()
# Integer to float
value = float(5)
print(value)  # 5.0
# String to float
value = float("3.14")
print(value)  # 3.14
value = float("10")
print(value)  # 10.0
# Boolean to float
value = float(True)
print(value)  # 1.0
# Scientific notation
value = float("1.5e3")
print(value)  # 1500.0
# This will fail
value = float("hello")  # ValueError


Converting to String: str()
# Number to string
value = str(42)
print(value)  # "42"
value = str(3.14)
print(value)  # "3.14"
# Boolean to string
value = str(True)
print(value)  # "True"
# List to string
value = str([1, 2, 3])
print(value)  # "[1, 2, 3]"
# Dictionary to string
value = str({"a": 1, "b": 2})
print(value)  # "{'a': 1, 'b': 2}"
# None to string
value = str(None)
print(value)  # "None"


Converting to Boolean: bool()
# Numbers
print(bool(0))    # False
print(bool(1))    # True
print(bool(-5))   # True
print(bool(0.0))  # False
print(bool(3.14)) # True
# Strings
print(bool(""))   # False
print(bool("hi")) # True
print(bool(" "))  # True (non-empty string)
# Collections
print(bool([]))   # False
print(bool([1]))  # True
print(bool({}))   # False
print(bool({"a": 1}))  # True
print(bool(()))   # False
print(bool((1,))) # True
# None
print(bool(None)) # False
The truth value testing rules are documented here: docs.python.org/3/library/stdtypes.html#truth-value-testing



Why Explicit Conversion Matters
Explicit conversion is essential when you're getting input from users or files. Everything that comes from input() is a string, even if the user typed a number:

# Without conversion
age = input("Enter your age: ")  # User types "25"
print(age + 1)  # TypeError: can only concatenate str (not "int") to str
# With conversion
age = int(input("Enter your age: "))  # User types "25"
print(age + 1)  # 26
The official Python documentation covers input handling here: docs.python.org/3/library/functions.html#input



A Note on Type Hints
Python 3.5 introduced type hints — a way to document what types you expect:

def add_numbers(a: int, b: int) -> int:
    return a + b
These hints don't change how Python works. They're documentation for developers (and tools like mypy can check them). They don't force conversion — if you pass a string, Python will still try to add it.

Type hints are covered in the official documentation: docs.python.org/3/library/typing.html