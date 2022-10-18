// Time : 5 ms	Space: 2.5 MB

func isIsomorphic(s string, t string) bool {
	mp1 := make(map[byte]byte)
	mp2 := make(map[byte]bool)

	for i := 0; i < len(s); i++ {
		value1, ok1 := mp1[s[i]]
		value2, ok2 := mp2[t[i]]

		if ok1 != true && ok2 != true {
			mp1[s[i]] = t[i]
			mp2[t[i]] = true
		} else if ok1 != true && ok2 == true {
			return false
		} else if ok1 && value1 != t[i] {
			return false
		} else if ok2 == false && value2 {
			return false
		}

	}

	return true

}
