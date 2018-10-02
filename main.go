package main

import (
	"echo-rest-api/api"
	"echo-rest-api/config"
	"echo-rest-api/service"
	"echo-rest-api/store"
	"flag"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func main() {
	var err error
	log.SetFormatter(&log.JSONFormatter{})
	// Получаем флаги запуска приложения
	configFile := flag.String("c", "config.yaml", "Path to config file")
	flag.Parse()
	// Загружаем конфиг
	var conf *config.Config
	if conf, err = config.NewConfig(*configFile); err != nil {
		log.Fatal(err)
	}
	// Конфигурируем логгер
	log.SetLevel(log.Level(conf.LogLevel))
	log.Info("Starting service with configuration: ", conf.ConfigFile)
	// Создаем сторедж
	store1, err := store.NewStore(conf)
	if err != nil {
		log.Fatal(err)
	}
	defer store1.Close()
	log.Info("Store created successfully")
	// Создаем сервисы
	cs := service.NewCategoryService(store1)
	ps := service.NewProductService(store1)
	log.Info("Services created successfully")
	// Создаем  Api
	api1 := api.NewApi(conf, cs, ps)
	log.WithField("address", api1.GetApiInfo().Address).
		WithField("mw", api1.GetApiInfo().MW).
		WithField("routs", api1.GetApiInfo().Routs).
		Info("Starting api")
	log.Fatal(api1.Start())
}
