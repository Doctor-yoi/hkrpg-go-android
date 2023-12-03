package Game

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *Game) GetRogueScoreRewardInfoCsReq() {
	rsp := new(proto.GetRogueScoreRewardInfoScRsp)
	rsp.ScoreRewardInfo = &proto.RogueScoreRewardInfo{
		HasTakenInitialScore: true,
		PoolRefreshed:        true,
		PoolId:               20,
	}

	g.send(cmd.GetRogueScoreRewardInfoScRsp, rsp)
}

func (g *Game) HandleGetCurSceneInfoCsReq(payloadMsg []byte) {
	rsp := new(proto.GetCurSceneInfoScRsp)
	pos := g.Player.Pos
	rot := g.Player.Rot
	rsp.Scene = &proto.SceneInfo{
		WorldId:         101,
		LeaderEntityId:  1,
		FloorId:         20001001,
		GameModeType:    2,
		PlaneId:         20001,
		EntryId:         2000101,
		EntityGroupList: make([]*proto.SceneEntityGroupInfo, 0),
	}
	entityGroup := &proto.SceneEntityGroupInfo{
		EntityList: make([]*proto.SceneEntityInfo, 0),
	}
	for _, avatarid := range g.Player.DbLineUp.LineUpList[g.Player.DbLineUp.MainLineUp].AvatarIdList {
		if avatarid == 0 {
			continue
		}
		entityList := &proto.SceneEntityInfo{
			EntityCase: &proto.SceneEntityInfo_Actor{Actor: &proto.SceneActorInfo{
				AvatarType:   proto.AvatarType_AVATAR_FORMAL_TYPE,
				BaseAvatarId: avatarid,
			}},
			Motion: &proto.MotionInfo{
				Pos: &proto.Vector{
					X: int32(pos.X),
					Y: int32(pos.Y),
					Z: int32(pos.Z),
				},
				Rot: &proto.Vector{
					X: int32(rot.X),
					Y: int32(rot.Y),
					Z: int32(rot.Z),
				},
			},
			EntityId: uint32(g.GetNextGameObjectGuid()),
		}
		entityGroup.EntityList = append(entityGroup.EntityList, entityList)
	}
	rsp.Scene.EntityGroupList = append(rsp.Scene.EntityGroupList, entityGroup)

	g.send(cmd.GetCurSceneInfoScRsp, rsp)
}
