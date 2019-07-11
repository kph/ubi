// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package ubi

import (
	"syscall"
	"unsafe"

	"github.com/paypal/gatt/linux/gioctl"
)

type attachReq struct {
	ubiNum        [4]byte
	mtdNum        [4]byte
	vidHdrOffset  [4]byte
	maxPebPer1024 [2]byte
	padding       [10]byte
}

// Attach is used to associate a specified UBI device with a specified
// MTD device. This makes all of the logical volumes available.
func Attach(ubiNum, mtdNum, vidHdrOffset int32,
	maxPebPer1024 int16) (err error) {
	u, err := syscall.Open("/dev/ubi_ctrl", syscall.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer syscall.Close(u)

	r := attachReq{}
	*(*int32)(unsafe.Pointer(&r.ubiNum)) = ubiNum
	*(*int32)(unsafe.Pointer(&r.mtdNum)) = mtdNum
	*(*int32)(unsafe.Pointer(&r.vidHdrOffset)) = vidHdrOffset
	*(*int16)(unsafe.Pointer(&r.maxPebPer1024)) = maxPebPer1024

	err = gioctl.Ioctl(uintptr(u), gioctl.IoW(iocCtrlMagic, 64,
		unsafe.Sizeof(r)),
		uintptr(unsafe.Pointer(&r)))

	return
}
