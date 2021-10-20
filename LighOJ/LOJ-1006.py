class Solution:
    def __init__(self, a, b, c, d, e, f):
        self.value_list = []
        self.value_list.append(a)
        self.value_list.append(b)
        self.value_list.append(c)
        self.value_list.append(d)
        self.value_list.append(e)
        self.value_list.append(f)


    def fn(self, n):
        for i in range(6, n+1):
            self.value_list.append(self.value_list[i-1] + self.value_list[i-2] + self.value_list[i-3] + self.value_list[i-4] + self.value_list[i-5] + self.value_list[i-6])
        return self.value_list[n]



def main():
    cases = int(input())
    for caseno in range(0, cases):
        str = input().split()
        a = int(str[0])
        b = int(str[1])
        c = int(str[2])
        d = int(str[3])
        e = int(str[4])
        f = int(str[5])
        n = int(str[6])
        solution = Solution(a, b, c, d, e, f)
        print("Case {0}: {1}".format(caseno+1, (solution.fn(n) % 10000007)))

# Run the program
main()
