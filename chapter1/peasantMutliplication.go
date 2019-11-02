package chapter1

import "math"

/*
	Peasant Multiplication Algorithm
	x * y ==  	-> 0 if x == 0
			  	-> floor(x/2)*( y + y) if x is even
				-> floor(x/2)*( y + y) + y if x is odd

	Complexity: O(log x*y) or O(log(min(x,y))
    if we swap the x and y when x < y or y < x
 */
func PeasantMultiplication(x, y int64) int64  {
	prod := int64(0)
	for x > 0 {
		if x % 2 == 1 {
			prod = prod + y
		}
		x = int64(math.Floor(float64(x>>1)))
		y += y
	}
	return prod
}

func PeasantMultiplicationOptimal(x, y int64) int64  {
	prod := int64(0)
	for x > 0 {
		if x > y {
			y,x = x,y
		}
		if x % 2 == 1 {
			prod = prod + y
		}
		x = int64(math.Floor(float64(x>>1)))
		y += y
	}
	return prod
}