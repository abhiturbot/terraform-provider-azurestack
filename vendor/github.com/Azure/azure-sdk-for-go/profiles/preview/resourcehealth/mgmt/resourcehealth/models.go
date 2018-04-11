// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package resourcehealth

import original "github.com/Azure/azure-sdk-for-go/services/resourcehealth/mgmt/2017-07-01/resourcehealth"

type AvailabilityStatusesClient = original.AvailabilityStatusesClient

func NewAvailabilityStatusesClient(subscriptionID string) AvailabilityStatusesClient {
	return original.NewAvailabilityStatusesClient(subscriptionID)
}
func NewAvailabilityStatusesClientWithBaseURI(baseURI string, subscriptionID string) AvailabilityStatusesClient {
	return original.NewAvailabilityStatusesClientWithBaseURI(baseURI, subscriptionID)
}

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}

type AvailabilityStateValues = original.AvailabilityStateValues

const (
	Available   AvailabilityStateValues = original.Available
	Unavailable AvailabilityStateValues = original.Unavailable
	Unknown     AvailabilityStateValues = original.Unknown
)

func PossibleAvailabilityStateValuesValues() []AvailabilityStateValues {
	return original.PossibleAvailabilityStateValuesValues()
}

type ReasonChronicityTypes = original.ReasonChronicityTypes

const (
	Persistent ReasonChronicityTypes = original.Persistent
	Transient  ReasonChronicityTypes = original.Transient
)

func PossibleReasonChronicityTypesValues() []ReasonChronicityTypes {
	return original.PossibleReasonChronicityTypesValues()
}

type AvailabilityStatus = original.AvailabilityStatus
type AvailabilityStatusListResult = original.AvailabilityStatusListResult
type AvailabilityStatusListResultIterator = original.AvailabilityStatusListResultIterator
type AvailabilityStatusListResultPage = original.AvailabilityStatusListResultPage
type AvailabilityStatusProperties = original.AvailabilityStatusProperties
type AvailabilityStatusPropertiesRecentlyResolvedState = original.AvailabilityStatusPropertiesRecentlyResolvedState
type ErrorResponse = original.ErrorResponse
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type RecommendedAction = original.RecommendedAction
type ServiceImpactingEvent = original.ServiceImpactingEvent
type ServiceImpactingEventIncidentProperties = original.ServiceImpactingEventIncidentProperties
type ServiceImpactingEventStatus = original.ServiceImpactingEventStatus
type OperationsClient = original.OperationsClient

func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
