package main

import "github.com/mt3hr/nimar/nimar/cmd"

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
