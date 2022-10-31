// 1st Solution not much optimized
func longestCommonPrefix(strs []string)  string {
    if len(strs)<2{
        return strs[0]
    }

	result:= commonTwo(strs[0],strs[1])
    if result == "" || len(strs) == 2{
		return result
	}
	for i:=2; i<len(strs); i++{
		result=commonTwo(result,strs[i])
		if result == ""{
			return result
		}

	}
	return result


}
func commonTwo(s, t string) string{
	l:=0
	if len(s)<len(t){
		l=len(s)
	}else{
		l=len(t)
	}
	result :=""

	for i:=0; i<l; i++{
		if s[i] == t[i]{
			result+=string(s[i])
        }else{
            break
        }
	}
	return result

}
