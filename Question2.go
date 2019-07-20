package LintCode

/*
Question URL :  https://www.lintcode.com/problem/trailing-zeros/description

Helper URLs: https://www.purplemath.com/modules/factzero.htm
*/

func TrailingZeros(n int64) int64 {
	var result int64 = 0
	var multiplier int64 = 5
	for n/multiplier >= 1 {
		result = AplusBInt64(result, n/multiplier)
		multiplier = AplusBInt64(multiplier<<2, multiplier)
	}
	return result
}

/**
 * @param a: An integer
 * @param b: An integer
 * @return: The sum of a and b
 */

func AplusBInt64(a int64, b int64) int64 {
	for {
		if b == 0 {
			break
		}
		carryBit := a & b //this will give the remainder bit when you add the two numbers e.g 1 & 1 = 1, 1&0=0; 0&0=0
		a = a ^ b         // This will give the addition of the two numbers
		b = carryBit << 1 // you want to add the remainder to the left of the bit added.
	}
	return a
}
