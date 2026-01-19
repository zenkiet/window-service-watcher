package domain

type ServiceStatus struct {
	Name      string `json:"name"`
	Status    string `json:"status"`
	IsHealthy bool   `json:"is_healthy"`
	// CPU       float64 `json:"cpu_usage"`
}

type ServiceManager interface {
	Connect() error
	Disconnect() error

	CheckStatus() (ServiceStatus, error)
	StartLogWatcher(filePath string, onLog func(string), onError func(error))
	StopLogWatcher()
}
