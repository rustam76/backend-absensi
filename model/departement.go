package model

type Departement struct {
	ID              uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	DepartementName string `gorm:"type:varchar(255);not null" json:"departement_name"`
	MaxClockInTime  string `gorm:"type:time" json:"max_clock_in_time"`
	MaxClockOutTime string `gorm:"type:time" json:"max_clock_out_time"`
}
