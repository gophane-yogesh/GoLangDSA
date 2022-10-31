// Time Complexity : O(n) Space Complexity O(1)
func lastStoneWeight(stones []int) int {

	for len(stones)>1{
		sort.Ints(stones)
		x,y := stones[len(stones)-1], stones[len(stones)-2]
		stones = stones[:len(stones)-2]
		if x == y {
			continue
		}else{
			result:=x-y
			result=abs(result)
			stones=append(stones,result)
		}


	}
	if len(stones)== 0{
		return  0
	}

	return stones[0]

}
func abs(num int)int{
	if num<0{
		return -num
	}
	return num

}

