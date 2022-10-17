#first Solution 20 ms	6.6 MB
func pivotIndex(nums []int) int {
    
	 var leftsum  []int
	 var rightsum []int

 	 x:=0


	 for _, value := range nums {
	 	x+=value
		 leftsum = append(leftsum,x)
	 }
	x=0
	 for i:=len(nums)-1; i>=0; i-- {
	 	x+=nums[i]
		 rightsum = append(rightsum, x)
	 }

	left:=0
	right:=len(nums)-1
	for left<len(nums){
		if leftsum[left] == rightsum[right] {
			return left
		}
		left++
		right--
	}
    return -1
}

#2nd Solution 	18 ms	6.2 MB
func pivotIndex(nums []int) int {
    
    sum:=0
    leftsum:=0
    
    for _, value :=range nums{
        sum+=value
}
    for i:=0; i<len(nums); i++ {
        if leftsum == sum - leftsum - nums[i] {
            return i
        }
        leftsum += nums[i]
}
    return -1
