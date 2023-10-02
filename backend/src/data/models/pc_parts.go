package models

type OpticalDrive struct {
	BaseModel
	Type             string  `gorm:"not null;unique;size:40;type:string"`
	PowerConsumption float64 `gorm:"not null"`
	Manufacturer     string  `gorm:"null;size:255"`
}

type HardDrive struct {
	BaseModel
	Size             float32 `gorm:"type:DECIMAL(5,2)"`
	Rpm              int     `gorm:"not null"`
	PowerConsumption float64 `gorm:"not null"`
}

type Ram struct {
	BaseModel
	Type             string  `gorm:"not null;type:string;size:6"`
	RamSize          int     `gorm:"not null"`
	PowerConsumption float64 `gorm:"not null"`
}

type Graphic struct {
	BaseModel
	Brand            string  `gorm:"unique;not null"`
	GpuName          string  `gorm:"not null;unique"`
	PowerConsumption float64 `gorm:"not null"`
}

type Motherboard struct {
	BaseModel
	FormFactor       string  `gorm:"not null;unique;size:120"`
	PowerConsumption float64 `gorm:"not null"`
}

type Cpu struct {
	BaseModel
	Brand            string  `gorm:"unique;not null"`
	Socket           string  `gorm:"not null"`
	Model            string  `gorm:"not null"`
	PowerConsumption float64 `gorm:"not null"`
}

type Ssd struct {
	BaseModel
	Min              int     `gorm:"not null"`
	Max              int     `gorm:"not null"`
	PowerConsumption float64 `gorm:"not null"`
}
