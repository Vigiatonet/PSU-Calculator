package dto

type CreateOpticalDriveRequest struct {
	Type             string  `json:"type"`
	PowerConsumption float64 `json:"powerConsumption"`
	Manufacturer     string  `json:"manufacturer"`
}

type UpdateOpticalDriveRequest struct {
	Type         string `json:"type,omitempty"`
	Manufacturer string `json:"manufacturer,omitempty"`
}

type OpticalDriveResponse struct {
	Id               int     `json:"id"`
	Type             string  `json:"type"`
	PowerConsumption float64 `json:"powerConsumption"`
	Manufacturer     string  `json:"manufacturer"`
}

type CreateHardDriveRequest struct {
	Size             float32 `json:"size"`
	Rpm              int     `json:"rpm"`
	PowerConsumption float64 `json:"powerConsumption"`
}

type UpdateHardDriveRequest struct {
	Size float32 `json:"size"`
	Rpm  int     `json:"rpm"`
}

type HardDriveResponse struct {
	Id               int     `json:"id"`
	Size             float32 `json:"size"`
	Rpm              int     `json:"rpm"`
	PowerConsumption float64 `json:"powerConsumption"`
}

type CreateRamModelRequest struct {
	Type             string  `json:"type"`
	RamSize          int     `json:"RamSize"`
	PowerConsumption float64 `json:"powerConsumption"`
}

type UpdateRamModelRequest struct {
	Type    string `json:"type"`
	RamSize int    `json:"ramSize"`
}

type RamModelResponse struct {
	Id               int     `json:"id"`
	Type             string  `json:"type"`
	RamSize          int     `json:"ramSize"`
	PowerConsumption float64 `json:"powerConsumption"`
}

type CreateMotherboardRequest struct {
	FormFactor       string  `json:"formFactor"`
	PowerConsumption float64 `json:"powerConsumption"`
}

type UpdateMotherboardRequest struct {
	FormFactor string `json:"formFactor"`
}

type MotherboardResponse struct {
	Id               int     `json:"id"`
	FormFactor       string  `json:"formFactor"`
	PowerConsumption float64 `json:"powerConsumption"`
}

type CreateGraphicRequest struct {
	GpuBrandId       int     `json:"gpuBrandId"`
	GpuName          string  `json:"gpuName"`
	PowerConsumption float64 `json:"powerConsumption"`
}

type UpdateGraphicRequest struct {
	GpuName string `json:"gpuName"`
}

type GraphicResponse struct {
	Id               int     `json:"id"`
	GpuBrandId       int     `json:"gpuBrandId"`
	GpuName          string  `json:"gpuName"`
	PowerConsumption float64 `json:"powerConsumption"`
}

type CreateCpuRequest struct {
	CpuBrandId       int     `json:"cpuBrandId"`
	Socket           string  `json:"socket"`
	Model            string  `json:"model"`
	PowerConsumption float64 `json:"powerConsumption"`
}

type UpdateCpuRequest struct {
	Socket string `json:"socket"`
	Model  string `json:"model"`
}

type CpuResponse struct {
	Id               int     `json:"id"`
	CpuBrandId       int     `json:"cpuBrandId"`
	Socket           string  `json:"socket"`
	Model            string  `json:"model"`
	PowerConsumption float64 `json:"powerConsumption"`
}

type CreateSsdRequest struct {
	SizeRange        string  `json:"sizeRange"`
	PowerConsumption float64 `json:"powerConsumption"`
}

type UpdateSsdRequest struct {
	SizeRange string `json:"sizeRange"`
}

type SsdResponse struct {
	Id               int     `json:"id"`
	SizeRange        string  `json:"sizeRange"`
	PowerConsumption float64 `json:"powerConsumption"`
}

// type CalculatePowerRequest struct {
// 	CpuId int `json:"cpuId" binding:"required"`
// 	// rams
// 	RamId    int `json:"ramId" binding:"required"`
// 	RamCount int `json:"ramCount"`
// 	// Gpu
// 	GraphicId    int `json:"graphicId" binding:"required"`
// 	GraphicCount int `json:"graphicCount"`
// 	// Ssd
// 	SsdId    int `json:"ssdId" binding:"required"`
// 	SsdCount int `json:"ssdCount"`
// 	// hdd
// 	HardDriveId    int `json:"hardDriveId" binding:"required"`
// 	HardDriveCount int `json:"hardDriveCount"`
// 	// od
// 	OpticalDriveId    int `json:"opticalDriveId" binding:"required"`
// 	OpticalDriveCount int `json:"opticalDriveCount"`
// 	// mb
// 	MotherboardId int `json:"motherboardId" binding:"required"`
// }

type CalculatePowerRequest struct {
	CpuPower float64 `json:"cpuId" binding:"required"`

	RamPower float64 `json:"ramPower" binding:"required"`
	RamCount int     `json:"ramCount"`

	GraphicPower float64 `json:"graphicPower" binding:"required"`
	GraphicCount int     `json:"graphicCount"`

	SsdPower float64 `json:"ssdPower" binding:"required"`
	SsdCount int     `json:"ssdCount"`

	HardDrivePower float64 `json:"hardDrivePower" binding:"required"`
	HardDriveCount int     `json:"hardDriveCount"`

	OpticalDrivePower float64 `json:"opticalDrivePower" binding:"required"`
	OpticalDriveCount int     `json:"opticalDriveCount"`

	MotherboardPower float64 `json:"motherboardPower" binding:"required"`
}

type CalculatePowerResponse struct {
	TotalPowerConsumption float64 `json:"totalPowerConsumption"`
}
