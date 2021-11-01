//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ElasticPoolActivitiesClient contains the methods for the ElasticPoolActivities group.
// Don't use this type directly, use NewElasticPoolActivitiesClient() instead.
type ElasticPoolActivitiesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewElasticPoolActivitiesClient creates a new instance of ElasticPoolActivitiesClient with the specified values.
func NewElasticPoolActivitiesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ElasticPoolActivitiesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ElasticPoolActivitiesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// ListByElasticPool - Returns elastic pool activities.
// If the operation fails it returns a generic error.
func (client *ElasticPoolActivitiesClient) ListByElasticPool(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, options *ElasticPoolActivitiesListByElasticPoolOptions) (ElasticPoolActivitiesListByElasticPoolResponse, error) {
	req, err := client.listByElasticPoolCreateRequest(ctx, resourceGroupName, serverName, elasticPoolName, options)
	if err != nil {
		return ElasticPoolActivitiesListByElasticPoolResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ElasticPoolActivitiesListByElasticPoolResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ElasticPoolActivitiesListByElasticPoolResponse{}, client.listByElasticPoolHandleError(resp)
	}
	return client.listByElasticPoolHandleResponse(resp)
}

// listByElasticPoolCreateRequest creates the ListByElasticPool request.
func (client *ElasticPoolActivitiesClient) listByElasticPoolCreateRequest(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, options *ElasticPoolActivitiesListByElasticPoolOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}/elasticPoolActivity"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if elasticPoolName == "" {
		return nil, errors.New("parameter elasticPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{elasticPoolName}", url.PathEscape(elasticPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2014-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByElasticPoolHandleResponse handles the ListByElasticPool response.
func (client *ElasticPoolActivitiesClient) listByElasticPoolHandleResponse(resp *http.Response) (ElasticPoolActivitiesListByElasticPoolResponse, error) {
	result := ElasticPoolActivitiesListByElasticPoolResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ElasticPoolActivityListResult); err != nil {
		return ElasticPoolActivitiesListByElasticPoolResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByElasticPoolHandleError handles the ListByElasticPool error response.
func (client *ElasticPoolActivitiesClient) listByElasticPoolHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
