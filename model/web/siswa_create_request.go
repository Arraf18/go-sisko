package web

type SiswaCreateRequest struct {
	Nama          string `validate:"required,min=1,max=100" json:"nama"`
	Alamat        string `validate:"required,min=1,max=200" json:"alamat"`
	TanggalLahir  string `validate:"required,min=10,max=36" json:"tanggal_lahir"`
	TempatLahir   string `validate:"required,min=1,max=100" json:"tempat_lahir"`
	JenisKelamin  string `validate:"required,min=1,max=10" json:"jenis_kelamin"`
	Agama         string `validate:"required,min=1,max=20" json:"agama"`
	GolonganDarah string `validate:"required,min=1,max=2" json:"golongan_darah"`
	NoTelepon     string `validate:"required,min=1,max=20" json:"no_telepon"`
}
