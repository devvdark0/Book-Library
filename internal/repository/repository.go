package repository

type BookRepository interface {
	Create()
	List()
	Get()
	Update()
	Delete()
}
