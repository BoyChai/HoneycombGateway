package config

var Cfg config

type config struct {
	// 全局配置
	Global struct {
		Encryption string `yaml:"encryption"`
	} `yaml:"global"`
	// 内网穿透配置
	Penetrate struct {
		Listen AddrName `yaml:"listen"`
		Nodes  []struct {
			Name     string   `yaml:"name"`
			Protocol Protocol `yaml:"protocol" ant:"network_transport_protocol_common"`
			Addr     AddrName `yaml:"addr"`
		} `yaml:"nodes"`
	} `yaml:"penetrate"`
	// 内网穿透客户端配置
	Client struct {
		Server  AddrName `yaml:"server"`
		Service []struct {
			Name     string   `yaml:"name"`
			Protocol Protocol `yaml:"protocol" ant:"network_transport_protocol_common"`
			Addr     AddrName `yaml:"addr"`
		} `yaml:"service"`
	} `yaml:"client"`
	// 流量转发配置
	Forward []struct {
		Name     string   `yaml:"name"`
		Protocol Protocol `yaml:"protocol" ant:"network_transport_protocol_common"`
		Listen   AddrName `yaml:"listen"`
		Target   AddrName `yaml:"target"`
	} `yaml:"forward"`
	Load []struct {
		Name     string     `yaml:"name"`
		Protocol Protocol   `yaml:"protocol" ant:"network_transport_protocol_common"`
		Servers  []AddrName `yaml:"servers"`
	} `yaml:"load"`
	Address []Addr
}

type AddrName string

type Protocol string

const TCP Protocol = "TCP"
const UDP Protocol = "UDP"

type Addr struct {
	Name string `yaml:"name"`
	Addr string `yaml:"addr"`
	Port int    `yaml:"port" ant:"port"`
}
