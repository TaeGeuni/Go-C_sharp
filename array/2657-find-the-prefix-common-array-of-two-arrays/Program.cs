// 2657-find-the-prefix-common-array-of-two-arrays
using System;

public class Solution
{
    public int[] FindThePrefixCommonArray(int[] A, int[] B)
    {
        int n = A.Length;
        int[] res = new int[n];

        bool[] freq = new bool[n + 1];
        int num = 0;

        for (int i = 0; i < n; i++)
        {
            if (!freq[A[i]])
            {
                freq[A[i]] = true;
            }
            else
            {
                num++;
            }

            if (!freq[B[i]])
            {
                freq[B[i]] = true;
            }
            else
            {
                num++;
            }

            res[i] = num;
        }
        return res;
    }
}