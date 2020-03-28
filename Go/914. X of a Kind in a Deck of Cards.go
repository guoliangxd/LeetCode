package leetcode

/*In a deck of cards, each card has an integer written on it.

Return true if and only if you can choose X >= 2 such that it is possible to split the entire deck into 1 or more groups of cards, where:

Each group has exactly X cards.
All the cards in each group have the same integer.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/x-of-a-kind-in-a-deck-of-cards
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func hasGroupsSizeX(deck []int) bool {
	cnts := make([]int, 10001)
	size := len(deck)
	for i := 0; i < size; i++ {
		cnts[deck[i]]++
	}
	GCD := 0
	for i := 0; i < 10001; i++ {
		if cnts[i] != 0 {
			if GCD == 0 {
				GCD = cnts[i]
				continue
			}
			temp := gcd(GCD, cnts[i])
			if temp > 1 {
				GCD = temp
			} else {
				return false
			}
		}
	}
	return GCD > 1
}

func gcd(x, y int) int {
	for x%y != 0 {
		x, y = y, x%y
	}
	return y
}
