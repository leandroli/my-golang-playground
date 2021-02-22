package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	var findNumber func(int, int) bool
	var binarySearchHorizontal func(int, int, int) bool
	var binarySearchVertical func(int, int, int) bool
	binarySearchVertical = func(start, end, posY int) bool {
		if start > end {
			return false
		}
		mid := (start + end) / 2
		if matrix[mid][posY] == target {
			return true
		} else if matrix[mid][posY] > target {
			return binarySearchVertical(start, mid-1, posY)
		} else {
			return binarySearchVertical(mid+1, end, posY)
		}
	}
	binarySearchHorizontal = func(start, end, posX int) bool {
		if start > end {
			return false
		}
		mid := (start + end) / 2
		if matrix[posX][mid] == target {
			return true
		} else if matrix[posX][mid] > target {
			return binarySearchHorizontal(start, mid-1, posX)
		} else {
			return binarySearchHorizontal(mid+1, end, posX)
		}
	}
	findNumber = func(posX, posY int) bool {
		flag := false
		if matrix[posX][posY] == target {
			return true
		}
		if matrix[posX][posY] > target {
			flag = binarySearchHorizontal(0, posY-1, posX)
		} else {
			flag = binarySearchVertical(posX+1, len(matrix)-1, posY)
		}
		if flag {
			return true
		}
		if posY == 0 || posX == len(matrix)-1 {
			return false
		}
		return findNumber(posX+1, posY-1)
	}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	return findNumber(0, len(matrix[0])-1)
}
