package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/mitchellh/go-ps"
)

func main() {
	displayPIDInfo()

	if len(os.Args) < 3 {
		fmt.Println("lack of command args")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "run":
		parent()
	case "child":
		child()
	default:
		panic("wat should I do")
	}
}

func parent() {
	if err := execCMD(append([]string{"child"}, os.Args[2:]...)...); err != nil {
		fmt.Println("ERROR parent", err)
		os.Exit(1)
	}

	fmt.Println("Finish parent")
}

func child() {
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR child", err)
		os.Exit(1)
	}
	fmt.Println("Finish child")
}

func displayPIDInfo() {
	pid := os.Getpid()

	pidInfo, _ := ps.FindProcess(pid)

	fmt.Printf("> PID          : %d\n", pidInfo.Pid())
	fmt.Printf("> PPID         : %d\n", pidInfo.PPid())
	fmt.Printf("> Process name : %s\n", pidInfo.Executable())
}
