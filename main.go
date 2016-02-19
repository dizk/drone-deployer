package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/drone/drone-plugin-go/plugin"
	"github.com/kr/pty"
)

// SSHConfig the config used on the test runner
var SSHConfig = `Host *
    StrictHostKeyChecking no`

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

	if err := ioutil.WriteFile("/root/.ssh/config", []byte(SSHConfig), 0644); err != nil {
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

	f, err := pty.Start(c)
	if err != nil {
		log.Fatal(err)
	}

	// Seems to always return an error
	// read /dev/ptmx: input/output error
	// So no error checking
	io.Copy(os.Stdout, f)

	log.Println("Command completed successfully")
}
