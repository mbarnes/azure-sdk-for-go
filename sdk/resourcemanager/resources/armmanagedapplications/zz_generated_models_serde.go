//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagedapplications

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type Application.
func (a Application) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", a.ID)
	populate(objectMap, "identity", a.Identity)
	populate(objectMap, "kind", a.Kind)
	populate(objectMap, "location", a.Location)
	populate(objectMap, "managedBy", a.ManagedBy)
	populate(objectMap, "name", a.Name)
	populate(objectMap, "plan", a.Plan)
	populate(objectMap, "properties", a.Properties)
	populate(objectMap, "sku", a.SKU)
	populate(objectMap, "tags", a.Tags)
	populate(objectMap, "type", a.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ApplicationDefinition.
func (a ApplicationDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", a.ID)
	populate(objectMap, "identity", a.Identity)
	populate(objectMap, "location", a.Location)
	populate(objectMap, "managedBy", a.ManagedBy)
	populate(objectMap, "name", a.Name)
	populate(objectMap, "properties", a.Properties)
	populate(objectMap, "sku", a.SKU)
	populate(objectMap, "tags", a.Tags)
	populate(objectMap, "type", a.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ApplicationDefinitionProperties.
func (a ApplicationDefinitionProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "artifacts", a.Artifacts)
	populate(objectMap, "authorizations", a.Authorizations)
	populate(objectMap, "createUiDefinition", &a.CreateUIDefinition)
	populate(objectMap, "description", a.Description)
	populate(objectMap, "displayName", a.DisplayName)
	populate(objectMap, "isEnabled", a.IsEnabled)
	populate(objectMap, "lockLevel", a.LockLevel)
	populate(objectMap, "mainTemplate", &a.MainTemplate)
	populate(objectMap, "packageFileUri", a.PackageFileURI)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ApplicationPatchable.
func (a ApplicationPatchable) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", a.ID)
	populate(objectMap, "identity", a.Identity)
	populate(objectMap, "kind", a.Kind)
	populate(objectMap, "location", a.Location)
	populate(objectMap, "managedBy", a.ManagedBy)
	populate(objectMap, "name", a.Name)
	populate(objectMap, "plan", a.Plan)
	populate(objectMap, "properties", a.Properties)
	populate(objectMap, "sku", a.SKU)
	populate(objectMap, "tags", a.Tags)
	populate(objectMap, "type", a.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type GenericResource.
func (g GenericResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", g.ID)
	populate(objectMap, "identity", g.Identity)
	populate(objectMap, "location", g.Location)
	populate(objectMap, "managedBy", g.ManagedBy)
	populate(objectMap, "name", g.Name)
	populate(objectMap, "sku", g.SKU)
	populate(objectMap, "tags", g.Tags)
	populate(objectMap, "type", g.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", r.ID)
	populate(objectMap, "location", r.Location)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "tags", r.Tags)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}
