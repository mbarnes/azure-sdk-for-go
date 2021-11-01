//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// RestorableMongodbResourcesClient contains the methods for the RestorableMongodbResources group.
// Don't use this type directly, use NewRestorableMongodbResourcesClient() instead.
type RestorableMongodbResourcesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewRestorableMongodbResourcesClient creates a new instance of RestorableMongodbResourcesClient with the specified values.
func NewRestorableMongodbResourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *RestorableMongodbResourcesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &RestorableMongodbResourcesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// List - Return a list of database and collection combo that exist on the account at the given timestamp and location. This helps in scenarios to validate
// what resources exist at given timestamp and location.
// This API requires 'Microsoft.DocumentDB/locations/restorableDatabaseAccounts/…/read' permission.
// If the operation fails it returns the *CloudError error type.
func (client *RestorableMongodbResourcesClient) List(ctx context.Context, location string, instanceID string, options *RestorableMongodbResourcesListOptions) (RestorableMongodbResourcesListResponse, error) {
	req, err := client.listCreateRequest(ctx, location, instanceID, options)
	if err != nil {
		return RestorableMongodbResourcesListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RestorableMongodbResourcesListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RestorableMongodbResourcesListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *RestorableMongodbResourcesClient) listCreateRequest(ctx context.Context, location string, instanceID string, options *RestorableMongodbResourcesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts/{instanceId}/restorableMongodbResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15")
	if options != nil && options.RestoreLocation != nil {
		reqQP.Set("restoreLocation", *options.RestoreLocation)
	}
	if options != nil && options.RestoreTimestampInUTC != nil {
		reqQP.Set("restoreTimestampInUtc", *options.RestoreTimestampInUTC)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RestorableMongodbResourcesClient) listHandleResponse(resp *http.Response) (RestorableMongodbResourcesListResponse, error) {
	result := RestorableMongodbResourcesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorableMongodbResourcesListResult); err != nil {
		return RestorableMongodbResourcesListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *RestorableMongodbResourcesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
