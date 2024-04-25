package lightning

type LNDConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Macaroon string `mapstructure:"macaroon"`
}
