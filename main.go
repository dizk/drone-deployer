package main

import (
	"bytes"
	"log"
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

	c := exec.Command("/usr/bin/php", "/bin/dep", d.Task, d.Stage)
	c.Dir = w.Path

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	c.Stdout = &stdout
	c.Stderr = &stderr

	if err := c.Run(); err != nil {
		log.Fatal(stderr.String())
		log.Fatal(err)
	}

	log.Println(stdout.String())
	log.Println("Command completed successfully")
}
