def check_number(text):
    for character in text:
        if character .isdigit():
            return True
    return False
print(check_number("Hello123"))
print(check_number("11234"))
print(check_number("hellothere"))
print(check_number("hi"))