package has

import "regexp"

// ReEmail : RegExp Compilation
func ReEmail() *regexp.Regexp {
	r := regexp.MustCompile("^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+.[a-zA-Z0-9-.]+$")
	return r
}

// ReIPv4 : RegExp Compilation
func ReIPv4() string {
	return ""
}

// ReIPv6 : RegExp Compilation
func ReIPv6() string {
	return ""
}

// ReMAC : RegExp Compilation
func ReMAC() string {
	return ""
}

// ReURL : RegExp Compilation
func ReURL() string {
	return ""
}

// ReURI : RegExp Compilation
func ReURI() string {
	return ""
}

// ReHostname : RegExp Compilation
func ReHostname() string {
	return ""
}

// ReDomain : RegExp Compilation
func ReDomain() string {
	return ""
}

// ReDNS : RegExp Compilation
func ReDNS() string {
	return ""
}

// ReMD5 : RegExp Compilation
func ReMD5() string {
	return ""
}

// ReSHA1 : RegExp Compilation
func ReSHA1() string {
	return ""
}

// ReSHA2 : RegExp Compilation
func ReSHA2() string {
	return ""
}

// ReSHA256 : RegExp Compilation
func ReSHA256() string {
	return ""
}

// ReSHA512 : RegExp Compilation
func ReSHA512() string {
	return ""
}

// ReSSDeep : RegExp Compilation
func ReSSDeep() string {
	return ""
}

// ReUUID : RegExp Compilation
func ReUUID() string {
	return ""
}

// ReUUID3 : RegExp Compilation
func ReUUID3() string {
	return ""
}

// ReUUID4 : RegExp Compilation
func ReUUID4() string {
	return ""
}

// ReUUID5 : RegExp Compilation
func ReUUID5() string {
	return ""
}

// ReDataURI : RegExp Compilation
func ReDataURI() string {
	return ""
}

// ReBitcoin : RegExp Compilation
func ReBitcoin() *regexp.Regexp {
	r := regexp.MustCompile("([123mn][a-km-zA-HJ-NP-Z1-9]{25,34})")
	return r
}

// ReSSN : RegExp Compilation
func ReSSN() string {
	return ""
}

// ReCreditCard : RegExp Compilation
func ReCreditCard() string {
	return ""
}

// ReWinPath : RegExp Compilation
func ReWinPath() string {
	return ""
}

// ReUnixPath : RegExp Compilation
func ReUnixPath() string {
	return ""
}