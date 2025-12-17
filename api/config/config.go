package config

import (
	"time"

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
	Strategy        string
	Pass            string
	Port            string
	Host            string
	DB              int
	Protocol        int
	Expiration      time.Duration
	MaxRetries      int
	MinRetryBackoff time.Duration
	MaxRetryBackoff time.Duration
	DialTimeout     time.Duration
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	PoolSize        int
	MinIdleConns    int
}

type InMemoryDatabaseConverter interface {
	ToInMemoryDB() (InMemoryDatabase, error)
}

type Trigger struct {
	Strategy        string        `mapstructure:"TRIGGER_IN_MEMORY_STRATEGY"`
	Pass            string        `mapstructure:"TRIGGER_IN_MEMORY_PASSWORD"`
	Port            string        `mapstructure:"TRIGGER_IN_MEMORY_PORT"`
	Host            string        `mapstructure:"TRIGGER_IN_MEMORY_HOST"`
	DB              int           `mapstructure:"TRIGGER_IN_MEMORY_DB"`
	Protocol        int           `mapstructure:"TRIGGER_IN_MEMORY_PROTOCOL"`
	Expiration      time.Duration `mapstructure:"TRIGGER_IN_MEMORY_EXPIRATION_DEFAULT_IN_MS"`
	MaxRetries      int           `mapstructure:"TRIGGER_MAX_RETRIES"`
	MinRetryBackoff time.Duration `mapstructure:"TRIGGER_MIN_RETRY_BACKOFF_IN_MS"`
	MaxRetryBackoff time.Duration `mapstructure:"TRIGGER_MAX_RETRY_BACKOFF_IN_S"`
	DialTimeout     time.Duration `mapstructure:"TRIGGER_DIAL_TIMEOUT_IN_S"`
	ReadTimeout     time.Duration `mapstructure:"TRIGGER_READ_TIMEOUT_IN_S"`
	WriteTimeout    time.Duration `mapstructure:"TRIGGER_WRITE_TIMEOUT_IN_S"`
	PoolSize        int           `mapstructure:"TRIGGER_POOL_SIZE"`
	MinIdleConns    int           `mapstructure:"TRIGGER_MIN_IDDLE_CONNS"`
}

func (t *Trigger) ToInMemoryDB() InMemoryDatabase {
	return InMemoryDatabase{
		Strategy:        t.Strategy,
		Pass:            t.Pass,
		Port:            t.Port,
		Host:            t.Host,
		DB:              t.DB,
		Protocol:        t.Protocol,
		Expiration:      t.Expiration,
		MaxRetries:      t.MaxRetries,
		MinRetryBackoff: t.MinRetryBackoff,
		MaxRetryBackoff: t.MaxRetryBackoff,
		DialTimeout:     t.DialTimeout,
		ReadTimeout:     t.ReadTimeout,
		WriteTimeout:    t.WriteTimeout,
		PoolSize:        t.PoolSize,
		MinIdleConns:    t.MinIdleConns,
	}
}

type ShadowKey struct {
	Strategy        string        `mapstructure:"SHADOWKEY_IN_MEMORY_STRATEGY"`
	Pass            string        `mapstructure:"SHADOWKEY_IN_MEMORY_PASSWORD"`
	Port            string        `mapstructure:"SHADOWKEY_IN_MEMORY_PORT"`
	Host            string        `mapstructure:"SHADOWKEY_IN_MEMORY_HOST"`
	DB              int           `mapstructure:"SHADOWKEY_IN_MEMORY_DB"`
	Protocol        int           `mapstructure:"SHADOWKEY_IN_MEMORY_PROTOCOL"`
	Expiration      time.Duration `mapstructure:"SHADOWKEY_IN_MEMORY_EXPIRATION_DEFAULT_IN_MS"`
	MaxRetries      int           `mapstructure:"SHADOWKEY_MAX_RETRIES"`
	MinRetryBackoff time.Duration `mapstructure:"SHADOWKEY_MIN_RETRY_BACKOFF_IN_MS"`
	MaxRetryBackoff time.Duration `mapstructure:"SHADOWKEY_MAX_RETRY_BACKOFF_IN_S"`
	DialTimeout     time.Duration `mapstructure:"SHADOWKEY_DIAL_TIMEOUT_IN_S"`
	ReadTimeout     time.Duration `mapstructure:"SHADOWKEY_READ_TIMEOUT_IN_S"`
	WriteTimeout    time.Duration `mapstructure:"SHADOWKEY_WRITE_TIMEOUT_IN_S"`
	PoolSize        int           `mapstructure:"SHADOWKEY_POOL_SIZE"`
	MinIdleConns    int           `mapstructure:"SHADOWKEY_MIN_IDDLE_CONNS"`
}

func (s *ShadowKey) ToInMemoryDB() InMemoryDatabase {
	return InMemoryDatabase{
		Strategy:        s.Strategy,
		Pass:            s.Pass,
		Port:            s.Port,
		Host:            s.Host,
		DB:              s.DB,
		Protocol:        s.Protocol,
		Expiration:      s.Expiration,
		MaxRetries:      s.MaxRetries,
		MinRetryBackoff: s.MinRetryBackoff,
		MaxRetryBackoff: s.MaxRetryBackoff,
		DialTimeout:     s.DialTimeout,
		ReadTimeout:     s.ReadTimeout,
		WriteTimeout:    s.WriteTimeout,
		PoolSize:        s.PoolSize,
		MinIdleConns:    s.MinIdleConns,
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
	TriggerDB        Trigger          `mapstructure:",squash"`
	ShadowKeyDB      ShadowKey        `mapstructure:",squash"`
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
