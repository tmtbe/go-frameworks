package usercase

type HealthyUsercaseImpl struct {
}

func NewHealthyUsercaseImpl() *HealthyUsercaseImpl {
	return &HealthyUsercaseImpl{}
}
func (h HealthyUsercaseImpl) Healthy() bool {
	return true
}
