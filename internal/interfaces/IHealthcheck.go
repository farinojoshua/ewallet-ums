package interfaces

type IHealthCheckService interface {
	HealthCheckServices() (string, error)
}

type IHealthCheckRepo interface {
}
