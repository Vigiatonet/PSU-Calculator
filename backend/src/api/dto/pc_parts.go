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
