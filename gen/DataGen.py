import base64
import os
import time

GameDataFileNameList = ["ActivityScheduling.json", "Banners.json", "localhost.crt", "localhost.key", "RogueMapGen.json"]


def readFileToB64(filepath: str):
    with open(filepath, "rb") as file:
        b64str = base64.b64encode(file.read()).decode()
    return b64str


if __name__ == '__main__':
    print("hkrpg-go gameData2go")
    print("author: github@Doctor-yoi")
    print("outFile: $ProjectRoot$/gameData/GameData_b64.go")

    goCode = "///    Generate By DataGen.py"
    goCode += "\n///    DO NOT EDIT!!!"
    goCode += "\n///    Generate Time: " + time.asctime(time.localtime())
    goCode += "\npackage gameData"
    goCode += "\n\nvar ("
    goCode += "\n    Ec2b = \"\""

    for file in GameDataFileNameList:
        print("Translate {}".format(file))
        (fileName, fileSuffix) = os.path.splitext(file)
        dataFilePath = os.path.abspath('..') + "\\data\\" + file
        b64Content = readFileToB64(dataFilePath)
        if fileName == "localhost":
            goLine = "\n    " + fileName + "_" + fileSuffix.replace(".", "") + " = \"" + str(b64Content) + "\""
        else:
            goLine = "\n    " + fileName + " = \"" + str(b64Content) + "\""
        goCode += goLine
    goCode += "\n)"
    with open(os.path.abspath("..") + "\\gameData\\GameData_b64.go", "w") as goFile:
        goFile.write(goCode)
    exit(0)