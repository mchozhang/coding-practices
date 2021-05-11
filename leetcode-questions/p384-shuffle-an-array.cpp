/**
 * 384. Shuffle an Array
 * https://leetcode.com/problems/shuffle-an-array/
 * 
 * submission: faster than 90%
 */

#include <vector>
using namespace std;

class Solution
{
private:
    vector<int> nums;

    void swap(int &i, int &j)
    {
        int t = i;
        i = j;
        j = t;
    }

public:
    Solution(vector<int> &nums)
    {
        this->nums = nums;
    }

    /** Resets the array to its original configuration and return it. */
    vector<int> reset()
    {
        return nums;
    }

    /** Returns a random shuffling of the array. */
    vector<int> shuffle()
    {
        vector<int> arr(nums);
        for (int i = 0; i < nums.size(); i++)
        {
            int j = rand() % (nums.size() - i);
            swap(arr[i], arr[i + j]);
        }
        return arr;
    }
};