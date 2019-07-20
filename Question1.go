package LintCode

/*
Question URL :  https://www.lintcode.com/problem/a-b-problem/description
*/

/**
 * @param a: An integer
 * @param b: An integer
 * @return: The sum of a and b
 */
func AplusB(a int, b int) int {
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
