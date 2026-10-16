LESSON CONTENT

Python uses indentation to show structure
In many human languages, punctuation shows where one thought ends and another begins. In Python, indentation shows which lines belong together.

Indentation is the whitespace at the beginning of a line. It is usually four spaces per level, but you can use other amounts as long as you are consistent. Python uses this whitespace to understand the logical structure of your code.

This is different from many languages. Languages like C, Java, and JavaScript use curly braces {} to mark blocks. Python uses indentation instead. This makes Python code visually clean, but it also means that indentation is not optional.



Statements
A statement is a single instruction in Python. It is a line of code that performs an action.

Examples of simple statements:

name = "Amaka"
print(name)
age = 25
Each line is a statement. When Python runs the script, it executes these statements in order from top to bottom.

A simple statement usually occupies one line. It does not require indentation beyond the baseline. A script with only simple statements has all lines starting at the left margin.

Compound statements
A compound statement is a statement that contains other statements. It usually has a header line ending with a colon, followed by an indented block of code.

The if statement is a compound statement.

if age >= 18:
    print("You can open the account")
    print("Please provide your details")
Here:

if age >= 18: is the header.
The colon at the end says, “A block is coming.”
The next two lines are indented. They form the block.
Both indented lines run only if the condition is true.
Compound statements allow you to group multiple actions under one condition, loop, function, or class.

Other compound statements include for, while, def, and class.



Colons and code blocks
In Python, a colon : at the end of a header line introduces a code block. The indented lines after the colon are the body of that block.

A code block is a group of statements that belong together. When the condition is true, the whole block runs. When the condition is false, the whole block is skipped.

For example:

balance = 2000
if balance < 5000:
    print("Balance is low")
    print("Please recharge")
else:
    print("Balance is enough")
The if block contains two indented lines. The else block contains one indented line. The colon after else: also introduces a block.

Without the colons, Python cannot tell where the block begins.

If you write:

if balance < 5000
    print("Balance is low")
Python raises a SyntaxError because the colon is missing. It expected a colon after the condition.

Think of the colon as the phrase, “Now listen carefully.” It signals that the next indented lines are part of this structure.



Indentation as structural logic
Indentation is not decoration. It is the structure.

Consider this code:

print("Start")
if True:
    print("Inside block")
    print("Still inside block")
print("End")
The output is:

Start
Inside block
Still inside block
End
The two middle lines are inside the if block because they are indented. The last line is not indented, so it is outside the block and always runs.

Now consider:

print("Start")
if False:
    print("Inside block")
    print("Still inside block")
print("End")
Output:

Start
End
The block is skipped because the condition is false. The indentation tells Python which lines are controlled by the condition.

If the indentation is wrong, the structure changes.

print("Start")
if True:
    print("Inside block")
print("Still inside block")
print("End")
Output:

Start
Inside block
Still inside block
End
Now Still inside block is not controlled by the if. It runs regardless. The indentation change altered the logic.

This is why Python programmers are careful about indentation. A single space can change what a program does.



Tabs, spaces, and TabError
The most dangerous indentation issue is mixing tabs and spaces.

Python allows you to indent with spaces or tabs, but you must be consistent. If you mix them in a way that Python cannot resolve, it raises a TabError or IndentationError.

A tab is a special character. A space is another character. They look similar in some editors but are different to Python.

Imagine a street where some houses use meters and others use feet. If you say “Go three units,” the meaning depends on the unit. Mixing units creates confusion.

In Python, mixing tabs and spaces is like mixing units. The interpreter may not know how many spaces your tab represents.

Most style guides recommend using four spaces per indentation level. You can configure your editor to insert four spaces when you press Tab. This avoids the problem entirely.

If you copy code from different sources, tabs may sneak in. That is why many errors happen after pasting code from a website into your editor. The original code may use tabs, while your editor uses spaces.

To avoid this, many programmers use an editor that displays whitespace and converts tabs to spaces automatically. Tools like python -m tabnanny can also check for ambiguous indentation.



Indentation levels can be nested
You can have blocks inside blocks. Each nested block is indented further.

For example:

balance = 8000
account_active = True
if account_active:
    print("Account is active")
    if balance >= 5000:
        print("Balance is sufficient")
    else:
        print("Insufficient balance")
else:
    print("Account is inactive")
The outer if controls a block. Inside that block, there is another if statement. The inner if has its own block, indented further. This nesting shows which condition controls which lines.

Python uses the indentation to keep track of nesting. The inner block is indented two levels: eight spaces if you use four per level.



Common mistakes and misconceptions
A common mistake is to forget the colon after an if, else, for, while, def, or class header. This immediately raises a SyntaxError.

Another mistake is to indent a line that should not be indented. For example, an else should be at the same level as its matching if. If you indent else differently, Python may raise an IndentationError or associate it incorrectly.

A third mistake is to use inconsistent indentation within the same block. For example, the first line in a block is indented four spaces, and the second is indented two spaces. Python may still accept it if the block is consistent in its own way, but it is confusing and can lead to logic errors.

A fourth mistake is to mix tabs and spaces. This can cause a TabError even if the code looks fine on screen.

A fifth mistake is to think indentation is just for appearance. It is not. Indentation defines which lines belong to which block. Changing indentation can change program behavior.



Summary
Python uses indentation to define code blocks.
A statement is a single instruction.
A compound statement has a header line ending with a colon and an indented body.
The colon signals the start of a block.
Indentation is structural logic, not decoration.
Mixing tabs and spaces can cause TabError.
Use four spaces per indentation level and be consistent.
Nested blocks are indented further.


Additional Resources
Python Documentation: Compound statements
Python PEP 8: Indentation
Real Python: Python Indentation
GeeksforGeeks: Indentation in Python
Lessons in this topic
01
LESSON 5.1: PYTHON KEYWORDS, SYNTAX & THE STRICT COMPILER
02
LESSON 5.1.1 RESERVED KEYWORDS
03
LESSON 5.1.2 THE INDENTATION ENGINE
04
LESSON 5.1.3 SYNTAX ERRORS VS. RUNTIME ERRORS
05
LESSON 5.1.4 CODE AS DOCUMENTATION
In this lesson

LESSON CONTENTPython uses indentation to show structure

Statements

Compound statements

Colons and code blocks

Indentation as structural logic

Tabs, spaces, and TabError

Indentation levels can be nested

Common mistakes and misconceptions
