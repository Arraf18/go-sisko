package service

import (
	"context"
	"database/sql"
	"github.com/Arraf18/go-sisko/exception"
	"github.com/Arraf18/go-sisko/helper"
	"github.com/Arraf18/go-sisko/model/domain"
	"github.com/Arraf18/go-sisko/model/web"
	"github.com/Arraf18/go-sisko/repository"
	"github.com/go-playground/validator"
)

type SiswaServiceImpl struct {
	SiswaRepository repository.SiswaRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewSiswaService(siswaRepository repository.SiswaRepository, DB *sql.DB, validate *validator.Validate) SiswaService {
	return &SiswaServiceImpl{
		SiswaRepository: siswaRepository,
		DB:              DB,
		Validate:        validate,
	}
}

func (service *SiswaServiceImpl) Create(ctx context.Context, request web.SiswaCreateRequest) web.SiswaResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	siswa := domain.Siswa{
		Nama:          request.Nama,
		Alamat:        request.Alamat,
		TanggalLahir:  request.TanggalLahir,
		TempatLahir:   request.TempatLahir,
		JenisKelamin:  request.JenisKelamin,
		Agama:         request.Agama,
		GolonganDarah: request.GolonganDarah,
		NoTelepon:     request.NoTelepon,
	}

	siswa = service.SiswaRepository.Save(ctx, tx, siswa)

	return helper.ToSiswaResponse(siswa)
}

func (service *SiswaServiceImpl) Update(ctx context.Context, request web.SiswaUpdateRequest) web.SiswaResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	siswa, err := service.SiswaRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	siswa.Nama = request.Nama
	siswa.Alamat = request.Alamat
	siswa.TanggalLahir = request.TanggalLahir
	siswa.TempatLahir = request.TempatLahir
	siswa.JenisKelamin = request.JenisKelamin
	siswa.Agama = request.Agama
	siswa.GolonganDarah = request.GolonganDarah
	siswa.NoTelepon = request.NoTelepon

	siswa = service.SiswaRepository.Update(ctx, tx, siswa)

	return helper.ToSiswaResponse(siswa)
}

func (service *SiswaServiceImpl) Delete(ctx context.Context, siswaId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	siswa, err := service.SiswaRepository.FindById(ctx, tx, siswaId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.SiswaRepository.Delete(ctx, tx, siswa)
}

func (service *SiswaServiceImpl) FindById(ctx context.Context, siswaId int) web.SiswaResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	siswa, err := service.SiswaRepository.FindById(ctx, tx, siswaId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToSiswaResponse(siswa)
}

func (service *SiswaServiceImpl) FindAll(ctx context.Context) []web.SiswaResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	siswas := service.SiswaRepository.FindByAll(ctx, tx)

	return helper.ToSiswaResponses(siswas)
}
