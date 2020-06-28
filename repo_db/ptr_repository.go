package repo_db

import (
	"database/sql"
	"git.rappet.de/rappet/netmap-api/model"
	"git.rappet.de/rappet/netmap-api/repo"
)

type PtrRepository struct {
	db *sql.DB
}

func NewPtrRepository(db *sql.DB) PtrRepository {
	return PtrRepository{
		db: db,
	}
}

func (repo PtrRepository) GetPtrs(parameters repo.GetPtrsParameters) (model.Ptrs, error) {
	sqlStatement := "SELECT address, record, timestamp FROM ptrs WHERE record ILIKE $1 LIMIT 1000;"
	ptrs := make(model.Ptrs, 0)
	rows, err := repo.db.Query(sqlStatement, parameters.Like)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var ptr model.Ptr
		if err := rows.Scan(&ptr.Address, &ptr.Record, &ptr.Timestamp); err != nil {
			return nil, err
		}
		ptrs = append(ptrs, ptr)
	}
	return ptrs, nil
}

func (repo PtrRepository) GetPtr(address string) (*model.Ptr, error) {
	sqlStatement := "SELECT address, record, timestamp FROM ptrs WHERE address=$1;"
	var ptr model.Ptr
	row := repo.db.QueryRow(sqlStatement, address)
	switch err := row.Scan(&ptr.Address, &ptr.Record, &ptr.Timestamp); err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return &ptr, nil
	default:
		return nil, err
	}
}