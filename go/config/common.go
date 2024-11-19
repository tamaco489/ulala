package config

func IsLocal() bool {
	return EnvConfig.Environment == "local"
}

func IsDevelopment() bool {
	return EnvConfig.Environment == "development"
}

func IsStaging() bool {
	return EnvConfig.Environment == "staging"
}

func IsProduction() bool {
	return EnvConfig.Environment == "production"
}
