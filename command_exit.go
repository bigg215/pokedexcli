package main

func commandExit(cfg *config) (string, error) {
	cfg.appConfig.Stop()
	return "", nil
}
