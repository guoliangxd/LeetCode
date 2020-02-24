package leetcode

/*Given a valid (IPv4) IP address, return a defanged version of that IP address.
A defanged IP address replaces every period "." with "[.]".

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/defanging-an-ip-address
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

import "strings"
func defangIPaddr(address string) string {
    subStr := strings.Split(address, ".")
    return strings.Join(subStr, "[.]")
}