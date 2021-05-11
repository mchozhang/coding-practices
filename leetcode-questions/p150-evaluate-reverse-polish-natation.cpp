/**
 * 150. Evaluate Reverse Polish Notation
 * https://leetcode.com/problems/evaluate-reverse-polish-notation/
 * 
 * submission: faster than 51%
 */
#include <iostream>
#include <vector>
#include <stack>
#include <string>
#include <cctype>
using namespace std;

int evalRPN(vector<string> &tokens)
{
    stack<int> stack;
    for (auto token : tokens)
    {
        if (isdigit(token[token.length()-1]))
        {
            int num = stoi(token, nullptr, 0);
            stack.push(num);
        }
        else
        {
            int num1 = stack.top();
            stack.pop();
            int num2 = stack.top();
            stack.pop();
            if (token == "+")
            {
                stack.push(num1 + num2);
            }
            else if (token == "-")
            {
                stack.push(num2 - num1);
            }
            else if (token == "*")
            {
                stack.push(num1 * num2);
            }
            else
            {
                stack.push(num2 / num1);
            }
        }
    }
    return stack.top();
}