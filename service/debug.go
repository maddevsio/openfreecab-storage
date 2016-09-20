package service

import (
	"net/http"

	_ "net/http/pprof"

	"github.com/gen1us2k/log"
)

type DebugService struct {
	BaseService

	os     *OpenStorage
	logger log.Logger
}

func (h *DebugService) Name() string {
	return "debug_api"
}

func (h *DebugService) Init(os *OpenStorage) error {
	h.os = os
	h.logger = log.NewLogger(h.Name())
	return nil
}

func (h *DebugService) Run() error {
	http.ListenAndServe("localhost:6060", nil)
	return nil
}
