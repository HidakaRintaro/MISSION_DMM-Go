package dao

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/pkg/errors"

	"github.com/jmoiron/sqlx"
)

type (
	status struct {
		db *sqlx.DB
	}
)

func NewStatus(db *sqlx.DB) repository.Status {
	return &status{db: db}
}

// CreateStatus : statusとメディアIdからstatusを作成
func (r *status) CreateStatus(_ context.Context, accountId int64, content string) (int64, error) {
	stmt, err := r.db.Preparex("insert into status (account_id, content) values (?, ?)")
	if err != nil {
		return 0, err
	}
	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	result, err := stmt.Exec(accountId, content)
	if err != nil {
		return 0, err
	}

	var id int64
	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// FindById : idからstatusを取得
func (r *status) FindById(ctx context.Context, id int64) (*object.Status, error) {
	entity := new(object.Status)
	err := r.db.QueryRowxContext(ctx, "select * from status where id = ?", id).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w", err)
	}

	return entity, nil
}

func (r *status) DeleteById(_ context.Context, id int64) error {
	stmt, err := r.db.Preparex("delete from status where id = ?")
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	// TODO 削除できるものがない時のエラーハンドリングをする
	// そもそも削除できなかったらDBのエラーになる？？
	// row, err := result.RowsAffected()
	// if err != nil {
	//	return err
	// }
	// if row == 0 {
	//	return errors.New("here were no status that could be deleted")
	// }

	return nil
}

// FindByQuery : クエリパラメータに一致するstatusを取得(TimeLine)
func (r *status) FindByQuery(ctx context.Context, qp map[string]int) ([]*object.Status, error) {
	entity := []*object.Status{}
	query := "select * from status"
	var args []interface{}
	var WhereAry []string

	if len(qp) > 1 {
		if v, ok := qp["max_id"]; ok {
			WhereAry = append(WhereAry, "id < ?")
			args = append(args, v)
		}
		if v, ok := qp["since_id"]; ok {
			WhereAry = append(WhereAry, "id > ?")
			args = append(args, v)
		}
		query += " where " + strings.Join(WhereAry, " and ")
	}

	args = append(args, qp["limit"])
	query += " limit ?"

	err := r.db.SelectContext(ctx, &entity, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("%w", err)
	}

	return entity, nil
}
