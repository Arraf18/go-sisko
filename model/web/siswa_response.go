package web

type SiswaResponse struct {
	Id            int    `json:"id"`
	Nama          string `json:"nama"`
	Alamat        string `json:"alamat"`
	TanggalLahir  string `json:"tanggal_lahir"`
	TempatLahir   string `json:"tempat_lahir"`
	JenisKelamin  string `json:"jenis_kelamin"`
	Agama         string `json:"Agama"`
	GolonganDarah string `json:"golongan_darah"`
	NoTelepon     string `json:"no_telepon"`
}
