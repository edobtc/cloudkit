package lightning

type EclairConfig struct {
	Scheme string `mapstructure:"scheme"`
	Host   string `mapstructure:"host"`
	Port   int    `mapstructure:"port"`
}
