package main

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jgkawell/galactus/pkg/chassis"
	ct "github.com/jgkawell/galactus/pkg/chassis/context"

	rgpb "github.com/jgkawell/galactus/api/gen/go/core/registry/v1"
)

func main() {

	b := chassis.NewMainBuilder(&chassis.MainBuilderConfig{
		ApplicationName: "eventer",
		EventConfig:     &chassis.EventConfig{},
		OnRun: func(b chassis.MainBuilder) {

			for {
				ctx, span := ct.New(context.Background(), b.GetTelemetry()).Span()
				time.Sleep(5 * time.Second)
				err := b.GetEventManager().CreateAndSendEvent(ctx, &rgpb.Consumer{
					Id: uuid.New().String(),
				})
				if err != nil {
					ctx.WithError(err).Logger().Error("failed to emit event")
					span.End()
					continue
				}
				ctx.Logger().Info("event emitted")
				span.End()
			}
		},
	})
	defer b.Close()
	b.Run()
}
