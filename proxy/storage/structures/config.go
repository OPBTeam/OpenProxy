package structures

type Config struct {
	Listener *ConfigListener
	Logging  *ConfigLogging
}

type ConfigListener struct {
	Address string
	Auth    bool
	Motd    string
	Secret  string
}

type ConfigLogging struct {
	Level string
}
