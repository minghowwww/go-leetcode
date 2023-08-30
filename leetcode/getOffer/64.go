package getOffer

/*
*
求 1+2+...+n ，
要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。
*/
func SumNums(n int) int {
	sum := n
	doSum(&sum, n-1)
	return sum
}
func doSum(sum *int, a int) bool {
	*sum += a
	bo := a > 0 && doSum(sum, a-1)
	return bo
}
