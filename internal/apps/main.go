// Copyright 2023 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package apps

import teaconst "github.com/oy1978/EdgeNode/internal/const"

func RunMain(f func()) {
	if teaconst.IsMain {
		f()
	}
}
