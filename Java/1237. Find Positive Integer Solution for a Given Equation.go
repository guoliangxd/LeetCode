package leetcode

/*Given a function  f(x, y) and a value z, return all positive integer pairs x and y where f(x,y) == z.

The function is constantly increasing, i.e.:

f(x, y) < f(x + 1, y)
f(x, y) < f(x, y + 1)

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-positive-integer-solution-for-a-given-equation
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

/*
 * // This is the custom function interface.
 * // You should not implement it, or speculate about its implementation
 * class CustomFunction {
 *     // Returns f(x, y) for any given positive integers x and y.
 *     // Note that f(x, y) is increasing with respect to both x and y.
 *     // i.e. f(x, y) < f(x + 1, y), f(x, y) < f(x, y + 1)
 *     public int f(int x, int y);
 * };
 */
 class Solution {
    public List<List<Integer>> findSolution(CustomFunction customfunction, int z) {
        List<List<Integer>> rlt = new ArrayList<>();
        for(int i = 1; i <= 1000; i++) {
            int res = binarySearch(customfunction, i, z);
            if(res != -1) {
                List<Integer> temp = new ArrayList<>();
                temp.add(i);
                temp.add(res);
                rlt.add(temp);
            }
        }
        return rlt;
    }

    private int binarySearch(CustomFunction customfunction, int first, int target) {
        int begin = 1;
        int end = 1000;
        while(begin <= end) {
            int mid = (begin + end) / 2;
            int temp = customfunction.f(first, mid);
            if(temp == target) {
                return mid;
            } else if(temp < target) {
                begin = mid + 1;
            } else {
                end = mid - 1;
            }
        }
        return -1;
    }
}