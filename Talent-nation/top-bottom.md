LESSON CONTENT
Python checks branches in order
When you write an if/elif/else chain, Python does not check all conditions simultaneously. It checks them from top to bottom. The first condition that evaluates to true wins. Once a branch is selected, the rest of the chain is skipped.

This is called top-to-bottom resolution.

For example:

temperature = 29
if temperature > 40:
    print("Extreme heat")
elif temperature > 30:
    print("Hot day")
elif temperature > 20:
    print("Warm day")
else:
    print("Cool day")
Let us trace:

Check temperature > 40. 29 > 40 is false. Move on.
Check temperature > 30. 29 > 30 is false. Move on.
Check temperature > 20. 29 > 20 is true. Run print("Warm day").
Skip the else.
Output:

Warm day
Even though temperature > 20 is true, Python had to pass through the earlier false conditions first. The order matters.

 
 
First matching branch wins
If two conditions could both be true, the one that appears first wins.

score = 85
if score >= 80:
    print("Distinction")
elif score >= 50:
    print("Pass")
else:
    print("Fail")
Both score >= 80 and score >= 50 are true for 85. But Python checks score >= 80 first. It is true, so the block prints "Distinction". The elif score >= 50 is never checked.

Output:

Distinction
This is exactly the source of the grade-ordering bug you will explore in the drills. If you place a broad condition before a narrow one, the narrow one may never run.

 
 
Multiple elif branches
You can have many elif branches. They are checked in the order you write them.

day = "Wednesday"
if day == "Monday":
    print("Start of week")
elif day == "Wednesday":
    print("Midweek")
elif day == "Friday":
    print("Weekend is near")
else:
    print("Another day")
Python checks each condition in sequence. When it finds day == "Wednesday", it runs the midweek message and stops.

 
 
Independent if statements
Sometimes you do not want mutual exclusivity. You want several independent checks to run one after another. In that case, use separate if statements rather than elif.

For example, a customer can have multiple tags:

is_vip = True
has_coupon = True
is_new = False
if is_vip:
    print("VIP treatment")
if has_coupon:
    print("Coupon applied")
if is_new:
    print("New customer bonus")
Output:

VIP treatment
Coupon applied
All three if statements are independent. Each condition is checked. Because is_new is false, only the first two blocks run.

If you had written the same logic with elif, only the first true branch would run, and the coupon would not apply. That would be wrong for this use case.

So:

Use if/elif when exactly one branch should run.
Use multiple if statements when multiple branches can run independently.
 
 
Combined logical conditions vs. sequential branches
Sometimes you can replace multiple branches with a single compound condition.

Instead of:

if age >= 18:
    if has_id:
        print("Can vote")
You can write:

if age >= 18 and has_id:
    print("Can vote")
Both are correct. The compound version is flatter and often clearer.

But sometimes sequential if/elif branches express a priority. For example, if you want to check VIP first, then has_coupon, then new_customer, you need separate branches.

 
 
Shadowed or unreachable branches
An unreachable branch is one that can never run because an earlier branch always catches the condition.

Example:

age = 25
if age > 0:
    print("Positive age")
elif age > 18:
    print("Adult")
Here, the second elif age > 18 is unreachable for any positive age because the first condition age > 0 is broader. For age = 25, the first branch runs. For age = -5, neither runs. There is no case where the second branch can run if the first is false and the second is true. This is a shadowed branch.

A more subtle example:

score = 95
if score > 50:
    print("Above 50")
elif score > 90:
    print("Above 90")
The elif score > 90 is unreachable because any score above 90 is also above 50, so the first branch would catch it first. The second branch is shadowed.

To fix, order from the most specific or highest priority to the most general:

if score > 90:
    print("Above 90")
elif score > 50:
    print("Above 50")
Now both branches can run for appropriate values.

 
 
Why top-to-bottom resolution matters
Top-to-bottom resolution gives you control. You decide the priority of conditions by the order you write them. If you write them in the wrong order, your logic will break.

For example, a bank fraud system might check:
```py
if amount < 1000:
    print("No review needed")
elif amount < 100000:
    print("Standard review")
elif amount >= 100000:
    print("High-value review")
This is ordered from small to large. It is fine because the conditions are mutually exclusive by range. But if you wrote:

if amount < 100000:
    print("Standard review")
elif amount < 1000:
    print("No review needed")
```
The second branch is shadowed for amounts under 1000, because they are also under 100000. The system would never say “No review needed.” This could flood the review team with unnecessary work.

 
 
Tracing a chain manually
When tracing an if/elif/else chain, do the following:

Start at the first if.
Evaluate the condition.
If true, run its block and skip the rest of the chain.
If false, go to the next elif.
Repeat.
If no if or elif is true, run the else block if it exists.
This manual tracing skill is essential. It is exactly what the interpreter does.

Let us trace a more complex chain:
```py
balance = 4500
account_locked = False
if account_locked:
    print("Account locked")
elif balance <= 0:
    print("No funds")
elif balance < 5000:
    print("Low balance warning")
else:
    print("Balance is healthy")
account_locked is false. Skip the first block.
balance <= 0 is false. Skip.
balance < 5000 is true. Run print("Low balance warning").
Skip the else.
Output:

Low balance warning
The order ensured that account lock is checked first, then zero balance, then low balance. This is a logical priority.

```
 
Common mistakes and misconceptions
A common mistake is to write conditions in the wrong order, causing a broad condition to shadow a narrower one.

Another mistake is to use elif when multiple independent checks should run. If you want both a VIP discount and a coupon, use separate if statements.

A third mistake is to assume Python checks all conditions and runs all true ones. It does not. Only the first true branch runs in an if/elif/else chain.

A fourth mistake is to forget that else is the fallback. If you place else before elif, it is a syntax error. else must come last.

A fifth mistake is to write compound conditions that are logically impossible, such as if age < 18 and age > 65:. No one can be both under 18 and over 65. This branch is unreachable.

A sixth mistake is to trace an if/elif chain as if all elif conditions are evaluated even after a match. They are not.

 
 
Summary
Python resolves if/elif/else chains from top to bottom.
The first true condition wins.
Later branches are skipped once a match is found.
Use elif for mutually exclusive paths.
Use multiple if statements for independent checks.
Shadowed branches occur when a broad condition appears before a narrow one.
Manual tracing mirrors the interpreter’s behavior.
 
 
Additional Resources
Python Documentation: The if statement
Real Python: Conditional Logic and Flow
GeeksforGeeks: Python if-else
W3Schools: Python Conditions