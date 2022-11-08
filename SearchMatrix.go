// Time Complexity : O(n) Space Complexity : O(1)
func searchMatrix(matrix [][]int, target int) bool {
    
    n := len(matrix[0])-1
   return  helper(matrix,0,n,target)
    
    
}
func helper(matrix[][] int, m, n, target int) bool{

	if m >= len(matrix) || n < 0 {
		return false
	}

	if matrix[m][n]  == target {
		return true
	}
	if matrix[m][n] < target {
		return helper(matrix, m+1, n , target)
	}
		return helper(matrix, m, n-1, target)


}
