package domain

type ServiceStatus struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	IsHealthy bool   `json:"is_healthy"`
}

type ServiceManager interface {
	Connect() error
	Disconnect() error

	GetServiceState(serviceName string) (string, bool, error)

	StartService(serviceName string) error
	StopService(serviceName string) error
	RestartService(serviceName string) error

	// StartLogWatcher(filePath string, onLog func(string), onError func(error))
	// StopLogWatcher()
}
