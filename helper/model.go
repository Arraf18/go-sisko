package helper

import (
	"github.com/Arraf18/go-sisko/model/domain"
	"github.com/Arraf18/go-sisko/model/web"
)

func ToSiswaResponse(siswa domain.Siswa) web.SiswaResponse {
	return web.SiswaResponse{
		Id:            siswa.Id,
		Nama:          siswa.Nama,
		Alamat:        siswa.Alamat,
		TanggalLahir:  siswa.TanggalLahir,
		TempatLahir:   siswa.TempatLahir,
		JenisKelamin:  siswa.JenisKelamin,
		Agama:         siswa.Agama,
		GolonganDarah: siswa.GolonganDarah,
		NoTelepon:     siswa.NoTelepon,
	}
}

func ToSiswaResponses(siswas []domain.Siswa) []web.SiswaResponse {
	var siswaResponses []web.SiswaResponse
	for _, siswa := range siswas {
		siswaResponses = append(siswaResponses, ToSiswaResponse(siswa))
	}
	return siswaResponses
}
