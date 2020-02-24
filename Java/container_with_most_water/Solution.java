package container_with_most_water;
/*
Given n non-negative integers a1 , a2 , ..., an , where each represents a point at coordinate (i, ai ).
n vertical lines are drawn such that the two endpoints of line i is at (i, ai ) and (i, 0).
Find two lines, which together with x-axis forms a container, such that the container contains the most water.
Note: You may not slant the container.
* */
public class Solution {
    public int maxArea(int[] height) {
        int maxArea = 0;
        int k = 0;
        int area;
        for(int i = 0;i < height.length; i ++){
            for(int j = i + 1;j < height.length; j ++){
                area = Math.min(height[i],height[j]) * (j - i);
                if(area > maxArea)
                    maxArea = area;
            }
        }
        return maxArea;
    }

    public static void main(String[] args){
        Solution i = new Solution();
        int[] test = {1,2,3,4,5,6,7,8,9};
        System.out.print(i.maxArea(test));
    }
}
