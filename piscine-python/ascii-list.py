import string
def ascii_char(text):
    ascii_list = []
    for char in text:
        ascii_list.append(ord(char))
    return ascii_list
        #print(ord("A"))
       # print(f"{char}: {ord(char)}")


ascii_result = ascii_char(string.ascii_letters)
print(ascii_result)
