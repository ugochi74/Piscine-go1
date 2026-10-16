LESSON CONTENT

Two kinds of problems
When Python code fails, the error usually falls into one of two categories: a syntax error or a runtime error.

A syntax error happens when Python cannot understand the structure of your code. The code breaks the grammar of the language. Python refuses to run any part of the program because it cannot parse the file.

A runtime error, also called an exception, happens while the program is running. The syntax is valid, but the program attempts an invalid operation, such as dividing by zero or using an undefined variable.

Understanding the difference saves time. Syntax errors are caught before execution. Runtime errors appear during execution. Both are important, but they require different fixes.



SyntaxError
A SyntaxError occurs when the Python parser encounters invalid syntax. Common causes include:

Missing colons.
Unclosed strings.
Missing parentheses.
Using a keyword as a variable.
Invalid assignment.
Example of missing colon:

if balance > 1000
    print("High balance")
When you run this, Python prints something like:

 File "script.py", line 1
    if balance > 1000
                     ^
SyntaxError: expected ':'
Python points to the location where it expected a colon. The caret ^ shows the approximate position.

Example of unclosed string:

print("Hello)
Python may report:

 File "script.py", line 1
    print("Hello)
               ^
SyntaxError: unterminated string literal
Python noticed that the string did not end before the line ended.

Syntax errors are usually easy to fix once you read the message. Python tells you what it expected and where.



IndentationError / TabError
An IndentationError occurs when the indentation of a block is inconsistent or unexpected.

For example, if you indent the first line of a block with four spaces and the next with two spaces, Python may still accept it if the block is internally consistent, but if the indentation does not match any expected level, it raises an error.

A common IndentationError happens when you write an if block but the body is not indented:

if True:
print("No indentation")
Python reports:

 File "script.py", line 2
    print("No indentation")
    ^
IndentationError: expected an indented block
A TabError is a specific type of IndentationError that occurs when tabs and spaces are mixed inconsistently.

For example:

if True:
    print("Spaces")   # four spaces
\tprint("Tab")         # a tab
Depending on how Python interprets the tab, this may cause a TabError: inconsistent use of tabs and spaces in indentation.

Python enforces consistent indentation. When it cannot decide whether your indentation level should be four spaces or one tab, it stops.



Reading error messages
Python error messages are informative. They contain:

The file name.
The line number.
The offending line.
A caret pointing to the location.
The type of error.
A brief description.
Let us examine a full traceback:

Traceback (most recent call last):
  File "script.py", line 3, in <module>
    result = 10 / 0
ZeroDivisionError: division by zero
This tells us:

The program ran until line 3.
The line result = 10 / 0 caused the error.
The error is ZeroDivisionError.
The reason is division by zero.
For syntax errors, the traceback may be shorter because the program never starts. But it still includes file, line, and caret.

Learning to read these messages is a superpower. Instead of guessing, you can go directly to the line and fix the issue.



File and line information
Every error message includes the file name and line number.

For example:

 File "bank.py", line 5, in <module>
This means the problem is in the file bank.py, on line 5, at the top level of the module.

If the error is inside a function, the traceback may show multiple frames. Each frame tells you the path from the main program to the failing line.

For now, focus on the bottom line of the traceback. That is usually where the actual error occurred. The lines above show the call chain.

The line number may not always be exactly the line where the mistake is, especially for missing colons. Python may point to the line where it noticed the problem. But the line number is still your best starting point.



Locating the source of the error
When Python reports an error, do the following:

Read the error type and message.
Find the file and line number.
Open the file and go to that line.
Look for the obvious issue: missing colon, wrong indentation, undefined name, unclosed string.
Fix the error.
Run the script again.
If the error is a SyntaxError with a caret, look at the character before the caret. That is often where the problem is.

For example:

if age >= 18
         ^
SyntaxError: expected ':'
The caret is under the end of the condition. The fix is to add a colon:

if age >= 18:
    print("Adult")
If the error is a NameError, the problem is that you used a variable before assigning it, or you misspelled it.

print(account_balance)
If account_balance was never assigned, Python raises:

NameError: name 'account_balance' is not defined
The fix is to define it before use or correct the spelling.



Syntax errors are compile-time errors
A syntax error is detected during the compilation phase, before the program runs. Python compiles your source to bytecode before executing. If the syntax is invalid, compilation fails, and the program never starts.

This is why syntax errors prevent any output from the script. Even if the first line is a correct print, the whole file is parsed first. If there is a syntax error later in the file, nothing runs.

This is different from a runtime error, which occurs after the program has started and executed some lines.

Think of a teacher marking a script. If the student wrote complete nonsense, the teacher fails it before reading. That is a syntax error. If the student wrote correctly but attempted an impossible calculation, the teacher reads it and then marks the calculation wrong. That is a runtime error.



Runtime errors are execution-time errors
A runtime error happens when the program is already running. The syntax was valid, but an operation failed.

Common runtime errors include:

ZeroDivisionError
NameError
TypeError
ValueError
FileNotFoundError
For example:

age = int(input("Enter age: "))
print("Next year you will be", age + 1)
If the user enters abc, the line int("abc") raises a ValueError at runtime. The syntax was perfect, but the conversion failed because abc cannot become an integer.

Runtime errors can be handled with try and except, but that is a later topic. For now, focus on understanding that they happen during execution.



Common mistakes and misconceptions
A common mistake is to think all errors are syntax errors. Many are runtime errors. Read the error type.

Another mistake is to ignore the line number. The line number is your map. Use it.

A third mistake is to try to fix the error by changing the line above the reported line without checking the actual message. Python often points to the exact line with the problem.

A fourth mistake is to assume that if the error message mentions a line, that line is always wrong. For missing colons, the line before the reported line may be the culprit. But usually the message is precise.

A fifth mistake is to copy an error message without reading it. Read it. It contains the answer.

A sixth mistake is to be afraid of error messages. Errors are not personal. They are Python’s way of telling you exactly what to fix.



Summary
Syntax errors occur when Python cannot parse the code.
Runtime errors occur during execution.
SyntaxError, IndentationError, and TabError are common syntax-related errors.
Python error messages include file, line, caret, and description.
Read the error type and message first.
Use the line number to locate the problem.
Syntax errors happen during compilation, before execution.
Runtime errors happen after some code has run.


Additional Resources
Python Documentation: Errors and Exceptions
Real Python: Understanding Python Tracebacks
GeeksforGeeks: Python Syntax Errors
W3Schools: Python Syntax Errors