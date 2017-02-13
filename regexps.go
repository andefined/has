package has

import "regexp"

// ReEmail : RegExp Compilation
func ReEmail() *regexp.Regexp {
	r := regexp.MustCompile(`\b[a-z0-9!#$%&'*+/=?^_{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_{|}~-]+)*(@| *\[ *at *\] *| *\( *at *\) *| *\{ *at *\} *| *- *at *- *| *\[ *@ *\] *| *\( *@ *\) *| *\{ *@ *\} *| *- *@ *- *)(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?(\.| *\[ *dot *\] *| *\( *dot *\) *| *\{ *dot *\} *| *- *dot *- *| *\[ *\. *\] *| *\( *\. *\) *| *\{ *\. *\} *| *- *. *- *))+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\b`)
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
func ReMAC() *regexp.Regexp {
	r := regexp.MustCompile(`(([0-9A-Fa-f]{2}[-:]){5}[0-9A-Fa-f]{2})|(([0-9A-Fa-f]{4}\.){2}[0-9A-Fa-f]{4}|([0-9A-Fa-f]{12}))`)
	return r
}

// ReURL : RegExp Compilation
func ReURL() *regexp.Regexp {
	r := regexp.MustCompile(`^((http[s]?|ftp):\/)?\/?([^:\/\s]+)((\/\w+)*\/)([\w\-\.]+[^#?\s]+)(.*)?(#[\w\-]+)?$`)
	return r
}

// ReHostname : RegExp Compilation
func ReHostname() *regexp.Regexp {
	r := regexp.MustCompile(`^(?!http|https|https:\/\/|http:\/\/)([?a-zA-Z0-9-.\+]{2,256}\.[a-z]{2,4}\b)`)
	return r
}

// ReDomain : RegExp Compilation
func ReDomain() *regexp.Regexp {
	r := regexp.MustCompile(`^((?:([a-z0-9]\.|[a-z0-9][a-z0-9\-]{0,61}[a-z0-9])\.)+)([a-z0-9]{2,63}|(?:[a-z0-9][a-z0-9\-]{0,61}[a-z0-9]))\.?$`)
	return r
}

// ReDNS : RegExp Compilation
func ReDNS() *regexp.Regexp {
	r := regexp.MustCompile(`^((?:([a-z0-9]\.|[a-z0-9][a-z0-9\-]{0,61}[a-z0-9])\.)+)([a-z0-9]{2,63}|(?:[a-z0-9][a-z0-9\-]{0,61}[a-z0-9]))\.?$`)
	return r
}

// ReMD5 : RegExp Compilation
func ReMD5() *regexp.Regexp {
	r := regexp.MustCompile(`^[a-fA-F0-9]{32}$`)
	return r
}

// ReSHA1 : RegExp Compilation
func ReSHA1() *regexp.Regexp {
	r := regexp.MustCompile(`\b([a-fA-F0-9]{40})\b`)
	return r
}

// ReSHA256 : RegExp Compilation
func ReSHA256() *regexp.Regexp {
	r := regexp.MustCompile(`\b([a-fA-F0-9]{64})\b`)
	return r
}

// ReSHA512 : RegExp Compilation
func ReSHA512() *regexp.Regexp {
	r := regexp.MustCompile(`\b([a-fA-F0-9]{128})\b`)
	return r
}

// ReSSDeep : RegExp Compilation
func ReSSDeep() *regexp.Regexp {
	r := regexp.MustCompile(`\b([\d]{1,10}:[a-zA-Z\d+/]{1,100}:[a-zA-Z\d+/]{1,100})\b`)
	return r
}

// ReUUID : RegExp Compilation
func ReUUID() *regexp.Regexp {
	r := regexp.MustCompile(`^([0-9a-f]{8}-?[0-9a-f]{4}-?[1-5][0-9a-f]{3}-?[0-9a-f]{4}-?[0-9a-f]{12}|\{[0-9a-f]{8}-?[0-9a-f]{4}-?[1-5][0-9a-f]{3}-?[0-9a-f]{4}-?[0-9a-f]{12}\}|[A-Za-z0-9+/]{22})$`)
	return r
}

// ReBitcoinAddress : RegExp Compilation
func ReBitcoinAddress() *regexp.Regexp {
	r := regexp.MustCompile("([123mn][a-km-zA-HJ-NP-Z1-9]{25,34})")
	return r
}

// ReCreditCard : RegExp Compilation
func ReCreditCard() *regexp.Regexp {
	r := regexp.MustCompile(`^(?:\d[ -]*?){13,16}$`)
	return r
}

// ReWinPath : RegExp Compilation
func ReWinPath() *regexp.Regexp {
	r := regexp.MustCompile(`^([a-zA-Z]):[\\\/]((?:[^<>:"\\\/\|\?\*]+[\\\/])*)([^<>:"\\\/\|\?\*]+)\.([^<>:"\\\/\|\?\*\s]+)$`)
	return r
}

// ReUnixPath : RegExp Compilation
func ReUnixPath() *regexp.Regexp {
	r := regexp.MustCompile(`^(((?:\./|\.\./|/)?(?:\.?\w+/)*)(\.?\w+\.?\w+))$`)
	return r
}

// ReShellShock : RegExp Compilation
func ReShellShock() *regexp.Regexp {
	r := regexp.MustCompile(`\s*\(\s*\)\s*\{.*?;\s*\}\s*;`)
	return r
}
