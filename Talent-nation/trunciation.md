
LESSON OVERVIEW
This next section is where things get a little interesting — and a little dangerous. You see, when you convert a float to an integer, you're telling Python to get rid of the decimal part. And when you get rid of the decimal part, you're losing information.

Think about it like this: imagine you have ₦1,500.50 in your account. Someone asks you how much money you have. If you say "I have about ₦1,500," you're giving an approximate value. If you say "I have ₦1,500.50," you're being precise. The approximate value might be fine for casual conversation, but if you're trying to balance your account, precision matters.

When you convert a float to an integer, you're essentially doing the same thing — you're approximating. And sometimes that approximation can have serious consequences.


LESSON CONTENT
What Is Truncation?
Truncation is what happens when you convert a float to an integer using int(). Python doesn't round the number — it just cuts off everything after the decimal point. This is called truncation, and it's a distinct concept from rounding.

Let me show you:

# Truncation examples with int()
print(int(3.14))    # Output: 3
print(int(3.99))    # Output: 3
print(int(5.0))     # Output: 5
print(int(-2.5))    # Output: -2 (truncation moves toward zero)
print(int(7.999))   # Output: 7
print(int(-3.2))    # Output: -3
Notice that Python always truncates toward zero. It doesn't round up or down — it just moves toward zero. This means 3.99 becomes 3, not 4. -2.5 becomes -2, not -3.

This is one of those subtle differences that can trip you up if you're not careful. If you're expecting rounding, you're going to be disappointed.

Information Loss During Conversion
When you truncate a float to an integer, you lose information — specifically, the fractional part. This information loss can be significant depending on the context.

# Demonstrating information loss
original_value = 3.999
truncated_value = int(original_value)
lost_value = original_value - truncated_value
print(f"Original: {original_value}")
print(f"Truncated: {truncated_value}")
print(f"Lost: {lost_value:.3f}")
Output:

Original: 3.999
Truncated: 3
Lost: 0.999
In this case, you lost almost 1 whole unit of value. That might not seem like much, but if you're working with currency or scientific calculations, it can be significant.


Real-World Example: Why Truncation Matters
Let me give you a practical example. Imagine you're building a payroll system. An employee works 8.7 hours and earns ₦500 per hour. The total pay is 8.7 * 500 = ₦4,350. That's correct.

But what if you truncate that 8.7 to an integer because you don't want decimal hours? The employee would get paid for only 8 hours, losing ₦350. That's not a small amount.

Here's the code:

hours = 8.7
rate = 500
# Truncating incorrectly
pay = int(hours) * rate  # 8 * 500 = 4000
print(f"Incorrect pay: ₦{pay}")
# The correct way
pay = hours * rate  # 8.7 * 500 = 4350
print(f"Correct pay: ₦{pay:.2f}")
Output:

Incorrect pay: ₦4000
Correct pay: ₦4350.00
In this case, the truncation cost the employee ₦350. That's a serious problem in a real-world application.



Converting Does Not Alter the Original Value
One important thing to understand: when you convert a value to a different type, Python does not change the original value. It creates a new value of the new type.

# The original value remains unchanged
price = 99.99
price_as_int = int(price)
print(f"Original price: {price}")        # 99.99 (still a float)
print(f"Price as integer: {price_as_int}") # 99 (new integer)
This means you can safely convert values without worrying about destroying your original data.



When Is Truncation Actually Useful?
Truncation isn't always bad. Sometimes you actually want to truncate values. For example, imagine you're building a game where the player's score is calculated as a floating point number, but you only display integer values. Truncation is perfect for that. Or maybe you're working with arrays and you need an index, which has to be an integer. Truncation can be a safe and effective way to convert your float to an integer.

Here's a case where truncation is exactly what you want:

# Calculating how many full hours you can work without taking a break
total_minutes = 135
full_hours = int(total_minutes / 60)  # 135 / 60 = 2.25, truncated to 2
print(f"Full hours: {full_hours}")  # Output: 2
In this case, you're calculating how many full hours you have. The 0.25 hour (15 minutes) is important, but if you specifically want only full hours, truncation is correct.



The Difference Between Truncation and Rounding
One of the most common mistakes programmers make is confusing truncation with rounding. They're fundamentally different operations:

Truncation: Cuts off the decimal part (moves toward zero)
Rounding: Finds the nearest integer (up or down depending on the decimal)
import math
# Truncation vs Rounding
print(f"int(3.6): {int(3.6)}")        # 3 (truncation)
print(f"round(3.6): {round(3.6)}")    # 4 (rounding)
print(f"math.floor(3.6): {math.floor(3.6)}") # 3 (floor - rounds down)
print(f"math.ceil(3.6): {math.ceil(3.6)}")   # 4 (ceiling - rounds up)
# Negative numbers show the difference clearly
print(f"int(-3.6): {int(-3.6)}")      # -3 (truncation toward zero)
print(f"round(-3.6): {round(-3.6)}")  # -4 (rounding)
print(f"math.floor(-3.6): {math.floor(-3.6)}") # -4 (floor - always down)
print(f"math.ceil(-3.6): {math.ceil(-3.6)}")   # -3 (ceiling - always up)