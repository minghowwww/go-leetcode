package hot100

func IsValid(s string) bool {

	hash := map[byte]byte{}
	hash['('] = ')'
	hash['['] = ']'
	hash['{'] = '}'

	size := len(s)
	// 奇数个元素
	if size&1 == 1 {
		return false
	}
	// 尾元素为开始元素
	_, ok := hash[s[size-1]]
	if ok {
		return false
	}

	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if _, ok := hash[s[i]]; ok {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			if hash[stack[len(stack)-1]] != s[i] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
