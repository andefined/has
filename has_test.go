package has_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/andefined/has"
)

func TestAll(t *testing.T) {
	dat, err := ioutil.ReadFile("README.md")
	if err != nil {
		panic(err)
	}

	emails := has.Email(string(dat))
	fmt.Print("Email: ", len(emails), emails, "\n")

	ipv4s := has.IPv4(string(dat))
	fmt.Print("IPv4: ", len(ipv4s), ipv4s, "\n")

	ipv6s := has.IPv6(string(dat))
	fmt.Print("IPv6: ", len(ipv6s), ipv6s, "\n")

	macs := has.MAC(string(dat))
	fmt.Print("MAC: ", len(macs), macs, "\n")

	urls := has.URL(string(dat))
	fmt.Print("URL: ", len(urls), urls, "\n")

	hostnames := has.Hostname(string(dat))
	fmt.Print("Hostname: ", len(hostnames), hostnames, "\n")

	domains := has.Domain(string(dat))
	fmt.Print("Domain: ", len(domains), domains, "\n")

	dnss := has.DNS(string(dat))
	fmt.Print("DNS: ", len(dnss), dnss, "\n")

	md5s := has.MD5(string(dat))
	fmt.Print("MD5: ", len(md5s), md5s, "\n")

	sha1s := has.SHA1(string(dat))
	fmt.Print("SHA1: ", len(sha1s), sha1s, "\n")

	sha256s := has.SHA256(string(dat))
	fmt.Print("SHA256: ", len(sha256s), sha256s, "\n")

	sha512s := has.SHA512(string(dat))
	fmt.Print("SHA512: ", len(sha512s), sha512s, "\n")

	ssdeeps := has.SSDeep(string(dat))
	fmt.Print("SSDeep: ", len(ssdeeps), ssdeeps, "\n")

	uuids := has.UUID(string(dat))
	fmt.Print("UUID: ", len(uuids), uuids, "\n")

	bitcoin := has.Bitcoin(string(dat))
	fmt.Print("Bitcoin: ", len(bitcoin), bitcoin, "\n")

	creditcards := has.CreditCard(string(dat))
	fmt.Print("CreditCard: ", len(creditcards), creditcards, "\n")

	winpaths := has.WinPath(string(dat))
	fmt.Print("WinPath: ", len(winpaths), winpaths, "\n")

	unixpaths := has.UnixPath(string(dat))
	fmt.Print("UnixPath: ", len(unixpaths), unixpaths, "\n")

	shellshocks := has.ShellShock(string(dat))
	fmt.Print("ShellShock: ", len(shellshocks), shellshocks, "\n")

	datauris := has.DataURI(string(dat))
	fmt.Print("DataURI: ", len(datauris), datauris, "\n")

	cyrillics := has.Cyrillic(string(dat))
	fmt.Print("Cyrillic: ", len(cyrillics), cyrillics, "\n")
}
