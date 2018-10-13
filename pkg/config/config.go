package config

import (
	"net/url"
	"time"

	"github.com/kamilsk/go-kit/pkg/config"
)

// ApplicationConfig holds all configurations of the application.
type ApplicationConfig struct {
	Union struct {
		CPUs int `json:"cpus" xml:"cpus" yaml:"cpus"`

		DatabaseConfig   `json:"db"         xml:"db"         yaml:"db"`
		GRPCConfig       `json:"grpc"       xml:"grpc"       yaml:"grpc"`
		MigrationConfig  `json:"migration"  xml:"migration"  yaml:"migration"`
		MonitoringConfig `json:"monitoring" xml:"monitoring" yaml:"monitoring"`
		ProfilingConfig  `json:"profiling"  xml:"profiling"  yaml:"profiling"`
		ServerConfig     `json:"server"     xml:"server"     yaml:"server"`
		ServiceConfig    `json:"service"    xml:"service"    yaml:"service"`
	} `json:"config" xml:"config" yaml:"config"`
}

// DatabaseConfig contains configuration related to database.
type DatabaseConfig struct {
	DSN         config.Secret `json:"dsn"      xml:"dsn"      yaml:"dsn"`
	MaxIdle     int           `json:"idle"     xml:"idle"     yaml:"idle"`
	MaxOpen     int           `json:"open"     xml:"open"     yaml:"open"`
	MaxLifetime time.Duration `json:"lifetime" xml:"lifetime" yaml:"lifetime"`

	dsn *url.URL
}

// DriverName returns database driver name.
// Not thread-safe call.
func (cnf *DatabaseConfig) DriverName() string {
	if cnf.dsn == nil {
		cnf.dsn, _ = url.Parse(string(cnf.DSN))
	}
	return cnf.dsn.Scheme
}

// GRPCConfig contains configuration related to gRPC server.
type GRPCConfig struct {
	Interface string        `json:"interface" xml:"interface" yaml:"interface"`
	Timeout   time.Duration `json:"timeout"   xml:"timeout"   yaml:"timeout"`
	Token     config.Secret `json:"token"     xml:"token"     yaml:"token"`
}

// MigrationConfig contains configuration related to migrations.
type MigrationConfig struct {
	Table    string `json:"table"     xml:"table"     yaml:"table"`
	Schema   string `json:"schema"    xml:"schema"    yaml:"schema"`
	Limit    uint   `json:"limit"     xml:"limit"     yaml:"limit"`
	DryRun   bool   `json:"dry-run"   xml:"dry-run"   yaml:"dry-run"`
	WithDemo bool   `json:"with-demo" xml:"with-demo" yaml:"with-demo"`
}

// MonitoringConfig contains configuration related to monitoring.
type MonitoringConfig struct {
	Enabled   bool   `json:"enabled"   xml:"enabled"   yaml:"enabled"`
	Interface string `json:"interface" xml:"interface" yaml:"interface"`
}

// ProfilingConfig contains configuration related to profiler.
type ProfilingConfig struct {
	Enabled   bool   `json:"enabled"   xml:"enabled"   yaml:"enabled"`
	Interface string `json:"interface" xml:"interface" yaml:"interface"`
}

// ServerConfig contains configuration related to the server.
type ServerConfig struct {
	Interface         string        `json:"interface"           xml:"interface"           yaml:"interface"`
	ReadTimeout       time.Duration `json:"read-timeout"        xml:"read-timeout"        yaml:"read-timeout"`
	ReadHeaderTimeout time.Duration `json:"read-header-timeout" xml:"read-header-timeout" yaml:"read-header-timeout"`
	WriteTimeout      time.Duration `json:"write-timeout"       xml:"write-timeout"       yaml:"write-timeout"`
	IdleTimeout       time.Duration `json:"idle-timeout"        xml:"idle-timeout"        yaml:"idle-timeout"`
}

// ServiceConfig contains configuration related to the service.
type ServiceConfig struct {
	BaseURL     string `json:"base-url"            xml:"base-url"            yaml:"base-url"`
	TemplateDir string `json:"tpl-dir"             xml:"tpl-dir"             yaml:"tpl-dir"`
}
