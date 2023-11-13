package checkpoints

import (
	"github.com/iwind/TeaGo/maps"
	"github.com/oy1978/EdgeNode/internal/waf/requests"
	"github.com/oy1978/EdgeNode/internal/waf/utils"
)

type RequestUserAgentCheckpoint struct {
	Checkpoint
}

func (this *RequestUserAgentCheckpoint) RequestValue(req requests.Request, param string, options maps.Map, ruleId int64) (value any, hasRequestBody bool, sysErr error, userErr error) {
	value = req.WAFRaw().UserAgent()
	return
}

func (this *RequestUserAgentCheckpoint) ResponseValue(req requests.Request, resp *requests.Response, param string, options maps.Map, ruleId int64) (value any, hasRequestBody bool, sysErr error, userErr error) {
	if this.IsRequest() {
		return this.RequestValue(req, param, options, ruleId)
	}
	return
}

func (this *RequestUserAgentCheckpoint) CacheLife() utils.CacheLife {
	return utils.CacheShortLife
}
