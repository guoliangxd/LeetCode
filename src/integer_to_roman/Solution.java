package integer_to_roman;

/*
* Given an integer, convert it to a roman numeral.

Input is guaranteed to be within the range from 1 to 3999.*/

public class Solution {
    public String intToRoman(int num) {
        StringBuffer roman = new StringBuffer();
        String[][] ch = {
            {"","I","II","III","IV","V","VI","VII","VIII","IX"},
                {"","X","XX","XXX","XL","L","LX","LXX","LXXX","XC"},
                {"","C","CC","CCC","CD","D","DC","DCC","DCCC","CM"},
                {"","M","MM","MMM"}
        };
        roman.append(ch[3][num / 1000 % 10]);
        roman.append(ch[2][num / 100 % 10]);
        roman.append(ch[1][num / 10 % 10]);
        roman.append(ch[0][num % 10]);
        return roman.toString();
    }

    public static void main(String[] args){
        Solution i = new Solution();
        System.out.print(i.intToRoman(1000));
    }
}
