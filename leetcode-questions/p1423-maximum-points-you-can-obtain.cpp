/**
 * 1423. Maximum Points You Can Obtain from Cards
 * https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards/
 * 
 * submission: faster than 64%
 */

#include <vector>
using namespace std;

int maxScore(vector<int>& cardPoints, int k) {
    int sum = 0;
    for (int i = 0; i < k; i++)
    {
        sum += cardPoints[i];
    }
    int max = sum;
    for (int i = 0; i < k;i++) {
        sum += cardPoints[cardPoints.size()-1-i] - cardPoints[k-i-1];
        if (sum > max) {
            max = sum;
        }
    }
    return max;
}