//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

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

// SoftwareUpdateConfigurationRunsClient contains the methods for the SoftwareUpdateConfigurationRuns group.
// Don't use this type directly, use NewSoftwareUpdateConfigurationRunsClient() instead.
type SoftwareUpdateConfigurationRunsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewSoftwareUpdateConfigurationRunsClient creates a new instance of SoftwareUpdateConfigurationRunsClient with the specified values.
func NewSoftwareUpdateConfigurationRunsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SoftwareUpdateConfigurationRunsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &SoftwareUpdateConfigurationRunsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// GetByID - Get a single software update configuration Run by Id.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SoftwareUpdateConfigurationRunsClient) GetByID(ctx context.Context, resourceGroupName string, automationAccountName string, softwareUpdateConfigurationRunID string, options *SoftwareUpdateConfigurationRunsGetByIDOptions) (SoftwareUpdateConfigurationRunsGetByIDResponse, error) {
	req, err := client.getByIDCreateRequest(ctx, resourceGroupName, automationAccountName, softwareUpdateConfigurationRunID, options)
	if err != nil {
		return SoftwareUpdateConfigurationRunsGetByIDResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SoftwareUpdateConfigurationRunsGetByIDResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SoftwareUpdateConfigurationRunsGetByIDResponse{}, client.getByIDHandleError(resp)
	}
	return client.getByIDHandleResponse(resp)
}

// getByIDCreateRequest creates the GetByID request.
func (client *SoftwareUpdateConfigurationRunsClient) getByIDCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, softwareUpdateConfigurationRunID string, options *SoftwareUpdateConfigurationRunsGetByIDOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/softwareUpdateConfigurationRuns/{softwareUpdateConfigurationRunId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	urlPath = strings.ReplaceAll(urlPath, "{softwareUpdateConfigurationRunId}", url.PathEscape(softwareUpdateConfigurationRunID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.ClientRequestID != nil {
		req.Raw().Header.Set("clientRequestId", *options.ClientRequestID)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getByIDHandleResponse handles the GetByID response.
func (client *SoftwareUpdateConfigurationRunsClient) getByIDHandleResponse(resp *http.Response) (SoftwareUpdateConfigurationRunsGetByIDResponse, error) {
	result := SoftwareUpdateConfigurationRunsGetByIDResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SoftwareUpdateConfigurationRun); err != nil {
		return SoftwareUpdateConfigurationRunsGetByIDResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getByIDHandleError handles the GetByID error response.
func (client *SoftwareUpdateConfigurationRunsClient) getByIDHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Return list of software update configuration runs
// If the operation fails it returns the *ErrorResponse error type.
func (client *SoftwareUpdateConfigurationRunsClient) List(ctx context.Context, resourceGroupName string, automationAccountName string, options *SoftwareUpdateConfigurationRunsListOptions) (SoftwareUpdateConfigurationRunsListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, automationAccountName, options)
	if err != nil {
		return SoftwareUpdateConfigurationRunsListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SoftwareUpdateConfigurationRunsListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SoftwareUpdateConfigurationRunsListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *SoftwareUpdateConfigurationRunsClient) listCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *SoftwareUpdateConfigurationRunsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/softwareUpdateConfigurationRuns"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", *options.Skip)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", *options.Top)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.ClientRequestID != nil {
		req.Raw().Header.Set("clientRequestId", *options.ClientRequestID)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SoftwareUpdateConfigurationRunsClient) listHandleResponse(resp *http.Response) (SoftwareUpdateConfigurationRunsListResponse, error) {
	result := SoftwareUpdateConfigurationRunsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SoftwareUpdateConfigurationRunListResult); err != nil {
		return SoftwareUpdateConfigurationRunsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *SoftwareUpdateConfigurationRunsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
