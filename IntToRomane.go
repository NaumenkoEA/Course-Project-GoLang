package main

func main() {

	intToRoman(110)

}

func intToRoman(num int) string {
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	res, i := "", 0
	for num != 0 {
		for num < nums[i] {
			i++

		}
		num -= nums[i]
		res += symbols[i]

	}
	return res
}
