func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    oldcolor := image[sr][sc]
    if color!= oldcolor{
        dfs(image,sr,sc,oldcolor,color)
    }
    return image
}

func dfs (image [][]int, r int, c int, oldcolor int, color int){
    if image[r][c]==oldcolor{
        image[r][c] = color
        if r >=1 {
            dfs(image,r-1,c, oldcolor, color)
        }
        if c>=1 {
            dfs(image,r, c-1, oldcolor, color)
        }
        if r+1<len(image){
            dfs(image, r+1, c, oldcolor, color)
        }
        if c+1 <len(image[0]){
            dfs(image, r, c+1, oldcolor, color)
        }
    }
}
