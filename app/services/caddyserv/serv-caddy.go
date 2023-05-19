package caddyserv

import "cadigo-api/app/ports/caddyport"

type Handler struct {
	repo *caddyport.CaddyRepository
}

func NewHandler(repo caddyport.CaddyRepository) *Handler {
	return &Handler{
		repo: &repo,
	}
}
