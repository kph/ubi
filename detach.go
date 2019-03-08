// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package ubi

import (
	"unsafe"

	"github.com/paypal/gatt/linux/gioctl"
)

func (u Ubictrl) Detach(ubiNum int32) (err error) {
	err = gioctl.Ioctl(u.f.Fd(), gioctl.IoW(iocMagic, 65,
		unsafe.Sizeof(ubiNum)),
		uintptr(unsafe.Pointer(&ubiNum)))

	return
}
