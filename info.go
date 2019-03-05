// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package ubi

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type DevInfo struct {
	Avail_eraseblocks int
	Bad_peb_count     int
	Eraseblock_size   int
	Total_eraseblocks int
	Volumes_count     int
}

func sysfileToI(f string) (i int, err error) {
	v, err := ioutil.ReadFile(f)
	if err != nil {
		return
	}

	i, err = strconv.Atoi(strings.TrimSuffix(string(v), "\n"))
	return
}

func (Ubictrl) Info(ubiNum int) (di DevInfo, err error) {
	d := "/sys/devices/virtual/ubi/ubi" + strconv.Itoa(ubiNum) + "/"
	di = DevInfo{}

	di.Avail_eraseblocks, err = sysfileToI(d + "avail_eraseblocks")
	if err != nil {
		return
	}
	di.Bad_peb_count, err = sysfileToI(d + "bad_peb_count")
	if err != nil {
		return
	}
	di.Eraseblock_size, err = sysfileToI(d + "eraseblock_size")
	if err != nil {
		return
	}
	di.Total_eraseblocks, err = sysfileToI(d + "total_eraseblocks")
	if err != nil {
		return
	}
	di.Volumes_count, err = sysfileToI(d + "volumes_count")
	return
}
