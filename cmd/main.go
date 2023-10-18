package main

import (
	"context"
	"github.com/linqcod/orders-manager/internal/repository"
	"github.com/linqcod/orders-manager/internal/service"
	"github.com/linqcod/orders-manager/pkg/config"
	"github.com/linqcod/orders-manager/pkg/database"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func init() {
	config.LoadConfig(".env")
}

func main() {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	baseLogger, err := loggerConfig.Build()
	if err != nil {
		log.Fatalf("error while building zap logger: %v", err)
	}

	logger := baseLogger.Sugar()

	// init db connection
	db, err := database.InitDB()
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		logger.Fatalf("error while trying to ping db: %v", err)
	}

	//init repo and service
	repo := repository.NewOrderRepository(context.Background(), db)
	serv := service.NewOrderService(repo)

	//main logic
	input := os.Args[1]

	idStrings := strings.Split(input, ",")
	ids := make([]int64, len(idStrings))
	for i, v := range idStrings {
		id, err := strconv.Atoi(v)
		if err != nil {
			logger.Fatalf("error while parsing user input: %v", err)
		}
		ids[i] = int64(id)
	}

	mainRacksInfo, err := serv.GetMainRacksInfo(ids)
	if err != nil {
		logger.Fatal(err)
	}

	serv.PrintMainRacksInfo(mainRacksInfo, input)
}
