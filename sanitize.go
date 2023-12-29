package sanitize

// Integer returns only numbers
func Integer(input string) string {
	runes := []rune{}
	n := 0
	for _, r := range input {
		if !(r == 45) && !(r >= 48 && r <= 57) { // regexp.MustCompile(`[^-\d]`).ReplaceAllString(input, "$1")
			continue
		}
		switch r {
		case 45: // "-"
			if n > 0 {
				continue
			}
		default: // Digits
			n++
		}
		runes = append(runes, r)
	}
	if n > 0 {
		return string(runes)
	}
	return ""
}

// Float returns sanitized floating-point numbers
func Float(input string) string {
	runes := []rune{}
	dot := false
	n := 0
	for i, r := range input {
		if !(r >= 44 && r <= 46) && !(r >= 48 && r <= 57) { // regexp.MustCompile(`[^-\.,\d]`).ReplaceAllString(input, "$1")
			continue
		}
		switch r {
		case 45: // "-"
			if n > 0 {
				continue
			}
		case 44: // ","
			fallthrough
		case 46: // "."
			if dot || i == len(input)-1 {
				continue
			}
			dot = true
			if n == 0 {
				runes = append(runes, 48) // "0"
			}
		default: // Digits
			n++
		}
		runes = append(runes, r)
	}
	if n > 0 {
		return string(runes)
	}
	return ""
}
