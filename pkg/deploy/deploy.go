package deploy

import (
	"fmt"
	"os"
)

// Run our cli
func Run() {
	config, err := GetConfig()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	err = config.Validate()
	if err != nil {
		fmt.Printf("We are missing a configuration option: %v\n", err)
		os.Exit(1)
	}
}
