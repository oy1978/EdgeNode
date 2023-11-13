// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package nodes_test

import (
	"github.com/oy1978/EdgeCommon/pkg/nodeconfigs"
	"github.com/oy1978/EdgeNode/internal/caches"
	"github.com/oy1978/EdgeNode/internal/nodes"
	"testing"
)

func TestHTTPCacheTaskManager_Loop(t *testing.T) {
	// initialize cache policies
	config, err := nodeconfigs.SharedNodeConfig()
	if err != nil {
		t.Fatal(err)
	}
	caches.SharedManager.UpdatePolicies(config.HTTPCachePolicies)

	var manager = nodes.NewHTTPCacheTaskManager()
	err = manager.Loop()
	if err != nil {
		t.Fatal(err)
	}
}
