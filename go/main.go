package main

import (
	"app/types"
	"fmt"
)

func greet(u *types.User) {
	fmt.Printf("Hello, %s!\n", u.Name)
}

func main() {
	u := &types.User{
		Name: "Dan",
		Age:  22,
	}

	greet(u)
}
