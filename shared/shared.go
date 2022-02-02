// You can create modules at this level and they will be
// interpreted as under module birc.au.dk, so to import
// package `shared` you need `import "birc.au.dk/gsa/shared"`

package shared

import "fmt"

func Todo(genome, reads string) string {
	return fmt.Sprintf("Search for reads from %s in the genome in %s", reads, reads)
}
