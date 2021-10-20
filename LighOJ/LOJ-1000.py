class Solution:

    def sum(self, a, b):
        """
        Returns the total number of problems.

        @param a denotes the number of problems in the first computer
        @param b denotes the number of problems in the second computer
        """
        # Your implementation here
        return a + b


def main():
    """
    Takes care of the problem input and output.
    """
    cases = int(input())
    for caseno in range(1, cases + 1):
        a, b = map(int, input().split())
        result = Solution().sum(a, b)
        print(f"Case {caseno}: {result}")

# Run the program
main()
