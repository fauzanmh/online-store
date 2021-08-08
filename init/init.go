package init

// StartAppInit init all server needs
func StartAppInit() (config *Config) {
	setupSwagger()
	setupLogger()

	config = setupMainConfig()

	return
}
