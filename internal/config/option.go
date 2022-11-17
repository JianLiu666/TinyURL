package config

type ServerOpts struct {
	Name                string `mapstructure:"name" yaml:"name"`
	Domain              string `mapstructure:"domain" yaml:"domain"`
	Port                string `mapstructure:"port" yaml:"port"`
	TinyUrlCacheExpired int    `mapstructure:"tinyurl_cache_expired" yaml:"tinyurl_cache_expired"`
	TinyUrlRetry        int    `mapstructure:"tinyurl_retry" yaml:"tinyurl_retry"`
}

type MysqlOpts struct {
	Address         string `mapstructure:"address" yaml:"address"`
	UserName        string `mapstructure:"username" yaml:"username"`
	Password        string `mapstructure:"password" yaml:"password"`
	DBName          string `mapstructure:"dbname" yaml:"dbname"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns" yaml:"max_idle_conns"`
	MaxOpenConns    int    `mapstructure:"max_open_conns" yaml:"max_open_conns"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime" yaml:"conn_max_lifetime"`
}

type RedisOpts struct {
	Address  string `mapstructure:"address" yaml:"address"`
	Password string `mapstructure:"password" yaml:"password"`
	DB       int    `mapstructure:"db" yaml:"db"`
}

type JaegerOpts struct {
	RPCMetrics bool           `mapstructure:"rpc_metrics" yaml:"rpc_metrics"`
	Sampler    jaegerSampler  `mapstructure:"sampler" yaml:"sampler"`
	Reporter   jaegerReporter `mapstructure:"reporter" yaml:"reporter"`
	Headers    jaegerHeaders  `mapstructure:"headers" yaml:"headers"`
}

type jaegerSampler struct {
	Type  string `mapstructure:"type" yaml:"type"`
	Param int    `mapstructure:"param" yaml:"param"`
}

type jaegerReporter struct {
	LogSpans            bool   `mapstructure:"log_spans" yaml:"log_spans"`
	BufferFlushInterval int    `mapstructure:"buffer_flush_interval" yaml:"buffer_flush_interval"`
	LocalAgentHostPort  string `mapstructure:"local_agent_host_port" yaml:"local_agent_host_port"`
}

type jaegerHeaders struct {
	TraceBaggageHeaderPrefix string `mapstructure:"trace_baggage_header_prefix" yaml:"trace_baggage_header_prefix"`
	TraceContextHeaderName   string `mapstructure:"trace_context_header_name" yaml:"trace_context_header_name"`
}
