package service

import (
	"context"
	"subkegiatanServices/model/web"
)

type SubkegiatanService interface {
	Create(ctx context.Context, request web.SubkegiatanCreateRequest) (web.SubkegiatanResponse, error)
	Update(ctx context.Context, request web.SubkegiatanUpdateRequest) (web.SubkegiatanResponse, error)
	Delete(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (web.SubkegiatanResponse, error)
	FindAll(ctx context.Context) ([]web.SubkegiatanResponse, error)
	FindByKodeSubKegiatan(ctx context.Context, kodeSubKegiatan string) (web.SubkegiatanResponse, error)
}
