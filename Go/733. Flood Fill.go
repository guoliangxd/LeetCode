package leetcode

/*An image is represented by a 2-D array of integers, each integer representing the pixel value of the image (from 0 to 65535).

Given a coordinate (sr, sc) representing the starting pixel (row and column) of the flood fill, and a pixel value newColor, "flood fill" the image.

To perform a "flood fill", consider the starting pixel, plus any pixels connected 4-directionally to the starting pixel of the same color as the starting pixel, plus any pixels connected 4-directionally to those pixels (also with the same color as the starting pixel), and so on. Replace the color of all of the aforementioned pixels with the newColor.

At the end, return the modified image.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/flood-fill
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/


func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    if image[sr][sc] == newColor {
        return image
    }

    fill(image, sr, sc, image[sr][sc], newColor)
    
    return image
}

func fill(image [][]int, sr int, sc int, oldColor int, newColor int) {
    if oldColor == newColor || sr < 0 || sr >= len(image) || sc < 0 || sc >= len(image[0]) {
        return 
    }

    if oldColor == image[sr][sc] {
        image[sr][sc] = newColor
        fill(image, sr - 1, sc, oldColor, newColor)
        fill(image, sr + 1, sc, oldColor, newColor)
        fill(image, sr, sc - 1, oldColor, newColor)
        fill(image, sr, sc + 1, oldColor, newColor)
    }
}
