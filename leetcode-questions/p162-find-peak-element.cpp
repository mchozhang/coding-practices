/**
 * 162. Find Peak Element
 * https://leetcode.com/problems/find-peak-element/
 * 
 * submission: faster than 67%
 */
#include <vector>
#include <iostream>
using namespace std;

int find(vector<int> &nums, int start, int end)
{
    if (start > end)
    {
        return -1;
    }

    int mid = (start + end) / 2;
    if ((mid == 0 || nums[mid] > nums[mid-1]) && 
        (mid == nums.size()-1 || nums[mid] > nums[mid+1])) {
        return mid;
    }
    
    int left = -1, right = -1;
    if (mid - 1 > -1)
    {
        if (nums[mid] > nums[mid - 1])
        {
            left = find(nums, start, mid-2);
        }
        else
        {
            left = find(nums, start, mid-1);
        }
    }
    if (left != -1) {
        return left;
    }

    if (mid + 1 < nums.size())
    {
        if (nums[mid] > nums[mid + 1])
        {
            right = find(nums, mid+2, end);
        }
        else
        {
            right = find(nums, mid+1, end);
        }
    }
    return right;
}

int findPeakElement(vector<int> &nums)
{
    return find(nums, 0, nums.size() - 1);
}

int main()
{
    vector<int> arr{1, 2, 3, 4};
    cout << findPeakElement(arr) << endl;
    return 0;
}