package service

import (
	"fmt"
	"sync"

	"github.com/gen1us2k/log"
	"github.com/maddevsio/openfreecab-storage/conf"
)

// OpenStorage is main struct of daemon
// it stores all services that used by
type OpenStorage struct {
	config *conf.StorageConfig

	services  map[string]Service
	waitGroup sync.WaitGroup

	logger log.Logger
}

// NewOpenStorage creates and returns new OpenStorageInstance
func NewOpenStorage(config *conf.StorageConfig) *OpenStorage {
	os := new(OpenStorage)
	os.config = config
	os.logger = log.NewLogger("open_storage")
	os.services = make(map[string]Service)
	os.AddService(&RtreeService{})
	os.AddService(&HTTPService{})
	os.AddService(&DebugService{})
	return os
}

// Start starts all services in separate goroutine
func (os *OpenStorage) Start() error {
	os.logger.Info("Starting storage")
	for _, service := range os.services {
		os.logger.Infof("Initializing: %s\n", service.Name())
		if err := service.Init(os); err != nil {
			return fmt.Errorf("initialization of %q finished with error: %v", service.Name(), err)
		}
		os.waitGroup.Add(1)

		go func(srv Service) {
			defer os.waitGroup.Done()
			os.logger.Infof("running %q service\n", srv.Name())
			if err := srv.Run(); err != nil {
				os.logger.Errorf("error on run %q service, %v", srv.Name(), err)
			}
		}(service)
	}
	return nil
}

// AddService adds service into OpenStorage.services map
func (os *OpenStorage) AddService(srv Service) {
	os.services[srv.Name()] = srv

}

// Config returns current instance of StorageConfig
func (os *OpenStorage) Config() conf.StorageConfig {
	return *os.config
}

// Stop stops all services running
func (os *OpenStorage) Stop() {
	os.logger.Info("Worker is stopping...")
	for _, service := range os.services {
		service.Stop()
	}
}

// WaitStop blocks main thread and waits when all goroutines will be stopped
func (os *OpenStorage) WaitStop() {
	os.waitGroup.Wait()
}

func (os *OpenStorage) RtreeService() *RtreeService {
	service, ok := os.services["rtree_service"]
	if !ok {
		os.logger.Error("rtree_service not found")
	}
	return service.(*RtreeService)
}
