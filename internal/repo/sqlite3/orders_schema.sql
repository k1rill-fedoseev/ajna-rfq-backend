CREATE TABLE IF NOT EXISTS orders
(
    hash                  TEXT    NOT NULL PRIMARY KEY ON CONFLICT ABORT,
    lp_order              INTEGER NOT NULL,
    chain_id              TEXT    NOT NULL,
    rfq                   BLOB    NOT NULL CHECK ( length(rfq) == 20 ),
    pool                  BLOB    NOT NULL CHECK ( length(pool) == 20 ),
    maker                 BLOB    NOT NULL CHECK ( length(maker) == 20 ),
    taker                 BLOB             DEFAULT NULL CHECK ( length(taker) == 20 ),
    "index"               INTEGER NOT NULL CHECK ( "index" > 0 AND "index" <= 7388 ),
    make_amount           TEXT    NOT NULL,
    min_make_amount       TEXT    NOT NULL,
    expiry                INTEGER NOT NULL,
    price                 TEXT    NOT NULL,
    signature             BLOB    NOT NULL CHECK ( length(signature) == 64 OR length(signature) == 65 ),
    remaining_make_amount TEXT    NOT NULL,
    approved_amount       TEXT    NOT NULL,
    created_at            INTEGER NOT NULL DEFAULT (unixepoch('now')),
    updated_at            INTEGER NOT NULL DEFAULT (unixepoch('now'))
);

CREATE INDEX IF NOT EXISTS orders_active_orders_index ON orders (updated_at) WHERE (remaining_make_amount != '0' AND
                                                                                    remaining_make_amount != '-1' AND
                                                                                    approved_amount != '0');

CREATE INDEX IF NOT EXISTS orders_updated_at_index ON orders (updated_at);
