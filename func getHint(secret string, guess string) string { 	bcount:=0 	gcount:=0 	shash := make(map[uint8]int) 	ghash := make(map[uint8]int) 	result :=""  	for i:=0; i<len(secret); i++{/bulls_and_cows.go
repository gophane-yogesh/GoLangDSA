// Time Complexity : O(n) Space Complexity : O(n)
func getHint(secret string, guess string) string {
	bcount:=0
	gcount:=0
	shash := make(map[uint8]int)
	ghash := make(map[uint8]int)
	result :=""

	for i:=0; i<len(secret); i++{ // as both strings are same length
		if secret[i] == guess[i]{
			bcount++
		}else{
			shash[secret[i]]++
			ghash[guess[i]]++
		}
	}
	for key,value := range shash{
		value2 := ghash[key]
		if value2!=0{
			if value <= value2{
				gcount+=value
			}else{
				gcount+= value2
			}

		}


	}
	result = strconv.Itoa(bcount)+"A" + strconv.Itoa(gcount)+"B"
	return result
	




}

