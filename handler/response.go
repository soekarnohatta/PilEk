package handler

import (
	"PilEk/model"
	"PilEk/utils"
)

type userLoginResponse struct {
	User struct {
		Username   string `json:"username"`
		Nama       string `json:"nama"`
		IdKelas    int    `json:"idkelas"`
		IdKandidat int    `json:"idkandidat"`
		Status     string `json:"status"`
		Aktif      bool   `json:"aktif"`
		Token      string `json:"accessToken"`
		IsAdmin    bool   `json:"isAdmin"`
	} `json:"userlogin"`
}

func newUserLoginResponse(u *model.User) *userLoginResponse {
	r := new(userLoginResponse)
	r.User.Username = u.Username
	r.User.Nama = u.Nama
	r.User.IdKelas = u.IdKelas
	r.User.IdKandidat = u.IdKandidat
	if u.Status == model.SudahMemilih {
		r.User.Status = "Sudah Memilih"
	} else {
		r.User.Status = "Belum Memilih"
	}
	r.User.Aktif = u.Aktif
	r.User.IsAdmin = u.IsAdmin

	if r.User.Aktif {
		r.User.Token = utils.GenerateJWT(u.Username)
	}

	return r
}

func newUserCurrentResponse(u *model.User) *userLoginResponse {
	r := new(userLoginResponse)
	r.User.Username = u.Username
	r.User.Nama = u.Nama
	r.User.IdKelas = u.IdKelas
	r.User.IdKandidat = u.IdKandidat
	if u.Status == model.SudahMemilih {
		r.User.Status = "Sudah Memilih"
	} else {
		r.User.Status = "Belum Memilih"
	}
	r.User.Aktif = u.Aktif
	r.User.IsAdmin = u.IsAdmin
	r.User.Token = "Hidden"
	return r
}

type getKandidatResponse struct {
	Id          uint   `json:"idkandidat"`
	Nama        string `json:"nama"`
	NomorUrut   int    `json:"nomorurut"`
	JumlahSuara int    `json:"jumlahsuara"`
	Visi        string `json:"visi"`
	Misi        string `json:"misi"`
	Status      bool   `json:"status"`
	Image       string `json:"image"`
}

type singleKandidatResponse struct {
	SingleKandidat getKandidatResponse `json:"kandidat"`
}

type multipleKandidatResponse struct {
	MultipleKandidat []getKandidatResponse `json:"kandidat"`
}

func newGetKandidatResponse(kandidat *model.Kandidat) *singleKandidatResponse {
	newResp := new(getKandidatResponse)
	newResp.Id = kandidat.ID
	newResp.JumlahSuara = kandidat.JumlahSuara
	newResp.Nama = kandidat.Nama
	newResp.Visi = kandidat.Visi
	newResp.Misi = kandidat.Misi
	newResp.Status = kandidat.Status
	newResp.NomorUrut = kandidat.NomorUrut
	newResp.Image = kandidat.Image
	return &singleKandidatResponse{*newResp}
}

func newGetKandidatListResponse(kandidat []model.Kandidat) *multipleKandidatResponse {
	r := new(multipleKandidatResponse)
	cr := getKandidatResponse{}
	r.MultipleKandidat = make([]getKandidatResponse, 0)
	for _, val := range kandidat {
		cr.Id = val.ID
		cr.JumlahSuara = val.JumlahSuara
		cr.Nama = val.Nama
		cr.Visi = val.Visi
		cr.Misi = val.Misi
		cr.Status = val.Status
		cr.NomorUrut = val.NomorUrut
		cr.Image = val.Image
		r.MultipleKandidat = append(r.MultipleKandidat, cr)
	}

	return r
}

type getKelasResponse struct {
	Id     uint   `json:"idkelas"`
	Nama   string `json:"namakelas"`
	Jumlah int    `json:"jumlah"`
}

type singleKelasResponse struct {
	SingleKelas getKelasResponse `json:"kelas"`
}

type multipleKelasResponse struct {
	MultipleKelas []getKelasResponse `json:"kelas"`
}

func newGetKelasResponse(kelas *model.Kelas) *singleKelasResponse {
	newResp := new(getKelasResponse)
	newResp.Id = kelas.ID
	newResp.Jumlah = kelas.Jumlah
	newResp.Nama = kelas.Kelas
	return &singleKelasResponse{*newResp}
}

func newGetKelasListResponse(kelas []model.Kelas) *multipleKelasResponse {
	r := new(multipleKelasResponse)
	cr := getKelasResponse{}
	r.MultipleKelas = make([]getKelasResponse, 0)
	for _, val := range kelas {
		cr.Id = val.ID
		cr.Jumlah = val.Jumlah
		cr.Nama = val.Kelas
		r.MultipleKelas = append(r.MultipleKelas, cr)
	}

	return r
}
