package configs

import "flag"

type AppConfig struct {
	BaseRoute     string
	Port          int
	RunMode       string
	Routes        []Route
	EnableProfile bool
}

type Route struct {
	Method  string
	Path    string
	Version string
}

func appFlags() {
	flag.StringVar(&C.App.BaseRoute, "app_base_route", "api", "app base url")
	flag.IntVar(&C.App.Port, "app_port", 8080, "app expose port")
	flag.StringVar(&C.App.RunMode, "app_run_mode", "info", "app error level")
	flag.BoolVar(&C.App.EnableProfile, "app_enable_profile", false, "app enable profile")
}
