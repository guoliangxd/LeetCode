package leetcode

/*A website domain like "discuss.leetcode.com" consists of various subdomains. At the top level, we have "com", at the next level, we have "leetcode.com", and at the lowest level, "discuss.leetcode.com". When we visit a domain like "discuss.leetcode.com", we will also visit the parent domains "leetcode.com" and "com" implicitly.

Now, call a "count-paired domain" to be a count (representing the number of visits this domain received), followed by a space, followed by the address. An example of a count-paired domain might be "9001 discuss.leetcode.com".

We are given a list cpdomains of count-paired domains. We would like a list of count-paired domains, (in the same format as the input, and in any order), that explicitly counts the number of visits to each subdomain.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subdomain-visit-count
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

import (
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	mapVisit := make(map[string]int, 0)

	for i := 0; i < len(cpdomains); i++ {
		rawData := strings.Split(cpdomains[i], " ")
		time, _ := strconv.Atoi(rawData[0])
		subDomain := strings.Split(rawData[1], ".")
		parts := len(subDomain)
		temp := ""
		for j := 0; j < parts; j++ {
			if j == 0 {
				temp = subDomain[parts-j-1]
			} else {
				temp = subDomain[parts-j-1] + "." + temp
			}

			val, ok := mapVisit[temp]
			if ok {
				mapVisit[temp] = val + time
			} else {
				mapVisit[temp] = time
			}
		}
	}

	rlt := make([]string, 0, len(mapVisit))
	for k, v := range mapVisit {
		rlt = append(rlt, strconv.Itoa(v)+" "+k)
	}
	return rlt
}
