class Solution:
    def countAndSay(self, n: int) -> str:
        res = "1"
        for i in range(1, n):
            buf = []
            count = 1
            now = res[0]
            for j in range(1, len(res)):
                if now == res[j]:
                    count += 1
                else:
                    buf.append(str(count))
                    buf.append(now)
                    now = res[j]
                    count = 1
            buf.append(str(count))
            buf.append(now)

            res = "".join(buf)

        return res
