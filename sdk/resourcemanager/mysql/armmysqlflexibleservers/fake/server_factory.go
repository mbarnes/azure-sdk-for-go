//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// ServerFactory is a fake server for instances of the armmysqlflexibleservers.ClientFactory type.
type ServerFactory struct {
	AzureADAdministratorsServer                AzureADAdministratorsServer
	BackupAndExportServer                      BackupAndExportServer
	BackupsServer                              BackupsServer
	CheckNameAvailabilityServer                CheckNameAvailabilityServer
	CheckNameAvailabilityWithoutLocationServer CheckNameAvailabilityWithoutLocationServer
	CheckVirtualNetworkSubnetUsageServer       CheckVirtualNetworkSubnetUsageServer
	ConfigurationsServer                       ConfigurationsServer
	DatabasesServer                            DatabasesServer
	FirewallRulesServer                        FirewallRulesServer
	GetPrivateDNSZoneSuffixServer              GetPrivateDNSZoneSuffixServer
	LocationBasedCapabilitiesServer            LocationBasedCapabilitiesServer
	LogFilesServer                             LogFilesServer
	OperationsServer                           OperationsServer
	ReplicasServer                             ReplicasServer
	ServersServer                              ServersServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armmysqlflexibleservers.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armmysqlflexibleservers.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                          *ServerFactory
	trMu                                         sync.Mutex
	trAzureADAdministratorsServer                *AzureADAdministratorsServerTransport
	trBackupAndExportServer                      *BackupAndExportServerTransport
	trBackupsServer                              *BackupsServerTransport
	trCheckNameAvailabilityServer                *CheckNameAvailabilityServerTransport
	trCheckNameAvailabilityWithoutLocationServer *CheckNameAvailabilityWithoutLocationServerTransport
	trCheckVirtualNetworkSubnetUsageServer       *CheckVirtualNetworkSubnetUsageServerTransport
	trConfigurationsServer                       *ConfigurationsServerTransport
	trDatabasesServer                            *DatabasesServerTransport
	trFirewallRulesServer                        *FirewallRulesServerTransport
	trGetPrivateDNSZoneSuffixServer              *GetPrivateDNSZoneSuffixServerTransport
	trLocationBasedCapabilitiesServer            *LocationBasedCapabilitiesServerTransport
	trLogFilesServer                             *LogFilesServerTransport
	trOperationsServer                           *OperationsServerTransport
	trReplicasServer                             *ReplicasServerTransport
	trServersServer                              *ServersServerTransport
}

// Do implements the policy.Transporter interface for ServerFactoryTransport.
func (s *ServerFactoryTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	client := method[:strings.Index(method, ".")]
	var resp *http.Response
	var err error

	switch client {
	case "AzureADAdministratorsClient":
		initServer(s, &s.trAzureADAdministratorsServer, func() *AzureADAdministratorsServerTransport {
			return NewAzureADAdministratorsServerTransport(&s.srv.AzureADAdministratorsServer)
		})
		resp, err = s.trAzureADAdministratorsServer.Do(req)
	case "BackupAndExportClient":
		initServer(s, &s.trBackupAndExportServer, func() *BackupAndExportServerTransport {
			return NewBackupAndExportServerTransport(&s.srv.BackupAndExportServer)
		})
		resp, err = s.trBackupAndExportServer.Do(req)
	case "BackupsClient":
		initServer(s, &s.trBackupsServer, func() *BackupsServerTransport { return NewBackupsServerTransport(&s.srv.BackupsServer) })
		resp, err = s.trBackupsServer.Do(req)
	case "CheckNameAvailabilityClient":
		initServer(s, &s.trCheckNameAvailabilityServer, func() *CheckNameAvailabilityServerTransport {
			return NewCheckNameAvailabilityServerTransport(&s.srv.CheckNameAvailabilityServer)
		})
		resp, err = s.trCheckNameAvailabilityServer.Do(req)
	case "CheckNameAvailabilityWithoutLocationClient":
		initServer(s, &s.trCheckNameAvailabilityWithoutLocationServer, func() *CheckNameAvailabilityWithoutLocationServerTransport {
			return NewCheckNameAvailabilityWithoutLocationServerTransport(&s.srv.CheckNameAvailabilityWithoutLocationServer)
		})
		resp, err = s.trCheckNameAvailabilityWithoutLocationServer.Do(req)
	case "CheckVirtualNetworkSubnetUsageClient":
		initServer(s, &s.trCheckVirtualNetworkSubnetUsageServer, func() *CheckVirtualNetworkSubnetUsageServerTransport {
			return NewCheckVirtualNetworkSubnetUsageServerTransport(&s.srv.CheckVirtualNetworkSubnetUsageServer)
		})
		resp, err = s.trCheckVirtualNetworkSubnetUsageServer.Do(req)
	case "ConfigurationsClient":
		initServer(s, &s.trConfigurationsServer, func() *ConfigurationsServerTransport {
			return NewConfigurationsServerTransport(&s.srv.ConfigurationsServer)
		})
		resp, err = s.trConfigurationsServer.Do(req)
	case "DatabasesClient":
		initServer(s, &s.trDatabasesServer, func() *DatabasesServerTransport { return NewDatabasesServerTransport(&s.srv.DatabasesServer) })
		resp, err = s.trDatabasesServer.Do(req)
	case "FirewallRulesClient":
		initServer(s, &s.trFirewallRulesServer, func() *FirewallRulesServerTransport {
			return NewFirewallRulesServerTransport(&s.srv.FirewallRulesServer)
		})
		resp, err = s.trFirewallRulesServer.Do(req)
	case "GetPrivateDNSZoneSuffixClient":
		initServer(s, &s.trGetPrivateDNSZoneSuffixServer, func() *GetPrivateDNSZoneSuffixServerTransport {
			return NewGetPrivateDNSZoneSuffixServerTransport(&s.srv.GetPrivateDNSZoneSuffixServer)
		})
		resp, err = s.trGetPrivateDNSZoneSuffixServer.Do(req)
	case "LocationBasedCapabilitiesClient":
		initServer(s, &s.trLocationBasedCapabilitiesServer, func() *LocationBasedCapabilitiesServerTransport {
			return NewLocationBasedCapabilitiesServerTransport(&s.srv.LocationBasedCapabilitiesServer)
		})
		resp, err = s.trLocationBasedCapabilitiesServer.Do(req)
	case "LogFilesClient":
		initServer(s, &s.trLogFilesServer, func() *LogFilesServerTransport { return NewLogFilesServerTransport(&s.srv.LogFilesServer) })
		resp, err = s.trLogFilesServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "ReplicasClient":
		initServer(s, &s.trReplicasServer, func() *ReplicasServerTransport { return NewReplicasServerTransport(&s.srv.ReplicasServer) })
		resp, err = s.trReplicasServer.Do(req)
	case "ServersClient":
		initServer(s, &s.trServersServer, func() *ServersServerTransport { return NewServersServerTransport(&s.srv.ServersServer) })
		resp, err = s.trServersServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func initServer[T any](s *ServerFactoryTransport, dst **T, src func() *T) {
	s.trMu.Lock()
	if *dst == nil {
		*dst = src()
	}
	s.trMu.Unlock()
}
