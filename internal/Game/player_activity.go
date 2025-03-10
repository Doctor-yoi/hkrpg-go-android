package Game

import (
	"hkrpg/gdconf"
	"hkrpg/protocol/cmd"
	"hkrpg/protocol/proto"
)

func (g *Game) HandleGetActivityScheduleConfigCsReq(payloadMsg []byte) {
	rsp := new(proto.GetActivityScheduleConfigScRsp)
	rsp.ActivityScheduleList = make([]*proto.ActivityScheduleInfo, 0)
	for _, activity := range gdconf.GetActivitySchedulingMap() {
		activityScheduleList := &proto.ActivityScheduleInfo{
			ActivityId: activity.ActivityId,
			EndTime:    activity.EndTime,
			ModuleId:   activity.ModuleId,
			BeginTime:  activity.BeginTime,
		}
		rsp.ActivityScheduleList = append(rsp.ActivityScheduleList, activityScheduleList)
	}

	g.Send(cmd.GetActivityScheduleConfigScRsp, rsp)
}

func (g *Game) GetLoginActivityCsReq() {
	rsp := &proto.GetLoginActivityScRsp{
		LoginActivityList: make([]*proto.LoginActivityData, 0),
	}

	loginActivity := g.GetLoginActivity()
	idList := gdconf.GetActivityLoginListById()

	for _, id := range idList {
		if loginActivity[id] == 0 {
			loginActivity[id] = 1
		}
	}

	for id, loginDays := range loginActivity {
		loginActivityData := &proto.LoginActivityData{
			Id:        id,
			LoginDays: loginDays,
		}
		rsp.LoginActivityList = append(rsp.LoginActivityList, loginActivityData)
	}

	g.Send(cmd.GetLoginActivityScRsp, rsp)
}

func (g *Game) TakeLoginActivityRewardCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.TakeLoginActivityRewardCsReq, payloadMsg)
	req := msg.(*proto.TakeLoginActivityRewardCsReq)

	activityLoginConfig := gdconf.GetActivityLoginConfigById(req.Id)
	rewardData := gdconf.GetRewardDataById(activityLoginConfig.RewardList[req.LoginDays-1])

	rsp := &proto.TakeLoginActivityRewardScRsp{
		TakeDays: req.LoginDays,
		Id:       req.Id,
		Reward: &proto.ItemList{
			ItemList: make([]*proto.Item, 0),
		},
	}
	if rewardData.Count_1 != 0 {
		item := &proto.Item{
			ItemId: rewardData.ItemID_1,
			Num:    rewardData.Count_1,
		}
		rsp.Reward.ItemList = append(rsp.Reward.ItemList, item)
		g.AddMaterial(rewardData.ItemID_1, rewardData.Count_1)
	}

	g.Send(cmd.TakeLoginActivityRewardScRsp, rsp)
}

func (g *Game) GetTrialActivityDataCsReq() {
	rsp := &proto.GetTrialActivityDataScRsp{
		TrialActivityList: make([]*proto.TrialActivityInfo, 0),
	}

	for _, id := range g.GetTrialActivity() {
		trialActivityInfo := &proto.TrialActivityInfo{TrialActivityId: id}
		rsp.TrialActivityList = append(rsp.TrialActivityList, trialActivityInfo)
	}

	g.Send(cmd.GetTrialActivityDataScRsp, rsp)

}

func (g *Game) TakeTrialActivityRewardCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.TakeTrialActivityRewardCsReq, payloadMsg)
	req := msg.(*proto.TakeTrialActivityRewardCsReq)

	rsp := &proto.TakeTrialActivityRewardScRsp{
		TrialActivityId: req.TrialActivityId,
		Reward: &proto.ItemList{
			ItemList: make([]*proto.Item, 0),
		},
	}
	item := &proto.Item{
		ItemId: 102,
		Num:    100,
	}
	rsp.Reward.ItemList = append(rsp.Reward.ItemList, item)
	g.AddMaterial(102, 100)

	g.Send(cmd.TakeTrialActivityRewardScRsp, rsp)
}
