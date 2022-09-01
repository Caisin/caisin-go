package osutl

import (
	"fmt"
	"testing"
)

func TestHome(t *testing.T) {
	home, err := Home()
	fmt.Println(home, err)
}
