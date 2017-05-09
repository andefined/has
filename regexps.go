package has

import "regexp"

const (
	at             string = `*\[ *at *\] *| *\( *at *\) *| *\{ *at *\} *| *- *at *- *| *\[ *@ *\] *| *\( *@ *\) *| *\{ *@ *\} *| *- *@ *- *`        // @ symbol
	dot            string = `*\[ *dot *\] *| *\( *dot *\) *| *\{ *dot *\} *| *- *dot *- *| *\[ *\. *\] *| *\( *\. *\) *| *\{ *\. *\} *| *- *. *- *` // . symbol
	email          string = `(?i)(?m)\b[a-z0-9!#$%&'*+/=?^_{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_{|}~-]+)*(\@` + at + `)(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?(\.` + dot + `))+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\b`
	ipv4           string = `(?i)(?m)\b(?:(?:2(?:[0-4][0-9]|5[0-5])|[0-1]?[0-9]?[0-9])?(\.| ` + dot + `)){3}(?:(?:2([0-4][0-9]|5[0-5])|[0-1]?[0-9]?[0-9]))\b`
	ipv6           string = `(?i)(?m)\b(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))\b`
	mac            string = `(?i)(?m)\b(([0-9A-Fa-f]{2}[-:]){5}[0-9A-Fa-f]{2})|(([0-9A-Fa-f]{4}\.){2}[0-9A-Fa-f]{4}|([0-9A-Fa-f]{12}))`
	urlschema      string = `((ftp|tcp|udp|ws|http|):\/\/)`
	urlcredentials string = `(\S+(:\S*)?@)`
	urlpath        string = `((\/|\?|#)[^\s]*)`
	urlport        string = `(:(\d{1,5}))`
	urlip          string = `([1-9]\d?|1\d\d|2[01]\d|22[0-3])(\.(1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-4]))`
	urlsubdomain   string = `((www\.)|([a-zA-Z0-9]([-\.][-\._a-zA-Z0-9]+)*))`
	url            string = `(?i)(?m)\b(` + urlschema + `?` + urlcredentials + `?` + `((` + urlip + `|(\[` + ipv4 + `\])|(([a-zA-Z0-9]([a-zA-Z0-9-_]+)?[a-zA-Z0-9]([-\.][a-zA-Z0-9]+)*)|(` + urlsubdomain + `?))?(([a-zA-Z\x{00a1}-\x{ffff}0-9]+-?-?)*[a-zA-Z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-zA-Z\x{00a1}-\x{ffff}]{1,}))?))\.?` + urlport + `?` + urlpath + `)?/` // `\b((http[s]?|ftp):\/)?\/?([^:\/\s]+)((\/\w+)*\/)([\w\-\.]+[^#?\s]+)(.*)?(#[\w\-]+)?\b`
	hostname       string = `^((http[s]?|ftp):\/)?\/?([^:\/\s]+)([?a-zA-Z0-9-.\+]{2,256}\.[a-z]{2,4}\b)`
	domain         string = `^((?:([a-z0-9]\.|[a-z0-9][a-z0-9\-]{0,61}[a-z0-9])\.)+)([a-z0-9]{2,63}|(?:[a-z0-9][a-z0-9\-]{0,61}[a-z0-9]))\.?$`
	dns            string = `^((?:([a-z0-9]\.|[a-z0-9][a-z0-9\-]{0,61}[a-z0-9])\.)+)([a-z0-9]{2,63}|(?:[a-z0-9][a-z0-9\-]{0,61}[a-z0-9]))\.?$`
	md5            string = `(?m)\b([a-fA-F0-9]){32}?\b`
	sha1           string = `(?m)\b([a-fA-F0-9]){40}?\b`
	sha256         string = `(?m)\b([a-fA-F0-9]){64}?\b`
	sha512         string = `(?m)\b([a-fA-F0-9]){128}?\b`
	ssdeep         string = `(?m)\b([\d]){1,10}:([a-zA-Z\d+/]){1,100}:([a-zA-Z\d+/]){1,100}?\b`
	uuid           string = `(?m)\b([0-9a-f]{8}-?[0-9a-f]{4}-?[1-5][0-9a-f]{3}-?[0-9a-f]{4}-?[0-9a-f]{12}|\{[0-9a-f]{8}-?[0-9a-f]{4}-?[1-5][0-9a-f]{3}-?[0-9a-f]{4}-?[0-9a-f]{12}\}|[A-Za-z0-9+/]{22})\b`
	bitcoin        string = `(?m)\b[13][a-km-zA-HJ-NP-Z1-9]{25,34}?\b` // r := regexp.MustCompile("([123mn][a-km-zA-HJ-NP-Z1-9]{25,34})")
	creditcard     string = `(?m)\b(?:\d[ -]*?){13,16}\b`
	winpath        string = `(?m)\b[a-zA-Z]:\\(?:[^\\/:*?"<>|\r\n]+\\)*[^\\/:*?"<>|\r\n]*?`
	unixpath       string = `(?m)\b((/[^/\x00]*)+/)*?`
	shellshock     string = `\s*\(\s*\)\s*\{.*?;\s*\}\s*;`
	cyrillic       string = `([\x{0400}-\x{052F}]+)`
)

var (
	ReEmail      = regexp.MustCompile(email)
	ReIPv4       = regexp.MustCompile(ipv4)
	ReIPv6       = regexp.MustCompile(ipv6)
	ReMAC        = regexp.MustCompile(mac)
	ReURL        = regexp.MustCompile(url)
	ReHostname   = regexp.MustCompile(hostname)
	ReDomain     = regexp.MustCompile(domain)
	ReDNS        = regexp.MustCompile(dns)
	ReMD5        = regexp.MustCompile(md5)
	ReSHA1       = regexp.MustCompile(sha1)
	ReSHA256     = regexp.MustCompile(sha256)
	ReSHA512     = regexp.MustCompile(sha512)
	ReSSDeep     = regexp.MustCompile(ssdeep)
	ReUUID       = regexp.MustCompile(uuid)
	ReBitcoin    = regexp.MustCompile(bitcoin)
	ReCreditCard = regexp.MustCompile(creditcard)
	ReWinPath    = regexp.MustCompile(winpath)
	ReUnixPath   = regexp.MustCompile(unixpath)
	ReShellShock = regexp.MustCompile(shellshock)
	ReCyrillic   = regexp.MustCompile(cyrillic)
)
