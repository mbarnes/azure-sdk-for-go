//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmonitor

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ActivityLogAlertsClient contains the methods for the ActivityLogAlerts group.
// Don't use this type directly, use NewActivityLogAlertsClient() instead.
type ActivityLogAlertsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewActivityLogAlertsClient creates a new instance of ActivityLogAlertsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewActivityLogAlertsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ActivityLogAlertsClient, error) {
	cl, err := arm.NewClient(moduleName+".ActivityLogAlertsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ActivityLogAlertsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create a new Activity Log Alert rule or update an existing one.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - activityLogAlertName - The name of the Activity Log Alert rule.
//   - activityLogAlertRule - The Activity Log Alert rule to create or use for the update.
//   - options - ActivityLogAlertsClientCreateOrUpdateOptions contains the optional parameters for the ActivityLogAlertsClient.CreateOrUpdate
//     method.
func (client *ActivityLogAlertsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, activityLogAlertName string, activityLogAlertRule ActivityLogAlertResource, options *ActivityLogAlertsClientCreateOrUpdateOptions) (ActivityLogAlertsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, activityLogAlertName, activityLogAlertRule, options)
	if err != nil {
		return ActivityLogAlertsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ActivityLogAlertsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ActivityLogAlertsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ActivityLogAlertsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, activityLogAlertName string, activityLogAlertRule ActivityLogAlertResource, options *ActivityLogAlertsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/activityLogAlerts/{activityLogAlertName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if activityLogAlertName == "" {
		return nil, errors.New("parameter activityLogAlertName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{activityLogAlertName}", url.PathEscape(activityLogAlertName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, activityLogAlertRule)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ActivityLogAlertsClient) createOrUpdateHandleResponse(resp *http.Response) (ActivityLogAlertsClientCreateOrUpdateResponse, error) {
	result := ActivityLogAlertsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ActivityLogAlertResource); err != nil {
		return ActivityLogAlertsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete an Activity Log Alert rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - activityLogAlertName - The name of the Activity Log Alert rule.
//   - options - ActivityLogAlertsClientDeleteOptions contains the optional parameters for the ActivityLogAlertsClient.Delete
//     method.
func (client *ActivityLogAlertsClient) Delete(ctx context.Context, resourceGroupName string, activityLogAlertName string, options *ActivityLogAlertsClientDeleteOptions) (ActivityLogAlertsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, activityLogAlertName, options)
	if err != nil {
		return ActivityLogAlertsClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ActivityLogAlertsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ActivityLogAlertsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ActivityLogAlertsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ActivityLogAlertsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, activityLogAlertName string, options *ActivityLogAlertsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/activityLogAlerts/{activityLogAlertName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if activityLogAlertName == "" {
		return nil, errors.New("parameter activityLogAlertName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{activityLogAlertName}", url.PathEscape(activityLogAlertName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get an Activity Log Alert rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - activityLogAlertName - The name of the Activity Log Alert rule.
//   - options - ActivityLogAlertsClientGetOptions contains the optional parameters for the ActivityLogAlertsClient.Get method.
func (client *ActivityLogAlertsClient) Get(ctx context.Context, resourceGroupName string, activityLogAlertName string, options *ActivityLogAlertsClientGetOptions) (ActivityLogAlertsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, activityLogAlertName, options)
	if err != nil {
		return ActivityLogAlertsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ActivityLogAlertsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ActivityLogAlertsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ActivityLogAlertsClient) getCreateRequest(ctx context.Context, resourceGroupName string, activityLogAlertName string, options *ActivityLogAlertsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/activityLogAlerts/{activityLogAlertName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if activityLogAlertName == "" {
		return nil, errors.New("parameter activityLogAlertName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{activityLogAlertName}", url.PathEscape(activityLogAlertName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ActivityLogAlertsClient) getHandleResponse(resp *http.Response) (ActivityLogAlertsClientGetResponse, error) {
	result := ActivityLogAlertsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ActivityLogAlertResource); err != nil {
		return ActivityLogAlertsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Get a list of all Activity Log Alert rules in a resource group.
//
// Generated from API version 2020-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - ActivityLogAlertsClientListByResourceGroupOptions contains the optional parameters for the ActivityLogAlertsClient.NewListByResourceGroupPager
//     method.
func (client *ActivityLogAlertsClient) NewListByResourceGroupPager(resourceGroupName string, options *ActivityLogAlertsClientListByResourceGroupOptions) *runtime.Pager[ActivityLogAlertsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ActivityLogAlertsClientListByResourceGroupResponse]{
		More: func(page ActivityLogAlertsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ActivityLogAlertsClientListByResourceGroupResponse) (ActivityLogAlertsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ActivityLogAlertsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ActivityLogAlertsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ActivityLogAlertsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ActivityLogAlertsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ActivityLogAlertsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/activityLogAlerts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ActivityLogAlertsClient) listByResourceGroupHandleResponse(resp *http.Response) (ActivityLogAlertsClientListByResourceGroupResponse, error) {
	result := ActivityLogAlertsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertRuleList); err != nil {
		return ActivityLogAlertsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionIDPager - Get a list of all Activity Log Alert rules in a subscription.
//
// Generated from API version 2020-10-01
//   - options - ActivityLogAlertsClientListBySubscriptionIDOptions contains the optional parameters for the ActivityLogAlertsClient.NewListBySubscriptionIDPager
//     method.
func (client *ActivityLogAlertsClient) NewListBySubscriptionIDPager(options *ActivityLogAlertsClientListBySubscriptionIDOptions) *runtime.Pager[ActivityLogAlertsClientListBySubscriptionIDResponse] {
	return runtime.NewPager(runtime.PagingHandler[ActivityLogAlertsClientListBySubscriptionIDResponse]{
		More: func(page ActivityLogAlertsClientListBySubscriptionIDResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ActivityLogAlertsClientListBySubscriptionIDResponse) (ActivityLogAlertsClientListBySubscriptionIDResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionIDCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ActivityLogAlertsClientListBySubscriptionIDResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ActivityLogAlertsClientListBySubscriptionIDResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ActivityLogAlertsClientListBySubscriptionIDResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionIDHandleResponse(resp)
		},
	})
}

// listBySubscriptionIDCreateRequest creates the ListBySubscriptionID request.
func (client *ActivityLogAlertsClient) listBySubscriptionIDCreateRequest(ctx context.Context, options *ActivityLogAlertsClientListBySubscriptionIDOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/activityLogAlerts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionIDHandleResponse handles the ListBySubscriptionID response.
func (client *ActivityLogAlertsClient) listBySubscriptionIDHandleResponse(resp *http.Response) (ActivityLogAlertsClientListBySubscriptionIDResponse, error) {
	result := ActivityLogAlertsClientListBySubscriptionIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertRuleList); err != nil {
		return ActivityLogAlertsClientListBySubscriptionIDResponse{}, err
	}
	return result, nil
}

// Update - Updates 'tags' and 'enabled' fields in an existing Alert rule. This method is used to update the Alert rule tags,
// and to enable or disable the Alert rule. To update other fields use CreateOrUpdate
// operation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - activityLogAlertName - The name of the Activity Log Alert rule.
//   - activityLogAlertRulePatch - Parameters supplied to the operation.
//   - options - ActivityLogAlertsClientUpdateOptions contains the optional parameters for the ActivityLogAlertsClient.Update
//     method.
func (client *ActivityLogAlertsClient) Update(ctx context.Context, resourceGroupName string, activityLogAlertName string, activityLogAlertRulePatch AlertRulePatchObject, options *ActivityLogAlertsClientUpdateOptions) (ActivityLogAlertsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, activityLogAlertName, activityLogAlertRulePatch, options)
	if err != nil {
		return ActivityLogAlertsClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ActivityLogAlertsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ActivityLogAlertsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ActivityLogAlertsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, activityLogAlertName string, activityLogAlertRulePatch AlertRulePatchObject, options *ActivityLogAlertsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/activityLogAlerts/{activityLogAlertName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if activityLogAlertName == "" {
		return nil, errors.New("parameter activityLogAlertName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{activityLogAlertName}", url.PathEscape(activityLogAlertName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, activityLogAlertRulePatch)
}

// updateHandleResponse handles the Update response.
func (client *ActivityLogAlertsClient) updateHandleResponse(resp *http.Response) (ActivityLogAlertsClientUpdateResponse, error) {
	result := ActivityLogAlertsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ActivityLogAlertResource); err != nil {
		return ActivityLogAlertsClientUpdateResponse{}, err
	}
	return result, nil
}
