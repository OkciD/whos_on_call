package db

import "github.com/OkciD/whos_on_call/internal/pkg/duration"

type Config struct {
	Driver          string                        `json:"driver"`
	DSN             string                        `json:"dsn"`
	MaxIdleConns    int                           `json:"maxIdleConns"`
	MaxOpenConns    int                           `json:"maxOpenConns"`
	ConnMaxLifetime duration.MarshallableDuration `json:"connMaxLifetime"`
	ConnMaxIdleTime duration.MarshallableDuration `json:"connMaxIdleTime"`
	PingTimeout     duration.MarshallableDuration `json:"pingTimeout"`

	Stats struct {
		Enabled        bool                          `json:"enabled"`
		TickerDuration duration.MarshallableDuration `json:"tickerDuration"`
	} `json:"stats"`
}
