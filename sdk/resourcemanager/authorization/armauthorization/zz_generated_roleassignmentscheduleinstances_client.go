//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

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

// RoleAssignmentScheduleInstancesClient contains the methods for the RoleAssignmentScheduleInstances group.
// Don't use this type directly, use NewRoleAssignmentScheduleInstancesClient() instead.
type RoleAssignmentScheduleInstancesClient struct {
	ep string
	pl runtime.Pipeline
}

// NewRoleAssignmentScheduleInstancesClient creates a new instance of RoleAssignmentScheduleInstancesClient with the specified values.
func NewRoleAssignmentScheduleInstancesClient(credential azcore.TokenCredential, options *arm.ClientOptions) *RoleAssignmentScheduleInstancesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &RoleAssignmentScheduleInstancesClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Gets the specified role assignment schedule instance.
// If the operation fails it returns the *CloudError error type.
func (client *RoleAssignmentScheduleInstancesClient) Get(ctx context.Context, scope string, roleAssignmentScheduleInstanceName string, options *RoleAssignmentScheduleInstancesGetOptions) (RoleAssignmentScheduleInstancesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, roleAssignmentScheduleInstanceName, options)
	if err != nil {
		return RoleAssignmentScheduleInstancesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleAssignmentScheduleInstancesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleAssignmentScheduleInstancesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RoleAssignmentScheduleInstancesClient) getCreateRequest(ctx context.Context, scope string, roleAssignmentScheduleInstanceName string, options *RoleAssignmentScheduleInstancesGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignmentScheduleInstances/{roleAssignmentScheduleInstanceName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleAssignmentScheduleInstanceName == "" {
		return nil, errors.New("parameter roleAssignmentScheduleInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentScheduleInstanceName}", url.PathEscape(roleAssignmentScheduleInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RoleAssignmentScheduleInstancesClient) getHandleResponse(resp *http.Response) (RoleAssignmentScheduleInstancesGetResponse, error) {
	result := RoleAssignmentScheduleInstancesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignmentScheduleInstance); err != nil {
		return RoleAssignmentScheduleInstancesGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *RoleAssignmentScheduleInstancesClient) getHandleError(resp *http.Response) error {
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

// ListForScope - Gets role assignment schedule instances of a role assignment schedule.
// If the operation fails it returns the *CloudError error type.
func (client *RoleAssignmentScheduleInstancesClient) ListForScope(scope string, options *RoleAssignmentScheduleInstancesListForScopeOptions) *RoleAssignmentScheduleInstancesListForScopePager {
	return &RoleAssignmentScheduleInstancesListForScopePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listForScopeCreateRequest(ctx, scope, options)
		},
		advancer: func(ctx context.Context, resp RoleAssignmentScheduleInstancesListForScopeResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RoleAssignmentScheduleInstanceListResult.NextLink)
		},
	}
}

// listForScopeCreateRequest creates the ListForScope request.
func (client *RoleAssignmentScheduleInstancesClient) listForScopeCreateRequest(ctx context.Context, scope string, options *RoleAssignmentScheduleInstancesListForScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignmentScheduleInstances"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2020-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listForScopeHandleResponse handles the ListForScope response.
func (client *RoleAssignmentScheduleInstancesClient) listForScopeHandleResponse(resp *http.Response) (RoleAssignmentScheduleInstancesListForScopeResponse, error) {
	result := RoleAssignmentScheduleInstancesListForScopeResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignmentScheduleInstanceListResult); err != nil {
		return RoleAssignmentScheduleInstancesListForScopeResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listForScopeHandleError handles the ListForScope error response.
func (client *RoleAssignmentScheduleInstancesClient) listForScopeHandleError(resp *http.Response) error {
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
