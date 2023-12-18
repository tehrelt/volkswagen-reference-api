package repository

type Store interface {
	Car() CarRepository
}

type CarRepository interface {
	Create()
	Find()
	Delete(id int)
	
	GetAll()
	Get(id int)

}