class Solution:

    def find_solution(self, n):
        if (n > 10):
            a = 10
            b = n - 10
            print("{0} {1}".format(a, b))
        else:
            print("{0} {1}".format(n, 0))


def main():
    """
    Takes care of the problem input and output.
    """
    cases = int(input())
    solution = Solution()
    for caseno in range(0, cases):
        n = int(input())
        solution.find_solution(n)

# Run the program
main()
