// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package clock_test

import (
	"github.com/oy1978/EdgeNode/internal/utils/clock"
	"testing"
)

func TestReadServer(t *testing.T) {
	t.Log(clock.NewClockManager().ReadServer("pool.ntp.org"))
}
