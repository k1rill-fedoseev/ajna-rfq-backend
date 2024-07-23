package sqlite3

import (
	"database/sql"
	_ "embed"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/ethereum/go-ethereum/common"
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

func (db *OrdersRepoSQLite3) ListMakeOrders(chainId string, maker common.Address, pools []common.Address, createdBefore *uint64) ([]*repo.StoredOrder, error) {
	q := sq.Select("hash", "chain_id", "rfq", "pool", "maker", "taker", "\"index\"", "make_amount", "min_make_amount", "expiry", "price", "signature", "remaining_make_amount", "approved_amount", "updated_at").
		From("orders").
		Where(sq.Eq{"chain_id": chainId}).
		Where(sq.Eq{"maker": maker}).
		OrderBy("created_at DESC").
		Limit(10)
	if len(pools) > 0 {
		q = q.Where(sq.Eq{"pool": pools})
	}
	if createdBefore != nil {
		q = q.Where(sq.Lt{"created_at": createdBefore})
	}

	rows, err := q.RunWith(db.db).Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]*repo.StoredOrder, 0, 50)
	for rows.Next() {
		order := new(repo.StoredOrder)
		err = rows.Scan(&order.Hash, &order.ChainId, &order.RFQ, &order.Pool, &order.Maker, &order.Taker, &order.Index, &order.MakeAmount, &order.MinMakeAmount, &order.Expiry, &order.Price, &order.Signature, &order.RemainingMakeAmount, &order.ApprovedAmount, &order.UpdatedAt)
		if err != nil {
			return nil, err
		}
		res = append(res, order)
	}
	return res, nil
}

func (db *OrdersRepoSQLite3) ListTakeOrders(chainId string, taker *common.Address, pools []common.Address, createdBefore *uint64) ([]*repo.StoredOrder, error) {
	q := sq.Select("hash", "chain_id", "rfq", "pool", "maker", "taker", "\"index\"", "make_amount", "min_make_amount", "expiry", "price", "signature", "remaining_make_amount", "approved_amount", "updated_at").
		From("orders").
		Where(sq.Eq{"chain_id": chainId}).
		Where(sq.NotEq{"remaining_make_amount": []string{"0", "-1"}, "approved_amount": "0"}).
		OrderBy("created_at DESC").
		Limit(50)
	if taker != nil {
		q = q.Where(sq.Or{sq.Eq{"taker": taker}, sq.Eq{"taker": nil}})
	} else {
		q = q.Where(sq.Eq{"taker": nil})
	}
	if len(pools) > 0 {
		q = q.Where(sq.Eq{"pool": pools})
	}
	if createdBefore != nil {
		q = q.Where(sq.Lt{"created_at": createdBefore})
	}

	rows, err := q.RunWith(db.db).Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]*repo.StoredOrder, 0, 50)
	for rows.Next() {
		order := new(repo.StoredOrder)
		err = rows.Scan(&order.Hash, &order.ChainId, &order.RFQ, &order.Pool, &order.Maker, &order.Taker, &order.Index, &order.MakeAmount, &order.MinMakeAmount, &order.Expiry, &order.Price, &order.Signature, &order.RemainingMakeAmount, &order.ApprovedAmount, &order.UpdatedAt)
		if err != nil {
			return nil, err
		}
		res = append(res, order)
	}
	return res, nil
}

func (db *OrdersRepoSQLite3) ListLastUpdatedActive() ([]*repo.StoredOrder, error) {
	rows, err := db.db.Query(`
		SELECT hash, chain_id, rfq, pool, maker, taker, "index", make_amount, min_make_amount, expiry, price, signature, remaining_make_amount, approved_amount
		FROM orders
		WHERE remaining_make_amount != '0'
	    AND remaining_make_amount != '-1'
		ORDER BY updated_at
		LIMIT 50`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]*repo.StoredOrder, 0, 50)
	for rows.Next() {
		log.Println("scan")
		order := new(repo.StoredOrder)
		err = rows.Scan(&order.Hash, &order.ChainId, &order.RFQ, &order.Pool, &order.Maker, &order.Taker, &order.Index, &order.MakeAmount, &order.MinMakeAmount, &order.Expiry, &order.Price, &order.Signature, &order.RemainingMakeAmount, &order.ApprovedAmount)
		log.Println("end scan")
		if err != nil {
			return nil, err
		}
		log.Println("end scan 2")
		res = append(res, order)
	}
	return res, nil
}

func (db *OrdersRepoSQLite3) Upsert(order *repo.StoredOrder) error {
	_, err := db.db.Exec(`
		INSERT INTO orders (hash, chain_id, rfq, pool, maker, taker, "index", make_amount, min_make_amount, expiry, price, signature, remaining_make_amount, approved_amount, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, unixepoch('now'), unixepoch('now'))
		ON CONFLICT (hash) DO UPDATE
		SET remaining_make_amount = excluded.remaining_make_amount,
		    updated_at = excluded.updated_at
		WHERE chain_id = excluded.chain_id`,
		order.Hash, order.ChainId, order.RFQ, order.Pool, order.Maker, order.Taker, order.Index, order.MakeAmount, order.MinMakeAmount, order.Expiry, order.Price, order.Signature, order.RemainingMakeAmount, order.ApprovedAmount,
	)
	return err
}
