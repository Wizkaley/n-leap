package leap

import (
	"log"

	"github.com/Wizkaley/n-leap.git/engine"
)

// this will initialize the leap and figure out the type of language to scan for
func InitializeLeap() error {

	err := engine.InitializeEngine()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
