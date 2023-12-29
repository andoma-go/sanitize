package sanitize

import (
	"net"
	"regexp"
)

var reIP = regexp.MustCompile(`[^a-zA-Z0-9:.]`) // IPv4 and IPv6 characters only

// Integer returns numbers only
func Integer(input string) string {
	runes := []rune{}
	n := 0
	for _, r := range input {
		if !(r == 45) && !(r >= 48 && r <= 57) { // Ignore any invalid characters first
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
		if !(r >= 44 && r <= 46) && !(r >= 48 && r <= 57) { // Ignore any invalid characters first
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

// IP returns the IP address in IPv4 or IPv6 format
func IP(input string) string {
	sanitized := reIP.ReplaceAllString(input, "")
	if ip := net.ParseIP(sanitized); ip != nil {
		return ip.String()
	}
	return ""
}
