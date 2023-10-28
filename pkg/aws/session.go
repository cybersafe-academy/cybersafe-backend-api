package aws

import (
	"context"
	"cybersafe-backend-api/internal/api/components"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func GetAWSConfig(c *components.Components) aws.Config {
	sdkConfig, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(c.Settings.String("aws.region")),
	)
	if err != nil {
		panic(err)
	}

	return sdkConfig
}
