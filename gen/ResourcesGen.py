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
    "ItemConfigEquipment.json",
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

ResourcesFilePath_Base = os.path.abspath("..") + "\\resources"
ResourcesFilePath_ExcelOutput = ResourcesFilePath_Base + "\\ExcelOutput"
ResourcesFilePath_Floor = ResourcesFilePath_Base + "\\Config\\LevelOutput\\Floor"
ResourcesFilePath_Group = ResourcesFilePath_Base + "\\Config\\LevelOutput\\Group"


def readFileToB64(filepath: str):
    with open(filepath, "rb") as file:
        b64str = base64.b64encode(file.read()).decode()
    return b64str



if __name__ == "__main__":
    print("hkrpg-go resources2go")
    print("author: github@Doctor-yoi")
    print("outFile: $ProjectRoot$/gameData/Resources_b64.go")
    print("!!!WARNING! DO NOT OPEN IT IN ANY CODE EDITOR!!!")
    startTime = time.perf_counter()

    goCode = "///   Generate By ResourcesGen.py"
    goCode += "\n///    DO NOT EDIT!!!"
    goCode += "\n///    Generate Time: " + time.asctime(time.localtime())
    goCode += "\npackage gameData"

    print("Translate ExcelOutput...")
    goCode += "\n\nvar ("
    for file in ResourcesFileNameList_ExcelOutput:
        (fileName, fileSuffix) = os.path.splitext(file)
        filePath = ResourcesFilePath_ExcelOutput + "\\" + file
        b64Content = readFileToB64(filePath)
        goLine = "\n    " + fileName + " = \"" + str(b64Content) + "\""
        goCode += goLine
    goCode += "\n)"

    print("Translate Floor...")
    goCode += "\n\nfunc FloorList() map[string]string {"
    goCode += "\n    floorMap = make(map[string]string,0)"
    for root, dirs, files in os.walk(ResourcesFilePath_Floor, topdown=True):
        if root == ResourcesFilePath_Floor:
            continue

        for file in files:
            (fileName, fileSuffix) = os.path.splitext(file)
            filePath = root + "\\" + file
            b64Content = readFileToB64(filePath)
            goLine = "\n    floorMap[\"" + file + "\"] = \"" + str(b64Content) + "\""
            goCode += goLine
    goCode += "\n}"

    print("Translate Group...")
    goCode += "\n\nfunc GroupList() map[string]string {"
    goCode += "\n    groupMap = make(map[string]string,0)"
    for root, dirs, files in os.walk(ResourcesFilePath_Group, topdown=True):
        if root == ResourcesFilePath_Group:
            continue

        for file in files:
            (fileName, fileSuffix) = os.path.splitext(file)
            filePath = root + "\\" + file
            b64Content = readFileToB64(filePath)
            goLine = "\n    groupMap[\"" + file + "\"] = \"" + str(b64Content) + "\""
            goCode += goLine
    goCode += "\n}"

    print("Saving...")
    with open(os.path.abspath("..") + "\\gameData\\Resources_b64.go", "w") as goFile:
        goFile.write(goCode)

    endTime = time.perf_counter()
    print('Done in {:.4f}s'.format(endTime-startTime))
    exit(0)
