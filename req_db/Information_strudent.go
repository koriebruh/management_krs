package req_db

import "koriebruh/try/domain"

type InformationStudent struct {
	MahasiswaDinus     domain.MahasiswaDinus
	HerregistMahasiswa domain.HerregistMahasiswa
	TagihanMahasiswa   []domain.TagihanMhs
}
