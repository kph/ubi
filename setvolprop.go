// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package ubi

import (
	"unsafe"

	"github.com/paypal/gatt/linux/gioctl"
)

type setVolPropReq struct {
	property byte
	padding  [7]byte
	value    [8]byte
}

const (
	VolPropDirectWrite = 1
)

func SetVolProp(fd uintptr, property byte, value uint64) (err error) {
	r := setVolPropReq{}
	r.property = property
	*(*uint64)(unsafe.Pointer(&r.value)) = value

	err = gioctl.Ioctl(fd, gioctl.IoW(iocVolMagic, 6,
		unsafe.Sizeof(r)),
		uintptr(unsafe.Pointer(&r)))

	return
}
