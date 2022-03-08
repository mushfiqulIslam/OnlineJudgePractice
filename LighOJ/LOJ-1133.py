def multiply(val_list, multiplier):
    for x in range(len(val_list)):
        val_list[x] = int(val_list[x] * multiplier)
    return val_list


def adder(val_list, multiplier):
    for x in range(len(val_list)):
        val_list[x] = val_list[x] + multiplier
    return val_list


def divisor(val_list, multiplier):
    for x in range(len(val_list)):
        val_list[x] = int(val_list[x] / multiplier)
    return val_list


def main():
    num_of_case = int(input())
    for i in range(1, num_of_case+1):
        num_of_op = input()
        num_of_op_list = num_of_op.split(" ")
        values = input()
        values_list = values.split(" ")
        values_int_list = []
        for j in range(int(num_of_op_list[0])):
            values_int_list.append(int(values_list[j]))

        for j in range(int(num_of_op_list[1])):
            operation = input()
            operation_list = operation.split(" ")
            if operation_list[0] == "S":
                values_int_list = adder(values_int_list, int(operation_list[1]))
            elif operation_list[0] == "M":
                values_int_list = multiply(values_int_list, int(operation_list[1]))
            elif operation_list[0] == "D":
                values_int_list = divisor(values_int_list, int(operation_list[1]))
            elif operation_list[0] == "R":
                values_int_list.reverse()
            else:  # Means P
                y = int(operation_list[1])
                z = int(operation_list[2])
                values_int_list[y], values_int_list[z] = values_int_list[z], values_int_list[y]
            # print(values_int_list)

        print("Case %d:" % i)
        output = ""
        for o in range(len(values_int_list)):
            if o == 0:
                output += "%d" % values_int_list[o]
            else:
                output += " %d" % values_int_list[o]

        print(output)


if __name__ == '__main__':
    main()
