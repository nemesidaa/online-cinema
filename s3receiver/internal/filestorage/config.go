package filestorage

import "time"

type FilestorageConfig struct {
	Creds struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
	Timeout            time.Time `yaml:"timeout"`
	DefaultBucket      string    `yaml:"default_bucket"`
	Address            string    `yaml:"address"`
	ChannelPayloadSize int       `yaml:"channel_payload_size"`
}
