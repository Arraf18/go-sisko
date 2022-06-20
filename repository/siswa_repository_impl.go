package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Arraf18/go-sisko/helper"
	"github.com/Arraf18/go-sisko/model/domain"
)

type SiswaRepositoryImpl struct {
}

func NewSiswaRepository() SiswaRepository {
	return &SiswaRepositoryImpl{}
}

func (c SiswaRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, siswa domain.Siswa) domain.Siswa {
	SQL := "insert into siswa(nama, alamat, tanggal_lahir, tempat_lahir, jenis_kelamin, agama, golongan_darah, no_telepon) values (?,?,?,?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, siswa.Nama, siswa.Alamat, siswa.TanggalLahir, siswa.TempatLahir, siswa.JenisKelamin, siswa.Agama, siswa.GolonganDarah, siswa.NoTelepon)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	siswa.Id = int(id)
	return siswa
}

func (c SiswaRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, siswa domain.Siswa) domain.Siswa {
	SQL := "update siswa set nama = ?, alamat = ?, tanggal_lahir = ?, tempat_lahir = ?, jenis_kelamin = ?, agama = ?, golongan_darah = ?, no_telepon = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, siswa.Nama, siswa.Alamat, siswa.TanggalLahir, siswa.TempatLahir, siswa.JenisKelamin, siswa.Agama, siswa.GolonganDarah, siswa.NoTelepon, siswa.Id)
	helper.PanicIfError(err)

	return siswa
}

func (c SiswaRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, siswa domain.Siswa) {
	SQL := "delete from siswa where id = ?"
	_, err := tx.ExecContext(ctx, SQL, siswa.Id)
	helper.PanicIfError(err)
}

func (c SiswaRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, siswaId int) (domain.Siswa, error) {
	SQL := "select id, nama, alamat, tanggal_lahir, tempat_lahir, jenis_kelamin, agama, golongan_darah, no_telepon from siswa where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, siswaId)
	helper.PanicIfError(err)
	defer rows.Close()

	siswa := domain.Siswa{}
	if rows.Next() {
		err := rows.Scan(&siswa.Id, &siswa.Nama, &siswa.Alamat, &siswa.TanggalLahir, &siswa.TempatLahir, &siswa.JenisKelamin, &siswa.Agama, &siswa.GolonganDarah, &siswa.NoTelepon)
		helper.PanicIfError(err)
		return siswa, nil
	} else {
		return siswa, errors.New("siswa is not found")
	}
}

func (c SiswaRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Siswa {
	SQL := "select id, nama, alamat, tanggal_lahir, tempat_lahir, jenis_kelamin, agama, golongan_darah, no_telepon from siswa"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var siswas []domain.Siswa
	for rows.Next() {
		siswa := domain.Siswa{}
		err := rows.Scan(&siswa.Id, &siswa.Nama, &siswa.Alamat, &siswa.TanggalLahir, &siswa.TempatLahir, &siswa.JenisKelamin, &siswa.Agama, &siswa.GolonganDarah, &siswa.NoTelepon)
		helper.PanicIfError(err)
		siswas = append(siswas, siswa)
	}
	return siswas
}
