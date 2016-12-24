// Copyright 2016, Marc Lavergne <mlavergn@gmail.com>. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package goutil

import (
	. "golog"
	"runtime"
)

const (
	HostUnknown = iota
	HostLinuxX86
	HostLinuxARM
	HostLinuxMIPS32LE
	HostMacosX86
)

func Platform() int {
	result := HostUnknown

	switch runtime.GOOS {
	case "linux":
		switch runtime.GOARCH {
		case "mips32le":
			result = HostLinuxMIPS32LE
		case "arm":
			result = HostLinuxARM
		default:
			result = HostLinuxX86
		}
	case "darwin":
		switch runtime.GOARCH {
		default:
			result = HostMacosX86
		}
	default:
		LogWarn("Unknown platform")
	}

	return result
}
