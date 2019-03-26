package config

type LogConfig struct {
	logPath  string
	logLevel string
	format   string
}

func (lc LogConfig) Path() string {
	return lc.logPath
}

func (lc LogConfig) Level() string {
	return lc.logLevel
}

func (lc LogConfig) Format() string {
	return lc.format
}
