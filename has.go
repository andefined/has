package has

import (
	"net"
	"regexp"
)

// Email : has.Email
func Email(str string) []string {
	elements := ReEmail.FindAllString(str, -1)
	var c []string
	for i := range elements {
		email := elements[i]
		reAT := regexp.MustCompile(`(\.` + at + `\b)`)
		reDOT := regexp.MustCompile(`(\.` + dot + `\b)`)
		email = reAT.ReplaceAllString(email, `@`)
		email = reDOT.ReplaceAllString(email, `.`)
		c = append(c, email)
	}
	return removeDuplicates(c)
}

// IPv4 : has.IPv4
func IPv4(s string) []string {
	elements := ReIPv4.FindAllString(s, -1)
	var c []string
	for i := range elements {
		ip := net.ParseIP(elements[i])
		if ip != nil {
			c = append(c, elements[i])
		}
	}
	return removeDuplicates(c)
}

// IPv6 : has.IPv6
func IPv6(s string) []string {
	elements := ReIPv6.FindAllString(s, -1)
	var c []string
	for i := range elements {
		ip := net.ParseIP(elements[i])
		if ip != nil {
			c = append(c, elements[i])
		}
	}
	return removeDuplicates(c)
}

// MAC : has.MAC
func MAC(s string) []string {
	elements := ReMAC.FindAllString(s, -1)
	var c []string
	for i := range elements {
		_, err := net.ParseMAC(elements[i])
		if err == nil {
			c = append(c, elements[i])
		}
	}
	return removeDuplicates(c)
}

// URL : has.URL
func URL(s string) []string {
	elements := ReURL.FindAllString(s, -1)
	return removeDuplicates(elements)
}

// Hostname : has.Hostname
func Hostname(s string) []string {
	elements := ReHostname.FindAllString(s, -1)
	return removeDuplicates(elements)
}

// Domain : has.Domain
func Domain(s string) []string {
	elements := ReDomain.FindAllString(s, -1)
	return removeDuplicates(elements)
}

// DNS : has.DNS
func DNS(s string) []string {
	elements := ReDNS.FindAllString(s, -1)
	return removeDuplicates(elements)
}

// MD5 : has.MD5
func MD5(s string) []string {
	elements := ReMD5.FindAllString(s, -1)
	return removeDuplicates(elements)
}

// SHA1 : has.SHA1
func SHA1(s string) []string {
	elements := ReSHA1.FindAllString(s, -1)
	return removeDuplicates(elements)
}

// SHA256 : has.SHA256
func SHA256(s string) []string {
	elements := ReSHA256.FindAllString(s, -1)
	return removeDuplicates(elements)
}

// SHA512 : has.SHA512
func SHA512(s string) []string {
	elements := ReSHA512.FindAllString(s, -1)
	return removeDuplicates(elements)
}

// SSDeep : has.SSDeep
func SSDeep(s string) []string {
	elements := ReSSDeep.FindAllString(s, -1)
	return removeDuplicates(elements)
}

// UUID : has.UUID
func UUID(s string) []string {
	elements := ReUUID.FindAllString(s, -1)
	return removeDuplicates(elements)
}

// Bitcoin : has.Bitcoin
func Bitcoin(s string) []string {
	elements := ReBitcoin.FindAllString(s, -1)
	return removeDuplicates(elements)
}

// CreditCard : has.CreditCard
func CreditCard(s string) []string {
	elements := ReCreditCard.FindAllString(s, -1)
	return removeDuplicates(elements)
}

// WinPath : has.WinPath
func WinPath(s string) []string {
	elements := ReWinPath.FindAllString(s, -1)
	return removeDuplicates(elements)
}

// UnixPath : has.UnixPath
func UnixPath(s string) []string {
	elements := ReUnixPath.FindAllString(s, -1)
	return removeDuplicates(elements)
}

// ShellShock : has.ShellShock
func ShellShock(s string) []string {
	elements := ReShellShock.FindAllString(s, -1)
	return removeDuplicates(elements)
}

// DataURI : has.DataURI
func DataURI(s string) []string {
	elements := ReDataURI.FindAllString(s, -1)
	return removeDuplicates(elements)
}

// Cyrillic : has.Cyrillic
func Cyrillic(s string) []string {
	elements := ReCyrillic.FindAllString(s, -1)
	return removeDuplicates(elements)
}
