package handler

import (
	"server2/internal/storages"
	"server2/internal/token"
)

// Handler wrap software controller.
type Handler struct {
	TokenController    *token.Controller
	StoragesController *storages.Controller
}

// NewHandler creates a new Handler.
func NewHandler(
	tokenController *token.Controller,
	storagesController *storages.Controller,
) *Handler {
	return &Handler{
		TokenController:    tokenController,
		StoragesController: storagesController,
	}
}
