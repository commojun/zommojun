package config

import "os"

func getEnvDefault(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

var (
	ServerPort        = getEnvDefault("SERVER_PORT", "30011")
	SlackToken        = os.Getenv("SLACK_BOT_TOKEN")
	SlackSigninSecret = os.Getenv("SLACK_SIGNING_SECRET")
)
