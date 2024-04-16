package config

type LoggerConfig struct {
	FilePath string
	Encoding string
	Level    string
	Logger   string
}

func LogConfig() LoggerConfig {
	return LoggerConfig{
		FilePath: "./log/",
		Encoding: "json",
		Level:    "Info",
		Logger:   "zap",
	}
}
