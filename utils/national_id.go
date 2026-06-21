package utils

func IsValidIranianNationalID(code string) bool {
	if len(code) != 10 {
		return false
	}

	for _, ch := range code {
		if ch < '0' || ch > '9' {
			return false
		}
	}

	allSame := true
	for i := 1; i < 10; i++ {
		if code[i] != code[0] {
			allSame = false
			break
		}
	}
	if allSame {
		return false
	}

	sum := 0
	for i := 0; i < 9; i++ {
		sum += int(code[i]-'0') * (10 - i)
	}

	remainder := sum % 11
	checkDigit := int(code[9] - '0')

	if remainder < 2 {
		return checkDigit == remainder
	}

	return checkDigit == 11-remainder
}