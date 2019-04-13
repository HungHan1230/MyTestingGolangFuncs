package transposematrix

import "fmt"

// Transpose comment.
// golang 有一個很有趣的設計，它並沒有「顯示」的宣告 public、private，
// 它的變數、func，可存取範圍是以 package 為單位，
// 也就是同一個 package 下，並沒有區分 public、private，
// 那不同 package 間要如何區分 public、private呢？
// 答案很簡單就是用首字的大小寫來做辨識，只要首字大寫，golang complie，就會把它視為 public ，
// 外部 package 也都能存取，小寫就是反之。
func Transpose(slice [][]string) [][]string {
	fmt.Println("Origin Tranpose", slice)

	return privateTranspose(slice)

	// xl := len(slice[0])
	// yl := len(slice)
	// result := make([][]string, xl)
	// for i := range result {
	// 	result[i] = make([]string, yl)
	// }
	// for i := 0; i < xl; i++ {
	// 	for j := 0; j < yl; j++ {
	// 		result[i][j] = slice[j][i]
	// 	}
	// }
	// return result
}
func privateTranspose(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}
