def main():
    num_of_case = int(input())
    for i in range(1, num_of_case+1):
        output = "Yes"
        ip1 = input()
        ip2 = input()
        ip1_list = ip1.split(".")
        ip2_list = ip2.split(".")
        for index in range(len(ip1_list)):
            if ip2_list[index] != format(int(ip1_list[index]), '08b'):
                output = "No"
                break

        print("Case %d: %s" % (i, output))


if __name__ == '__main__':
    main()
