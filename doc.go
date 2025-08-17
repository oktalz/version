// Package version provides utilities for retrieving and setting build version information.
//
// there is not init() function  you need to activate it manually
// Usage:
//
//	package main
//
//	import (
//		"fmt"
//		"github.com/oktalz/version"
//	)
//
//	func main() {
//		_ = version.Set()
//
//		fmt.Println(version.Version)
//		// Possible output examples
//		// v1.0.3
//		// v1.0.3+dirty
//		// v1.0.4-0.20250817123704-22f3a25a9bb1
//		// v1.0.4-0.20250817123704-22f3a25a9bb1+dirty
//	}
package version
