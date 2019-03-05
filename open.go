// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package ubi

import (
	"os"
	"syscall"
)

func Open() (u Ubictrl, err error) {
	u = Ubictrl{}
	u.f, err = os.OpenFile("/dev/ubi_ctrl", syscall.O_RDWR, 0666)
	return
}
