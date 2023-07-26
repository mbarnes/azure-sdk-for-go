//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

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

// ManagerCommitsClient contains the methods for the NetworkManagerCommits group.
// Don't use this type directly, use NewManagerCommitsClient() instead.
type ManagerCommitsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewManagerCommitsClient creates a new instance of ManagerCommitsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagerCommitsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagerCommitsClient, error) {
	cl, err := arm.NewClient(moduleName+".ManagerCommitsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagerCommitsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginPost - Post a Network Manager Commit.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01
//   - resourceGroupName - The name of the resource group.
//   - networkManagerName - The name of the network manager.
//   - parameters - Parameters supplied to specify which Managed Network commit is.
//   - options - ManagerCommitsClientBeginPostOptions contains the optional parameters for the ManagerCommitsClient.BeginPost
//     method.
func (client *ManagerCommitsClient) BeginPost(ctx context.Context, resourceGroupName string, networkManagerName string, parameters ManagerCommit, options *ManagerCommitsClientBeginPostOptions) (*runtime.Poller[ManagerCommitsClientPostResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.post(ctx, resourceGroupName, networkManagerName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ManagerCommitsClientPostResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ManagerCommitsClientPostResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Post - Post a Network Manager Commit.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01
func (client *ManagerCommitsClient) post(ctx context.Context, resourceGroupName string, networkManagerName string, parameters ManagerCommit, options *ManagerCommitsClientBeginPostOptions) (*http.Response, error) {
	var err error
	req, err := client.postCreateRequest(ctx, resourceGroupName, networkManagerName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// postCreateRequest creates the Post request.
func (client *ManagerCommitsClient) postCreateRequest(ctx context.Context, resourceGroupName string, networkManagerName string, parameters ManagerCommit, options *ManagerCommitsClientBeginPostOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/commit"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkManagerName == "" {
		return nil, errors.New("parameter networkManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkManagerName}", url.PathEscape(networkManagerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}
