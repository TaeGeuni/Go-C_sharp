public class Solution
{
    public string ProcessStr(string s)
    {
        List<char> btArr = new List<char> { };

        for (int i = 0; i < s.Length; i++)
        {
            if (s[i] == '*')
            {
                if (btArr.Count != 0)
                {
                    btArr.RemoveAt(btArr.Count - 1);
                }
            }
            else if (s[i] == '#')
            {
                btArr.AddRange(btArr);
            }
            else if (s[i] == '%')
            {
                btArr.Reverse();
            }
            else
            {
                btArr.Add(s[i]);
            }
        }
        string res = new string(btArr.ToArray());
        return res;
    }
}