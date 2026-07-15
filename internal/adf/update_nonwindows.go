//go:build !windows

package adf

func startDetachedWindows(string) error { return nil }
