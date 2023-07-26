//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9975d3476c05bcc6bd9535ad3dfb564e6a168fa5/specification/network/resource-manager/Microsoft.Network/stable/2023-02-01/examples/NetworkManagerCommitPost.json
func ExampleManagerCommitsClient_BeginPost() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagerCommitsClient().BeginPost(ctx, "resoureGroupSample", "testNetworkManager", armnetwork.ManagerCommit{
		CommitType: to.Ptr(armnetwork.ConfigurationTypeSecurityAdmin),
		ConfigurationIDs: []*string{
			to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resoureGroupSample/providers/Microsoft.Network/networkManagers/testNetworkManager/securityAdminConfigurations/SampleSecurityAdminConfig")},
		TargetLocations: []*string{
			to.Ptr("useast")},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagerCommit = armnetwork.ManagerCommit{
	// 	CommitID: to.Ptr("testCommitId"),
	// 	CommitType: to.Ptr(armnetwork.ConfigurationTypeSecurityAdmin),
	// 	ConfigurationIDs: []*string{
	// 		to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resoureGroupSample/providers/Microsoft.Network/networkManagers/testNetworkManager/securityAdminConfigurations/SampleSecurityAdminConfig")},
	// 		TargetLocations: []*string{
	// 			to.Ptr("useast")},
	// 		}
}
