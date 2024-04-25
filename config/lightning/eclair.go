package lightning

type EclairConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
