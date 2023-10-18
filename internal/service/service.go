package service

import (
	"fmt"
	"github.com/linqcod/orders-manager/internal/model"
	"strings"
)

type OrderRepository interface {
	GetMainRacksInfo(orderIds []int64) (map[int64]*model.MainRack, error)
}

type OrderService struct {
	repo OrderRepository
}

func NewOrderService(repo OrderRepository) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

func (s OrderService) GetMainRacksInfo(orderIds []int64) (map[int64]*model.MainRack, error) {
	return s.repo.GetMainRacksInfo(orderIds)
}

func (s OrderService) PrintMainRacksInfo(info map[int64]*model.MainRack, orders string) {
	fmt.Printf("=+=+=+=\nСтраница сборки заказов %s\n", orders)

	var sb strings.Builder
	for _, rack := range info {
		fmt.Printf("\n===Стеллаж %s", rack.Title)

		for _, product := range rack.OrderProducts {
			fmt.Printf("\n%s (id=%d)\nЗаказ: %d, %d шт\n",
				product.ProductType,
				product.ProductId,
				product.OrderId,
				product.ProductCount,
			)
			if product.SecondaryRacks != nil {
				fmt.Print("доп стеллаж: ")
				for _, secRack := range product.SecondaryRacks {
					sb.WriteString(fmt.Sprintf("%s,", secRack.Title))
				}
				fmt.Printf("%s\n", strings.TrimSuffix(sb.String(), ","))
				sb.Reset()
			}
		}
	}
}
