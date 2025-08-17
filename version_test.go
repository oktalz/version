package version_test

import (
	"fmt"

	"github.com/oktalz/version"
)

func Example() {
	_ = version.Set()

	fmt.Println(version.Version)
	fmt.Println(version.Repo)
	fmt.Println(version.CommitDate)
}
