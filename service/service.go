package service

import (
	"sync/atomic"
)

// StatusCreated, StatusRunned, StatusStopping, StatusStopped are constants
// that defines all status of Service
const (
	StatusCreated = iota
	StatusRunned
	StatusStopping
	StatusStopped
)

// Service interface is base service, with simple API
type Service interface {
	Init(*OpenStorage) error
	Run() error
	Name() string
	Stop()
}

// BaseService is base service used by all services, attached to OpenStorage
type BaseService struct {
	Service

	status uint32
}

// Status returns current status of service
func (s *BaseService) Status() uint32 {
	return atomic.LoadUint32(&s.status)
}

// SetStatus sets status of Service
func (s *BaseService) SetStatus(v uint32) {
	atomic.StoreUint32(&s.status, v)
}

// Stop stops BaseService
func (s *BaseService) Stop() {
	s.SetStatus(StatusStopping)
}

// IsNeedStop returns if BaseService is need to stop
func (s *BaseService) IsNeedStop() bool {
	return s.Status() == StatusStopping
}
