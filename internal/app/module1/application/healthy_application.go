package application

type HealthyApplicationImpl struct {
}

func NewHealthyApplicationImpl() *HealthyApplicationImpl {
	return &HealthyApplicationImpl{}
}
func (h HealthyApplicationImpl) Healthy() bool {
	return true
}
