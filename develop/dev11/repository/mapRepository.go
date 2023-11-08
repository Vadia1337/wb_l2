package repository

type MapRepository struct {
}

func NewMapRepository() Repository {
	return &MapRepository{}
}
