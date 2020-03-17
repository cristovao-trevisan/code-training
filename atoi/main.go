package atoi

const maxInt = 2147483647
const minInt = -2147483648

func myAtoi(str string) int {
	number := ""
	plus := true
	anySet := false

	for _, c := range str {
		if len(number) > 11 {
			break
		}
		if c == '-' || c == '+' {
			if len(number) != 0 {
				break
			}
			if anySet {
				return 0
			}
			if c == '-' {
				plus = false
			}
			anySet = true
			continue
		}

		ok := c >= '0' && c <= '9'
		if ok {
			anySet = true
			if !(len(number) == 0 && c == '0') {
				number = number + string(c)
			}
		} else {
			if len(number) == 0 {
				if c == ' ' && !anySet {
					continue
				}
				return 0
			}
			break
		}
	}

	var result int = 0
	var mult int = 1

	if !plus {
		mult = -1
	}
	for i := len(number) - 1; i >= 0; i-- {
		result += mult * int(number[i]-'0')

		mult *= 10

		if result > maxInt {
			return maxInt
		}
		if result < minInt {
			return minInt
		}
	}

	return int(result)
}
