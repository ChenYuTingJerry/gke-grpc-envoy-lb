package configs

import "flag"

type ComponentConfig struct {
	Echo EchoCompConfig
}

type EchoCompConfig struct {
	Host string
	Port string
}

func componentFlags() {
	flag.StringVar(&C.Component.Echo.Host, "component_echo_host", "", "echo component host")
	flag.StringVar(&C.Component.Echo.Port, "component_echo_port", "", "echo component port")
}
