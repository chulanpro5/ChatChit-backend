package main

import (
	"go.uber.org/zap"
	"test-chat/internal/broadcaster"
	"test-chat/internal/old-client"
	"test-chat/internal/router"
	"test-chat/pkg/common"
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

	clientHandler := old_client.NewHandler(common.GetCommon())

	broadcaster := broadcaster.NewBroadcasterService(common.GetCommon())
	go broadcaster.Run()

	zap.L().Debug("Starting server")
	router.InitRouter(clientHandler)
	router.Start("localhost:8000")
}
