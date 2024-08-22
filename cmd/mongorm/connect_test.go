package mongorm

import (
	"context"
	"log"
	"testing"

	"github.com/testcontainers/testcontainers-go/modules/mongodb"
)

func TestConnect(t *testing.T) {
	type args struct {
		uri string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Successful Connection", args{uri: "mongodb://localhost:27017"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			mongodbContainer, err := mongodb.Run(ctx, "mongo:latest")
			if err != nil {
				log.Fatalf("failed to start container: %s", err)
			}

			// Clean up the container
			defer func() {
				if err := mongodbContainer.Terminate(ctx); err != nil {
					log.Fatalf("failed to terminate container: %s", err)
				}
			}()

			got, err := Connect(tt.args.uri)
			if (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err = got.Ping(ctx, nil); err != nil {
				t.Fatalf("Failed to Ping Mongodb: %v", err)
			}
		})
	}
}
