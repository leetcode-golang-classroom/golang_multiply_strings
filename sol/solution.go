package sol

func multiply(num1 string, num2 string) string {
	lenNum1, lenNum2 := len(num1), len(num2)
	fixLen := lenNum1 + lenNum2 - 1
	carry := 0
	tempDigits := make([]int, fixLen)
	for digit := 0; digit < fixLen; digit++ {
		sum := 0
		// digit = lenNum1 - 1 - pos + lenNum2 - 1 - c_pos
		// c_pos = lenNum1 - 1 - pos + lenNum2 - 1 - digit
		for pos := lenNum1 - 1; pos >= 0; pos-- {
			c_pos := lenNum1 - 1 - pos + lenNum2 - 1 - digit
			if c_pos >= 0 && c_pos < lenNum2 {
				sum += int(num1[pos]-'0') * int(num2[c_pos]-'0')
			}
		}
		sum += carry
		tempDigits[fixLen-1-digit] = sum % 10
		carry = sum / 10
	}
	res := []byte{}
	if carry > 0 {
		ch := byte(carry + '0')
		res = append(res, ch)
	}
	for idx, val := range tempDigits {
		ch := byte(val + '0')
		res = append(res, ch)
		if ch-'0' == 0 && idx == 0 && carry == 0 {
			break
		}
	}
	return string(res)
}
