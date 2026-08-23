package logger

type LogFormat string

const (
	LogFormatJSON LogFormat = "json"
	LogFormatText LogFormat = "text"
)

type Config struct {
	Level      string    `json:"level"`
	Format     LogFormat `json:"format"`
	OutputFile string    `json:"outputFile"`
}
