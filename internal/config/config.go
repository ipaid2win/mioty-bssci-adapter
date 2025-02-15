package config

import (
	"time"
)

// C holds the global configuration.
var C Config

// Config defines the configuration structure.
type Config struct {
	General struct {
		LogJSON     bool `mapstructure:"log_json"`
		LogLevel    int  `mapstructure:"log_level"`
		LogToSyslog bool `mapstructure:"log_to_syslog"`
	} `mapstructure:"general"`

	Backend struct {
		Type string `mapstructure:"type"`

		BssciV1 struct {
			Bind          string        `mapstructure:"bind"`
			TLSCert       string        `mapstructure:"tls_cert"`
			TLSKey        string        `mapstructure:"tls_key"`
			CACert        string        `mapstructure:"ca_cert"`
			StatsInterval time.Duration `mapstructure:"stats_interval"`
			PingInterval  time.Duration `mapstructure:"ping_interval"`
			ReadTimeout   time.Duration `mapstructure:"read_timeout"`
			WriteTimeout  time.Duration `mapstructure:"write_timeout"`
		} `mapstructure:"bssci_v1"`
	} `mapstructure:"backend"`

	Integration struct {
		Marshaler string `mapstructure:"marshaler"`
		MQTTV3 struct {
			EventTopicTemplate      string        `mapstructure:"event_topic_template"`
			CommandTopicTemplate    string        `mapstructure:"command_topic_template"`
			StateTopicTemplate      string        `mapstructure:"state_topic_template"`
			StateRetained           bool          `mapstructure:"state_retained"`
			KeepAlive               time.Duration `mapstructure:"keep_alive"`
			MaxReconnectInterval    time.Duration `mapstructure:"max_reconnect_interval"`
			TerminateOnConnectError bool          `mapstructure:"terminate_on_connect_error"`
			MaxTokenWait            time.Duration `mapstructure:"max_token_wait"`

			Auth struct {
				Type string `mapstructure:"type"`
				Server       string   `mapstructure:"server"`
				Servers      []string `mapstructure:"servers"`
				Username     string   `mapstructure:"username"`
				Password     string   `mapstrucure:"password"`
				CACert       string   `mapstructure:"ca_cert"`
				TLSCert      string   `mapstructure:"tls_cert"`
				TLSKey       string   `mapstructure:"tls_key"`
				QOS          uint8    `mapstructure:"qos"`
				CleanSession bool     `mapstructure:"clean_session"`
				ClientID     string   `mapstructure:"client_id"`
			} `mapstructure:"auth"`
		} `mapstructure:"mqtt_v3"`
	} `mapstructure:"integration"`

	Metrics struct {
		Prometheus struct {
			EndpointEnabled bool   `mapstructure:"endpoint_enabled"`
			Bind            string `mapstructure:"bind"`
		} `mapstructure:"prometheus"`
	} `mapstructure:"metrics"`
}
