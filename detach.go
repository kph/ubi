// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package ubi

import (
	"syscall"
	"unsafe"

	"github.com/paypal/gatt/linux/gioctl"
)

// Detach is used to disassociate the given UBI device from its
// associated MTD device.
func Detach(ubiNum int32) (err error) {
	u, err := syscall.Open("/dev/ubi_ctrl", syscall.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer syscall.Close(u)

	err = gioctl.Ioctl(uintptr(u), gioctl.IoW(iocCtrlMagic, 65,
		unsafe.Sizeof(ubiNum)),
		uintptr(unsafe.Pointer(&ubiNum)))

	return
}
