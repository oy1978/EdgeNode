package checkpoints

import (
	"github.com/iwind/TeaGo/maps"
	"github.com/oy1978/EdgeNode/internal/waf/requests"
	"github.com/oy1978/EdgeNode/internal/waf/utils"
)

type RequestRemoteUserCheckpoint struct {
	Checkpoint
}

func (this *RequestRemoteUserCheckpoint) RequestValue(req requests.Request, param string, options maps.Map, ruleId int64) (value any, hasRequestBody bool, sysErr error, userErr error) {
	username, _, ok := req.WAFRaw().BasicAuth()
	if !ok {
		value = ""
		return
	}
	value = username
	return
}

func (this *RequestRemoteUserCheckpoint) ResponseValue(req requests.Request, resp *requests.Response, param string, options maps.Map, ruleId int64) (value any, hasRequestBody bool, sysErr error, userErr error) {
	if this.IsRequest() {
		return this.RequestValue(req, param, options, ruleId)
	}
	return
}

func (this *RequestRemoteUserCheckpoint) CacheLife() utils.CacheLife {
	return utils.CacheMiddleLife
}
