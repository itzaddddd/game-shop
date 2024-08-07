package repository

type (
	MiddlewareRepositoryService interface{}

	middlewareRepository struct{}
)

func NewMiddlewareRepository() MiddlewareRepositoryService {
	return &middlewareRepository{}
}
