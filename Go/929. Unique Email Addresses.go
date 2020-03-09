package leetcode

/*Every email consists of a local name and a domain name, separated by the @ sign.

For example, in alice@leetcode.com, alice is the local name, and leetcode.com is the domain name.

Besides lowercase letters, these emails may contain '.'s or '+'s.

If you add periods ('.') between some characters in the local name part of an email address, mail sent there will be forwarded to the same address without dots in the local name.  For example, "alice.z@leetcode.com" and "alicez@leetcode.com" forward to the same email address.  (Note that this rule does not apply for domain names.)

If you add a plus ('+') in the local name, everything after the first plus sign will be ignored. This allows certain emails to be filtered, for example m.y+name@email.com will be forwarded to my@email.com.  (Again, this rule does not apply for domain names.)

It is possible to use both of these rules at the same time.

Given a list of emails, we send one email to each address in the list.  How many different addresses actually receive mails?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-email-addresses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

import "strings"

func numUniqueEmails(emails []string) int {
	mapEmail := make(map[string]int, 0)
	for i := 0; i < len(emails); i++ {
		subEmail := strings.Split(emails[i], "@")
		subLocalName := strings.Split(subEmail[0], "+")
		address := strings.Join(strings.Split(subLocalName[0], "."), "") + "@" + subEmail[1]
		_, ok := mapEmail[address]
		if !ok {
			mapEmail[address] = 1
		}
	}
	return len(mapEmail)
}
