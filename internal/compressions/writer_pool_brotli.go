// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package compressions

import (
	"github.com/andybalholm/brotli"
	teaconst "github.com/oy1978/EdgeNode/internal/const"
	"github.com/oy1978/EdgeNode/internal/utils"
	"io"
)

var sharedBrotliWriterPool *WriterPool

func init() {
	if !teaconst.IsMain {
		return
	}

	var maxSize = utils.SystemMemoryGB() * 256
	if maxSize == 0 {
		maxSize = 256
	}
	sharedBrotliWriterPool = NewWriterPool(maxSize, brotli.BestCompression, func(writer io.Writer, level int) (Writer, error) {
		return newBrotliWriter(writer, level)
	})
}
