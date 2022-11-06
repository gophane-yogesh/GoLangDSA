// Time Complexity : O(n) Space Complexity : O(N)

func longestPalindrome(words []string) int {

	hmp := make(map[string]int)
	result:=0

	for _,value := range  words{
		hmp[value] ++
	}

	central := false
	for index, value := range hmp{

		if index[0] == index[1]{
			if value % 2 == 0 {
				result += value
			}else{
				
				result += value -1
				central = true
			}
		}else{
			rindex := string(index[1])+string(index[0])

			if hmp[rindex]!= 0 {
				rvalue := hmp[rindex]
				if value <= rvalue{
					result+= value

				}else{
					result += rvalue

				}
			}

		}

	}
	if central{
		result+=1
	}




	return 2*result

}

