package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/edobtc/cloudkit/config/lightning"
	"github.com/edobtc/cloudkit/environment"
	"github.com/edobtc/cloudkit/namespace"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	cfg  Config
	once sync.Once
)

const (
	SettingsDirName = ".btc-cloudkit"
)

// Config stores all service configuration that can be injected via
// ENV var or falls back to defaults
type Config struct {
	Verbosity    string `mapstructure:"verbosity"`
	LogFormatter string `mapstructure:"logFormatter"`

	DefaultNamespace string `mapstructure:"defaultNamespace"`

	// DefaultPlatform (optional) is the default platform
	// to target when creating resources, vs defaulting the to
	// default enum (ie: TARGET_AWS_UNSPECIFIED) an override can be
	// configured  here by the deployment operator to target some other
	// target (eg: DigitalOcean)
	DefaultPlatform string `mapstructure:"digitalOceanToken"`

	DigitalOceanToken  string `mapstructure:"digitalOceanToken"`
	CloudflareAPIToken string `mapstructure:"cloudflareApiToken"`

	// Server stuff
	Listen             string `mapstructure:"listen"`
	GRPCListen         string `mapstructure:"grpcListen"`
	GRPCGatewayListen  string `mapstructure:"grpcGatewayListen"`
	GRPCGatewayEnabled bool   `mapstructure:"grpcGatewayEnabled"`

	Node    Node         `mapstructure:"node"`
	Streams StreamConfig `mapstructure:"streams"`

	// Notifications
	Notifications Notifications `mapstructure:"notifications"`

	// Environment
	Environment string `mapstructure:"environment"`

	// ApiKey
	EnableApiKey bool   `mapstructure:"enableApiKey"`
	APIKey       string `mapstructure:"apiKey"`
	SSHPrivKey   string `mapstructure:"sshPrivKey"`
	SSHKeyPath   string `mapstructure:"sshKeyPath"`
	SSHKeyName   string `mapstructure:"sshKeyName"`

	// Redis / MemoryDB stuff for presence and key management
	RedisHost     string `mapstructure:"redisHost"`
	RedisPassword string `mapstructure:"redisPassword"`
	RedisDB       int    `mapstructure:"redisDB"`

	EventPublisherName string `mapstructure:"eventPublisherName"`

	// Publishers
	RabbitMQ RabbitMQConfig `mapstructure:"rmq"`

	AWS AWS `mapstructure:"aws"`

	// Lightning implementations
	LND    lightning.LNDConfig    `mapstructure:"lnd"`
	Eclair lightning.EclairConfig `mapstructure:"eclair"`
}

type RabbitMQConfig struct {
	URL          string `mapstructure:"URL"`
	QueueName    string `mapstructure:"QueueName"`
	ExchangeName string `mapstructure:"ExchangeName"`
	Mandatory    bool   `mapstructure:"Mandatory"`
	Immediate    bool   `mapstructure:"Immediate"`
	DeliveryMode int    `mapstructure:"DeliveryMode"`
	ContentType  string `mapstructure:"ContentType"`

	// Optional
	Durable    bool `mapstructure:"Durable"`
	AutoDelete bool `mapstructure:"AutoDelete"`
	Exclusive  bool `mapstructure:"Exclusive"`
	NoWait     bool `mapstructure:"NoWait"`
}

type Notifications struct {
	WebhookURL                string `mapstructure:"webhookUrl"`
	SNSTopicArn               string `mapstructure:"snsTopicArn"`
	SQSEventsQueue            string `mapstructure:"sqsEventsQueue"`
	AllowWebsocketSubscribers bool   `mapstructure:"allowWebsocketSubscribers"`
}

type AWS struct {
	// Storage Configurations
	DynamoDBTablePrefix string `mapstructure:"dynamodbTablePrefix"`
}

type Node struct {
	Host        string `mapstructure:"host"`
	NodeType    string `mapstructure:"type"`
	RPCUser     string `mapstructure:"rpcUser"`
	RPCPassword string `mapstructure:"rpcPassword"`
}

type StreamConfig struct {
	ZeroMqListenAddr string `mapstructure:"zeroMQListenAddr"`
}

// Read returns an instance of config, initializing
// it only once ever, handling settings defaults and
// binding ENV vars to structural config values
func Read() *Config {
	once.Do(func() {
		path, err := settingsPath()
		if err != nil {
			log.Error(err)
			return
		}

		viper.SetConfigName("config") // name of config file (without extension)
		viper.SetConfigType("toml")   // REQUIRED if the config file does not have the extension in the name
		viper.AddConfigPath(path)     // call multiple times to add many search paths

		viper.SetDefault("verbosity", "INFO")
		viper.SetDefault("logFormatter", "")
		viper.SetDefault("settingsPath", "")

		// Server Settings
		viper.SetDefault("grpcListen", "0.0.0.0:8181")        // grpc server listen
		viper.SetDefault("grpcGatewayEnabled", true)          // enable grpc-gateway server
		viper.SetDefault("grpcGatewayListen", "0.0.0.0:8282") // grpc-gateway server listen
		viper.SetDefault("listen", "0.0.0.0:8081")            // http / websocket server listen

		viper.SetDefault("enableApiKey", false)

		viper.SetDefault("redisHost", "localhost:6379")
		viper.SetDefault("redisPassword", "")
		viper.SetDefault("redisDB", 0)

		viper.SetDefault("environment", environment.Local)
		viper.SetDefault("defaultNamespace", namespace.DefaultNamespace)

		// lightning

		// LND
		viper.SetDefault("lnd.host", "127.0.0.1")
		viper.SetDefault("lnd.port", 10009)

		// Eclair
		viper.SetDefault("eclair.host", "127.0.0.1")
		viper.SetDefault("eclair.port", 8282)

		// Node defaults
		viper.SetDefault("node.host", DefaultNodeHost)
		viper.SetDefault("node.rpcUser", "admin")
		viper.SetDefault("node.rpcPassword", "admin")

		viper.SetDefault("aws.dynamodbTablePrefix", "edobtc_cloudkit_")

		// Event publishers
		viper.SetDefault("eventPublisherName", "demo-stream")

		viper.SetDefault("notifications.AllowWebsocketSubscribers", true)
		viper.SetDefault("notifications.WebhookUrl", "https://127.0.0.1:8081/webhook")
		viper.SetDefault("notifications.TopicArn", "arn:aws:sns:us-east-1:351249512935:demo-topic")
		viper.SetDefault("notifications.EventsQueue", "https://sqs.us-east-1.amazonaws.com/351249512935/demo-queue")

		// Default settings for RabbitMQ
		viper.SetDefault("rabbitMQ.URL", "amqp://guest:guest@localhost:5672/")
		viper.SetDefault("rabbitMQ.QueueName", "defaultQueue")
		viper.SetDefault("rabbitMQ.ExchangeName", "defaultExchange")
		viper.SetDefault("rabbitMQ.Durable", true)
		viper.SetDefault("rabbitMQ.AutoDelete", false)
		viper.SetDefault("rabbitMQ.Exclusive", false)
		viper.SetDefault("rabbitMQ.NoWait", false)
		viper.SetDefault("rabbitMQ.Mandatory", false)
		viper.SetDefault("rabbitMQ.Immediate", false)
		viper.SetDefault("rabbitMQ.ContentType", "text/plain")

		// StreamConfig
		viper.SetDefault("streams.zeroMQListenAddr", "tcp://127.0.0.1:5558")

		// SSH Key Config
		viper.SetDefault("sshKeyName", DefaultSSHKeyName)

		// Storage Configuration Options

		// Environment
		_ = viper.BindEnv("streams.zeroMQListenAddr", "ZEROMQ_LISTEN_ADDR")

		_ = viper.BindEnv("digitalOceanToken", "DIGITALOCEAN_TOKEN")
		_ = viper.BindEnv("linodeToken", "LINODE_TOKEN")
		_ = viper.BindEnv("cloudflareApiToken", "CLOUDFLARE_API_TOKEN")

		_ = viper.BindEnv("environment", "ENVIRONMENT")

		_ = viper.BindEnv("node.host", "BITCOIND_HOST")
		_ = viper.BindEnv("node.rpcUser", "BITCOIND_RPC_USER")
		_ = viper.BindEnv("node.rpcPassword", "BITCOIND_RPC_PASSWORD")

		// Notifications bindings
		_ = viper.BindEnv("notifications.webhookUrl", "WEBHOOK_URL")
		_ = viper.BindEnv("notifications.topicArn", "SNS_TOPIC_ARN")
		_ = viper.BindEnv("notifications.eventsQueue", "SQS_EVENT_QUEUE_URL")

		// Middleware configs
		_ = viper.BindEnv("enableApiKey", "ENABLE_API_KEY")
		_ = viper.BindEnv("apiKey", "API_KEY")
		_ = viper.BindEnv("sshPrivKey", "SSH_PRIV_KEY")
		_ = viper.BindEnv("sshKeyPath", "SSH_KEY_PATH")
		_ = viper.BindEnv("sshKeyName", "SSH_KEY_NAME")

		_ = viper.BindEnv("redisHost", "REDIS_HOST")
		_ = viper.BindEnv("redisPassword", "REDIS_PASSWORD")
		_ = viper.BindEnv("redisDB", "REDIS_DB")

		_ = viper.BindEnv("verbosity", "VERBOSITY")
		_ = viper.BindEnv("logFormatter", "LOG_FORMATTER")
		_ = viper.BindEnv("settingsPath", "SETTINGS_PATH")
		_ = viper.BindEnv("listen", "LISTEN")
		_ = viper.BindEnv("grpcListen", "GRPC_LISTEN")

		_ = viper.BindEnv("defaultNamespace", "DEFAULT_NAMESPACE")

		// publishers
		// rabbit
		_ = viper.BindEnv("rabbitMQ.URL", "RMQ_URL")
		_ = viper.BindEnv("rabbitMQ.QueueName", "RMQ_QUEUE_NAME")
		_ = viper.BindEnv("rabbitMQ.ExchangeName", "RMQ_EXCHANGE_NAME")
		_ = viper.BindEnv("rabbitMQ.Durable", "RMQ_DURABLE")
		_ = viper.BindEnv("rabbitMQ.AutoDelete", "RMQ_AUTO_DELETE")
		_ = viper.BindEnv("rabbitMQ.Exclusive", "RMQ_EXCLUSIVE")
		_ = viper.BindEnv("rabbitMQ.NoWait", "RMQ_NO_WAIT")
		_ = viper.BindEnv("rabbitMQ.Mandatory", "RMQ_MANDATORY")
		_ = viper.BindEnv("rabbitMQ.Immediate", "RMQ_IMMEDIATE")
		_ = viper.BindEnv("rabbitMQ.ContentType", "RMQ_CONTENT_TYPE")

		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				// Config file not found; ignore error if desired
				log.Info("no config found, consider writing defaults")

				exists, err := Exists(path)
				if err != nil {
					log.Error(err)
					return
				}

				if !exists {
					err = os.Mkdir(path, 0700)
					if err != nil {
						log.Error(err)
						return
					}
				}

				err = viper.SafeWriteConfig()
				if err != nil {
					log.Error(err)
					return
				}
			} else {
				// Config file was found but another error was produced
				log.Error(err)
				return
			}
		}

		cfg = Config{}
		err = viper.Unmarshal(&cfg)

		if err != nil {
			log.Error(err)
			return
		}
	})

	return &cfg
}

// Exists test if any settings file exists
func Exists(path string) (bool, error) {
	if _, err := os.Stat(path); err == nil {
		return true, nil
	}

	return false, nil
}

// SavePath is a helper for setting
// up the user's home path
func settingsPath() (string, error) {
	// if path := config.Read().SettingsPath; path != "" {
	// 	return fmt.Sprintf("%s/%s", path, SettingsDirName), nil
	// }

	return defaultSettingsPath()
}

func defaultSettingsPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", home, SettingsDirName), nil
}
