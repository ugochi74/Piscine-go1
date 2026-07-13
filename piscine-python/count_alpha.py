def count_alpha(text):
    count = 0
    for character in text:
        if character .isalpha():
            count += 1
    return count
print(count_alpha("Hello123"))
print(count_alpha("he11234"))
print(count_alpha("hellothere"))
print(count_alpha("hi"))