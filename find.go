// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package ubi

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

var ErrVolumeNotFound = errors.New("Volume not found")

// FindVolume locates the named volume on the given UBI device and
// returns the associated Linux device.
func FindVolume(ubiNum int32, volName string) (devName string, err error) {
	prefix := "ubi" + strconv.Itoa(int(ubiNum))
	baseDir := "/sys/devices/virtual/ubi/ubi" + prefix + "/"

	files, err := ioutil.ReadDir(baseDir)
	for _, file := range files {
		name := file.Name()
		if !strings.HasPrefix(name, prefix+"_") {
			continue
		}
		ubiName, err := ioutil.ReadFile(baseDir + name + "/name")
		if err != nil {
			return "", err
		}
		if strings.TrimSpace(string(ubiName)) == volName {
			return "/dev/" + name, nil
		}
	}
	return "", ErrVolumeNotFound
}
