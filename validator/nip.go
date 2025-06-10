package validator

import (
	"fmt"
	"os"
)

func ValidateNip(nip string) bool {

	fmt.Fprintf(os.Stdout, "Nip: %s\n", nip)

	if len(nip) != 10 {
		return false
	}

	weights := []int{6, 5, 7, 2, 3, 4, 5, 6, 7}

	sum := 0

	for i := 0; i < 9; i++ {
		sum += int(nip[i]-'0') * weights[i]
	}

	checksum := sum % 11

	if checksum == 10 {
		checksum = 0
	}

	fmt.Fprintf(os.Stdout, "Checksum: %d, Nip: %d Sum: %d\n", checksum, int(nip[9]-'0'), sum)

	return checksum == int(nip[9]-'0')
}
