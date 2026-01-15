package main

import (
    "fmt"
    "strings"
)

// SOAL 1
func upperCaseExcept(arr []string, except string) []string {
    result := make([]string, len(arr))
    for i, v := range arr {
        if v == except {
            result[i] = v
        } else {
            result[i] = strings.ToUpper(v)
        }
    }
    return result
}

// SOAL 2
func findMinMax(arr []int) []int {
    min, max := arr[0], arr[0]
    for _, v := range arr {
        if v < min {
            min = v
        }
        if v > max {
            max = v
        }
    }
    return []int{min, max}
}

// SOAL 3
func findMinRange(arr []int, start, end int) []int {
    min := arr[start]
    index := start
    for i := start; i < end; i++ {
        if arr[i] < min {
            min = arr[i]
            index = i
        }
    }
    return []int{min, index}
}

func findMaxRange(arr []int, start, end int) []int {
    max := arr[start]
    index := start
    for i := start; i < end; i++ {
        if arr[i] > max {
            max = arr[i]
            index = i
        }
    }
    return []int{max, index}
}

// SOAL 4
func evenOddOrder(arr []int) []int {
    var even, odd []int
    for _, v := range arr {
        if v%2 == 0 {
            even = append(even, v)
        } else {
            odd = append(odd, v)
        }
    }
    return append(even, odd...)
}

// SOAL 5
func rotateArray(arr []int, n int) []int {
    n = n % len(arr)
    return append(arr[n:], arr[:n]...)
}

// SOAL 6
func diagonalLeft(size int) [][]int {
    matrix := make([][]int, size)
    for i := 0; i < size; i++ {
        matrix[i] = make([]int, size)
        for j := 0; j < size; j++ {
            if i == j {
                matrix[i][j] = i + 1
            } else {
                matrix[i][j] = 10
            }
        }
    }
    return matrix
}

// SOAL 7
func diagonalRight(size int) [][]int {
    matrix := make([][]int, size)
    for i := 0; i < size; i++ {
        matrix[i] = make([]int, size)
        for j := 0; j < size; j++ {
            if j == size-1-i {
                matrix[i][j] = size - i
            } else {
                matrix[i][j] = 20
            }
        }
    }
    return matrix
}

// SOAL 8
func borderMatrix(n int) [][]int {
    matrix := make([][]int, n)
    for i := 0; i < n; i++ {
        matrix[i] = make([]int, n)
        for j := 0; j < n; j++ {
            if i == 0 || i == n-1 || j == 0 || j == n-1 {
                matrix[i][j] = i + j
            }
        }
    }
    return matrix
}

// SOAL 9
func sumMatrix(n int) [][]int {
    matrix := make([][]int, n+1)
    for i := 0; i <= n; i++ {
        matrix[i] = make([]int, n+1)
    }

    for i := 0; i < n; i++ {
        rowSum := 0
        for j := 0; j < n; j++ {
            matrix[i][j] = i + j
            rowSum += matrix[i][j]
            matrix[n][j] += matrix[i][j]
        }
        matrix[i][n] = rowSum
        matrix[n][n] += rowSum
    }
    return matrix
}

// SOAL 10
func quizScore(siswa [][]string, kunci []string) []int {
    skor := make([]int, len(siswa))
    for i, jawaban := range siswa {
        for j, v := range jawaban {
            if v == kunci[j] {
                skor[i]++
            }
        }
    }
    return skor
}

// MAIN
func main() {
    // Soal 1
    fmt.Println("Soal 1:", upperCaseExcept([]string{"code", "java", "cool"}, "java"))
    fmt.Println("Soal 1:", upperCaseExcept([]string{"black", "pink", "venom"}, "venom"))

    // Soal 2
    fmt.Println("Soal 2:", findMinMax([]int{2, 3, 4, 5, 6, 7, 8, 9, 1, 10}))

    // Soal 3
    fmt.Println("Soal 3:", findMinRange([]int{5, 3, 4, 2, 6, 7, 8, 9, 1, 10}, 0, 10))
    fmt.Println("Soal 3:", findMinRange([]int{5, 3, 4, 2, 6, 7, 8, 9, 1, 10}, 0, 7))
    fmt.Println("Soal 3:", findMaxRange([]int{1, 22, 3, 4, 5, 10, 7, 8, 9, 49}, 0, 10))
    fmt.Println("Soal 3:", findMaxRange([]int{1, 22, 3, 4, 5, 10, 7, 8, 9, 49}, 2, 7))

    // Soal 4
    fmt.Println("Soal 4:", evenOddOrder([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))

    // Soal 5
    fmt.Println("Soal 5:", rotateArray([]int{12, 15, 1, 5, 20}, 3))

    // Soal 6
    fmt.Println("Soal 6:", diagonalLeft(5))

    // Soal 7
    fmt.Println("Soal 7:", diagonalRight(5))

    // Soal 8
    fmt.Println("Soal 8:", borderMatrix(7))

    // Soal 9
    fmt.Println("Soal 9:", sumMatrix(7))

    // Soal 10
    jawabanSiswa := [][]string{
        {"A", "B", "A", "C", "C", "D", "E", "E", "A", "D"},
        {"D", "B", "A", "B", "C", "A", "E", "E", "A", "D"},
        {"E", "D", "D", "A", "C", "B", "E", "E", "A", "D"},
        {"C", "B", "A", "E", "D", "C", "D", "E", "A", "D"},
        {"A", "B", "D", "C", "C", "D", "E", "E", "A", "D"},
        {"B", "B", "E", "C", "C", "D", "E", "E", "A", "D"},
        {"B", "B", "A", "C", "C", "D", "E", "E", "A", "D"},
        {"E", "B", "E", "C", "C", "D", "E", "E", "A", "D"},
    }
    kunci := []string{"D", "B", "D", "C", "C", "D", "A", "E", "A", "D"}
    skor := quizScore(jawabanSiswa, kunci)
    for i, s := range skor {
        fmt.Printf("Jawaban Siswa %d yang benar : %d\n", i, s)
    }
}
