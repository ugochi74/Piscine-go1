LESSON OVERVIEW
Let me tell you something that might sound harsh: your users will break your program. They won't mean to do it, but they will. They'll type letters instead of numbers. They'll add commas where they shouldn't. They'll try to convert "hello" to an integer. And when they do, your program will crash.

But here's the thing: you can prevent it. And this is what separates a beginner's program from a professional one: the ability to handle errors gracefully.

What Is a ValueError?
A ValueError is what happens when Python tries to convert a value to a specific type, but the value doesn't make sense for that type. In the context of type casting, this usually happens when you try to convert a string that doesn't represent a valid number.

# This will raise a ValueError
value = int("not a number")
The error message will look something like this:

ValueError: invalid literal for int() with base 10: 'not a number'
In simpler terms: "I tried to convert this string to an integer, but it doesn't look like a number."



Invalid Numeric Input: Common Causes of ValueError
Here are the most common scenarios that cause ValueError during type casting:

1. Letters Instead of Numbers

# This will raise ValueError
user_input = "twenty-five"
value = int(user_input)  # ValueError: invalid literal for int()
2. Decimal Points in int()

# This will raise ValueError
user_input = "3.14"
value = int(user_input)  # ValueError: invalid literal for int()
# Note: float("3.14") would work, but int("3.14") does not
3. Commas in Numbers

# This will raise ValueError
user_input = "2,500"
value = int(user_input)  # ValueError: invalid literal for int()
4. Currency Symbols

# This will raise ValueError
user_input = "₦2500"
value = int(user_input)  # ValueError: invalid literal for int()
5. Extra Characters

# This will raise ValueError
user_input = "25 years old"
value = int(user_input)  # ValueError: invalid literal for int()


Understanding ValueError
When Python encounters a ValueError, it's telling you that the value you provided has the correct format for the function signature but doesn't have the right content for the operation you're trying to perform.

Let me break this down:

TypeError: You're using the wrong type (e.g., trying to add a string and an integer)
ValueError: You're using the right type, but the value is invalid for the operation (e.g., trying to convert "hello" to an integer)
Distinguishing Valid and Invalid Conversions
Here's a quick reference for what works and what doesn't:

Valid Conversions:

int("42")        # Valid
int("-15")       # Valid
int("+7")        # Valid
int("  100  ")   # Valid (whitespace is ignored)
float("3.14")    # Valid
float("42")      # Valid (integer string to float)
float("-2.5")    # Valid
float("  7.5  ") # Valid
Invalid Conversions:

int("3.14")      # Invalid (float string cannot be converted to int)
int("hello")     # Invalid (letters)
int("2,500")     # Invalid (commas)
int("₦2500")     # Invalid (currency symbol)
int("25 years")  # Invalid (extra characters)
int("")          # Invalid (empty string)
float("")        # Invalid (empty string)


Handling ValueError with try/except
The most common way to handle ValueError is with a try/except block. This is a way of saying to Python: "Try this thing. If it works, great. If it raises a ValueError, don't crash — just do this instead."

Here's the basic structure:

try:
    # Code that might raise an error
    value = int(user_input)
    print(f"Successfully converted to: {value}")
except ValueError:
    # Code to run if the error happens
    print("That's not a valid number!")


A Complete Example: Safe Conversion Function
def safe_convert_to_int(user_input):
    """Safely convert a string to an integer, handling ValueError."""
    try:
        value = int(user_input)
        print(f"Successfully converted to: {value}")
        return value
    except ValueError:
        print(f"Error: '{user_input}' is not a valid integer")
        return None
# Test the function
safe_convert_to_int("25")           # Success: 25
safe_convert_to_int("3.14")         # Error: '3.14' is not a valid integer
safe_convert_to_int("hello")        # Error: 'hello' is not a valid integer
safe_convert_to_int("")             # Error: '' is not a valid integer
safe_convert_to_int("  42  ")       # Success: 42


A More Complex Example: Handling Multiple Formats
Here's a function that can handle different number formats and provides helpful error messages:

def parse_number_safely(user_input):
    """
    Parse a number from a string, handling common formatting issues.
    Returns a tuple: (parsed_value, error_message)
    """
    # Check for empty input
    if not user_input or user_input.strip() == "":
        return None, "Empty input provided."
    # Clean the input: remove spaces and commas
    cleaned = user_input.replace(",", "").strip()
    cleaned = cleaned.replace("₦", "")
    # Try to convert to float first (handles integers too)
    try:
        return float(cleaned), None
    except ValueError:
        return None, f"'{user_input}' is not a valid number."
# Test the function
test_inputs = ["42", "3.14", "2,500.50", "₦1500", "hello", "", "  100  "]
for inp in test_inputs:
    value, error = parse_number_safely(inp)
    if error:
        print(f"❌ Input: '{inp}' -> Error: {error}")
    else:
        print(f"✅ Input: '{inp}' -> Value: {value}")


The Danger of Silent Type Errors
One of the biggest mistakes you can make is assuming your type conversions will always work. When a user enters invalid input, your program should:

Tell the user what's wrong (with a helpful message)
Allow them to correct their input (in an interactive program)
Handle the error gracefully (in production code)
Here's a real-world example of handling user input in a loop:

def get_valid_age():
    """Keep asking for age until the user enters a valid number."""
    while True:
        user_input = input("Please enter your age: ")
        user_input = user_input.strip()
        if user_input == "":
            print("Please enter something. Age cannot be empty.")
            continue
        try:
            age = int(user_input)
            if age < 0:
                print("Please enter a positive number.")
                continue
            if age > 150:
                print("That seems unlikely. Please enter a realistic age.")
                continue
            return age
        except ValueError:
            print(f"Sorry, '{user_input}' is not a valid age. Please enter a number.")
# Usage
age = get_valid_age()
print(f"Your age is: {age}")