package checkpoints

import (
	"github.com/iwind/TeaGo/maps"
	"github.com/oy1978/EdgeNode/internal/waf/requests"
	"github.com/oy1978/EdgeNode/internal/waf/utils"
)

type RequestIsCNAMECheckpoint struct {
	Checkpoint
}

func (this *RequestIsCNAMECheckpoint) RequestValue(req requests.Request, param string, options maps.Map, ruleId int64) (value any, hasRequestBody bool, sysErr error, userErr error) {
	if req.Format("${cname}") == req.Format("${host}") {
		value = 1
	} else {
		value = 0
	}
	return
}

func (this *RequestIsCNAMECheckpoint) ResponseValue(req requests.Request, resp *requests.Response, param string, options maps.Map, ruleId int64) (value any, hasRequestBody bool, sysErr error, userErr error) {
	if this.IsRequest() {
		return this.RequestValue(req, param, options, ruleId)
	}
	return
}

func (this *RequestIsCNAMECheckpoint) CacheLife() utils.CacheLife {
	return utils.CacheLongLife
}
