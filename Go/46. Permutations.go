package leetcode

/*Given a collection of distinct integers, return all possible permutations.*/

var res [][]int

func permute(nums []int) [][]int {
	res = make([][]int, 0)
	track := make([]int, 0, len(nums))
	backtrack(nums, track)
	return res
}

func backtrack(left []int, track []int) {

	if len(track) == len(left) {
		temp := make([]int, len(left))
		copy(temp, track)
		res = append(res, temp)
		return
	}

	for i := 0; i < len(left); i++ {
		find := false
		for j := 0; j < len(track); j++ {
			if left[i] == track[j] {
				find = true
				break
			}
		}
		if find {
			continue
		}

		track = append(track, left[i])
		backtrack(left, track)

		track = track[:len(track)-1]
	}
}
