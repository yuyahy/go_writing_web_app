package main

import "fmt"

type MyErr struct{}

func (me *MyErr) Error() string { return "" }

func Apply() error {
	var err *MyErr = nil
	// 型情報が設定されているため、err自体はnilではない
	return err
}

func Apply2() error {
	var err error = nil
	// 結果的に問題ないが、return nilの様に
	// 明示的にnilを返す方が望ましい
	return err
}

func main() {
	fmt.Println(Apply() == nil)
	fmt.Println(Apply2() == nil)
}
