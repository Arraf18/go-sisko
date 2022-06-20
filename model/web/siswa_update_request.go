package web

type SiswaUpdateRequest struct {
	Id            int    `validate:"required"`
	Nama          string `validate:"required,max=200,min=1" json:"nama"`
	Alamat        string `validate:"required,max=100,min=1" json:"alamat"`
	TanggalLahir  string `validate:"required,max=36,min=10" json:"tanggal_lahir"`
	TempatLahir   string `validate:"required,max=100,min=1" json:"tempat_lahir"`
	JenisKelamin  string `validate:"required,max=10,min=1" json:"jenis_kelamin"`
	Agama         string `validate:"required,max=20,min=1" json:"agama"`
	GolonganDarah string `validate:"required,max=2,min=1" json:"golongan_darah"`
	NoTelepon     string `validate:"required,max=20,min=1" json:"no_telepon"`
}
