/*
package makeOddGenerator returns odd numbers
*/

package main

func makeOddGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 1
		return
	}
}

func main() {
	nextOdd := makeOddGenerator()
	fm.Println(nextOdd())
	fm.Println(nextOdd())
	fm.Println(nextOdd())
}
