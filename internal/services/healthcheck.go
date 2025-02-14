package services

import "ewallet-ums/internal/interfaces"

type HealthCheck struct {
	HealthCheckRepository interfaces.IHealthCheckRepo
}

func (s *HealthCheck) HealthCheckServices() (string, error) {
	return "service is running", nil
}
