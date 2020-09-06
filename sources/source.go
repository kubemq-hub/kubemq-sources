package sources

import (
	"context"
	"fmt"
	"github.com/kubemq-hub/kubemq-sources/config"
	"github.com/kubemq-hub/kubemq-sources/middleware"
	"github.com/kubemq-hub/kubemq-sources/sources/aws/sqs"
	"github.com/kubemq-hub/kubemq-sources/sources/messaging/rabbitmq"
)

type Source interface {
	Init(ctx context.Context, cfg config.Spec) error
	Start(ctx context.Context, target middleware.Middleware) error
	Stop() error
	Name() string
}

func Init(ctx context.Context, cfg config.Spec) (Source, error) {

	switch cfg.Kind {

	case "source.messaging.rabbitmq":
		source := rabbitmq.New()

		if err := source.Init(ctx, cfg); err != nil {
			return nil, err
		}
		return source, nil
	case "source.aws.sqs":
		source := sqs.New()

		if err := source.Init(ctx, cfg); err != nil {
			return nil, err
		}
		return source, nil
	default:
		return nil, fmt.Errorf("invalid kind %s for source %s", cfg.Kind, cfg.Name)
	}

}