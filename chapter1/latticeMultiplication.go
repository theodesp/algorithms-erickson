package chapter1

import (
	"fmt"
	"math"
	"strings"
)

/*
	Lattice Multiplication
 */
func LatticeMultiplication(x,y string) string  {
	hold := 0
	n := len(x)
	m := len(y)
	xRev := Reverse(x)
	yRev := Reverse(y)
	result := strings.Builder{}
	for k:=0;k<m+n-1;k+=1 {
		for i:=0;i<n;i+=1 {
			for j:=0;j<m;j+=1 {
				if i+j==k {
					hold = hold + int(xRev[i]-48) + int(yRev[j]-48)
				}
			}
		}
		result.WriteString(fmt.Sprintf("%d", hold % 10))
		hold = int(math.Floor(float64((hold)/10)))
	}
	res := result.String()
	return res
}


func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}