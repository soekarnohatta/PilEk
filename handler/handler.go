package handler

import (
	"PilEk/interfaces"
	"PilEk/store"
)

type Handler struct {
	us interfaces.IUser
	ks interfaces.IKandidat
	fs interfaces.IKelas
}

func NewHandler(us store.UserStore, ks store.KandidatStore, fs store.KelasStore) *Handler {
	return &Handler{
		us: &us,
		ks: &ks,
		fs: &fs,
	}
}
