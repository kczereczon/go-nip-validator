package validator

func ValidatePesel(pesel string) bool {
	if len(pesel) != 11 {
		return false
	}

	weights := []int{1, 3, 7, 9, 1, 3, 7, 9, 1, 3}

	sum := 0

	for i := 0; i < 10; i++ {
		sum += int(pesel[i]-'0') * weights[i]
	}

	checksum := 10 - sum%10

	if checksum == 10 {
		checksum = 0
	}

	return checksum == int(pesel[10]-'0')
}
