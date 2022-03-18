# sl

a toolbox for go with generics

### Require

require go version >= 1.18

### Example

##### Mapper

```go
package main

import "github.com/yguilai/sl"

type User struct {
	Id   int64
	Name string
}

func foo(users []User) {
	userIdSlice := sl.Mapper(users, func(u User) int64 { u.Id }).CollectSlice()
}
```