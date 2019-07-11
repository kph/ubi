// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package ubi

import (
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

func (u Ubictrl) Attach(ubiNum, mtdNum, vidHdrOffset int32,
	maxPebPer1024 int16) (err error) {
	r := attachReq{}
	*(*int32)(unsafe.Pointer(&r.ubiNum)) = ubiNum
	*(*int32)(unsafe.Pointer(&r.mtdNum)) = mtdNum
	*(*int32)(unsafe.Pointer(&r.vidHdrOffset)) = vidHdrOffset
	*(*int16)(unsafe.Pointer(&r.maxPebPer1024)) = maxPebPer1024

	err = gioctl.Ioctl(u.f.Fd(), gioctl.IoW(iocCtrlMagic, 64,
		unsafe.Sizeof(r)),
		uintptr(unsafe.Pointer(&r)))

	return
}
