/*
Copyright © 2023 NAME HERE <fabrzytech@gmail.com>
*/
package main

import (
	"errors"

	"github.com/FabioSebs/golangWrappers/cmd"
	"github.com/FabioSebs/golangWrappers/logs"
)

func main() {
	cmd.Execute()
	logs.LogInfo("testing")
	logs.LogError(errors.New("error"))
}
