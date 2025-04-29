package web

// @Description Response Subkegiatan
type SubkegiatanResponse struct {
	Id              string `json:"id,omitempty"`
	KodeSubKegiatan string `json:"kode_subkegiatan"`
	NamaSubKegiatan string `json:"nama_subkegiatan"`
	CreatedAt       string `json:"created_at,omitempty"`
	UpdatedAt       string `json:"updated_at,omitempty"`
}
