package has_test

import (
	"fmt"
	"testing"

	"github.com/andefined/has"
)

func TestHasEmail(t *testing.T) {
	val := has.Email("one@email.com, two@mail.com, harrypotter [at] hogwarts [dot] ac (dot) uk")
	fmt.Print("Email: ", len(val), val, "\n")
}

func TestHasIPv4(t *testing.T) {
	val := has.IPv4("252.168.1.1 252.168.1.2")
	fmt.Print("IPv4: ", len(val), val, "\n")
}

func TestHasIPv6(t *testing.T) {
	val := has.IPv6("2001:0db8:0000:0000:0000:ff00:0042:8329, 2001:db8:0:0:0:ff00:42:8329, 2001:db8::ff00:42:8329")
	fmt.Print("IPv6: ", len(val), val, "\n")
}

func TestHasBitcoinAddress(t *testing.T) {
	val := has.BitcoinAddress("2F1tAaz5x1HUXrCNLbtMDqcw6o5GNn4xqX, 1F1tAaz5x1HUXrCNLbtMDqcw6o5GNn4xqX")
	fmt.Print("Bitcoin: ", len(val), val, "\n")
}

func TestHasMD5(t *testing.T) {
	val := has.MD5("00236a2ae558018ed13b5222ef1bd977")
	fmt.Print("MD5: ", len(val), val, "\n")
}

func TestHasSHA1(t *testing.T) {
	val := has.SHA1("da39a3ee5e6b4b0d3255bfef95601890afd80709 e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855 cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e c672b8d1ef56ed28ab87c3622c5114069bdd3ad7b8f9737498d0c01ecef0967a 38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95b da39a3ee5e6b4b0d3255bfef95601890afd80709 13DFB2603BD6893F4EFA9F0B6205A32DB54D9E259C20C1AD35D4D8AA2962D3609EE73C3D26B28B5742A555B44C11E212CFFA887EC0A34342BE85C4B92CAA4CD2")
	fmt.Print("SHA1: ", len(val), val, "\n")
}

func TestHasSHA256(t *testing.T) {
	val := has.SHA256("da39a3ee5e6b4b0d3255bfef95601890afd80709 e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855 cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e c672b8d1ef56ed28ab87c3622c5114069bdd3ad7b8f9737498d0c01ecef0967a 38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95b da39a3ee5e6b4b0d3255bfef95601890afd80709 13DFB2603BD6893F4EFA9F0B6205A32DB54D9E259C20C1AD35D4D8AA2962D3609EE73C3D26B28B5742A555B44C11E212CFFA887EC0A34342BE85C4B92CAA4CD2")
	fmt.Print("SHA256: ", len(val), val, "\n")
}

func TestHasSHA512(t *testing.T) {
	val := has.SHA512("da39a3ee5e6b4b0d3255bfef95601890afd80709 e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855 cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e c672b8d1ef56ed28ab87c3622c5114069bdd3ad7b8f9737498d0c01ecef0967a 38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95b da39a3ee5e6b4b0d3255bfef95601890afd80709 13DFB2603BD6893F4EFA9F0B6205A32DB54D9E259C20C1AD35D4D8AA2962D3609EE73C3D26B28B5742A555B44C11E212CFFA887EC0A34342BE85C4B92CAA4CD2")
	fmt.Print("SHA512: ", len(val), val, "\n")
}

func TestHasSSDeep(t *testing.T) {
	val := has.SSDeep("125:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855:921d36ce9ce47d0d13c5d85f2b0ff")
	fmt.Print("SSDeep: ", len(val), val, "\n")
}

func TestHasMAC(t *testing.T) {
	val := has.MAC("12:23:45:55:77:88 12:23:45:5Z:77:88 12B5.1a34.4567 12B5.1G34.4567 969982549025 96998G549025 1B-F2-43-44-55-41 1B-F2-43-44-55-G1")
	fmt.Print("MAC: ", len(val), val, "\n")
}
