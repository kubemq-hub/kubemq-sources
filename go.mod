module github.com/kubemq-hub/kubemq-sources

go 1.14

require (
	cloud.google.com/go/pubsub v1.4.0
	github.com/Shopify/sarama v1.27.0
	github.com/aws/aws-sdk-go v1.32.4
	github.com/eclipse/paho.mqtt.golang v1.2.0
	github.com/fortytw2/leaktest v1.3.0
	github.com/ghodss/yaml v1.0.0
	github.com/go-resty/resty/v2 v2.3.0 // indirect
	github.com/go-stomp/stomp v2.0.6+incompatible
	github.com/json-iterator/go v1.1.10
	github.com/kubemq-io/kubemq-go v1.4.0
	github.com/labstack/echo/v4 v4.1.16
	github.com/nats-io/nuid v1.0.1
	github.com/pelletier/go-toml v1.4.0 // indirect
	github.com/prometheus/client_golang v0.9.3
	github.com/smartystreets/assertions v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.0
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.6.1
	go.uber.org/atomic v1.6.0
	go.uber.org/zap v1.10.0
	google.golang.org/genproto v0.0.0-20200624020401-64a14ca9d1ad // indirect
	google.golang.org/grpc v1.30.0 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)