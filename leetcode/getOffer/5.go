package getOffer

func replaceSpace(s string) string {

	t := []byte(s)
	res := []byte{}
	flag := 0
	for i := 0; i < len(t); i++ {
		if t[i] == ' ' {
			res = append(res, t[flag:i-1]...)
			res = append(res, '%', '2', '0')
			flag = i + 1
		}
	}
	res = append(res, t[flag:]...)
	return string(res)
}
