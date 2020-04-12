package leetcode

/*HTML entity parser is the parser that takes HTML code as input and replace all the entities of the special characters by the characters itself.

The special characters and their entities for HTML are:

Quotation Mark: the entity is &quot; and symbol character is ".
Single Quote Mark: the entity is &apos; and symbol character is '.
Ampersand: the entity is &amp; and symbol character is &.
Greater Than Sign: the entity is &gt; and symbol character is >.
Less Than Sign: the entity is &lt; and symbol character is <.
Slash: the entity is &frasl; and symbol character is /.
Given the input text string to the HTML parser, you have to implement the entity parser.

Return the text after replacing the entities by the special characters.*/

func entityParser(text string) string {
	size := len(text)
	res := make([]byte, 0, size)
	for i := 0; i < size; i++ {
		if text[i] == '&' {
			if size-i >= 4 && text[i+1:i+4] == "gt;" {
				res = append(res, '>')
				i += 3
				continue
			}
			if size-i >= 4 && text[i+1:i+4] == "lt;" {
				res = append(res, '<')
				i += 3
				continue
			}
			if size-i >= 5 && text[i+1:i+5] == "amp;" {
				res = append(res, '&')
				i += 4
				continue
			}
			if size-i >= 6 && text[i+1:i+6] == "quot;" {
				res = append(res, '"')
				i += 5
				continue
			}
			if size-i >= 6 && text[i+1:i+6] == "apos;" {
				res = append(res, '\'')
				i += 5
				continue
			}
			if size-i >= 7 && text[i+1:i+7] == "frasl;" {
				res = append(res, '/')
				i += 6
				continue
			}
			res = append(res, text[i])
		} else {
			res = append(res, text[i])
		}
	}
	return string(res)
}
