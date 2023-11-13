package iplibrary_test

import (
	"github.com/iwind/TeaGo/maps"
	"github.com/oy1978/EdgeCommon/pkg/rpc/pb"
	"github.com/oy1978/EdgeCommon/pkg/serverconfigs/firewallconfigs"
	"github.com/oy1978/EdgeNode/internal/iplibrary"
	"testing"
	"time"
)

func TestIPSetAction_Init(t *testing.T) {
	action := iplibrary.NewIPSetAction()
	err := action.Init(&firewallconfigs.FirewallActionConfig{
		Params: maps.Map{
			"path":      "/usr/bin/iptables",
			"whiteName": "white-list",
			"blackName": "black-list",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}

func TestIPSetAction_AddItem(t *testing.T) {
	var action = iplibrary.NewIPSetAction()
	action.SetConfig(&firewallconfigs.FirewallActionIPSetConfig{
		Path:          "/usr/bin/iptables",
		WhiteName:     "white-list",
		BlackName:     "black-list",
		WhiteNameIPv6: "white-list-ipv6",
		BlackNameIPv6: "black-list-ipv6",
	})
	{
		err := action.AddItem(iplibrary.IPListTypeWhite, &pb.IPItem{
			Type:      "ipv4",
			Id:        1,
			IpFrom:    "192.168.1.100",
			ExpiredAt: time.Now().Unix() + 30,
		})
		if err != nil {
			t.Fatal(err)
		}
		t.Log("ok")
	}
	{
		err := action.AddItem(iplibrary.IPListTypeWhite, &pb.IPItem{
			Type:      "ipv4",
			Id:        1,
			IpFrom:    "1:2:3:4",
			ExpiredAt: time.Now().Unix() + 30,
		})
		if err != nil {
			t.Fatal(err)
		}
		t.Log("ok")
	}
	{
		err := action.AddItem(iplibrary.IPListTypeBlack, &pb.IPItem{
			Type:      "ipv4",
			Id:        1,
			IpFrom:    "192.168.1.100",
			ExpiredAt: time.Now().Unix() + 30,
		})
		if err != nil {
			t.Fatal(err)
		}
		t.Log("ok")
	}
	{
		err := action.AddItem(iplibrary.IPListTypeBlack, &pb.IPItem{
			Type:      "ipv4",
			Id:        1,
			IpFrom:    "1:2:3:4",
			ExpiredAt: time.Now().Unix() + 30,
		})
		if err != nil {
			t.Fatal(err)
		}
		t.Log("ok")
	}
}

func TestIPSetAction_DeleteItem(t *testing.T) {
	action := iplibrary.NewIPSetAction()
	err := action.Init(&firewallconfigs.FirewallActionConfig{
		Params: maps.Map{
			"path":      "/usr/bin/firewalld",
			"whiteName": "white-list",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	err = action.DeleteItem(iplibrary.IPListTypeWhite, &pb.IPItem{
		Type:      "ipv4",
		Id:        1,
		IpFrom:    "192.168.1.100",
		ExpiredAt: time.Now().Unix() + 30,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}
