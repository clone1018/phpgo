package phpgo

import (
	"fmt"
)

func Echo(args ...interface{}) {
	fmt.Print(args)
}
