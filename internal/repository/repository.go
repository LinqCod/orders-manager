package repository

import (
	"context"
	"database/sql"
	"github.com/lib/pq"
	"github.com/linqcod/orders-manager/internal/model"
)

const (
	getOrdersInfoQuery = `
		SELECT o.id, o.product_id, o.count, p.type, pr.rack_id, pr.rack_type, r.title
		FROM orders o
		JOIN products p on o.product_id = p.id
		JOIN products_racks pr on p.type = pr.product_type
		JOIN racks r on r.id = pr.rack_id 
		WHERE o.id = ANY($1)
		ORDER BY r.id
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

func (r OrderRepository) GetOrdersInfo(orderIds []int64) ([]*model.OrderInfo, error) {
	var results []*model.OrderInfo

	rows, err := r.db.QueryContext(r.ctx, getOrdersInfoQuery, pq.Array(orderIds))
	if err != nil {
		return nil, err
	}
	if rows.Err() != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var result model.OrderInfo
		if err := rows.Scan(
			&result.OrderId,
			&result.ProductId,
			&result.Count,
			&result.ProductType,
			&result.RackId,
			&result.RackType,
			&result.RackTitle,
		); err != nil {
			return nil, err
		}
		results = append(results, &result)
	}

	return results, nil
}
