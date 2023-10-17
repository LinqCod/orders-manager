package main

import (
	"context"
	"fmt"
	"github.com/linqcod/orders-manager/internal/model"
	"github.com/linqcod/orders-manager/internal/repository"
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

	repo := repository.NewOrderRepository(context.Background(), db)

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

	ordersInfo, err := repo.GetOrdersInfo(ids)
	if err != nil {
		logger.Fatal(err)
	}

	printOrdersInfo(ordersInfo, input)
}

func printOrdersInfo(info []*model.OrderInfo, orders string) {
	productSecondaryRackTitles := make(map[int64][]string)

	for _, orderInfo := range info {
		if orderInfo.RackType == "Secondary" {
			if _, ok := productSecondaryRackTitles[orderInfo.ProductId]; !ok {
				productSecondaryRackTitles[orderInfo.ProductId] = make([]string, 0)
			}
			productSecondaryRackTitles[orderInfo.ProductId] = append(productSecondaryRackTitles[orderInfo.ProductId], orderInfo.RackTitle)
		}
	}

	fmt.Printf("=+=+=+=\nСтраница сборки заказов %s\n", orders)

	var currentRackIndex int64 = -1
	for _, orderInfo := range info {
		if orderInfo.RackType == "Main" {
			if orderInfo.RackId != currentRackIndex {
				fmt.Printf("\n===Стеллаж %s", orderInfo.RackTitle)
				currentRackIndex = orderInfo.RackId
			}

			fmt.Printf("\n%s (id=%d)\nЗаказ: %d, %d шт\n",
				orderInfo.ProductType,
				orderInfo.ProductId,
				orderInfo.OrderId,
				orderInfo.Count,
			)
			if secRackTitles, ok := productSecondaryRackTitles[orderInfo.ProductId]; ok {
				fmt.Printf("доп стеллаж: %s\n", strings.Join(secRackTitles, ","))
			}
		}
	}
}
