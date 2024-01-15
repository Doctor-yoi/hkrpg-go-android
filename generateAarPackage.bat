@echo off
go get golang.org/x/mobile/bind

cd ./gen
python DataGen.py
python ResourcesGen.py
cd ../

gomobile bind -target=android hkrpg

./generateEmptyGameData.bat