// TIme Complexity : O(n) Space Complexity : O(1)

func twoSum(nums []int, target int)  []int{
	var result []int
	temp := make(map[int]int)
	for i:=0; i<len(nums); i++{
		
		x:=target - nums[i]
		if temp[x]!=0{
			result = append(result, temp[x]-1,i)
            return result
        }else{
            temp[nums[i]] = i+1
        }

	}
return result

}
