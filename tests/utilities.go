package test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func getAWSAccountID(t *testing.T) string {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		t.Fatalf("Failed to get config: %v", err)
		return ""
	}
	svc := sts.NewFromConfig(cfg)

	result, err := svc.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
	if err != nil {
		t.Fatalf("Failed to get AWS Account ID: %v", err)
		return ""
	}
	return *result.Account
}
