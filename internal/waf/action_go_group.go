package waf

import (
	"github.com/iwind/TeaGo/types"
	"github.com/oy1978/EdgeNode/internal/remotelogs"
	"github.com/oy1978/EdgeNode/internal/waf/requests"
	"net/http"
)

type GoGroupAction struct {
	BaseAction

	GroupId string `yaml:"groupId" json:"groupId"`
}

func (this *GoGroupAction) Init(waf *WAF) error {
	return nil
}

func (this *GoGroupAction) Code() string {
	return ActionGoGroup
}

func (this *GoGroupAction) IsAttack() bool {
	return false
}

func (this *GoGroupAction) WillChange() bool {
	return true
}

func (this *GoGroupAction) Perform(waf *WAF, group *RuleGroup, set *RuleSet, request requests.Request, writer http.ResponseWriter) (continueRequest bool, goNextSet bool) {
	nextGroup := waf.FindRuleGroup(types.Int64(this.GroupId))
	if nextGroup == nil || !nextGroup.IsOn {
		return true, true
	}

	b, _, nextSet, err := nextGroup.MatchRequest(request)
	if err != nil {
		remotelogs.Error("WAF", "GO_GROUP_ACTION: "+err.Error())
		return true, false
	}

	if !b {
		return true, false
	}

	return nextSet.PerformActions(waf, nextGroup, request, writer)
}
