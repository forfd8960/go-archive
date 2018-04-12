package model

import (
	"context"
	"database/sql"
	"time"

	"github.com/forfd8960/go-archive/db"
)

type ArchiveItem struct {
	ID        int64
	FromHost  string
	Name      string
	URL       string
	CollectAt time.Time
}

type itemImpl struct {
	ctx context.Context
	db  db.DBer
}

func NewArchiveItemer(ctx context.Context, db db.DBer) *itemImpl {
	return &itemImpl{ctx: ctx, db: db}
}

func (item *itemImpl) GetItemList(offset, limit int32) ([]*ArchiveItem, error) {
	SQL := "select id, from_host, name, url, collect_at from archive_item order by collect_at desc limit ?,?"

	rows, err := item.db.Query(SQL, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*ArchiveItem
	if rows.Next() {
		var t = &ArchiveItem{}
		if err := rows.Scan(&t.ID, &t.FromHost, &t.Name, &t.URL, &t.CollectAt); err != nil {
			return nil, err
		}

		items = append(items, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (item *itemImpl) GetItemCount() (int64, error) {
	countSQL := "select count(id) from archive_item"

	var count int64
	if err := item.db.QueryRow(countSQL).Scan(&count); err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}

		return -1, err
	}

	return count, nil
}
