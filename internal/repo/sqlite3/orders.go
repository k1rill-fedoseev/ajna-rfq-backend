package sqlite3

import (
	"database/sql"
	_ "embed"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"

	"ajna-rfq/internal/repo"
)

//go:embed orders_schema.sql
var schema string

type OrdersRepoSQLite3 struct {
	db *sql.DB
}

func NewOrdersRepoSQLite3(file string) (*OrdersRepoSQLite3, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}
	if _, err := db.Exec(schema); err != nil {
		return nil, err
	}
	return &OrdersRepoSQLite3{
		db: db,
	}, nil
}

func (db *OrdersRepoSQLite3) ListOrders(filter repo.Filter, limit uint64) ([]*repo.StoredOrder, error) {
	q := sq.Select("hash", "lp_order", "chain_id", "rfq", "pool", "maker", "taker", "\"index\"", "make_amount", "min_make_amount", "expiry", "price", "signature", "remaining_make_amount", "approved_amount", "created_at", "updated_at").
		From("orders").
		Where(sq.Eq{"lp_order": filter.LpOrder}).
		Where(sq.Eq{"chain_id": filter.ChainId}).
		OrderBy("created_at DESC").
		Limit(limit)
	if len(filter.Pools) > 0 {
		q = q.Where(sq.Eq{"pool": filter.Pools})
	}
	if filter.CreatedBefore != nil {
		q = q.Where(sq.Lt{"created_at": filter.CreatedBefore})
	}
	if filter.Maker != nil && filter.Taker != nil {
		q = q.Where(sq.Eq{"maker": filter.Maker, "taker": filter.Taker})
	} else if filter.Maker != nil {
		q = q.Where(sq.Eq{"maker": filter.Maker})
	} else if filter.Taker != nil {
		q = q.Where(sq.Or{sq.Eq{"taker": filter.Taker}, sq.Eq{"taker": nil}})
	} else {
		q = q.Where(sq.Eq{"taker": nil})
	}
	if filter.Active {
		q = q.
			Where(sq.Gt{"expiry": time.Now().Add(time.Second * 15).Unix()}).
			Where(sq.NotEq{"remaining_make_amount": []string{"0", "-1"}})
	}

	rows, err := q.RunWith(db.db).Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]*repo.StoredOrder, 0, limit)
	for rows.Next() {
		order, err := scanOrder(rows)
		if err != nil {
			return nil, err
		}
		res = append(res, order)
	}
	return res, nil
}

func (db *OrdersRepoSQLite3) Upsert(order repo.StoredOrder) (*repo.StoredOrder, error) {
	row := db.db.QueryRow(`
		INSERT INTO orders (hash, lp_order, chain_id, rfq, pool, maker, taker, "index", make_amount, min_make_amount, expiry, price, signature, remaining_make_amount, approved_amount, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, unixepoch('now'), unixepoch('now'))
		ON CONFLICT (hash) DO UPDATE
		SET remaining_make_amount = excluded.remaining_make_amount,
		    approved_amount = excluded.approved_amount,
		    updated_at = excluded.updated_at
		WHERE chain_id = excluded.chain_id
		RETURNING *`,
		order.Hash, order.LpOrder, order.ChainId, order.RFQ, order.Pool, order.Maker, order.Taker, order.Index, order.MakeAmount, order.MinMakeAmount, order.Expiry, order.Price, order.Signature, order.RemainingMakeAmount, order.ApprovedAmount,
	)
	err := row.Err()
	if err != nil {
		return nil, err
	}
	return scanOrder(row)
}

func scanOrder(scanner sq.RowScanner) (*repo.StoredOrder, error) {
	order := new(repo.StoredOrder)
	err := scanner.Scan(&order.Hash, &order.LpOrder, &order.ChainId, &order.RFQ, &order.Pool, &order.Maker, &order.Taker, &order.Index, &order.MakeAmount, &order.MinMakeAmount, &order.Expiry, &order.Price, &order.Signature, &order.RemainingMakeAmount, &order.ApprovedAmount, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return order, nil
}
