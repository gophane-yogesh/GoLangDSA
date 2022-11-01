// Time Complexity : O(logn) Space Complexity O(logn)
func isHappy(n int) bool {
	result:=make(map[int]int)
	result[n] = 1
	for n!=1{
		n = helper(n)

		if result[n] == 0 {
			result[n] ++
		}else{
			return false
		}

	}
	return true
}
func helper(num int) int{
	result:=0
	for num>0{
		r:=num%10
		num = num/10
		result+=r*r
	}
	return result
}
