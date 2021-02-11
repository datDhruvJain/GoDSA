You may assume the stack only to deal with positive integers.

- Int size() //returns the total number of elements in the stack

- int isEmpty() //returns 1 of the stack is empty and 0 otherwise

- int isFull() //returns 1 of the stack is full and 0 otherwise

- top() //returns the topmost element on the stack without popping

- push(x) //push the value x to the top of the stack, returns -1 if in case the stack is full

- pop() //pops the value x from the stack, returns -1 if in case the stack is empty

implement a separate programs for each of the following problems using your stack implementation (you may have to alter your stack implementations accordingly).

1. Write a program to print the contents of an integer array in reverse order.

2. Check whether the given parenthesized arithmetic expression has balanced parenthesis or not, assuming the expression to have parenthesis types (), {}, or []. You may assume the input expression to be composed of constructs operators, operands and parentheses where each construct is a character.

3. Evaluate a given postfix expression for a given value assignment. The expression is composed of operators (‘+’,’-’,’*’,’ /’, or ‘^’) or single character operands (characters A to Z alone). Your procedure should read the expression, value of each operand from the user, compute result (error message if in case the expression is an invalid postfix expression), and output the result.