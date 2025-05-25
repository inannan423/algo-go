package main

func numIslands(grid [][]byte) int {
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i > len(grid)-1 || j < 0 || j > len(grid[0])-1 || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0' // 沉没
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	sum := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				sum += 1
				dfs(i, j)
			}
		}
	}

	return sum
}

func main() {
	// Example usage
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '1', '1'},
		{'0', '0', '0', '1', '0'},
		{'0', '1', '1', '0', '0'},
	}
	result := numIslands(grid)
	println(result) // Output: 3
}
