// Program go-template prints the version and exits
package main

import (
  "fmt"
  "os"
  "path/filepath"
  "runtime"

  "github.com/rasa/go-template/version"
)

func main() {
	fmt.Printf("%s: Version %s (%s)\n", filepath.Base(os.Args[0]), version.VERSION, version.GITCOMMIT)
	fmt.Printf("Built with go %s for %s/%s\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)
}
