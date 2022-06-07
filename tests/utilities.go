package test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/stretchr/testify/assert"
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

func getCloudFrontFunctionETag(t *testing.T, client *cloudfront.Client, functionName string) string {
	result, err := client.DescribeFunction(context.TODO(), &cloudfront.DescribeFunctionInput{
		Name: &functionName,
	})
	if err != nil {
		t.Fatalf("Failed to describe function: %v", err)
		return ""
	}
	return *result.ETag
}

func testCloudFrontFunction(t *testing.T, functionName string) {
	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		t.Fatalf("Failed to get config: %v", err)
		return
	}

	cloudfrontClient := cloudfront.NewFromConfig(cfg)

	eventObject := []byte(`{
		"version": "1.0",
		"context": {
		    "eventType": "viewer-request"
		},
		"viewer": {
		    "ip": "198.51.100.11"
		},
		"request": {
		    "method": "GET",
		    "uri": "/",
		    "headers": {
			"host": {"value": "example.org"}
		    }
		}
	}`)

	functionETag := getCloudFrontFunctionETag(t, cloudfrontClient, functionName)

	functionInput := &cloudfront.TestFunctionInput{
		Name:        &functionName,
		EventObject: eventObject,
		IfMatch:     &functionETag,
	}
	result, err := cloudfrontClient.TestFunction(context.TODO(), functionInput)

	if err != nil {
		t.Fatalf("Failed to test CloudFront function: %v", err)
	}

	type FunctionOutput struct {
		Version string `json:"version"`
		Context struct {
			EventType string `json:"eventType"`
		} `json:"context"`
		Viewer struct {
			IP string `json:"ip"`
		} `json:"viewer"`
		Request struct {
			Method  string `json:"method"`
			URI     string `json:"uri"`
			Headers struct {
				Host struct {
					Value string `json:"value"`
				} `json:"host"`
			} `json:"headers"`
		} `json:"request"`
		Response struct {
			Status int `json:"status"`
		} `json:"response"`
	}

	functionOutput := &FunctionOutput{}

	functionOutputString := *result.TestResult.FunctionOutput
	json.Unmarshal([]byte(functionOutputString), functionOutput)

	expectedURI := "/index.html"

	assert.Equal(t, expectedURI, functionOutput.Request.URI)
}
