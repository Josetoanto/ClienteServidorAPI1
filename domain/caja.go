package domain

type Caja struct {
	ID     uint    `json:"id" gorm:"primaryKey"`
	Dinero float64 `json:"dinero"`
}
