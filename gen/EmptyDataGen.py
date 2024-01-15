import base64
import os
import time

GameDataFileNameList = [
    "ActivityScheduling.json",
    "Banners.json",
    "localhost.crt",
    "localhost.key",
    "RogueMapGen.json"
]


if __name__ == '__main__':
    print("hkrpg-go gameData2go")
    print("author: github@Doctor-yoi")
    print("outFile: $ProjectRoot$/gameData/GameData_b64.go")

    goCode = "///    Generate By EmptyDataGen.py"
    goCode += "\n///    DO NOT EDIT!!!"
    goCode += "\n///    Generate Time: " + time.asctime(time.localtime())
    goCode += "\npackage gameData"
    goCode += "\n\nvar ("
    goCode += "\n    Ec2b = \"\""

    for file in GameDataFileNameList:
        print("Translate {}".format(file))
        (fileName, fileSuffix) = os.path.splitext(file)
        if fileName == "localhost":
            goLine = "\n    " + fileName + "_" + fileSuffix.replace(".", "") + " = \"" + "\""
        else:
            goLine = "\n    " + fileName + " = \"" + "\""
        goCode += goLine
    goCode += "\n)"
    with open(os.path.abspath("..") + "/gameData/GameData_b64.go", "w") as goFile:
        goFile.write(goCode)
    exit(0)