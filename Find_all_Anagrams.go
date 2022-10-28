// Slow Solution

func findAnagrams(s string, p string) []int {
	np:= len(p)
	var result []int
	for i:=0;i<=len(s)-np; i++{
		if solve(s[i:i+np],p){
			result = append(result,i)
		}

	}
	return result
}

func solve(s string, p string) bool{
	a:= make(map[uint8]int)
	b:= make(map[uint8]int)
	for i:=range s{
		a[s[i]] ++
		b[p[i]] ++

	}
	return reflect.DeepEqual(a,b)

}

// 

func findAnagrams(s string, p string) []int {
    n:=len(s)
    m:=len(p)
    var res []int
    
     dictOrig := [26]int{}  
    for i:=0; i < m ;i++ {
        dictOrig[p[i] -'a']++
    }
    
    for i:=0; i < n; i++ {
       
        dict := [26]int{} 
        for j:=i; j < i + m && j < n; j++ {
            dict[s[j]-'a']++
        }
        if dict == dictOrig { 
            res = append(res, i)
        }
    }
    return res
}

