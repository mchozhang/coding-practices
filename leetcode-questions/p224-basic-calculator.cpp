/**
 * 227. Basic Calculator
 * https://leetcode.com/problems/basic-calculator/
 * 
 * submission : faster than 45%
 */

#include <iostream>
#include <vector>
using namespace std;

void addNum(vector<int> &stack, char sign, int num)
{
    if (sign == '+')
    {
        stack.push_back(num);
    }
    else if (sign == '-')
    {
        stack.push_back(-num);
    }
    else if (sign == '*')
    {
        int num2 = stack.back();
        stack.pop_back();
        stack.push_back(num2 * num);
    }
    else if (sign == '/')
    {
        int num2 = stack.back();
        stack.pop_back();
        stack.push_back(num2 / num);
    }
}

int sum(vector<int> &stack)
{
    int res = 0;
    for (int value : stack)
    {
        res += value;
    }
    return res;
}

int calculate(string s)
{

    s += '+';
    vector<vector<int>> stack(1, vector<int>(0));
    vector<char> signStack(1, '+');
    char sign = '+';
    int num = 0;
    for (char c : s)
    {
        if (c == ' ')
        {
            continue;
        }

        if (c == '(')
        {
            signStack.push_back(sign);
            sign = '+';
            stack.push_back(vector<int>(0));
        }
        else if (c == ')')
        {
            addNum(stack.back(), sign, num);
            num = sum(stack.back());
            stack.pop_back();

            sign = signStack.back();
            signStack.pop_back();
        }
        else if ('0' <= c && c <= '9')
        {
            num = num * 10 - '0' + c;
        }
        else
        {
            addNum(stack.back(), sign, num);
            num = 0;
            sign = c;
        }
    }
    return sum(stack.back());
}

int main()
{
    // cout << calculate("1+23+(11+17)-(2*4)") << endl;
    cout << calculate("1+(23*32+98/14)/2+3*3/(11+17)") << endl;
    // cout<<calculate("1+(2+3+8-4)-2+3-3+(1+1)")<<endl;
    // cout<<calculate("(1+(4+5+2)-3)+(6+8)")<<endl;
    
    return 0;
}