package has

// Email : has.Email
// TODO: Use RFC 5322, Introduce SubMatches
// ex. harrypotter [at] hogwarts [dot] ac (dot) uk
func Email(str string) []string {
	elements := ReEmail().FindAllString(str, -1)
	return removeDuplicates(elements)
}

// IPv4 : has.IPv4
// TODO: Validate IP using net.ParseIP
func IPv4(s string) []string {
	elements := ReIPv4().FindAllString(s, -1)
	return removeDuplicates(elements)
}

// IPv6 : has.IPv6
// TODO: Validate IP using net.ParseIP
func IPv6(s string) []string {
	elements := ReIPv6().FindAllString(s, -1)
	return removeDuplicates(elements)
}

// BitcoinAddress : has.BitcoinAddress
func BitcoinAddress(s string) []string {
	elements := ReBitcoinAddress().FindAllString(s, -1)
	return removeDuplicates(elements)
}

// MD5 : has.MD5
func MD5(s string) []string {
	elements := ReMD5().FindAllString(s, -1)
	return removeDuplicates(elements)
}

// Helper Function
// removeDuplicates : Remove Duplicates
func removeDuplicates(elements []string) []string {
	encountered := map[string]bool{}
	for v := range elements {
		encountered[elements[v]] = true
	}

	result := []string{}
	for key := range encountered {
		result = append(result, key)
	}

	return result
}
