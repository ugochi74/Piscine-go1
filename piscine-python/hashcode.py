def Hash_Code(text):
    length = len(text)
    result = ""
    for char in text:
        new_ascii = (ord(char) + length) %127
        if new_ascii < 32:
            new_ascii += 33
        result += chr(new_ascii)
    return result
print(Hash_Code("A"))
print(Hash_Code("AB"))
print(Hash_Code("BAC"))
print(Hash_Code("Hello World"))