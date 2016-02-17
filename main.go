package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/drone/drone-plugin-go/plugin"
)

type deployer struct {
	Task  string `json:"task"`
	Stage string `json:"stage"`
}

func main() {
	d := new(deployer)
	w := new(plugin.Workspace)

	plugin.Param("vargs", d)
	plugin.Param("workspace", w)
	plugin.MustParse()

	c := exec.Command("/bin/dep", d.Task, d.Stage)
	c.Path = w.Path

	err := c.Run()
	if err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Command completed successfully")
}
