// time complexity : O(m*n)
//space complexity : O(1)

var r , c int
func numIslands(grid [][]byte) int {
    r,c  = len(grid), len(grid[0])
    
    count:=0
    for i:=0; i<r; i++{
        for j:=0; j<c; j++{
            if grid[i][j]=='1'{
                count++
                 dfs(grid,i,j)
            }
           
        }
    }
    return count
    
}

func dfs(inp [][]byte, x int, y int){
    if x<0 || y< 0 || x >= r || y >= c || inp[x][y] == '0'{
            return
    }
    inp[x][y] = '0'
    dfs(inp, x+1, y )
    dfs(inp, x-1,y)
    dfs(inp, x, y+1)
    dfs(inp, x, y-1)
}
