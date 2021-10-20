package problem.solving;

import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int num = sc.nextInt();
        long second, x, y, root, lackings;

        for (int i=1; i<=num; i++){
            second = sc.nextInt();
            root = (long) Math.ceil(Math.sqrt(second * 1.0));
            lackings = root * root - second;
            if (lackings < root){
                y = root;
                x = lackings + 1;
            }
            else{
                x = root;
                y = second - (root-1) * (root-1);
            }
            if(root % 2 == 0){
                long temp = x;
                x = y;
                y = temp;
            }
            System.out.printf("Case %1$d: %2$s %3$s\n", i, (int)x, (int)y);
        }
    }
}
