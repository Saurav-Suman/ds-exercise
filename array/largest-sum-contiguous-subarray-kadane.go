package main

import "fmt"

func main() {
	arr := [8]int{-2, -3, 4, -1, -2, 1, 5, -3}

	var meh int // max ending here
	var msf int // max so far

	meh = 0
	msf = 0

	for i := 0; i < 7; i++ {
		meh = meh + arr[i]
		if meh < 0 {
			meh = 0
		}
		if msf < meh {
			msf = meh
		}
	}

	fmt.Println(msf)

}
