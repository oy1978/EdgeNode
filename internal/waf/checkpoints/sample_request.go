package checkpoints

import (
	"github.com/iwind/TeaGo/maps"
	"github.com/oy1978/EdgeNode/internal/waf/requests"
	"github.com/oy1978/EdgeNode/internal/waf/utils"
)

// SampleRequestCheckpoint just a sample checkpoint, copy and change it for your new checkpoint
type SampleRequestCheckpoint struct {
	Checkpoint
}

func (this *SampleRequestCheckpoint) RequestValue(req requests.Request, param string, options maps.Map, ruleId int64) (value any, hasRequestBody bool, sysErr error, userErr error) {
	return
}

func (this *SampleRequestCheckpoint) ResponseValue(req requests.Request, resp *requests.Response, param string, options maps.Map, ruleId int64) (value any, hasRequestBody bool, sysErr error, userErr error) {
	if this.IsRequest() {
		return this.RequestValue(req, param, options, ruleId)
	}
	return
}

func (this *SampleRequestCheckpoint) CacheLife() utils.CacheLife {
	return utils.CacheMiddleLife
}
