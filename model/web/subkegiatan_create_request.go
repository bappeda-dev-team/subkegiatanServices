package web

// @Description Create Request Subkegiatan
type SubkegiatanCreateRequest struct {
	NamaSubKegiatan string `json:"nama_subkegiatan" validate:"required"`
	KodeSubKegiatan string `json:"kode_subkegiatan" validate:"required"`
}
