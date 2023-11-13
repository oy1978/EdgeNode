// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package compressions

import (
	"github.com/klauspost/compress/zstd"
	teaconst "github.com/oy1978/EdgeNode/internal/const"
	"github.com/oy1978/EdgeNode/internal/utils"
	"io"
)

var sharedZSTDWriterPool *WriterPool

func init() {
	if !teaconst.IsMain {
		return
	}

	var maxSize = utils.SystemMemoryGB() * 256
	if maxSize == 0 {
		maxSize = 256
	}
	sharedZSTDWriterPool = NewWriterPool(maxSize, int(zstd.SpeedBestCompression), func(writer io.Writer, level int) (Writer, error) {
		return newZSTDWriter(writer, level)
	})
}
