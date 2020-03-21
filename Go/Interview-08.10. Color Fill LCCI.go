package leetcode

/*Implement the "paint fill" function that one might see on many image editing programs. That is, given a screen (represented by a two-dimensional array of colors), a point, and a new color, fill in the surrounding area until the color changes from the original color.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/color-fill-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if newColor == image[sr][sc] {
		return image
	}

	fill(image, sr, sc, newColor, image[sr][sc])

	return image
}

func fill(image [][]int, sr int, sc int, nColor int, oColor int) {
	if sr < 0 || sr >= len(image) || sc < 0 || sc >= len(image[0]) {
		return
	}

	if image[sr][sc] == oColor {
		image[sr][sc] = nColor
		fill(image, sr-1, sc, nColor, oColor)
		fill(image, sr+1, sc, nColor, oColor)
		fill(image, sr, sc-1, nColor, oColor)
		fill(image, sr, sc+1, nColor, oColor)
	}
}
