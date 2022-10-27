// Time COmplexity : O(n) Space Complexity O(n)

func uniquePaths(m int, n int) int {
    memo := make(map[[2]int]int,m+n)
    return solve(m,n,memo)
}


func solve( m, n int ,memo map[[2]int]int) int{
    key:=[2]int{m,n}
    if m == 0 && n==0{
        return 0
    }
    if m==1 || n==1{
        return 1
    }
    if val,ok:=memo[key]; ok{
        return val
    }
    memo[key] = solve(m-1,n,memo) + solve (m, n-1, memo)
    return memo[key]
    
}
