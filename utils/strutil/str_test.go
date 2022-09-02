package strutil

import (
	"fmt"
	"testing"
)

func TestCamel2Case(t *testing.T) {
	camel2Case := Camel2Case("CaisinHelloMyNameIs_hekx")
	fmt.Println(camel2Case)
	camel2Case = Case2Camel(camel2Case)
	fmt.Println(camel2Case)
}
