package service

import (
	"context"
	"github.com/Arraf18/go-sisko/model/web"
)

type SiswaService interface {
	Create(ctx context.Context, request web.SiswaCreateRequest) web.SiswaResponse
	Update(ctx context.Context, request web.SiswaUpdateRequest) web.SiswaResponse
	Delete(ctx context.Context, siswaId int)
	FindById(ctx context.Context, siswaId int) web.SiswaResponse
	FindAll(ctx context.Context) []web.SiswaResponse
}
