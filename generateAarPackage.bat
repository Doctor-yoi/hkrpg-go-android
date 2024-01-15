@echo off
go get golang.org/x/mobile/bind
./generateGameData.bat
gomobile bind -target=android hkrpg
./generateEmptyGameData.bat