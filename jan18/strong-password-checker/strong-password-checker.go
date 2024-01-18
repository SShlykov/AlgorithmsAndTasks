package strong_password_checker

func strongPasswordChecker(password string) int {
	missingType := countCharacterTypes(password)
	replace, ones, twos := calculateReplaceAndSequences(password)

	if len(password) < 6 {
		return max(missingType, 6-len(password))
	}

	if len(password) <= 20 {
		return max(missingType, replace)
	}

	del := len(password) - 20
	replace -= min(del, ones)
	replace -= min(max(del-ones, 0), twos*2) / 2
	replace -= max(del-ones-2*twos, 0) / 3

	return del + max(missingType, replace)
}

func countCharacterTypes(password string) int {
	l, u, d, missingType := 1, 1, 1, 3
	for i := range password {
		if l > 0 && 'a' <= password[i] && password[i] <= 'z' {
			l--
			missingType--
		}
		if u > 0 && 'A' <= password[i] && password[i] <= 'Z' {
			u--
			missingType--
		}
		if d > 0 && '0' <= password[i] && password[i] <= '9' {
			d--
			missingType--
		}
	}
	return missingType
}

func calculateReplaceAndSequences(password string) (int, int, int) {
	var replace, ones, twos int
	for p := 0; p+2 < len(password); p++ {
		if password[p] == password[p+1] && password[p+1] == password[p+2] {
			repeatingLen := 2
			for p+2 < len(password) && password[p] == password[p+2] {
				repeatingLen++
				p++
			}
			replace += repeatingLen / 3
			if repeatingLen%3 == 0 {
				ones++
			} else if repeatingLen%3 == 1 {
				twos++
			}
		}
	}
	return replace, ones, twos
}
