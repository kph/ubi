// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package ubi

import (
	"os"
)

const (
	iocMagic = 'o'

	DynamicVol = 3
	StaticVol  = 4

	MaxVolName = 127
	MaxRnVol   = 32

	VolNumAuto = -1
)

type Ubictrl struct {
	f *os.File
}

// The types below are for future use. They should be moved into their
// own files when the methods are created.

type rsvolReq struct {
	bytes [4]byte
	volId [2]byte
}

type rnvolReq struct {
	count   [4]byte
	padding [12]byte
	ents    [MaxRnVol]struct {
		volId    [4]byte
		nameLen  [2]byte
		padding2 [2]byte
		name     [MaxVolName + 1]byte
	}
}

type lebChangeReq struct {
	lnum    [4]byte
	butes   [4]byte
	dtype   byte
	padding [7]byte
}

type mapReq struct {
	lnum    [4]byte
	dtype   byte
	padding [3]byte
}

type setVolPropReq struct {
	property byte
	padding  [7]byte
	value    [4]byte
}

type blkCreateReq struct {
	padding [128]byte
}
