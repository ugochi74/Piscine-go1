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