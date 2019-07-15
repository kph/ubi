// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package ubi

// Write is used to write to an underlying UBI device. We break up the
// caller's slice into 8KiB chunks to support large writes from our callers
// while controlling the kernel's memory use. The kernel needs to allocate
// contiguous pages of the size we request, so we pick a size that is
// 2 pages assuming 4KiB pages.
func (v Volume) Write(p []byte) (n int, err error) {
	for len(p) > 0 {
		l := len(p)
		if l > 8192 {
			l = 8192
		}
		q := p[:l]
		nn := 0
		nn, err = v.F.Write(q)
		n = n + nn
		if err != nil {
			return
		}
		p = p[l:]
	}
	return
}
