package dto

type CreateOpticalDriveRequest struct {
	Type             string  `json:"type"`
	PowerConsumption float64 `json:"powerConsumption"`
	Manufacturer     string  `json:"manufacturer"`
}

type UpdateOpticalDriveRequest struct {
	Type             string  `json:"type,omitempty"`
	PowerConsumption float64 `json:"powerConsumption,omitempty"`
	Manufacturer     string  `json:"manufacturer,omitempty"`
}

type OpticalDriveResponse struct {
	Id               int     `json:"id"`
	Type             string  `json:"type"`
	PowerConsumption float64 `json:"powerConsumption"`
	Manufacturer     string  `json:"manufacturer"`
}
