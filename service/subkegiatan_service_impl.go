package service

import (
	"context"
	"database/sql"
	"fmt"
	"subkegiatanServices/helper"
	"subkegiatanServices/model/domain"
	"subkegiatanServices/model/web"
	"subkegiatanServices/repository"

	"regexp"

	"github.com/go-playground/validator/v10"
)

type SubkegiatanServiceImpl struct {
	subkegiatanRepository repository.SubkegiatanRepository
	db                    *sql.DB
	validate              *validator.Validate
}

func NewSubkegiatanServiceImpl(subkegiatanRepository repository.SubkegiatanRepository, db *sql.DB, validate *validator.Validate) *SubkegiatanServiceImpl {
	return &SubkegiatanServiceImpl{
		subkegiatanRepository: subkegiatanRepository,
		db:                    db,
		validate:              validate,
	}
}

func (service *SubkegiatanServiceImpl) Create(ctx context.Context, request web.SubkegiatanCreateRequest) (web.SubkegiatanResponse, error) {
	err := service.validate.Struct(request)
	if err != nil {
		return web.SubkegiatanResponse{}, err
	}

	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		return web.SubkegiatanResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	if request.KodeSubKegiatan == "" {
		return web.SubkegiatanResponse{}, fmt.Errorf("kode subkegiatan wajib terisi")
	}

	regex := `^\d\.\d{2}\.\d{2}\.\d\.\d{2}\.\d{4}$`
	matched, err := regexp.MatchString(regex, request.KodeSubKegiatan)
	if err != nil {
		return web.SubkegiatanResponse{}, fmt.Errorf("error validasi format kode: %v", err)
	}
	if !matched {
		return web.SubkegiatanResponse{}, fmt.Errorf("format kode tidak valid")
	}

	Id := fmt.Sprintf("SUB-KEG-%v", request.KodeSubKegiatan)

	subkegiatan, err := service.subkegiatanRepository.Create(ctx, tx, domain.Subkegiatan{
		Id:              Id,
		NamaSubKegiatan: request.NamaSubKegiatan,
		KodeSubKegiatan: request.KodeSubKegiatan,
	})
	if err != nil {
		return web.SubkegiatanResponse{}, err
	}

	return web.SubkegiatanResponse{
		NamaSubKegiatan: subkegiatan.NamaSubKegiatan,
		KodeSubKegiatan: subkegiatan.KodeSubKegiatan,
	}, nil
}

func (service *SubkegiatanServiceImpl) Update(ctx context.Context, request web.SubkegiatanUpdateRequest) (web.SubkegiatanResponse, error) {
	err := service.validate.Struct(request)
	if err != nil {
		return web.SubkegiatanResponse{}, err
	}

	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		return web.SubkegiatanResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	if request.KodeSubKegiatan == "" {
		return web.SubkegiatanResponse{}, fmt.Errorf("kode subkegiatan wajib terisi")
	}

	regex := `^\d\.\d{2}\.\d{2}\.\d\.\d{2}\.\d{4}$`
	matched, err := regexp.MatchString(regex, request.KodeSubKegiatan)
	if err != nil {
		return web.SubkegiatanResponse{}, fmt.Errorf("error validasi format kode: %v", err)
	}
	if !matched {
		return web.SubkegiatanResponse{}, fmt.Errorf("format kode tidak valid")
	}

	subkegiatan := service.subkegiatanRepository.Update(ctx, tx, domain.Subkegiatan{
		Id:              request.Id,
		NamaSubKegiatan: request.NamaSubKegiatan,
		KodeSubKegiatan: request.KodeSubKegiatan,
	})

	return web.SubkegiatanResponse{
		NamaSubKegiatan: subkegiatan.NamaSubKegiatan,
		KodeSubKegiatan: subkegiatan.KodeSubKegiatan,
	}, nil
}

func (service *SubkegiatanServiceImpl) Delete(ctx context.Context, id string) error {
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	err = service.subkegiatanRepository.Delete(ctx, tx, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *SubkegiatanServiceImpl) FindById(ctx context.Context, id string) (web.SubkegiatanResponse, error) {
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		return web.SubkegiatanResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	subkegiatan, err := service.subkegiatanRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.SubkegiatanResponse{}, err
	}

	return web.SubkegiatanResponse{
		Id:              subkegiatan.Id,
		NamaSubKegiatan: subkegiatan.NamaSubKegiatan,
		KodeSubKegiatan: subkegiatan.KodeSubKegiatan,
	}, nil
}

func (service *SubkegiatanServiceImpl) FindAll(ctx context.Context) ([]web.SubkegiatanResponse, error) {
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		return []web.SubkegiatanResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	subkegiatans, err := service.subkegiatanRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.SubkegiatanResponse{}, err
	}

	subkegiatanResponses := []web.SubkegiatanResponse{}
	for _, subkegiatan := range subkegiatans {
		subkegiatanResponses = append(subkegiatanResponses, web.SubkegiatanResponse{
			Id:              subkegiatan.Id,
			NamaSubKegiatan: subkegiatan.NamaSubKegiatan,
			KodeSubKegiatan: subkegiatan.KodeSubKegiatan,
			CreatedAt:       subkegiatan.CreatedAt,
			UpdatedAt:       subkegiatan.UpdatedAt,
		})
	}

	return subkegiatanResponses, nil
}

func (service *SubkegiatanServiceImpl) FindByKodeSubKegiatan(ctx context.Context, kodeSubKegiatan string) (web.SubkegiatanResponse, error) {
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		return web.SubkegiatanResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	subkegiatan, err := service.subkegiatanRepository.FindByKodeSubKegiatan(ctx, tx, kodeSubKegiatan)
	if err != nil {
		return web.SubkegiatanResponse{}, err
	}

	return web.SubkegiatanResponse{
		NamaSubKegiatan: subkegiatan.NamaSubKegiatan,
		KodeSubKegiatan: subkegiatan.KodeSubKegiatan,
	}, nil
}
