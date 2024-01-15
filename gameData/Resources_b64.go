///    Generate By EmptyResourcesGen.py
///    DO NOT EDIT!!!
///    Generate Time: Mon Jan 15 21:51:24 2024
package gameData

var (
	AvatarConfig               = ""
	ActivityLoginConfig        = ""
	ActivityPanel              = ""
	AvatarDemoConfig           = ""
	AvatarExpItemConfig        = ""
	AvatarPromotionConfig      = ""
	AvatarSkillTreeConfig      = ""
	BackGroundMusic            = ""
	ChallengeMazeConfig        = ""
	ChallengeStoryMazeExtra    = ""
	ChallengeTargetConfig      = ""
	ChallengeStoryTargetConfig = ""
	CocoonConfig               = ""
	EquipmentConfig            = ""
	EquipmentExpItemConfig     = ""
	EquipmentExpType           = ""
	EquipmentPromotionConfig   = ""
	ExpType                    = ""
	ItemConfig                 = ""
	ItemConfigAvatar           = ""
	ItemConfigAvatarPlayerIcon = ""
	ItemConfigAvatarRank       = ""
	ItemConfigBook             = ""
	ItemConfigDisk             = ""
	ItemConfigEquipment        = ""
	ItemConfigRelic            = ""
	LoadingDesc                = ""
	MapEntrance                = ""
	MappingInfo                = ""
	MazeBuff                   = ""
	MazePlane                  = ""
	MazeProp                   = ""
	MonsterConfig              = ""
	NPCData                    = ""
	NPCMonsterData             = ""
	PlaneEvent                 = ""
	PlayerLevelConfig          = ""
	QuestData                  = ""
	RelicConfig                = ""
	RelicExpItem               = ""
	RelicExpType               = ""
	RelicMainAffixConfig       = ""
	RelicSubAffixConfig        = ""
	RewardData                 = ""
	RogueAreaConfig            = ""
	RogueMap                   = ""
	RogueRoom                  = ""
	RogueTalent                = ""
	ShopConfig                 = ""
	ShopGoodsConfig            = ""
	SpecialAvatar              = ""
	StageConfig                = ""
	TextJoinConfig             = ""
)
var floorMap map[string]string

func FloorMap() map[string]string {
	if floorMap == nil {
		floorMap = make(map[string]string)
	}
	return floorMap
}

var groupMap map[string]string

func GroupMap() map[string]string {
	if groupMap == nil {
		groupMap = make(map[string]string)
	}
	return groupMap
}
