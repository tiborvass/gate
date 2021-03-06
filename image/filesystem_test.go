// Copyright (c) 2019 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"testing"
)

func TestFilesystem(*testing.T) {
	var _ BackingStore = new(Filesystem)
	var _ LocalStorage = new(Filesystem)
	var _ Storage = new(Filesystem)
}
