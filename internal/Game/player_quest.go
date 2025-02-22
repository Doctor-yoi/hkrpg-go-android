package Game

import (
	"hkrpg/gdconf"
	"hkrpg/protocol/cmd"
	"hkrpg/protocol/proto"
)

func (g *Game) HandleGetMissionStatusCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.GetMissionStatusCsReq, payloadMsg)
	req := msg.(*proto.GetMissionStatusCsReq)

	rsp := new(proto.GetMissionStatusScRsp)
	rsp.FinishedMainMissionIdList = []uint32{}
	rsp.SubMissionStatusList = make([]*proto.Mission, 0)
	for _, id := range req.MainMissionIdList {
		rsp.FinishedMainMissionIdList = append(rsp.FinishedMainMissionIdList, id)
	}
	for _, id := range req.SubMissionIdList {
		rsp.SubMissionStatusList = append(rsp.SubMissionStatusList, &proto.Mission{
			Id:     id,
			Status: proto.MissionStatus_MISSION_FINISH,
		})
	}

	g.Send(cmd.GetMissionStatusScRsp, rsp)
}

func (g *Game) GetQuestDataCsReq(payloadMsg []byte) {
	rsp := new(proto.GetQuestDataScRsp)
	rsp.QuestList = make([]*proto.Quest, 0)
	for _, questList := range gdconf.GetQuestDataMap() {
		quest := &proto.Quest{
			FinishTime: 10000,
			Progress:   questList.QuestType,
			Status:     proto.QuestStatus_QUEST_CLOSE,
			Id:         questList.QuestID,
		}
		rsp.QuestList = append(rsp.QuestList, quest)
	}

	g.Send(cmd.GetQuestDataScRsp, rsp)
}

func (g *Game) GetDailyActiveInfoCsReq(payloadMsg []byte) {
	dailyActiveQuestIdList := []uint32{2100132, 2100133, 2100134, 2100139, 2100150, 2100152, 2100153, 2100154}
	rsp := &proto.GetDailyActiveInfoScRsp{
		DailyActiveLevelList:   make([]*proto.DailyActivityInfo, 0),
		DailyActiveQuestIdList: dailyActiveQuestIdList,
		DailyActivePoint:       500,
	}

	for i := 1; i < 5; i++ {
		dailyActivityInfo := &proto.DailyActivityInfo{
			WorldLevel:       g.PlayerPb.WorldLevel,
			Level:            uint32(i),
			DailyActivePoint: uint32(i * 100),
			IsHasTaken:       false,
		}
		rsp.DailyActiveLevelList = append(rsp.DailyActiveLevelList, dailyActivityInfo)
	}

	g.Send(cmd.GetDailyActiveInfoScRsp, rsp)
}
