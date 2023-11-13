// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package fnv_test

import (
	"github.com/iwind/TeaGo/types"
	"github.com/oy1978/EdgeNode/internal/utils/fnv"
	"testing"
)

func TestHash(t *testing.T) {
	for _, key := range []string{"costarring", "liquid", "hello"} {
		var h = fnv.HashString(key)
		t.Log(key + " => " + types.String(h))
	}
}

func BenchmarkHashString(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			fnv.HashString("abcdefh")
		}
	})
}
