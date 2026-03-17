package rideModels

import "time"

type Ride struct {
	ID string `json:"id"`

	// Usuarios
	ClientID string  `json:"client_id"`
	UberID   *string `json:"uber_id,omitempty"`

	// Coordenadas inicio
	StartLat float64 `json:"start_lat"`
	StartLng float64 `json:"start_lng"`

	// Coordenadas fin
	EndLat float64 `json:"end_lat"`
	EndLng float64 `json:"end_lng"`

	// Datos calculados
	DistanceKm float64 `json:"distance_km"`
	Price      float64 `json:"price"`

	// Estado del ride
	Status string `json:"status"` // requested | accepted | started | finished | cancelled

	// Tiempos
	RequestedAt time.Time  `json:"requested_at"`
	AcceptedAt  *time.Time `json:"accepted_at,omitempty"`
	StartedAt   *time.Time `json:"started_at,omitempty"`
	FinishedAt  *time.Time `json:"finished_at,omitempty"`
}
