package calculations

func New() Handler {
	svc := NewService()

	return NewHandler(svc)
}
