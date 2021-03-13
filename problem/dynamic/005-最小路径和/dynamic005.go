package dynamic

func minPath(arr [][]int) int {
	if len(arr) == 0 || len(arr[0]) == 0 {
		return 0
	}
	// 获取行/列
	rows := len(arr)
	columns := len(arr[0])

	// 创建返回数组
	d := make([][]int, rows)
	for i := 0; i < len(d); i++ {
		d[i] = make([]int, columns)
	}

	d[0][0] = arr[0][0]

	// 第一列赋值
	for i := 1; i < rows; i++ {
		d[i][0] = d[i-1][0] + arr[i][0]
	}
	// 第一列赋值
	for j := 1; j < columns; j++ {
		d[0][j] = d[0][j-1] + arr[0][j]
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			d[i][j] = min(d[i-1][j], d[i][j-1]) + arr[i][j]
		}
	}
	return d[rows-1][columns-1]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
