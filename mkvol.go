// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package ubi

import (
	"fmt"
	"os"
	"strconv"
	"syscall"
	"unsafe"

	"github.com/paypal/gatt/linux/gioctl"
)

type mkvolReq struct {
	volId     [4]byte
	alignment [4]byte
	bytes     [8]byte
	volType   byte
	padding1  byte
	nameLen   [2]byte
	padding2  [4]byte
	name      [MaxVolName + 1]byte
}

// Mkvol is used to create a new volume on the specified UBI device.
func Mkvol(ubiNum int, volId, alignment int32, static bool,
	bytes int64, name string) (err error) {
	l := len(name)
	if l > MaxVolName {
		return fmt.Errorf("Name length of %d exceeds maximum %d",
			l, MaxVolName)
	}
	r := mkvolReq{}
	t := DynamicVol
	if static {
		t = StaticVol
	}
	*(*int32)(unsafe.Pointer(&r.volId)) = volId
	*(*int32)(unsafe.Pointer(&r.alignment)) = alignment
	*(*int64)(unsafe.Pointer(&r.bytes)) = bytes
	*(*int8)(unsafe.Pointer(&r.volType)) = int8(t)
	*(*int16)(unsafe.Pointer(&r.nameLen)) = int16(l)
	copy(r.name[:], name)

	f, err := os.OpenFile("/dev/ubi"+strconv.Itoa(ubiNum),
		syscall.O_RDWR, 0666)
	if err != nil {
		return err
	}

	defer f.Close()

	err = gioctl.Ioctl(f.Fd(), gioctl.IoW(iocMagic, 0,
		unsafe.Sizeof(r)),
		uintptr(unsafe.Pointer(&r)))

	return
}
