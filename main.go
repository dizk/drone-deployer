package main

import (
	"bytes"
	"io/ioutil"
	"log"
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

	// Save ssh keys
	if err := os.MkdirAll("/root/.ssh", 0700); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("/root/.ssh/id_rsa", []byte(w.Keys.Private), 0600); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("/root/.ssh/id_rsa.pub", []byte(w.Keys.Public), 0644); err != nil {
		log.Fatal(err)
	}

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
