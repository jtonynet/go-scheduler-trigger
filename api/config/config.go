package config

import (
	"github.com/spf13/viper"
)

type API struct {
	Env string `mapstructure:"ENV"`

	Name       string `mapstructure:"API_NAME"`
	Port       string `mapstructure:"API_PORT"`
	RestHost   string `mapstructure:"API_REST_HOST"`
	TagVersion string `mapstructure:"API_TAG_VERSION"`
}

type PubSub struct {
	Strategy string `mapstructure:"PUBSUB_STRATEGY"`
	Pass     string `mapstructure:"PUBSUB_PASSWORD"`
	Port     string `mapstructure:"PUBSUB_PORT"`
	Host     string `mapstructure:"PUBSUB_HOST"`
	DB       int    `mapstructure:"PUBSUB_DB"`
	Protocol int    `mapstructure:"PUBSUB_PROTOCOL"`
}

type InMemoryDatabase struct {
	Strategy   string
	Pass       string
	Port       string
	Host       string
	DB         int
	Protocol   int
	Expiration int
}

type InMemoryDatabaseConverter interface {
	ToInMemoryDB() (InMemoryDatabase, error)
}

type Trigger struct {
	Strategy   string `mapstructure:"TRIGGER_IN_MEMORY_STRATEGY"`
	Pass       string `mapstructure:"TRIGGER_IN_MEMORY_PASSWORD"`
	Port       string `mapstructure:"TRIGGER_IN_MEMORY_PORT"`
	Host       string `mapstructure:"TRIGGER_IN_MEMORY_HOST"`
	DB         int    `mapstructure:"TRIGGER_IN_MEMORY_DB"`
	Protocol   int    `mapstructure:"TRIGGER_IN_MEMORY_PROTOCOL"`
	Expiration int    `mapstructure:"TRIGGER_IN_MEMORY_EXPIRATION_DEFAULT_IN_MS"`
}

func (t *Trigger) ToInMemoryDB() InMemoryDatabase {
	return InMemoryDatabase{
		Strategy:   t.Strategy,
		Pass:       t.Pass,
		Port:       t.Port,
		Host:       t.Host,
		DB:         t.DB,
		Protocol:   t.Protocol,
		Expiration: t.Expiration,
	}
}

type Cache struct {
	Strategy   string `mapstructure:"CACHE_IN_MEMORY_STRATEGY"`
	Pass       string `mapstructure:"CACHE_IN_MEMORY_PASSWORD"`
	Port       string `mapstructure:"CACHE_IN_MEMORY_PORT"`
	Host       string `mapstructure:"CACHE_IN_MEMORY_HOST"`
	DB         int    `mapstructure:"CACHE_IN_MEMORY_DB"`
	Protocol   int    `mapstructure:"CACHE_IN_MEMORY_PROTOCOL"`
	Expiration int    `mapstructure:"CACHE_IN_MEMORY_EXPIRATION_DEFAULT_IN_MS"`
}

func (c *Cache) ToInMemoryDB() InMemoryDatabase {
	return InMemoryDatabase{
		Strategy:   c.Strategy,
		Pass:       c.Pass,
		Port:       c.Port,
		Host:       c.Host,
		DB:         c.DB,
		Protocol:   c.Protocol,
		Expiration: c.Expiration,
	}
}

type MailNotification struct {
	SMTPHost         string `mapstructure:"MAIL_NOTIFICATION_SMTP_HOST"`
	SMTPPort         int    `mapstructure:"MAIL_NOTIFICATION_SMTP_PORT"`
	SMTPUser         string `mapstructure:"MAIL_NOTIFICATION_SMTP_USER"`
	SMTPPassword     string `mapstructure:"MAIL_NOTIFICATION_SMTP_PASSWORD"`
	EmailFromEmail   string `mapstructure:"MAIL_NOTIFICATION_EMAILS_FROM_EMAIL"`
	EmailFromName    string `mapstructure:"MAIL_NOTIFICATION_EMAILS_FROM_NAME"`
	UseMailTLS       bool   `mapstructure:"MAIL_NOTIFICATION_MAIL_TLS"`
	IsDevelopmentEnv bool   `mapstructure:"MAIL_NOTIFICATION_IS_DEVELOPMENT_ENV"`
}

type Config struct {
	API              API              `mapstructure:",squash"`
	PubSub           PubSub           `mapstructure:",squash"`
	Trigger          Trigger          `mapstructure:",squash"`
	Cache            Cache            `mapstructure:",squash"`
	MailNotification MailNotification `mapstructure:",squash"`
}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	env := viper.GetString("ENV")
	switch env {
	case "test":
		viper.SetConfigName(".env.TEST")
	case "dev", "":
		viper.SetConfigName(".env")
	}

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
