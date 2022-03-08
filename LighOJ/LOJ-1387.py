def main():
    num_of_case = int(input())
    for i in range(1, num_of_case+1):
        total_amount = 0
        num_of_op = int(input())
        print("Case %d:" % i)
        for j in range(num_of_op):
            operation = input()
            operation_list = operation.split(" ")
            if operation_list[0] == "donate":
                total_amount += int(operation_list[1])
            elif operation_list[0] == "report":
                print(total_amount)


if __name__ == '__main__':
    main()
