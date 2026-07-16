import string
def ascii_char(text):
    for char in text:
        #print(ord(char))
        #print(ord("A"))
        print(f"{char}: {ord(char)}")


ascii_char(string.ascii_letters)
