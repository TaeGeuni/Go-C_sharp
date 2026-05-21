// 0001-two-sum
using System;
using System.Collections.Generic;
public class Solution
{
    public int[] TwoSum(int[] nums, int target)
    {
        int[] res = new int[2];
        Dictionary<int, int> numMap = new Dictionary<int, int>();

        for (int i = 0; i < nums.Length; i++)
        {
            if (numMap.ContainsKey(target - nums[i]))
            {
                res[0] = i;
                res[1] = numMap[target - nums[i]];
                break;
            }
            numMap[nums[i]] = i;
        }
        return res;
    }
}