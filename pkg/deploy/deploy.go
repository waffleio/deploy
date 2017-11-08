package deploy

import (
	"fmt"
)

func getOutput() string {
	return "Done."
}

// Run our cli
func Run() {
	fmt.Println(getOutput())
}
