package repository

import (
	"context"
	"database/sql"
	"subkegiatanServices/helper"
	"subkegiatanServices/model/domain"
)

type SubkegiatanRepositoryImpl struct {
}

func NewSubkegiatanRepositoryImpl() *SubkegiatanRepositoryImpl {
	return &SubkegiatanRepositoryImpl{}
}

func (repository *SubkegiatanRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, subkegiatan domain.Subkegiatan) (domain.Subkegiatan, error) {
	script := "INSERT INTO tb_subkegiatan (id, nama_subkegiatan, kode_subkegiatan) VALUES ($1, $2, $3)"
	_, err := tx.ExecContext(ctx, script, subkegiatan.Id, subkegiatan.NamaSubKegiatan, subkegiatan.KodeSubKegiatan)
	if err != nil {
		return domain.Subkegiatan{}, err
	}
	return subkegiatan, nil
}

func (repository *SubkegiatanRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, subkegiatan domain.Subkegiatan) domain.Subkegiatan {
	script := "UPDATE tb_subkegiatan SET nama_subkegiatan = $1, kode_subkegiatan = $2 WHERE id = $3"
	_, err := tx.ExecContext(ctx, script, subkegiatan.NamaSubKegiatan, subkegiatan.KodeSubKegiatan, subkegiatan.Id)
	helper.PanicIfError(err)
	return subkegiatan
}

func (repository *SubkegiatanRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id string) error {
	script := "UPDATE tb_subkegiatan SET deleted_at = NOW() WHERE id = $1"
	_, err := tx.ExecContext(ctx, script, id)
	helper.PanicIfError(err)
	return nil
}

func (repository *SubkegiatanRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id string) (domain.Subkegiatan, error) {
	script := "SELECT id, nama_subkegiatan, kode_subkegiatan FROM tb_subkegiatan WHERE id = $1 AND deleted_at IS NULL"
	rows, err := tx.QueryContext(ctx, script, id)
	helper.PanicIfError(err)
	defer rows.Close()

	subkegiatan := domain.Subkegiatan{}
	if rows.Next() {
		err = rows.Scan(&subkegiatan.Id, &subkegiatan.NamaSubKegiatan, &subkegiatan.KodeSubKegiatan)
		helper.PanicIfError(err)
	}
	return subkegiatan, nil
}

func (repository *SubkegiatanRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Subkegiatan, error) {
	script := "SELECT id, nama_subkegiatan, kode_subkegiatan, created_at, updated_at FROM tb_subkegiatan WHERE deleted_at IS NULL"
	rows, err := tx.QueryContext(ctx, script)
	helper.PanicIfError(err)
	defer rows.Close()

	subkegiatans := []domain.Subkegiatan{}
	for rows.Next() {
		subkegiatan := domain.Subkegiatan{}
		err = rows.Scan(&subkegiatan.Id, &subkegiatan.NamaSubKegiatan, &subkegiatan.KodeSubKegiatan, &subkegiatan.CreatedAt, &subkegiatan.UpdatedAt)
		helper.PanicIfError(err)
		subkegiatans = append(subkegiatans, subkegiatan)
	}
	return subkegiatans, nil
}

func (repository *SubkegiatanRepositoryImpl) FindByKodeSubKegiatan(ctx context.Context, tx *sql.Tx, kodeSubKegiatan string) (domain.Subkegiatan, error) {
	script := "SELECT id, nama_subkegiatan, kode_subkegiatan FROM tb_subkegiatan WHERE kode_subkegiatan = $1 AND deleted_at IS NULL"
	rows, err := tx.QueryContext(ctx, script, kodeSubKegiatan)
	helper.PanicIfError(err)
	defer rows.Close()

	subkegiatan := domain.Subkegiatan{}
	if rows.Next() {
		err = rows.Scan(&subkegiatan.Id, &subkegiatan.NamaSubKegiatan, &subkegiatan.KodeSubKegiatan)
		helper.PanicIfError(err)
	}
	return subkegiatan, nil
}
