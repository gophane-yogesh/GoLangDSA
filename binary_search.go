// Time Complexity : O(logn)
// Space Complexity : O(1)

func search(nums []int, target int) int {
    low := 0
    high := len(nums)-1
   var mid int
    
    for low <= high {
        mid = (low +high)/2
        if nums[mid] == target{
            return mid
        }else if nums[mid] < target{
            low = mid+1
        }else{
            high = mid-1
        }
        

    }
    return -1
    
    
}
