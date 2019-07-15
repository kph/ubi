// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package ubi

// Read is used to read data into a slice. It returns the number of bytes
// read and any error.
func (v Volume) Read(p []byte) (n int, err error) {
	return v.F.Read(p)
}
