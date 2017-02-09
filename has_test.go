package has_test

import (
	"fmt"
	"testing"

	"github.com/andefined/has"
)

func TestHasEmail(t *testing.T) {
	val := has.Email("email@email.com.gr , s@email.com.gr")
	fmt.Print("Email: ", len(val), val, "\n")
}

func TestHasBitcoin(t *testing.T) {
	val := has.Bitcoin("1F1tAaz5x1HUXrCNLbtMDqcw6o5GNn4xqX")
	fmt.Print("Bitcoin: ", len(val), val, "\n")
}
