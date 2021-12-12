# rhythm
Package for testing Go 1.18 generics.

### Examples

```
package main

import "github.com/ncsurfus/rhythm/linked"

func test() {
	values := []int{0, 1, 2, 3, 4, 6}
	root := linked.FromSlice(values)
	root.Last().InsertAfter(7)
}
````