// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package ubi

// Close closes the underlying Linux device associated with a volume
func (v Volume) Close() (err error) {
	return v.F.Close()
}
