//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatalakeanalytics

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// AccountsListByResourceGroupPager provides operations for iterating over paged responses.
type AccountsListByResourceGroupPager struct {
	client    *AccountsClient
	current   AccountsListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AccountsListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AccountsListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AccountsListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DataLakeAnalyticsAccountListResult.NextLink == nil || len(*p.current.DataLakeAnalyticsAccountListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByResourceGroupHandleError(resp)
		return false
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current AccountsListByResourceGroupResponse page.
func (p *AccountsListByResourceGroupPager) PageResponse() AccountsListByResourceGroupResponse {
	return p.current
}

// AccountsListPager provides operations for iterating over paged responses.
type AccountsListPager struct {
	client    *AccountsClient
	current   AccountsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AccountsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AccountsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AccountsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DataLakeAnalyticsAccountListResult.NextLink == nil || len(*p.current.DataLakeAnalyticsAccountListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current AccountsListResponse page.
func (p *AccountsListPager) PageResponse() AccountsListResponse {
	return p.current
}

// ComputePoliciesListByAccountPager provides operations for iterating over paged responses.
type ComputePoliciesListByAccountPager struct {
	client    *ComputePoliciesClient
	current   ComputePoliciesListByAccountResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ComputePoliciesListByAccountResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ComputePoliciesListByAccountPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ComputePoliciesListByAccountPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ComputePolicyListResult.NextLink == nil || len(*p.current.ComputePolicyListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByAccountHandleError(resp)
		return false
	}
	result, err := p.client.listByAccountHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ComputePoliciesListByAccountResponse page.
func (p *ComputePoliciesListByAccountPager) PageResponse() ComputePoliciesListByAccountResponse {
	return p.current
}

// DataLakeStoreAccountsListByAccountPager provides operations for iterating over paged responses.
type DataLakeStoreAccountsListByAccountPager struct {
	client    *DataLakeStoreAccountsClient
	current   DataLakeStoreAccountsListByAccountResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, DataLakeStoreAccountsListByAccountResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DataLakeStoreAccountsListByAccountPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DataLakeStoreAccountsListByAccountPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DataLakeStoreAccountInformationListResult.NextLink == nil || len(*p.current.DataLakeStoreAccountInformationListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByAccountHandleError(resp)
		return false
	}
	result, err := p.client.listByAccountHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current DataLakeStoreAccountsListByAccountResponse page.
func (p *DataLakeStoreAccountsListByAccountPager) PageResponse() DataLakeStoreAccountsListByAccountResponse {
	return p.current
}

// FirewallRulesListByAccountPager provides operations for iterating over paged responses.
type FirewallRulesListByAccountPager struct {
	client    *FirewallRulesClient
	current   FirewallRulesListByAccountResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, FirewallRulesListByAccountResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *FirewallRulesListByAccountPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *FirewallRulesListByAccountPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.FirewallRuleListResult.NextLink == nil || len(*p.current.FirewallRuleListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByAccountHandleError(resp)
		return false
	}
	result, err := p.client.listByAccountHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current FirewallRulesListByAccountResponse page.
func (p *FirewallRulesListByAccountPager) PageResponse() FirewallRulesListByAccountResponse {
	return p.current
}

// StorageAccountsListByAccountPager provides operations for iterating over paged responses.
type StorageAccountsListByAccountPager struct {
	client    *StorageAccountsClient
	current   StorageAccountsListByAccountResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, StorageAccountsListByAccountResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *StorageAccountsListByAccountPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *StorageAccountsListByAccountPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.StorageAccountInformationListResult.NextLink == nil || len(*p.current.StorageAccountInformationListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByAccountHandleError(resp)
		return false
	}
	result, err := p.client.listByAccountHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current StorageAccountsListByAccountResponse page.
func (p *StorageAccountsListByAccountPager) PageResponse() StorageAccountsListByAccountResponse {
	return p.current
}

// StorageAccountsListSasTokensPager provides operations for iterating over paged responses.
type StorageAccountsListSasTokensPager struct {
	client    *StorageAccountsClient
	current   StorageAccountsListSasTokensResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, StorageAccountsListSasTokensResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *StorageAccountsListSasTokensPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *StorageAccountsListSasTokensPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SasTokenInformationListResult.NextLink == nil || len(*p.current.SasTokenInformationListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listSasTokensHandleError(resp)
		return false
	}
	result, err := p.client.listSasTokensHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current StorageAccountsListSasTokensResponse page.
func (p *StorageAccountsListSasTokensPager) PageResponse() StorageAccountsListSasTokensResponse {
	return p.current
}

// StorageAccountsListStorageContainersPager provides operations for iterating over paged responses.
type StorageAccountsListStorageContainersPager struct {
	client    *StorageAccountsClient
	current   StorageAccountsListStorageContainersResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, StorageAccountsListStorageContainersResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *StorageAccountsListStorageContainersPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *StorageAccountsListStorageContainersPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.StorageContainerListResult.NextLink == nil || len(*p.current.StorageContainerListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listStorageContainersHandleError(resp)
		return false
	}
	result, err := p.client.listStorageContainersHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current StorageAccountsListStorageContainersResponse page.
func (p *StorageAccountsListStorageContainersPager) PageResponse() StorageAccountsListStorageContainersResponse {
	return p.current
}
