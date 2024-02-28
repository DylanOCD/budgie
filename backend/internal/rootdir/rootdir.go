/*
 * Copyright (c) 2024 Dylan O' Connor Desmond
 */

package rootdir

import (
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)

	// Root folder of the project
	Root = filepath.Join(filepath.Dir(b), "../..")
)
