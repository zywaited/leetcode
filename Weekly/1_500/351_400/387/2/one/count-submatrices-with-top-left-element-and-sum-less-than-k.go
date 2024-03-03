package one

func countSubmatrices(grid [][]int, k int) int {
    num := 0
    m := len(grid)
    n := len(grid[0])
    prev := make([]int, n)
    curr := make([]int, n)
    for i := 0; i < m; i++ {
        sum := 0
        for j := 0; j < n; j++ {
            sum += grid[i][j]
            curr[j] = sum + prev[j]
            if curr[j] <= k {
                num++
            }
        }
        prev, curr = curr, prev
    }
    return num
}
