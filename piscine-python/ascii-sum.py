import string
def ascii_char_sum(text):
    total_sum = 0
    for char in text:
        total_sum += ord(char)
        print(f"{char}: {ord(char)}")
    return total_sum
        

grand_total = ascii_char_sum(string.ascii_letters)
print(f"The total sum:{grand_total}")
