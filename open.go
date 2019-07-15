// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package ubi

import (
	"os"
)

// Open finds the specified volume on the specified ubi device
// and returns a Volume structure for subsequent operations.
func Open(ubiNum int32, volName string) (v Volume, err error) {
	dev, err := FindVolume(ubiNum, volName)
	if err != nil {
		return
	}
	f, err := os.OpenFile(dev, os.O_RDWR, 0666)
	if err != nil {
		return
	}
	v = Volume{F: f}
	return
}
