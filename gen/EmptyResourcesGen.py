import os
import base64
import time

ResourcesFileNameList_ExcelOutput = [
    "AvatarConfig.json",
    "ActivityLoginConfig.json",
    "ActivityPanel.json",
    "AvatarDemoConfig.json",
    "AvatarExpItemConfig.json",
    "AvatarPromotionConfig.json",
    "AvatarSkillTreeConfig.json",
    "BackGroundMusic.json",
    "ChallengeMazeConfig.json",
    "ChallengeStoryMazeExtra.json",
    "ChallengeTargetConfig.json",
    "ChallengeStoryTargetConfig.json",
    "CocoonConfig.json",
    "EquipmentConfig.json",
    "EquipmentExpItemConfig.json",
    "EquipmentExpType.json",
    "EquipmentPromotionConfig.json",
    "ExpType.json",
    "ItemConfig.json",
    "ItemConfigAvatar.json",
    "ItemConfigAvatarPlayerIcon.json",
    "ItemConfigAvatarRank.json",
    "ItemConfigBook.json",
    "ItemConfigDisk.json",
    "ItemConfigEquipment.json",
    "ItemConfigRelic.json",
    "LoadingDesc.json",
    "MapEntrance.json",
    "MappingInfo.json",
    "MazeBuff.json",
    "MazePlane.json",
    "MazeProp.json",
    "MonsterConfig.json",
    "NPCData.json",
    "NPCMonsterData.json",
    "PlaneEvent.json",
    "PlayerLevelConfig.json",
    "QuestData.json",
    "RelicConfig.json",
    "RelicExpItem.json",
    "RelicExpType.json",
    "RelicMainAffixConfig.json",
    "RelicSubAffixConfig.json",
    "RewardData.json",
    "RogueAreaConfig.json",
    "RogueMap.json",
    "RogueRoom.json",
    "RogueTalent.json",
    "ShopConfig.json",
    "ShopGoodsConfig.json",
    "SpecialAvatar.json",
    "StageConfig.json",
    "TextJoinConfig.json"
]

if __name__ == "__main__":
    print("hkrpg-go resources2go")
    print("author: github@Doctor-yoi")
    print("outFile: $ProjectRoot$/gameData/Resources_b64.go")

    goCode = "///    Generate By EmptyResourcesGen.py"
    goCode += "\n///    DO NOT EDIT!!!"
    goCode += "\n///    Generate Time: " + time.asctime(time.localtime())
    goCode += "\npackage gameData"

    print("Translate ExcelOutput...")
    goCode += "\n\nvar ("
    for file in ResourcesFileNameList_ExcelOutput:
        (fileName, fileSuffix) = os.path.splitext(file)
        goLine = "\n    " + fileName + " = \"" + "\""
        goCode += goLine
    goCode += "\n)"

    print("Translate Floor...")
    goCode += "\nvar floorMap map[string]string"
    goCode += "\nfunc FloorMap() map[string]string {"
    goCode += "\n    if floorMap == nil {"
    goCode += "\n        floorMap = make(map[string]string)"
    goCode += "\n    }"
    goCode += "\n    return floorMap"
    goCode += "\n}"

    print("Translate Group...")
    goCode += "\nvar groupMap map[string]string"
    goCode += "\nfunc GroupMap() map[string]string {"
    goCode += "\n    if groupMap == nil {"
    goCode += "\n        groupMap = make(map[string]string)"
    goCode += "\n    }"
    goCode += "\n    return groupMap"
    goCode += "\n}"

    print("Saving...")
    with open(os.path.abspath("..") + "/gameData/Resources_b64.go", "w+") as goFile:
        goFile.write(goCode)

    print('Done.')
    exit(0)