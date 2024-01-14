package Game

import (
	"hkrpg/protocol/cmd"
	"hkrpg/protocol/proto"
	spb "hkrpg/protocol/server"
)

func (g *Game) HandleQueryProductInfoCsReq(payloadMsg []byte) {
	rsp := new(proto.QuitLineupCsReq)
	// TODO 是的，没错，还是同样的原因
	g.Send(cmd.QueryProductInfoScRsp, rsp)
}

func (g *Game) SceneEntityMoveCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SceneEntityMoveCsReq, payloadMsg)
	req := msg.(*proto.SceneEntityMoveCsReq)

	if g.GetBattleState().BattleType == spb.BattleType_Battle_NONE {
		for _, entryId := range req.EntityMotionList {
			if g.Player.EntityList[entryId.EntityId] == nil {
				break
			}
			if g.Player.EntityList[entryId.EntityId].Entity == g.GetSceneAvatarId() {
				g.PlayerPb.Pos = &spb.VectorBin{
					X: entryId.Motion.Pos.X,
					Y: entryId.Motion.Pos.Y,
					Z: entryId.Motion.Pos.Z,
				}

				g.PlayerPb.Rot = &spb.VectorBin{
					X: entryId.Motion.Rot.X,
					Y: entryId.Motion.Rot.Y,
					Z: entryId.Motion.Rot.Z,
				}
			}
		}
	}

	rsq := new(proto.SceneEntityMoveCsReq)
	// TODO 是的，没错，还是同样的原因
	g.Send(cmd.SceneEntityMoveScRsp, rsq)
}
