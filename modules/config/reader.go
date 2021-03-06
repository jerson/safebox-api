// Package config ...
package config

import (
	"path/filepath"

	"github.com/jinzhu/configor"
)

// Server ...
type Server struct {
	Port    int `toml:"port" default:"8000"`
	RPCPort int `toml:"rpc_port" default:"50051"`
}

// Cron ...
type Cron struct {
	TimeEmail string `toml:"time_email" default:"16:00"`
}

// Session ...
type Session struct {
	DurationMinutes int `toml:"duration_minutes" default:"5"`
}

// Purchase ...
type Purchase struct {
	PackageID      string `toml:"package_id" default:"dev.jerson.safebox"`
	GooglePlayFile string `toml:"google_play_file" required:"true"`
}

// SendGrid ...
type SendGrid struct {
	APIKey string `toml:"api_key" required:"true"`
	From   string `toml:"from"  default:"safebox@jerson.dev"`
}

//RabbitMQ ...
type RabbitMQ struct {
	Server string `toml:"server" default:"amqp://guest:guest@rabbitmq:5672"`
}

// Database ...
type Database struct {
	Name     string `toml:"name" default:"app"`
	User     string `toml:"user" default:"app"`
	Password string `toml:"password" default:"app"`
	Host     string `toml:"host" default:"mysql"`
	Port     int    `toml:"port" default:"3306"`
}

// Vars ...
var Vars = struct {
	Name     string   `toml:"name" default:"SafeBox"`
	Debug    bool     `toml:"debug" default:"false"`
	Version  string   `toml:"version" default:"latest"`
	Server   Server   `toml:"server"`
	SendGrid SendGrid `toml:"sendgrid"`
	RabbitMQ RabbitMQ `toml:"rabbitmq"`
	Database Database `toml:"database"`
	Purchase Purchase `toml:"purchase"`
	Session  Session  `toml:"session"`
	Cron     Cron     `toml:"cron"`
}{}

// ReadDefault ...
func ReadDefault() error {
	file, err := filepath.Abs("./config.toml")
	if err != nil {
		return err
	}
	return Read(file)
}

// Read ...
func Read(file string) error {
	return configor.New(&configor.Config{ENVPrefix: "APP", Debug: false, Verbose: false}).Load(&Vars, file)
}
