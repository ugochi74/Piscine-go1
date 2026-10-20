LESSON OVERVIEW
This is where things start to get interesting. You've probably noticed that when you use input() in Python, you get a string. But what if you want to do math with what the user typed? You need to parse that string — which means converting it from text into a number.

Think of it like this: a customer comes to your shop and says "I want to buy 5 oranges." You don't need to convert the word "five" to a number because the person has already done that for you. But imagine if they said "I want to buy five oranges" — you'd have to think about it and convert that word "five" to the number 5 in your head. That's basically what string parsing is — taking a string that represents a number and converting it into an actual number.



What Is String Parsing?
String parsing is the process of converting a string that contains a numeric value into an actual number. In Python, you do this using the int() and float() functions.

Here's a simple example:

# Parsing a string into an integer
user_input = "25"
value = int(user_input)
print(value + 10)  # Output: 35
# Parsing a string into a float
user_input = "3.14"
pi = float(user_input)
print(pi * 2)  # Output: 6.28


Valid Numeric Strings: What int() and float() Can Handle
int("42") — Converting Strings to Integers

The int() function can parse strings that represent whole numbers (integers):

# Valid integer strings
print(int("42"))      # 42
print(int("-15"))     # -15
print(int("0"))       # 0
print(int("  100  ")) # 100 (whitespace is ignored)
print(int("+7"))      # 7 (explicit positive sign)
float("3.14") — Converting Strings to Floats

The float() function can parse strings that represent numbers with decimal points:

# Valid float strings
print(float("3.14"))      # 3.14
print(float("-2.5"))      # -2.5
print(float("0.0"))       # 0.0
print(float("  7.5  "))   # 7.5 (whitespace is ignored)
print(float("1e-3"))      # 0.001 (scientific notation)
print(float("+3.0"))      # 3.0 (explicit positive sign)


Strings That Cannot Be Parsed as Numbers
Not every string can be converted to a number. Here are some invalid examples:

# These will cause ValueError
# int("3.14")      # Can't convert a float string to int
# int("42.0")      # Can't convert a float string to int
# int("hello")     # Letters are not numbers
# float("42")      # This actually works (int to float)
# float("3,14")    # Comma is not a decimal separator in Python
# int("1,000")     # Commas are not allowed in Python numbers
# int("42 spaces") # Extra characters cause errors
Here's an important distinction to remember:

# This works: converting an integer string to float
print(float("42"))  # 42.0 (works)
# This fails: converting a float string to int
# print(int("3.14"))  # ValueError: invalid literal for int() with base 10: '3.14'


The Importance of String Parsing
Why is string parsing so important? Because when data enters your program from the outside world, it almost always arrives as text (strings). And if you're building a real application that actually does useful work, you need to process that data.

Let me give you some examples:

Example 1: A Point-of-Sale System

Imagine you're building a Point-of-Sale (POS) system for a shop in Lagos. The cashier enters the price of an item as text. You need to parse that text into a number so you can calculate the total.

# Simulating user input
price_input = "2500.50"  # In a real system, this would come from an input field
quantity_input = "3"
# Parsing the strings into numbers
price = float(price_input)
quantity = int(quantity_input)
# Calculating the total
total = price * quantity
print(f"Total: ₦{total:.2f}")  # Output: Total: ₦7501.50
Example 2: A Banking Application

Let's say you're building an app that calculates loan interest. The user enters the principal amount and the interest rate as text, and you need to convert those to numbers to perform calculations.

# User input as strings
principal_input = "500000"  # ₦500,000 loan
rate_input = "7.5"          # 7.5% interest
# Parsing the strings
principal = float(principal_input)
rate = float(rate_input) / 100
# Calculating interest for one year
interest = principal * rate
print(f"Annual interest: ₦{interest:.2f}")  # Output: ₦37500.00
Handling Different Number Formats
Numbers can be formatted in different ways, and you need to clean them before parsing:

Removing Commas

# Numbers with commas
price_with_comma = "2,500.00"
# Clean the string first
cleaned = price_with_comma.replace(",", "")
value = float(cleaned)  # Output: 2500.0
Removing Currency Symbols

# Numbers with currency symbols
price_with_currency = "₦2500.00"
# Clean the string first
cleaned = price_with_currency.replace("₦", "")
value = float(cleaned)  # Output: 2500.0
Handling Whitespace

# Numbers with extra whitespace
price_with_spaces = "  2500.50  "
# Use strip() to remove whitespace
cleaned = price_with_spaces.strip()
value = float(cleaned)  # Output: 2500.5


A Function That Handles Multiple Formats
Here's a robust function that handles common number formats:

def parse_nigerian_number(input_string):
    """Convert a Nigerian-style number string to a float."""
    # Remove commas (used as thousand separators)
    cleaned = input_string.replace(",", "")
    # Remove currency symbols
    cleaned = cleaned.replace("₦", "")
    # Remove extra whitespace
    cleaned = cleaned.strip()
    # Convert to float
    try:
        return float(cleaned)
    except ValueError:
        # Try to handle other formatting issues
        cleaned = cleaned.replace(" ", "")
        return float(cleaned)
# Examples
print(parse_nigerian_number("₦2,500.00"))  # 2500.0
print(parse_nigerian_number("1,500.50"))   # 1500.5
print(parse_nigerian_number("3000"))       # 3000.0
print(parse_nigerian_number("  250  "))    # 250.0