package configs

import "flag"

type ComponentConfig struct {
	Echo EchoCompConfig
}

type EchoCompConfig struct {
	Host string
}

func componentFlags() {
	flag.StringVar(&C.Component.Echo.Host, "component_echo_host", "", "echo component host")
}
