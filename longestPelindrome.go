// Time Complexity : O(N) 
// Space Complexity : O(N)

func longestPalindrome(s string) int {

	var hash [128]int
	result := 0
	for _, value:= range s{
		hash[value]++
	}
	for _, value:= range hash{
		result+= (value/2)*2
		
		if result % 2 == 0 && value % 2 == 1{
			result++
		}
	}

	return result
}
