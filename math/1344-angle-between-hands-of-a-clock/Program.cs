public class Solution
{

    public double AngleClock(int hour, int minutes)
    {
        double h, m, res = 0;

        m = minutes * 6;
        h = hour * 30 + minutes * 0.5;

        res = Math.Max(h, m) - Math.Min(h, m);

        if (res > 180)
        {
            return 360 - res;
        }
        return res;
    }
}