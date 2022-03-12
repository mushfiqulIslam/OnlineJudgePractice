

def main():
    num_of_case = int(input())
    for i in range(1, num_of_case+1):
        length = int(input())
        l = {}
        ast = input()
        a_list = ast.split(" ")
        total = 0
        for j in range(length):
            a = int(a_list[j])
            b = a * -1
            if l.get(a, False):
                total -= l[a]
                l[a] = a
                total += a
            elif l.get(b, False):
                total -= l[b]
                l[b] = a
                total += a
            else:
                l[a] = a
                total += a
                
        print("Case %d: %d" % (i, total))


if __name__ == '__main__':
    main()