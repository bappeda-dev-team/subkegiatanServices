package repository

import (
	"context"
	"database/sql"
	"subkegiatanServices/model/domain"
)

type SubkegiatanRepository interface {
	Create(ctx context.Context, tx *sql.Tx, subkegiatan domain.Subkegiatan) (domain.Subkegiatan, error)
	Update(ctx context.Context, tx *sql.Tx, subkegiatan domain.Subkegiatan) domain.Subkegiatan
	Delete(ctx context.Context, tx *sql.Tx, id string) error
	FindById(ctx context.Context, tx *sql.Tx, id string) (domain.Subkegiatan, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Subkegiatan, error)
	FindByKodeSubKegiatan(ctx context.Context, tx *sql.Tx, kodeSubKegiatan string) (domain.Subkegiatan, error)
}
