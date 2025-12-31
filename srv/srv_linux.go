package srv

import (
	"os"
)

func runPath(service bool) (string, error) {
	return os.Args[0], nil
}
