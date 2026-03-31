package logger

type LogFormat string

const (
	LogFormatJson LogFormat = "json"
	LogFormatText LogFormat = "text"
)

type Config struct {
	Level  string    `json:"level"`
	Format LogFormat `json:"format"`
}
