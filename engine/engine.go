package engine

import "fmt"

type LeapEngine interface {
	Crawl() error
}

type SubjectRepository struct {
	Dependencies []SubjectRepository
}

func InitializeEngine() error {
	// the engine will
	// - figure out the type of language to scan for
	// -

	fmt.Println("Initializing engine")
	return nil
}
