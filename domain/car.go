package domain

// Car define el modelo de datos
type Car struct {
	ID     uint    `json:"id" gorm:"primaryKey"`
	Brand  string  `json:"brand"`
	Model  string  `json:"model"`
	Year   int     `json:"year"`
	Precio float64 `json:"precio"` // Nuevo campo
}

// CarRepository define la interfaz para acceder a los datos
type CarRepository interface {
	GetAll() ([]Car, error)
	GetByID(id uint) (Car, error)
	Create(car Car) (Car, error)
	Update(car Car) (Car, error)
	Delete(id uint) error
}
