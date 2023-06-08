package main

import (
	"envchecker/cmd"
	"envchecker/g"
)

func main() {
	if err := cmd.Executor(); err != nil {
		g.Log.Error("cmd.Executor", err)
	}
}
