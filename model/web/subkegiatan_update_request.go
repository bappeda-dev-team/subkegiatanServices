package web

// @Description Update Request Subkegiatan
type SubkegiatanUpdateRequest struct {
	Id              string `json:"id" `
	KodeSubKegiatan string `json:"kode_subkegiatan" validate:"required"`
	NamaSubKegiatan string `json:"nama_subkegiatan" validate:"required"`
}
