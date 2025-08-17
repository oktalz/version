# version

Package version provides utilities for retrieving and setting build version information.

Usage:

```go
package main

import (
  "fmt"
  "github.com/oktalz/version"
)

func main() {
  _ = version.Set()

  fmt.Println(version.Version)
  // Possible output examples
  // v1.0.3
  // v1.0.3+dirty
  // v1.0.4-0.20250817123704-22f3a25a9bb1
  // v1.0.4-0.20250817123704-22f3a25a9bb1+dirty

  // other available variables
  fmt.Println(version.Repo)
  fmt.Println(version.Commit)
  fmt.Println(version.CommitDate)
  fmt.Println(version.Modified)
}
```

## go build, go run commands

add `-buildvcs=true` in order to get proper version, `go install` does this automatically
