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
	GpuName          string   `gorm:"not null;unique"`
	PowerConsumption float64  `gorm:"not null"`
	GpuBrand         GpuBrand `gorm:"foreignKey:GpuBrandId;constraint:OnDelete:NO ACTION"`
	GpuBrandId       int
}
type GpuBrand struct {
	BaseModel
	Name string `gorm:"not null;unique;type:string;size:40"`
}

type Motherboard struct {
	BaseModel
	FormFactor       string  `gorm:"not null;unique;size:120"`
	PowerConsumption float64 `gorm:"not null"`
}

type Cpu struct {
	BaseModel
	CpuBrand         CpuBrand `gorm:"foreignKey:CpuBrandId;constraint:OnDelete:NO ACTION"`
	CpuBrandId       int
	Socket           string  `gorm:"not null"`
	Model            string  `gorm:"not null"`
	PowerConsumption float64 `gorm:"not null"`
}

type CpuBrand struct {
	BaseModel
	Name string `gorm:"not null;unique;type:string;size:40"`
}

type Ssd struct {
	BaseModel
	SizeRange        string  `gorm:"not null;unique"`
	PowerConsumption float64 `gorm:"not null"`
}

type PcModel struct {
	Ssd      Ssd `gorm:"foreignKey:SsdId;constraint:OnDelete:NO ACTION"`
	SsdId    int `gorm:"not null"`
	SsdCount int `gorm:"not null"`

	Cpu   Cpu `gorm:"foreignKey:CpuId;constraint:OnDelete:NO ACTION"`
	CpuId int `gorm:"not null"`

	Graphic      Graphic `gorm:"foreignKey:GraphicId;constraint:OnDelete:NO ACTION"`
	GraphicId    int     `gorm:"not null"`
	GraphicCount int     `gorm:"not null"`

	Motherboard   Motherboard `gorm:"foreignKey:MotherboardId;constraint:OnDelete:NO ACTION"`
	MotherboardId int         `gorm:"not null"`

	Ram      Ram `gorm:"foreignKey:RamId;constraint:OnDelete:NO ACTION"`
	RamId    int `gorm:"not null"`
	RamCount int `gorm:"not null"`

	HardDrive      HardDrive `gorm:"foreignKey:HardDriveId;constraint:OnDelete:NO ACTION"`
	HardDriveId    int       `gorm:"not null"`
	HardDriveCount int       `gorm:"not null"`

	OpticalDrive      OpticalDrive `gorm:"foreignKey:OpticalDriveId;constraint:OnDelete:NO ACTION"`
	OpticalDriveId    int          `gorm:"not null"`
	OpticalDriveCount int          `gorm:"not null"`
}
