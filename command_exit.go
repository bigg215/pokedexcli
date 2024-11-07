package main

func commandExit(cfg *config, args ...string) string {
	cfg.appConfig.Stop()
	return ""
}
