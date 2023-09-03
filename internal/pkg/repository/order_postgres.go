package repository

import (
	"OrderService/internal/models"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderPostgres struct {
	db *pgxpool.Pool
}

func NewOrderPostgres(db *pgxpool.Pool) *OrderPostgres {
	return &OrderPostgres{
		db: db,
	}
}

func (r *OrderPostgres) Create(order *models.Order) (string, error) {
	tx, err := r.db.Begin(context.Background())
	if err != nil {
		return "0", err
	}

	_, err = r.db.Exec(context.Background(), `insert into order_delivery
	(name, phone, zip, city, address, region, email) values ($1, $2, $3, $4, $5, $6, $7)`,
		order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip,
		order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email)
	if err != nil {
		return "0", err
	}

	_, err = r.db.Exec(context.Background(), `insert into order_payment
	("transaction", request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
	values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		order.Payment.Transaction, order.Payment.RequestID, order.Payment.Currency, order.Payment.Provider,
		order.Payment.Amount, order.Payment.PaymentDT, order.Payment.Bank, order.Payment.DeliveryCost,
		order.Payment.GoodsTotal, order.Payment.CustomFee)
	if err != nil {
		return "0", err
	}
	var uid string
	row := r.db.QueryRow(context.Background(), `insert into orders
	(order_uid, track_number, entry, delivery, payment, locale, internal_signature, customer_id,
	delivery_service, shardkey, sm_id, data_created, oof_shard)
	values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING order_uid`,
		order.OrderUid, order.TrackNumber, order.Entry, order.Delivery.Name,
		order.Payment.Transaction, order.Locale, order.InternalSignature,
		order.CustomerID, order.DeliveryService, order.ShardKey,
		order.SmID, order.DateCreated, order.OofShard)

	if err := row.Scan(&uid); err != nil {
		return "0", err
	}

	for _, item := range order.Items {
		_, err := r.db.Exec(context.Background(), `insert into order_items
		(chrt_id, track_number, price, rid, name, sale, "size", total_price, nm_id, brand, status, order_uid)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`,
			item.ChrtID, item.TrackNumber, item.Price, item.Rid, item.Name,
			item.Sale, item.Size, item.TotalPrice, item.NmId, item.Brand, item.Status, item.OrderUid)
		if err != nil {
			return "0", err
		}
	}
	return uid, tx.Commit(context.Background())
}

func (r *OrderPostgres) GetAll() ([]models.Order, error) {
	var orders []models.Order

	// Выполните SQL-запрос для извлечения всех заказов
	query := `
        SELECT
            o.order_uid,
            o.track_number,
            o.entry,
            o.locale,
            o.internal_signature,
            o.customer_id,
            o.delivery_service,
            o.shardkey,
            o.sm_id,
            o.data_created,
            o.oof_shard,
            d.name AS delivery_name,
            d.phone AS delivery_phone,
            d.zip AS delivery_zip,
            d.city AS delivery_city,
            d.address AS delivery_address,
            d.region AS delivery_region,
            d.email AS delivery_email,
            p.transaction AS payment_transaction,
            p.request_id AS payment_request_id,
            p.currency AS payment_currency,
            p.provider AS payment_provider,
            p.amount AS payment_amount,
            p.payment_dt AS payment_date,
            p.bank AS payment_bank,
            p.delivery_cost AS payment_delivery_cost,
            p.goods_total AS payment_goods_total,
            p.custom_fee AS payment_custom_fee
        FROM
            orders o
        LEFT JOIN
            order_delivery d ON o.delivery = d.name
        LEFT JOIN
            order_payment p ON o.payment = p.transaction
    `

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order models.Order
		err := rows.Scan(
			&order.OrderUid,
			&order.TrackNumber,
			&order.Entry,
			&order.Locale,
			&order.InternalSignature,
			&order.CustomerID,
			&order.DeliveryService,
			&order.ShardKey,
			&order.SmID,
			&order.DateCreated,
			&order.OofShard,
			&order.Delivery.Name,
			&order.Delivery.Phone,
			&order.Delivery.Zip,
			&order.Delivery.City,
			&order.Delivery.Address,
			&order.Delivery.Region,
			&order.Delivery.Email,
			&order.Payment.Transaction,
			&order.Payment.RequestID,
			&order.Payment.Currency,
			&order.Payment.Provider,
			&order.Payment.Amount,
			&order.Payment.PaymentDT,
			&order.Payment.Bank,
			&order.Payment.DeliveryCost,
			&order.Payment.GoodsTotal,
			&order.Payment.CustomFee,
		)
		if err != nil {
			return nil, err
		}

		// Теперь выполните запрос для извлечения данных о товарах и добавьте их к заказу
		itemQuery := `
            SELECT
                i.chrt_id,
                i.track_number,
                i.price,
                i.rid,
                i.name,
                i.sale,
                i.size,
                i.total_price,
                i.nm_id,
                i.brand,
                i.status
            FROM
                order_items i
            WHERE
                i.order_uid = $1
        `

		itemRows, err := r.db.Query(context.Background(), itemQuery, order.OrderUid)
		if err != nil {
			return nil, err
		}
		defer itemRows.Close()

		for itemRows.Next() {
			var item models.Item
			err := itemRows.Scan(
				&item.ChrtID,
				&item.TrackNumber,
				&item.Price,
				&item.Rid,
				&item.Name,
				&item.Sale,
				&item.Size,
				&item.TotalPrice,
				&item.NmId,
				&item.Brand,
				&item.Status,
			)
			if err != nil {
				return nil, err
			}

			order.Items = append(order.Items, item)
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (r *OrderPostgres) GetByUID(uid string) (models.Order, error) {
	var order models.Order

	// Выполните SQL-запрос для извлечения информации о заказе
	query := `
        SELECT
            o.order_uid,
            o.track_number,
            o.entry,
            o.locale,
            o.internal_signature,
            o.customer_id,
            o.delivery_service,
            o.shardkey,
            o.sm_id,
            o.data_created,
            o.oof_shard,
            d.name AS delivery_name,
            d.phone AS delivery_phone,
            d.zip AS delivery_zip,
            d.city AS delivery_city,
            d.address AS delivery_address,
            d.region AS delivery_region,
            d.email AS delivery_email,
            p.transaction AS payment_transaction,
            p.request_id AS payment_request_id,
            p.currency AS payment_currency,
            p.provider AS payment_provider,
            p.amount AS payment_amount,
            p.payment_dt AS payment_date,
            p.bank AS payment_bank,
            p.delivery_cost AS payment_delivery_cost,
            p.goods_total AS payment_goods_total,
            p.custom_fee AS payment_custom_fee
        FROM
            orders o
        LEFT JOIN
            order_delivery d ON o.delivery = d.name
        LEFT JOIN
            order_payment p ON o.payment = p.transaction
        WHERE
            o.order_uid = $1
    `

	row := r.db.QueryRow(context.Background(), query, uid)

	// Сканируем результаты запроса в структуру Order
	err := row.Scan(
		&order.OrderUid,
		&order.TrackNumber,
		&order.Entry,
		&order.Locale,
		&order.InternalSignature,
		&order.CustomerID,
		&order.DeliveryService,
		&order.ShardKey,
		&order.SmID,
		&order.DateCreated,
		&order.OofShard,
		&order.Delivery.Name,
		&order.Delivery.Phone,
		&order.Delivery.Zip,
		&order.Delivery.City,
		&order.Delivery.Address,
		&order.Delivery.Region,
		&order.Delivery.Email,
		&order.Payment.Transaction,
		&order.Payment.RequestID,
		&order.Payment.Currency,
		&order.Payment.Provider,
		&order.Payment.Amount,
		&order.Payment.PaymentDT,
		&order.Payment.Bank,
		&order.Payment.DeliveryCost,
		&order.Payment.GoodsTotal,
		&order.Payment.CustomFee,
	)

	if err != nil {
		return models.Order{}, err
	}

	// Выполняем второй запрос для извлечения данных о товарах и сканируем их в срез
	itemQuery := `
        SELECT
            i.chrt_id,
            i.track_number,
            i.price,
            i.rid,
            i.name,
            i.sale,
            i.size,
            i.total_price,
            i.nm_id,
            i.brand,
            i.status
        FROM
            order_items i
        WHERE
            i.order_uid = $1
    `

	rows, err := r.db.Query(context.Background(), itemQuery, uid)
	if err != nil {
		return models.Order{}, err
	}

	for rows.Next() {
		var item models.Item
		err := rows.Scan(
			&item.ChrtID,
			&item.TrackNumber,
			&item.Price,
			&item.Rid,
			&item.Name,
			&item.Sale,
			&item.Size,
			&item.TotalPrice,
			&item.NmId,
			&item.Brand,
			&item.Status,
		)

		if err != nil {
			return models.Order{}, err
		}

		order.Items = append(order.Items, item)
	}

	return order, nil
}
