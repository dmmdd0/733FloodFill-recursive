package main

func main() {

}

func floodFill(image [][]int, sr, sc, newColor int) [][]int {
	if image[sr][sc] != newColor {
		dfs(&image, sr, sc, image[sr][sc], newColor)
	}
	return image
}
func dfs(tmp *[][]int, a, b, c, d int) {
	if a < 0 || b < 0 || a >= len(*tmp) || b >= len((*tmp)[0]) || (*tmp)[a][b] != c {
		return
	}
	(*tmp)[a][b] = d
	dfs(tmp, a+1, b, c, d)
	dfs(tmp, a-1, b, c, d)
	dfs(tmp, a, b+1, c, d)
	dfs(tmp, a, b-1, c, d)
}
