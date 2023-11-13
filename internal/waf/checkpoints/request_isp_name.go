// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package checkpoints

import (
	"github.com/iwind/TeaGo/maps"
	"github.com/oy1978/EdgeNode/internal/waf/requests"
	"github.com/oy1978/EdgeNode/internal/waf/utils"
)

type RequestISPNameCheckpoint struct {
	Checkpoint
}

func (this *RequestISPNameCheckpoint) IsComposed() bool {
	return false
}

func (this *RequestISPNameCheckpoint) RequestValue(req requests.Request, param string, options maps.Map, ruleId int64) (value any, hasRequestBody bool, sysErr error, userErr error) {
	value = req.Format("${isp.name}")
	return
}

func (this *RequestISPNameCheckpoint) ResponseValue(req requests.Request, resp *requests.Response, param string, options maps.Map, ruleId int64) (value any, hasRequestBody bool, sysErr error, userErr error) {
	return this.RequestValue(req, param, options, ruleId)
}

func (this *RequestISPNameCheckpoint) CacheLife() utils.CacheLife {
	return utils.CacheLongLife
}
