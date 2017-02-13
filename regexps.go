package has

import "regexp"

// ReEmail : RegExp Compilation
func ReEmail() *regexp.Regexp {
	r := regexp.MustCompile(`\b[a-z0-9!#$%&'*+/=?^_{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\b$`)
	return r
}

// ReIPv4 : RegExp Compilation
func ReIPv4() *regexp.Regexp {
	r := regexp.MustCompile(`\b((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\b`)
	return r
}

// ReIPv6 : RegExp Compilation
func ReIPv6() *regexp.Regexp {
	r := regexp.MustCompile(`\b(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))\b`)
	return r
}

// ReMAC : RegExp Compilation
func ReMAC() string {
	return ""
}

// ReURL : RegExp Compilation
func ReURL() *regexp.Regexp {
	r := regexp.MustCompile(`^((http[s]?|ftp):\/)?\/?([^:\/\s]+)((\/\w+)*\/)([\w\-\.]+[^#?\s]+)(.*)?(#[\w\-]+)?$`)
	return r
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
func ReMD5() *regexp.Regexp {
	r := regexp.MustCompile(`^[a-fA-F0-9]{32}$`)
	return r
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

// ReDataURI : RegExp Compilation
func ReDataURI() string {
	return ""
}

// ReBitcoinAddress : RegExp Compilation
func ReBitcoinAddress() *regexp.Regexp {
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
