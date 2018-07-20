package config

import (
	"net/url"
	"time"
)

// ApplicationConfig holds all configurations of the application.
type ApplicationConfig struct {
	Union struct {
		DBConfig         `yaml:"db"`
		GRPCConfig       `yaml:"grpc"`
		MigrationConfig  `yaml:"migration"`
		MonitoringConfig `yaml:"monitoring"`
		ProfilerConfig   `yaml:"profiler"`
		ServerConfig     `yaml:"server"`
	} `yaml:"config"`
}

// DBConfig contains configuration related to database.
type DBConfig struct {
	DSN string `yaml:"dsn"`

	dsn *url.URL
}

// DriverName returns database driver name.
func (cnf *DBConfig) DriverName() string {
	if cnf.dsn == nil {
		cnf.dsn, _ = url.Parse(cnf.DSN)
	}
	return cnf.dsn.Scheme
}

// GRPCConfig contains configuration related to gRPC server.
type GRPCConfig struct {
	Interface string `yaml:"interface"`
}

// MigrationConfig contains configuration related to migrations.
type MigrationConfig struct {
	Table    string `yaml:"table"`
	Schema   string `yaml:"schema"`
	Limit    uint   `yaml:"limit"`
	DryRun   bool   `yaml:"dry-run"`
	WithDemo bool   `yaml:"with-demo"`
}

// MonitoringConfig contains configuration related to monitoring.
type MonitoringConfig struct {
	Enabled   bool   `yaml:"enabled"`
	Interface string `yaml:"interface"`
}

// ProfilerConfig contains configuration related to profiler.
type ProfilerConfig struct {
	Enabled   bool   `yaml:"enabled"`
	Interface string `yaml:"interface"`
}

// ServerConfig contains configuration related to the Forma server.
type ServerConfig struct {
	Interface         string        `yaml:"interface"`
	CPUCount          uint          `yaml:"cpus"`
	ReadTimeout       time.Duration `yaml:"read-timeout"`
	ReadHeaderTimeout time.Duration `yaml:"read-header-timeout"`
	WriteTimeout      time.Duration `yaml:"write-timeout"`
	IdleTimeout       time.Duration `yaml:"idle-timeout"`
	BaseURL           string        `yaml:"base-url"`
	TemplateDir       string        `yaml:"tpl-dir"`
}
