// Copyright 2016, Marc Lavergne <mlavergn@gmail.com>. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package goutil

import (
	. "golog"
)

func Spawn(path string, args string) {
	cmd := exec.Command(path, args)

	// prevent the child from being grouped with the parent and exited with it
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.SysProcAttr.Setpgid = true
	cmd.Start()
}
