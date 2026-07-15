//go:build windows

package adf

import (
	"os/exec"
	"syscall"
)

func startDetachedWindows(script string) error {
	cmd := exec.Command("cmd.exe", "/C", "start", "", "/B", script)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: 0x00000008,
	}
	return cmd.Start()
}
