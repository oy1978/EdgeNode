// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package metrics_test

import (
	timeutil "github.com/iwind/TeaGo/utils/time"
	"github.com/oy1978/EdgeNode/internal/metrics"
	"runtime"
	"testing"
)

func BenchmarkSumStat(b *testing.B) {
	runtime.GOMAXPROCS(2)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			metrics.SumStat(1, []string{"1.2.3.4"}, timeutil.Format("Ymd"), 1, 1)
		}
	})
}
