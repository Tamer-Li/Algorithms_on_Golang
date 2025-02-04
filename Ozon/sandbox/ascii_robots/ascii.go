// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	var in *bufio.Scanner
// 	var out *bufio.Writer

// 	in = bufio.NewScanner(os.Stdin)
// 	out = bufio.NewWriter(os.Stdout)
// 	defer out.Flush()

// 	t := scanInt(in)

// 	for i := 0; i < t; i++ {
// 		n, m := scanTwoNum(in)

// 		matrix := scanMatrix(in, n)

// 		xA, yA, xB, yB := locationRobots(matrix)

// 		xA, yA, xB, yB, pointA, pointB := locationUp(xA, yA, xB, yB)

// 		matrix, yA = refreshPointA(matrix, xA, yA, pointA)
// 		matrix, yB = refreshPointB(matrix, xB, yB, pointB)

// 		matrix = moveUp(matrix, xA, yA, pointA)
// 		matrix = moveLeft(matrix, yA, pointA)
// 		matrix = moveDown(matrix, xB, yB, pointB)
// 		matrix = moveRight(matrix, n, m, yB, pointB)

// 		output(matrix, out)
// 	}
// }

// func scanInt(scan *bufio.Scanner) int {
// 	scan.Scan()
// 	num, _ := strconv.Atoi(scan.Text())
// 	return num
// }

// func scanTwoNum(scan *bufio.Scanner) (int, int) {
// 	scan.Scan()
// 	row := scan.Text()
// 	strArr := strings.Fields(row)

// 	n, _ := strconv.Atoi(strArr[0])
// 	m, _ := strconv.Atoi(strArr[1])

// 	return n, m
// }

// func scanMatrix(scan *bufio.Scanner, rows int) []string {
// 	matrix := make([]string, 0, rows)

// 	for i := 0; i < rows; i++ {
// 		scan.Scan()
// 		matrix = append(matrix, scan.Text())
// 	}

// 	return matrix
// }

// func locationRobots(matrix []string) (xA int, yA int, xB int, yB int) {
// 	xA, yA, xB, yB = -1, -1, -1, -1

// 	for idxRow, row := range matrix {
// 		for idxCol, col := range row {
// 			if ((idxRow+1)%2 == 0) && ((idxCol+1)%2 == 0) {
// 				continue
// 			}
// 			if col == 'A' {
// 				xA = idxRow
// 				yA = idxCol
// 				continue
// 			}
// 			if col == 'B' {
// 				xB = idxRow
// 				yB = idxCol
// 				continue
// 			}

// 			if (xA != -1) && (xB != -1) {
// 				break
// 			}
// 		}
// 	}

// 	return xA, yA, xB, yB
// }

// func locationUp(xA int, yA int, xB int, yB int) (int, int, int, int, rune, rune) {
// 	pointA, pointB := 'a', 'b'
// 	Xa, Ya, Xb, Yb := xA+1, yA+1, xB+1, yB+1
// 	cA := Xa*Xa + Ya*Ya
// 	cB := Xb*Xb + Yb*Yb

// 	if cA < cB {
// 		return xA, yA, xB, yB, pointA, pointB
// 	}
// 	return xB, yB, xA, yA, pointB, pointA
// }

// func refreshPointA(matrix []string, xA int, yA int, pointA rune) ([]string, int) {
// 	if (yA+1)%2 == 0 {
// 		row := []rune(matrix[xA])
// 		row[yA+1] = pointA
// 		matrix[xA] = string(row)
// 		yA++
// 	}
// 	return matrix, yA
// }

// func refreshPointB(matrix []string, xB int, yB int, pointB rune) ([]string, int) {
// 	if (yB+1)%2 == 0 {
// 		row := []rune(matrix[xB])
// 		row[yB+1] = pointB
// 		matrix[xB] = string(row)
// 		yB++
// 	}
// 	return matrix, yB
// }

// func moveUp(matrix []string, x int, y int, point rune) []string {
// 	for idx, row := range matrix {
// 		if idx < x {
// 			matrix[idx] = row[:y] + string(point) + row[y+1:]
// 		} else {
// 			break
// 		}
// 	}
// 	return matrix
// }

// func moveDown(matrix []string, x int, y int, point rune) []string {
// 	for idx, row := range matrix {
// 		if idx <= x {
// 			continue
// 		}
// 		matrix[idx] = row[:y] + string(point) + row[y+1:]
// 	}
// 	return matrix
// }

// func moveLeft(matrix []string, y int, point rune) []string {
// 	firstRow := []rune(matrix[0])

// 	for i := 0; i < y; i++ {
// 		firstRow[i] = point
// 	}

// 	matrix[0] = string(firstRow)
// 	return matrix
// }

// func moveRight(matrix []string, n int, m int, y int, point rune) []string {
// 	lastRow := []rune(matrix[n-1])

// 	for i := y + 1; i < m; i++ {
// 		lastRow[i] = point
// 	}

// 	matrix[n-1] = string(lastRow)
// 	return matrix
// }

//	func output(matrix []string, out *bufio.Writer) {
//		for _, row := range matrix {
//			fmt.Fprintln(out, row)
//		}
//	}
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	t := scanInt(in)

	for i := 0; i < t; i++ {
		n, m := scanTwoNum(in)
		matrix := scanMatrix(in, n)

		// Находим координаты роботов
		xA, yA, xB, yB := findRobots(matrix)

		// Определяем, кто из роботов пойдет в (1, 1)
		if shouldARouteToTopLeft(xA, yA, xB, yB) {
			// Робот A идет в (1, 1), робот B идет в (n, m)
			matrix = buildRoute(matrix, xA, yA, 0, 0, 'a')
			matrix = buildRoute(matrix, xB, yB, n-1, m-1, 'b')
		} else {
			// Робот B идет в (1, 1), робот A идет в (n, m)
			matrix = buildRoute(matrix, xB, yB, 0, 0, 'b')
			matrix = buildRoute(matrix, xA, yA, n-1, m-1, 'a')
		}

		// Выводим результат
		for _, row := range matrix {
			fmt.Fprintln(out, row)
		}
	}
}

// Сканирует одно число
func scanInt(in *bufio.Scanner) int {
	in.Scan()
	num, _ := strconv.Atoi(in.Text())
	return num
}

// Сканирует два числа
func scanTwoNum(in *bufio.Scanner) (int, int) {
	in.Scan()
	row := in.Text()
	strArr := strings.Fields(row)
	n, _ := strconv.Atoi(strArr[0])
	m, _ := strconv.Atoi(strArr[1])
	return n, m
}

// Сканирует матрицу
func scanMatrix(in *bufio.Scanner, rows int) []string {
	matrix := make([]string, 0, rows)
	for i := 0; i < rows; i++ {
		in.Scan()
		matrix = append(matrix, in.Text())
	}
	return matrix
}

// Находит координаты роботов
func findRobots(matrix []string) (xA, yA, xB, yB int) {
	for i, row := range matrix {
		for j, cell := range row {
			if cell == 'A' {
				xA, yA = i, j
			} else if cell == 'B' {
				xB, yB = i, j
			}
		}
	}
	return
}

// Определяет, должен ли робот A идти в верхний левый угол
func shouldARouteToTopLeft(xA, yA, xB, yB int) bool {
	distanceA := abs(xA-0) + abs(yA-0) // Расстояние до (1, 1)
	distanceB := abs(xB-0) + abs(yB-0) // Расстояние до (1, 1)
	return distanceA < distanceB
}

// Строит маршрут для робота
func buildRoute(matrix []string, startX, startY, endX, endY int, pathChar rune) []string {
	// Преобразуем матрицу в слайс рун для удобства
	grid := make([][]rune, len(matrix))
	for i, row := range matrix {
		grid[i] = []rune(row)
	}

	// Двигаемся по вертикали
	x := startX
	for x != endX {
		if x < endX {
			x++
		} else {
			x--
		}
		if grid[x][startY] == '.' || grid[x][startY] == pathChar {
			grid[x][startY] = pathChar
		}
	}

	// Двигаемся по горизонтали
	y := startY
	for y != endY {
		if y < endY {
			y++
		} else {
			y--
		}
		if grid[x][y] == '.' || grid[x][y] == pathChar {
			grid[x][y] = pathChar
		}
	}

	// Преобразуем обратно в строки
	for i := range grid {
		matrix[i] = string(grid[i])
	}
	return matrix
}

// Возвращает модуль числа
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
