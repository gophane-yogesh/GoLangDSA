// Time Complexity : O(n) Space Complexity : O(1)

func searchRange(nums []int, target int) []int {
   
    if len(nums) == 0 {
        return []int {-1,-1}
    }
    
    return []int{leftFunc(nums,target), rightFunc(nums,target)}
    
}

func rightFunc(nums[] int, target int) int{
    
    left := 0
    right := len(nums)-1
    
    for (left <= right){
        mid := (left +right)/2
        
        if nums[mid] == target {
            left = mid +1
        }else if nums[mid] < target {
            left = mid + 1
            
        }else {
            right = mid -1
        }
    }
    
    if right < len(nums) && right >= 0 && nums[right] == target {
        return right
    }
    
    return -1
     
}

func leftFunc(nums[] int, target int) int {
    left :=0
    right := len(nums)-1
    
    for left <= right {
        mid := (left + right)/2
        
        if nums[mid] == target {
            right = mid-1
        }else if nums[mid] < target {
            left = mid +1
        }else{
            right = mid - 1
        }
    }
    
    if left >=0 && left < len(nums) && nums[left] == target {
        return left
    }
    return -1
}
