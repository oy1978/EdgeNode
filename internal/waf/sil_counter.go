// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package waf

import (
	"time"

	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/firewallconfigs"
	"github.com/TeaOSLab/EdgeNode/internal/utils"
	"github.com/TeaOSLab/EdgeNode/internal/utils/counters"
	"github.com/TeaOSLab/EdgeNode/internal/waf/requests"
	"github.com/iwind/TeaGo/types"
)

type silPageCode = string

const (
	SilPageCodeInit   silPageCode = "init"
	SilPageCodeShow   silPageCode = "show"
	SilPageCodeSubmit silPageCode = "submit"
)

// silIncreaseFails 增加sil失败次数，以便后续操作
func SilIncreaseFails(req requests.Request, actionConfig *SilAction, policyId int64, groupId int64, setId int64, pageCode silPageCode) (goNext bool) {
	var maxFails = actionConfig.MaxFails
	var failBlockTimeout = actionConfig.FailBlockTimeout
	if maxFails > 0 && failBlockTimeout > 0 {
		if maxFails <= 3 {
			maxFails = 3 // 不能小于3，防止意外刷新出现
		}
		var countFails = counters.SharedCounter.IncreaseKey(SilCacheKey(req, pageCode), 300)
		if int(countFails) >= maxFails {
			SharedIPBlackList.RecordIP(IPTypeAll, firewallconfigs.FirewallScopeService, req.WAFServerId(), req.WAFRemoteIP(), time.Now().Unix()+int64(failBlockTimeout), policyId, true, groupId, setId, "sil验证连续失败超过"+types.String(maxFails)+"次")
			return false
		}
	}
	return true
}

// silDeleteCacheKey 清除计数
func SilDeleteCacheKey(req requests.Request) {
	counters.SharedCounter.ResetKey(SilCacheKey(req, SilPageCodeInit))
	counters.SharedCounter.ResetKey(SilCacheKey(req, SilPageCodeShow))
	counters.SharedCounter.ResetKey(SilCacheKey(req, SilPageCodeSubmit))
}

// silCacheKey 获取sil缓存Key
func SilCacheKey(req requests.Request, pageCode silPageCode) string {
	var requestPath = req.WAFRaw().URL.Path

	if req.WAFRaw().URL.Path == SilPath {
		m, err := utils.SimpleDecryptMap(req.WAFRaw().URL.Query().Get("info"))
		if err == nil && m != nil {
			requestPath = m.GetString("url")
		}
	}

	return "WAF:sil:FAILS:" + pageCode + ":" + req.WAFRemoteIP() + ":" + types.String(req.WAFServerId()) + ":" + requestPath
}
