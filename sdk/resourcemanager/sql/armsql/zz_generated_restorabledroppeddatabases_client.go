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

// RestorableDroppedDatabasesClient contains the methods for the RestorableDroppedDatabases group.
// Don't use this type directly, use NewRestorableDroppedDatabasesClient() instead.
type RestorableDroppedDatabasesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewRestorableDroppedDatabasesClient creates a new instance of RestorableDroppedDatabasesClient with the specified values.
func NewRestorableDroppedDatabasesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *RestorableDroppedDatabasesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &RestorableDroppedDatabasesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Gets a restorable dropped database.
// If the operation fails it returns a generic error.
func (client *RestorableDroppedDatabasesClient) Get(ctx context.Context, resourceGroupName string, serverName string, restorableDroppedDatabaseID string, options *RestorableDroppedDatabasesGetOptions) (RestorableDroppedDatabasesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, restorableDroppedDatabaseID, options)
	if err != nil {
		return RestorableDroppedDatabasesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RestorableDroppedDatabasesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RestorableDroppedDatabasesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RestorableDroppedDatabasesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, restorableDroppedDatabaseID string, options *RestorableDroppedDatabasesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/restorableDroppedDatabases/{restorableDroppedDatabaseId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if restorableDroppedDatabaseID == "" {
		return nil, errors.New("parameter restorableDroppedDatabaseID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorableDroppedDatabaseId}", url.PathEscape(restorableDroppedDatabaseID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RestorableDroppedDatabasesClient) getHandleResponse(resp *http.Response) (RestorableDroppedDatabasesGetResponse, error) {
	result := RestorableDroppedDatabasesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorableDroppedDatabase); err != nil {
		return RestorableDroppedDatabasesGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *RestorableDroppedDatabasesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByServer - Gets a list of restorable dropped databases.
// If the operation fails it returns a generic error.
func (client *RestorableDroppedDatabasesClient) ListByServer(resourceGroupName string, serverName string, options *RestorableDroppedDatabasesListByServerOptions) *RestorableDroppedDatabasesListByServerPager {
	return &RestorableDroppedDatabasesListByServerPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
		},
		advancer: func(ctx context.Context, resp RestorableDroppedDatabasesListByServerResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RestorableDroppedDatabaseListResult.NextLink)
		},
	}
}

// listByServerCreateRequest creates the ListByServer request.
func (client *RestorableDroppedDatabasesClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *RestorableDroppedDatabasesListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/restorableDroppedDatabases"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *RestorableDroppedDatabasesClient) listByServerHandleResponse(resp *http.Response) (RestorableDroppedDatabasesListByServerResponse, error) {
	result := RestorableDroppedDatabasesListByServerResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorableDroppedDatabaseListResult); err != nil {
		return RestorableDroppedDatabasesListByServerResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByServerHandleError handles the ListByServer error response.
func (client *RestorableDroppedDatabasesClient) listByServerHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
