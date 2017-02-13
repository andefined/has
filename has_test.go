package has_test

import (
	"fmt"
	"testing"

	"github.com/andefined/has"
)

func TestHasEmail(t *testing.T) {
	val := has.Email("one@email.com, two@mail.com")
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
