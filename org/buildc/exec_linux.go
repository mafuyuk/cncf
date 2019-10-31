package main

import (
	"os"
	"os/exec"
	"syscall"
)

const LINUX_EXE_CMD = "/proc/self/exe"

func execCMD(args ...string) error {
	cmd := exec.Command(LINUX_EXE_CMD, args...)

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
