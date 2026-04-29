package entity

// Pasar represents the pasar (market) table in the database.
type Pasar struct {
	ID        string  `gorm:"column:id;type:varchar(50);primaryKey"`
	Nama      string  `gorm:"column:nama;type:varchar(255)"`
	Longitude float64 `gorm:"column:longitude;type:double precision"`
	Latitude  float64 `gorm:"column:latitude;type:double precision"`
	Alamat    string  `gorm:"column:alamat;type:text"`
	IsActive  int     `gorm:"column:is_active;type:smallint;default:1"`
}

func (p *Pasar) TableName() string {
	return "monogo.pasar"
}
