package a

import (
	"fmt"

	"github.com/cybersamx/c"
)

func Version() string {
	return fmt.Sprintf("a: v1.2.0\n%s", c.Version())
}
