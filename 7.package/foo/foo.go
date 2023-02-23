package foo

// １文字目が大文字の場合パブリックに、小文字の場合はプライベートに
// 関数も同様
const (
	MAX = 100
	min = 1
)

func ReturnMin() int {
	return min
}
