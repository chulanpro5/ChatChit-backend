package main

import (
	"go.uber.org/zap"
	"test-chat/internal/broadcaster"
	"test-chat/internal/client"
	"test-chat/internal/common"
	"test-chat/internal/router"
)

func main() {
	logger := zap.Must(zap.NewDevelopment())
	zap.ReplaceGlobals(logger)

	c := common.NewCommon()
	defer c.CleanUp()

	//roomService := room.NewRoomService(c)
	//users, err := roomService.GetUsersByRoomID(1)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(users)

	clientHandler := client.NewHandler(common.GetCommon())

	broadcaster := broadcaster.NewBroadcasterService(common.GetCommon())
	go broadcaster.Run()

	zap.L().Debug("Starting server")
	router.InitRouter(clientHandler)
	router.Start("localhost:8000")
}
