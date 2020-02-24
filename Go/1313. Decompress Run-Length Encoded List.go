package leetcode

/*We are given a list nums of integers representing a list compressed with run-length encoding.
Consider each adjacent pair of elements [a, b] = [nums[2*i], nums[2*i+1]] (with i >= 0).
For each such pair, there are a elements with value b in the decompressed list.
Return the decompressed list.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/decompress-run-length-encoded-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func decompressRLElist(nums []int) []int {
	//第一种方法略快，占用内存偏大，第二种反之
	/*动态增加
	  rlt := make([]int, 0, 50)
	  length := len(nums)
	  for i := 0; i < length; i += 2 {
	      for j := 0; j < nums[i]; j++ {
	          rlt = append(rlt, nums[i + 1])
	      }
	  }*/
	//固定数组
	length := len(nums)
	newLength := 0
	for i := 0; i < length; i += 2 {
		newLength += nums[i]
	}
	rlt := make([]int, newLength)
	index := 0
	for i := 0; i < length; i += 2 {
		for j := 0; j < nums[i]; j++ {
			rlt[index] = nums[i+1]
			index++
		}
	}
	return rlt
}
