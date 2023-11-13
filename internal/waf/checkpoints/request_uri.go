package checkpoints

import (
	"github.com/iwind/TeaGo/maps"
	"github.com/oy1978/EdgeNode/internal/waf/requests"
	"github.com/oy1978/EdgeNode/internal/waf/utils"
)

type RequestURICheckpoint struct {
	Checkpoint
}

func (this *RequestURICheckpoint) RequestValue(req requests.Request, param string, options maps.Map, ruleId int64) (value any, hasRequestBody bool, sysErr error, userErr error) {
	if len(req.WAFRaw().RequestURI) > 0 {
		value = req.WAFRaw().RequestURI
	} else if req.WAFRaw().URL != nil {
		value = req.WAFRaw().URL.RequestURI()
	}
	return
}

func (this *RequestURICheckpoint) ResponseValue(req requests.Request, resp *requests.Response, param string, options maps.Map, ruleId int64) (value any, hasRequestBody bool, sysErr error, userErr error) {
	if this.IsRequest() {
		return this.RequestValue(req, param, options, ruleId)
	}
	return
}

func (this *RequestURICheckpoint) CacheLife() utils.CacheLife {
	return utils.CacheShortLife
}
