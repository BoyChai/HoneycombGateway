package config

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
			Protocol Protocol `yaml:"protocol"`
			Addr     AddrName `yaml:"addr"`
		} `yaml:"nodes"`
	} `yaml:"penetrate"`
	// 内网穿透客户端配置
	Client struct {
		Server  AddrName `yaml:"server"`
		Service []struct {
			Name     string   `yaml:"name"`
			Protocol Protocol `yaml:"protocol"`
			Addr     Addr     `yaml:"addr"`
		} `yaml:"service"`
	} `yaml:"client"`
	// 流量转发配置
	Forward struct {
		Name     string   `yaml:"name"`
		Protocol Protocol `yaml:"protocol"`
		Listen   Addr     `yaml:"listen"`
		Target   Addr     `yaml:"target"`
	} `yaml:"forward"`
	Load struct {
		Name     string   `yaml:"name"`
		Protocol Protocol `yaml:"protocol"`
		Services []Addr   `yaml:"services"`
	} `yaml:"load"`
	Address []Addr
}

type AddrName string

type Protocol string

var TCP Protocol = "TCP"
var UDP Protocol = "UDP"

type Addr struct {
	Name string `yaml:"name"`
	Addr string `yaml:"addr"`
	Port int    `yaml:"port"`
}
