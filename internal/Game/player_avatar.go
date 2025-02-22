package Game

import (
	"strconv"

	"hkrpg/gdconf"
	"hkrpg/protocol/cmd"
	"hkrpg/protocol/proto"
)

func (g *Game) HandleGetHeroBasicTypeInfoCsReq(payloadMsg []byte) {
	rsp := new(proto.GetHeroBasicTypeInfoScRsp)
	rsp.Gender = proto.Gender(g.PlayerPb.Avatar.Gender)
	rsp.CurBasicType = proto.HeroBasicType(g.PlayerPb.Avatar.CurMainAvatar)
	for _, heroBasic := range g.GetHeroBasicTypeInfo() {
		basicTypeInfoList := &proto.HeroBasicTypeInfo{
			BasicType:     proto.HeroBasicType(heroBasic.BasicType),
			SkillTreeList: make([]*proto.AvatarSkillTree, 0),
			Rank:          heroBasic.Rank,
		}
		for _, skill := range heroBasic.SkillTreeList {
			if skill.Level == 0 {
				continue
			}
			avatarSkillTree := &proto.AvatarSkillTree{
				PointId: skill.PointId,
				Level:   skill.Level,
			}
			basicTypeInfoList.SkillTreeList = append(basicTypeInfoList.SkillTreeList, avatarSkillTree)
		}
		rsp.BasicTypeInfoList = append(rsp.BasicTypeInfoList, basicTypeInfoList)
	}

	g.Send(cmd.GetHeroBasicTypeInfoScRsp, rsp)
}

func (g *Game) HandleGetAvatarDataCsReq(payloadMsg []byte) {
	rsp := new(proto.GetAvatarDataScRsp)
	rsp.IsGetAll = true
	rsp.AvatarList = make([]*proto.Avatar, 0)

	avatarDb := g.GetAvatar()

	for avatarId, _ := range avatarDb.Avatar {
		avatarList := g.GetAvatarById(avatarId)
		if avatarId/1000 == 8 {
			avatarList.SkilltreeList = make([]*proto.AvatarSkillTree, 0)
		}
		rsp.AvatarList = append(rsp.AvatarList, avatarList)
	}

	g.Send(cmd.GetAvatarDataScRsp, rsp)
}

func (g *Game) RankUpAvatarCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.RankUpAvatarCsReq, payloadMsg)
	req := msg.(*proto.RankUpAvatarCsReq)

	g.GetAvatar().Avatar[req.BaseAvatarId].Rank++
	g.SubtractMaterial(req.BaseAvatarId+10000, 1)
	g.AvatarPlayerSyncScNotify(req.BaseAvatarId)

	rsp := new(proto.GetChallengeScRsp)
	g.Send(cmd.RankUpAvatarScRsp, rsp)
}

func (g *Game) AvatarExpUpCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.AvatarExpUpCsReq, payloadMsg)
	req := msg.(*proto.AvatarExpUpCsReq)
	if req.BaseAvatarId == 0 {
		rsp := &proto.AvatarExpUpScRsp{}
		g.Send(cmd.AvatarExpUpScRsp, rsp)
		return
	}

	var pileItem []*Material // 需要删除的升级材料
	var delScoin uint32      // 扣除的信用点
	var addExp uint32        // 增加的经验

	// 从背包获取需要升级的角色
	dbAvatar := g.GetAvatar().Avatar[req.BaseAvatarId]
	if dbAvatar == nil {
		rsp := &proto.AvatarExpUpScRsp{}
		g.Send(cmd.AvatarExpUpScRsp, rsp)
		return
	}

	gdconfAvatar := gdconf.GetAvatarDataById(strconv.Itoa(int(req.BaseAvatarId)))

	// 遍历用来升级的材料
	for _, pileList := range req.ItemCostList.ItemList {
		// 如果没有则退出
		if pileList.GetPileItem() == nil {
			continue
		}
		pile := new(Material)
		pile.Tid = pileList.GetPileItem().ItemId
		pile.Num = pileList.GetPileItem().ItemNum

		pileItem = append(pileItem, pile)
		// 获取材料配置
		pileconf := gdconf.GetAvatarExpItemConfigById(strconv.Itoa(int(pileList.GetPileItem().ItemId)))
		if pileconf == nil {
			rsp := &proto.AvatarExpUpScRsp{}
			g.Send(cmd.AvatarExpUpScRsp, rsp)
			return
		}
		// 获取要扣多少信用点
		delScoin += pileconf.Exp / 10 * pileList.GetPileItem().ItemNum
		// 获取能添加多少经验
		addExp += pileconf.Exp * pileList.GetPileItem().ItemNum
	}

	// 计算添加后有多少经验
	exp := addExp + dbAvatar.Exp

	// 获取能升级到的等级和升级后经验
	level, exp, newExp := gdconf.GetExpTypeByLevel(gdconfAvatar.ExpGroup, exp, dbAvatar.Level, dbAvatar.PromoteLevel, dbAvatar.AvatarId)
	if level == 0 && exp == 0 {
		rsp := &proto.AvatarExpUpScRsp{}
		g.Send(cmd.AvatarExpUpScRsp, rsp)
	}

	dbAvatar.Exp = exp
	dbAvatar.Level = level

	// 扣除本次升级需要的信用点
	g.GetItem().MaterialMap[2] -= delScoin

	// 删除用来升级的材料
	if len(pileItem) != 0 {
		g.DelMaterialPlayerSyncScNotify(pileItem)
	}

	// 通知角色还有多少信用点
	g.PlayerPlayerSyncScNotify()
	// 返还升级材料
	if newExp >= 1000 {
		num := (newExp / 1000) % 10
		if num >= 5 {
			g.AddMaterial(212, num/5)
		}
		g.AddMaterial(211, num%5)
	}
	// 通知升级后角色消息
	g.AvatarPlayerSyncScNotify(req.BaseAvatarId)
	rsp := &proto.AvatarExpUpScRsp{}
	g.Send(cmd.AvatarExpUpScRsp, rsp)
}

func (g *Game) PromoteAvatarCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.PromoteAvatarCsReq, payloadMsg)
	req := msg.(*proto.PromoteAvatarCsReq)

	var pileItem []*Material // 需要删除的突破材料
	var delScoin uint32      // 扣除的信用点

	// 从背包获取需要升级的角色
	dbAvatar := g.GetAvatar().Avatar[req.BaseAvatarId]
	if dbAvatar == nil {
		rsp := &proto.AvatarExpUpScRsp{}
		g.Send(cmd.AvatarExpUpScRsp, rsp)
		return
	}

	// 遍历用来突破的材料
	for _, pileList := range req.ItemList {
		// 如果没有则退出
		if pileList.GetPileItem() == nil {
			continue
		}
		pile := new(Material)
		pile.Tid = pileList.GetPileItem().ItemId
		pile.Num = pileList.GetPileItem().ItemNum
		pileItem = append(pileItem, pile)
	}

	// 删除用来突破的材料
	if len(pileItem) != 0 {
		g.DelMaterialPlayerSyncScNotify(pileItem)
	}
	// 计算需要扣除的信用点
	delScoin = gdconf.GetAvatarPromotionConfigByLevel(dbAvatar.AvatarId, dbAvatar.PromoteLevel)
	// 增加突破等级
	dbAvatar.PromoteLevel++
	// 扣除本次升级需要的信用点
	g.GetItem().MaterialMap[2] -= delScoin
	// 通知升级后角色消息
	g.AvatarPlayerSyncScNotify(req.BaseAvatarId)
	rsp := new(proto.GetChallengeScRsp)
	// TODO 是的，没错，还是同样的原因
	g.Send(cmd.PromoteAvatarScRsp, rsp)
}

func (g *Game) UnlockSkilltreeCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.UnlockSkilltreeCsReq, payloadMsg)
	req := msg.(*proto.UnlockSkilltreeCsReq)

	var pileItem []*Material // 需要删除的升级材料

	avatarId := req.PointId / 1000 // 获取要升级技能的角色Id
	// TODO 此处要做主角特殊处理
	avatarDb := g.GetAvatar().Avatar[avatarId]
	if avatarDb == nil {
		rsp := &proto.UnlockSkilltreeScRsp{
			Retcode: uint32(proto.Retcode_RET_FAIL),
		}
		g.Send(cmd.UnlockSkilltreeScRsp, rsp)
	}

	// 遍历用来升级的材料
	for _, pileList := range req.ItemList {
		// 如果没有则退出
		if pileList.GetPileItem() == nil {
			continue
		}
		pile := new(Material)
		pile.Tid = pileList.GetPileItem().ItemId
		pile.Num = pileList.GetPileItem().ItemNum
		pileItem = append(pileItem, pile)
	}

	// 删除用来突破的材料
	if len(pileItem) != 0 {
		g.DelMaterialPlayerSyncScNotify(pileItem)
	}
	// 升级
	for id, skilltree := range g.PlayerPb.Avatar.Avatar[avatarId].SkilltreeList {
		if skilltree.PointId == req.PointId {
			avatarDb.SkilltreeList[id].Level = req.Level
		}
	}
	// 通知升级后角色消息
	g.AvatarPlayerSyncScNotify(avatarId)
	rsp := &proto.UnlockSkilltreeScRsp{
		BaseAvatarId: avatarId,
		PointId:      req.PointId,
		Level:        req.Level,
	}
	g.Send(cmd.UnlockSkilltreeScRsp, rsp)
}

func (g *Game) TakePromotionRewardCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.TakePromotionRewardCsReq, payloadMsg)
	req := msg.(*proto.TakePromotionRewardCsReq)
	avatarDb := g.GetAvatar().Avatar[req.BaseAvatarId]
	if avatarDb == nil {
		rsp := &proto.TakePromotionRewardScRsp{
			Retcode: uint32(proto.Retcode_RET_FAIL),
		}
		g.Send(cmd.TakePromotionRewardScRsp, rsp)
	}
	avatarDb.TakenRewards = append(avatarDb.TakenRewards, req.Promotion)
	// 通知升级后角色信息
	g.AvatarPlayerSyncScNotify(req.BaseAvatarId)

	item := &proto.Item{
		ItemId:      101,
		Level:       0,
		Num:         1,
		MainAffixId: 0,
		Rank:        0,
		Promotion:   0,
		UniqueId:    0,
	}

	g.AddMaterial(101, 1)

	rsq := &proto.TakePromotionRewardScRsp{
		RewardList: &proto.ItemList{ItemList: []*proto.Item{
			item,
		}},
	}
	g.Send(cmd.TakePromotionRewardScRsp, rsq)
}
