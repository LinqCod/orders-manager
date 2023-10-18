package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"github.com/linqcod/orders-manager/internal/model"
)

const (
	getOrdersByIdsQuery = `
		SELECT id, product_id, count
		FROM orders
		WHERE id = ANY($1)
	`
	getProductTypeByIdQuery = `
		SELECT type
		FROM products
		WHERE id = $1
	`
	getRacksQuery = `
		SELECT r.id, r.title, pr.rack_type
		FROM racks r 
		JOIN products_racks pr on $1 = pr.product_type AND r.id = pr.rack_id
	`
)

type OrderRepository struct {
	ctx context.Context
	db  *sql.DB
}

func NewOrderRepository(ctx context.Context, db *sql.DB) *OrderRepository {
	return &OrderRepository{
		ctx: ctx,
		db:  db,
	}
}

func (r OrderRepository) GetMainRacksInfo(orderIds []int64) (map[int64]*model.MainRack, error) {
	var orderProducts []*model.OrderProduct

	rows, err := r.db.QueryContext(r.ctx, getOrdersByIdsQuery, pq.Array(orderIds))
	if err != nil {
		return nil, err
	}
	if rows.Err() != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var orderProduct model.OrderProduct
		if err := rows.Scan(
			&orderProduct.OrderId,
			&orderProduct.ProductId,
			&orderProduct.ProductCount,
		); err != nil {
			return nil, err
		}
		orderProducts = append(orderProducts, &orderProduct)
	}

	mainRacks := make(map[int64]*model.MainRack)
	for i := range orderProducts {
		err = r.db.QueryRowContext(r.ctx, getProductTypeByIdQuery, orderProducts[i].ProductId).Scan(&orderProducts[i].ProductType)
		if err != nil {
			return nil, err
		}

		rows, err := r.db.QueryContext(r.ctx, getRacksQuery, orderProducts[i].ProductType)
		if err != nil {
			return nil, err
		}
		if rows.Err() != nil {
			return nil, err
		}
		defer rows.Close()

		var mainRackId int64 = -1
		mainRackTitle := ""
		for rows.Next() {
			var rack model.Rack
			if err := rows.Scan(
				&rack.Id,
				&rack.Title,
				&rack.Type,
			); err != nil {
				return nil, err
			}

			if rack.Type == "Main" {
				mainRackId = rack.Id
				mainRackTitle = rack.Title
			} else {
				if orderProducts[i].SecondaryRacks == nil {
					orderProducts[i].SecondaryRacks = make([]*model.SecondaryRack, 0)
				}
				orderProducts[i].SecondaryRacks = append(orderProducts[i].SecondaryRacks, &model.SecondaryRack{
					Id:    rack.Id,
					Title: rack.Title,
				})
			}
		}

		if mainRackId != -1 && mainRackTitle != "" {
			if _, ok := mainRacks[mainRackId]; !ok {
				mainRacks[mainRackId] = &model.MainRack{
					Title:         mainRackTitle,
					OrderProducts: make([]*model.OrderProduct, 0),
				}
			}
			mainRacks[mainRackId].OrderProducts = append(mainRacks[mainRackId].OrderProducts, orderProducts[i])
		} else {
			return nil, errors.New("error: no main rack for order product")
		}
	}

	return mainRacks, nil
}
