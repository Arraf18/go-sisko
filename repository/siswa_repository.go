package repository

import (
	"context"
	"database/sql"
	"github.com/Arraf18/go-sisko/model/domain"
)

type SiswaRepository interface {
	Save(ctx context.Context, tx *sql.Tx, siswa domain.Siswa) domain.Siswa
	Update(ctx context.Context, tx *sql.Tx, siswa domain.Siswa) domain.Siswa
	Delete(ctx context.Context, tx *sql.Tx, siswa domain.Siswa)
	FindById(ctx context.Context, tx *sql.Tx, siswaId int) (domain.Siswa, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Siswa
}
