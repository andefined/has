[![Go Report Card](https://goreportcard.com/badge/github.com/andefined/has)](https://goreportcard.com/report/github.com/andefined/has)

# has

Find patterns in strings (wip)

## Installation

```bash
go get github.com/andefined/has
```

## Usage

```go
import (
    "github.com/andefined/has"
)
```

## Methods

```go
has.Email(s string) []string // email@example.com, email[at]example[dot]com, email [ at ] example [ dot ] com
has.IPv4(s string) []string // 192.168.1.1, 10.0.0.0, 0.0.0.0, 256.256.256.256
has.IPv6(s string) []string // 2001:0db8:0000:0000:0000:ff00:0042:8329
has.MAC(s string) []string // 1a-cd-3f-ac-42-34-ed-ab, 1a:cd:3f:ac:42:34:ed:ab
has.URL(s string) []string // https://www.domain.com, http://domain.com, https://github.com/andefined/has
has.Hostname(s string) []string // domain.com
has.Domain(s string) []string // domain.com
has.DNS(s string) []string // localhost.local
has.MD5(s string) []string // d131dd02c5e6eec4693d9a0698aff95c 2fcab58712467eab4004583eb8fb7f89
has.SHA1(s string) []string // de9f2c7fd25e1b3afad3e85a0bd17d9b100db4b3
has.SHA256(s string) []string // e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
has.SHA512(s string) []string // cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e
has.SSDeep(s string) []string // 01:username:password
has.UUID(s string) []string // 123e4567-e89b-12d3-a456-426655440000
has.Bitcoin(s string) []string // 1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2, 3J98t1WpEZ73CNmQviecrnyiWrnqRhWNLy, 4J98t1WpEZ73CNmQviecrnyiWrnqRhWNLy
has.CreditCard(s string) []string // 1234-4567-8901-2345, 1234 4567 8901 2345
has.WinPath(s string) []string // C:\Users\user\Desktop
has.UnixPath(s string) []string // /usr/src/bin
has.ShellShock(s string) []string // env x='() { :;}; echo vulnerable' bash -c "echo this is a test"
has.DataURI(s string) []string // data:text/html,%3Ch1%3EHello%2C%20World!%3C%2Fh1%3E, data:text/html,<script>alert('hi');</script>, data:text/plain;base64,SGVsbG8sIFdvcmxkIQ%3D%3D
has.Cyrillic(s string) []string // Привіт привет, Пламенный п. защитникам Родины
```
