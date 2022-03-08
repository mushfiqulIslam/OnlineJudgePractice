def main():
    num_of_case = int(input())
    for i in range(1, num_of_case+1):
        output = "No"
        first_line = input()
        second_line = input()
        first_line = first_line.replace(" ", "")
        second_line = second_line.replace(" ", "")
        first_line = first_line.lower()
        second_line = second_line.lower()
        first = list(first_line)
        second = list(second_line)
        first.sort()
        second.sort()
        print(first)
        if first == second:
            output = "Yes"
        print("Case %d: %s" % (i, output))


if __name__ == '__main__':
    main()
