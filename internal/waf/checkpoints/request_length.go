package checkpoints

import (
	"github.com/iwind/TeaGo/maps"
	"github.com/oy1978/EdgeNode/internal/waf/requests"
	"github.com/oy1978/EdgeNode/internal/waf/utils"
)

type RequestLengthCheckpoint struct {
	Checkpoint
}

func (this *RequestLengthCheckpoint) RequestValue(req requests.Request, param string, options maps.Map, ruleId int64) (value any, hasRequestBody bool, sysErr error, userErr error) {
	value = req.WAFRaw().ContentLength
	return
}

func (this *RequestLengthCheckpoint) ResponseValue(req requests.Request, resp *requests.Response, param string, options maps.Map, ruleId int64) (value any, hasRequestBody bool, sysErr error, userErr error) {
	if this.IsRequest() {
		return this.RequestValue(req, param, options, ruleId)
	}
	return
}

func (this *RequestLengthCheckpoint) CacheLife() utils.CacheLife {
	return utils.CacheShortLife
}
