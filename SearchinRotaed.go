// Time Complexity : O(logn) Space Complexity : O(1)

func search(nums []int, target int) int {
	low :=0
	high := len(nums)-1

	for low<=high{
		mid:= (low +high)/2
		if nums[mid] == target{
			return mid
		}
		if nums[low] <= nums[mid] {
			if target>=nums[low] && target<=nums[mid]{
				high =mid-1
            }else{
                low = mid+1
            }
		}else{
			if target<=nums[high] && target>=nums[mid]{
				low = mid+1
            }else{
                high = mid-1
            }
		}



	}
	return -1
}
