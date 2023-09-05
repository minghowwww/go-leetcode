package getOffer

func ReverseLeftWords(s string, n int) string {
	res := []byte(s)
	for i := n; i < len(s); i++ {
		res[i-n] = s[i]
	}
	for i := 0; i < n; i++ {
		res[len(s)-n+i] = s[i]
	}
	return string(res)
}
