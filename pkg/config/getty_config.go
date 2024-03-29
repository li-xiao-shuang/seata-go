package config

import (
	"time"
)

// GettyConfig
//Config holds supported types by the multiconfig package
type GettyConfig struct {
	ReconnectInterval int `default:"0" yaml:"reconnect_interval" json:"reconnect_interval,omitempty"`
	// getty_session pool
	ConnectionNum int `default:"16" yaml:"connection_number" json:"connection_number,omitempty"`

	// heartbeat
	HeartbeatPeriod time.Duration `default:"15s" yaml:"heartbeat_period" json:"heartbeat_period,omitempty"`

	// getty_session tcp parameters
	GettySessionParam GettySessionParam `required:"true" yaml:"getty_session_param" json:"getty_session_param,omitempty"`
}

// GetDefaultGettyConfig ...
func GetDefaultGettyConfig() GettyConfig {
	return GettyConfig{
		ReconnectInterval: 0,
		ConnectionNum:     2,
		HeartbeatPeriod:   10 * time.Second,
		GettySessionParam: GettySessionParam{
			CompressEncoding: false,
			TCPNoDelay:       true,
			TCPKeepAlive:     true,
			KeepAlivePeriod:  180 * time.Second,
			TCPRBufSize:      262144,
			TCPWBufSize:      65536,
			TCPReadTimeout:   time.Second,
			TCPWriteTimeout:  5 * time.Second,
			WaitTimeout:      time.Second,
			MaxMsgLen:        4096,
			SessionName:      "rpc_client",
		},
	}
}

// GettySessionParam getty session param
type GettySessionParam struct {
	CompressEncoding bool          `default:"false" yaml:"compress_encoding" json:"compress_encoding,omitempty"`
	TCPNoDelay       bool          `default:"true" yaml:"tcp_no_delay" json:"tcp_no_delay,omitempty"`
	TCPKeepAlive     bool          `default:"true" yaml:"tcp_keep_alive" json:"tcp_keep_alive,omitempty"`
	KeepAlivePeriod  time.Duration `default:"180s" yaml:"keep_alive_period" json:"keep_alive_period,omitempty"`
	TCPRBufSize      int           `default:"262144" yaml:"tcp_r_buf_size" json:"tcp_r_buf_size,omitempty"`
	TCPWBufSize      int           `default:"65536" yaml:"tcp_w_buf_size" json:"tcp_w_buf_size,omitempty"`
	TCPReadTimeout   time.Duration `default:"1s" yaml:"tcp_read_timeout" json:"tcp_read_timeout,omitempty"`
	TCPWriteTimeout  time.Duration `default:"5s" yaml:"tcp_write_timeout" json:"tcp_write_timeout,omitempty"`
	WaitTimeout      time.Duration `default:"7s" yaml:"wait_timeout" json:"wait_timeout,omitempty"`
	MaxMsgLen        int           `default:"4096" yaml:"max_msg_len" json:"max_msg_len,omitempty"`
	SessionName      string        `default:"rpc" yaml:"session_name" json:"session_name,omitempty"`
}
