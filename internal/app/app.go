package app

import (
	"context"
	"net"
	"os"
	"os/signal"
	constant "sober_driver/const"
	"sober_driver/internal/config"
	"sober_driver/internal/repository"

	configs "sober_driver/configs"
	"sober_driver/internal/service"

	handler "sober_driver/internal/transport/http"
	"sober_driver/pkg/utils"

	"github.com/ArenAzibekyan/logrus-helper/logger"
	"github.com/sirupsen/logrus"
)

func App() {

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()
	log := logger.Default()

	db, err := utils.Open(ctx, &utils.DBconfig{
		AppName: constant.AppName,
		DBName:  constant.DBName,
		URL:     constant.DBURL,
	})
	if err != nil {
		log.Fatal("error DB")
	}

	repository := repository.NewRepository(db.DB)
	service := service.NewService(&service.Deps{Repo: repository})
	router := handler.NewHandler(service, &config.SwaggerConfig{
		Enabled: configs.Enabled,
		DirPath: "",
		URL:     configs.Url,
	}).Init()
	listen, err := net.Listen("tcp", configs.Adress)
	if err != nil {
		log.Fatal("Error read addres:" + err.Error())
	}

	go close(ctx, log, listen)

	if err = router.RunListener(listen); err != nil {
		log.Fatal("error RunListener")
	}
}

func close(ctx context.Context, log *logrus.Entry, listen net.Listener) {
	<-ctx.Done()

	if err := listen.Close(); err != nil {
		log.Fatal("Close" + err.Error())
	}
}
