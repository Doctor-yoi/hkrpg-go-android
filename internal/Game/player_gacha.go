package Game

import (
	"math/rand"
	"time"

	"hkrpg/gdconf"
	"hkrpg/protocol/cmd"
	"hkrpg/protocol/proto"
)

func (g *Game) HandleGetGachaInfoCsReq(payloadMsg []byte) {
	rsp := new(proto.GetGachaInfoScRsp)
	rsp.GachaInfoList = make([]*proto.GachaInfo, 0)

	for _, bannerslist := range gdconf.GetBannersMap() {
		gacha := g.GetDbGacha(bannerslist.Id)
		gachaInfoList := &proto.GachaInfo{
			HistoryUrl: "http://127.0.0.1:8080/api/gacha/history",      // 历史记录
			DetailUrl:  "https://www.bilibili.com/video/BV1X94y177QK/", // 卡池详情
			Featured:   bannerslist.RateUpItems5,                       // 五星up
			UpInfo:     bannerslist.RateUpItems4,                       // 四星up
			GachaId:    bannerslist.Id,
		}
		if bannerslist.GachaType == "Normal" {
			list := []uint32{1003, 1101, 1211}
			gachaInfoList.GachaCeiling = &proto.GachaCeiling{
				IsClaimed:  false,
				AvatarList: make([]*proto.GachaCeilingAvatar, 0),
				CeilingNum: gacha.CeilingNum,
			}
			for _, id := range list {
				avatarlist := &proto.GachaCeilingAvatar{
					RepeatedCnt: 0,
					AvatarId:    id,
				}
				gachaInfoList.GachaCeiling.AvatarList = append(gachaInfoList.GachaCeiling.AvatarList, avatarlist)
			}
		} else {
			gachaInfoList.BeginTime = bannerslist.BeginTime // 开始时间
			gachaInfoList.EndTime = bannerslist.EndTime     // 结束时间
		}

		rsp.GachaInfoList = append(rsp.GachaInfoList, gachaInfoList)
	}

	g.Send(cmd.GetGachaInfoScRsp, rsp)
}

func (g *Game) HandleGetGachaCeilingCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.GetGachaCeilingCsReq, payloadMsg)
	req := msg.(*proto.GetGachaCeilingCsReq)

	rsp := &proto.GetGachaCeilingScRsp{
		GachaType: req.GachaType,
	}
	list := []uint32{1003, 1101, 1211}
	rsp.GachaCeiling = &proto.GachaCeiling{
		IsClaimed:  false,
		AvatarList: make([]*proto.GachaCeilingAvatar, 0),
		CeilingNum: g.GetDbGacha(1001).CeilingNum,
	}
	for _, id := range list {
		avatarlist := &proto.GachaCeilingAvatar{
			RepeatedCnt: 0,
			AvatarId:    id,
		}
		rsp.GachaCeiling.AvatarList = append(rsp.GachaCeiling.AvatarList, avatarlist)
	}

	g.Send(cmd.GetGachaCeilingScRsp, rsp)
}

func (g *Game) DoGachaCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.DoGachaCsReq, payloadMsg)
	req := msg.(*proto.DoGachaCsReq)

	if req.GachaNum != 10 && req.GachaNum != 1 {
		return
	}
	// 先扣球再抽卡
	upBanners := gdconf.GetBannersMap()[req.GachaId]
	if upBanners.GachaType == "Normal" {
		g.SubtractMaterial(101, req.GachaNum)
	} else {
		g.SubtractMaterial(102, req.GachaNum)
	}

	rsp := &proto.DoGachaScRsp{
		GachaId:       req.GachaId,
		CeilingNum:    0,
		GachaItemList: make([]*proto.GachaItem, 0),
		GachaNum:      req.GachaNum,
	}
	for i := 0; i < int(req.GachaNum); i++ {
		id := g.GachaRandom(req.GachaId)
		isAvatar, isNew := g.AddGachaItem(id)
		gachaItemList := &proto.GachaItem{
			TransferItemList: &proto.ItemList{ItemList: make([]*proto.Item, 0)},
			IsNew:            isNew,
			GachaItem:        nil,
			TokenItem:        &proto.ItemList{ItemList: make([]*proto.Item, 0)},
		}
		gachaItem := &proto.Item{
			ItemId:      id,
			Level:       1,
			Num:         1,
			MainAffixId: 0,
			Rank:        1,
			Promotion:   0,
			UniqueId:    0,
		}
		if isAvatar {
			if isNew {

			} else {
				tokenItemList := &proto.Item{
					Num:    8,
					ItemId: 252,
				}
				gachaItemList.TokenItem.ItemList = append(gachaItemList.TokenItem.ItemList, tokenItemList)

				transferItemList := &proto.Item{
					Num:    1,
					ItemId: 10000 + id,
				}
				gachaItemList.TransferItemList.ItemList = append(gachaItemList.TransferItemList.ItemList, transferItemList)

				g.AddMaterial(252, 8)
				g.AddMaterial(10000+id, 1)
			}
		} else {
			tokenItemList := &proto.Item{
				Num:    42,
				ItemId: 251,
			}
			gachaItemList.TokenItem.ItemList = append(gachaItemList.TokenItem.ItemList, tokenItemList)
		}
		gachaItemList.GachaItem = gachaItem

		rsp.GachaItemList = append(rsp.GachaItemList, gachaItemList)
	}

	g.AddMaterial(251, req.GachaNum*42)

	g.Send(cmd.DoGachaScRsp, rsp)
}

func (g *Game) GachaRandom(gachaId uint32) uint32 {
	var (
		list3 []uint32 // 三星池
		list4 []uint32 // 四星池
		list5 []uint32 // 五星池
	)

	probability5, probability4 := g.GetProbability(gachaId)

	upBanners := gdconf.GetBannersMap()[gachaId]

	for _, equi := range gdconf.GetEquipmentConfigMap() {
		switch equi.Rarity {
		case "CombatPowerLightconeRarity3":
			list3 = append(list3, equi.EquipmentID)
		case "CombatPowerLightconeRarity4":
			list4 = append(list4, equi.EquipmentID)
		case "CombatPowerLightconeRarity5":
			list5 = append(list5, equi.EquipmentID)
		}
	}

	for _, avatar := range gdconf.GetAvatarDataMap() {
		// 过滤主角
		if avatar.AvatarId/100 == 80 {
			continue
		}
		switch avatar.Rarity {
		case "CombatPowerAvatarRarityType4":
			list4 = append(list4, avatar.AvatarId)
		case "CombatPowerAvatarRarityType5":
			list5 = append(list5, avatar.AvatarId)
		}
	}

	// 特殊情况处理
	gachaFb := g.GetDbGacha(gachaId)
	if gachaFb.Pity4 == 8 && gachaFb.CeilingNum == 88 {
		// 五星
		if gachaFb.FailedFeaturedItemPulls5 {
			idIndex := rand.Intn(len(upBanners.RateUpItems5))
			gachaFb.CeilingNum = 0
			gachaFb.FailedFeaturedItemPulls5 = false
			return upBanners.RateUpItems5[idIndex]
		} else {
			idIndex := rand.Intn(len(list5))
			gachaFb.CeilingNum = 0
			for _, id := range upBanners.RateUpItems5 {
				if list5[idIndex] == id {
					gachaFb.FailedFeaturedItemPulls5 = false
					break
				} else {
					gachaFb.FailedFeaturedItemPulls5 = true
				}
			}
			gachaFb.Pity4++
			return list5[idIndex]
		}
	}

	// 保底四星
	if gachaFb.Pity4 == 9 && !gachaFb.FailedFeaturedItemPulls4 {
		idIndex := rand.Intn(len(list4))

		for _, id := range upBanners.RateUpItems4 {
			if list4[idIndex] == id {
				gachaFb.FailedFeaturedItemPulls4 = false
				break
			} else {
				gachaFb.FailedFeaturedItemPulls4 = true
			}
		}
		gachaFb.Pity4 = 0
		gachaFb.CeilingNum++
		return list4[idIndex]
	}

	// 大保底四星
	if gachaFb.Pity4 == 9 && gachaFb.FailedFeaturedItemPulls4 {
		idIndex := rand.Intn(len(upBanners.RateUpItems4))
		gachaFb.Pity4 = 0
		gachaFb.FailedFeaturedItemPulls4 = false
		return upBanners.RateUpItems4[idIndex]
	}

	// 保底五星
	if gachaFb.CeilingNum == 89 && !gachaFb.FailedFeaturedItemPulls5 {
		idIndex := rand.Intn(len(list5))
		gachaFb.CeilingNum = 0
		for _, id := range upBanners.RateUpItems5 {
			if list5[idIndex] == id {
				gachaFb.FailedFeaturedItemPulls5 = false
				break
			} else {
				gachaFb.FailedFeaturedItemPulls5 = true
			}
		}
		gachaFb.Pity4++
		return list5[idIndex]
	}

	// 大保底五星
	if gachaFb.CeilingNum == 89 && gachaFb.FailedFeaturedItemPulls5 {
		idIndex := rand.Intn(len(upBanners.RateUpItems5))
		gachaFb.CeilingNum = 0
		gachaFb.FailedFeaturedItemPulls5 = false
		return upBanners.RateUpItems5[idIndex]
	}

	// 下面是概率
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumbe := rand.Intn(10000) + 1
	randomNumber := uint32(randomNumbe)

	if randomNumber >= probability5 {
		// 五星
		if gachaFb.FailedFeaturedItemPulls5 {
			idIndex := rand.Intn(len(upBanners.RateUpItems5))
			gachaFb.CeilingNum = 0
			gachaFb.FailedFeaturedItemPulls5 = false
			return upBanners.RateUpItems5[idIndex]
		} else {
			idIndex := rand.Intn(len(list5))
			gachaFb.CeilingNum = 0
			for _, id := range upBanners.RateUpItems5 {
				if list5[idIndex] == id {
					gachaFb.FailedFeaturedItemPulls5 = false
					break
				} else {
					gachaFb.FailedFeaturedItemPulls5 = true
				}
			}
			gachaFb.Pity4++
			return list5[idIndex]
		}
	}
	if randomNumber >= probability4 {
		// 四星
		if gachaFb.FailedFeaturedItemPulls4 {
			idIndex := rand.Intn(len(upBanners.RateUpItems4))
			gachaFb.Pity4 = 0
			gachaFb.FailedFeaturedItemPulls4 = false
			return upBanners.RateUpItems4[idIndex]
		} else {
			idIndex := rand.Intn(len(list4))
			for _, id := range upBanners.RateUpItems4 {
				if list4[idIndex] == id {
					gachaFb.FailedFeaturedItemPulls4 = false
					break
				} else {
					gachaFb.FailedFeaturedItemPulls4 = true
				}
			}
			gachaFb.Pity4 = 0
			gachaFb.CeilingNum++
			return list4[idIndex]
		}
	}
	// 三星
	idIndex := rand.Intn(len(list3))
	gachaFb.CeilingNum++
	gachaFb.Pity4++
	return list3[idIndex]
}

func (g *Game) GetProbability(gachaId uint32) (uint32, uint32) {
	var probability5 uint32
	var probability4 uint32
	probability5 = 60
	probability4 = 510

	gaCha := g.GetDbGacha(gachaId)

	if gaCha.CeilingNum >= 73 {
		probability5 += (gaCha.CeilingNum - 73) * 622
		return 10000 - probability5, 10000 - probability5 - probability4
	}

	return 10000 - probability5, 10000 - probability5 - probability4
}
