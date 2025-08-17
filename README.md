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

  // available variables
  fmt.Println(version.Version)
  fmt.Println(version.Repo)
  fmt.Println(version.CommitDate)
}
```
