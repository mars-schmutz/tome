//go:build linux

package main

import "os"

// Disable WebKitGTK's DMA-BUF renderer, which shows black views on hybrid NVIDIA/Intel setups.
func init() {
	if _, ok := os.LookupEnv("WEBKIT_DISABLE_DMABUF_RENDERER"); !ok {
		os.Setenv("WEBKIT_DISABLE_DMABUF_RENDERER", "1")
	}
}
