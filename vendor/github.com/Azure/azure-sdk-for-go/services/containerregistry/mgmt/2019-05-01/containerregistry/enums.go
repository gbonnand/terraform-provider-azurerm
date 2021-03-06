package containerregistry

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Action enumerates the values for action.
type Action string

const (
	// Allow ...
	Allow Action = "Allow"
)

// PossibleActionValues returns an array of possible values for the Action const type.
func PossibleActionValues() []Action {
	return []Action{Allow}
}

// Architecture enumerates the values for architecture.
type Architecture string

const (
	// Amd64 ...
	Amd64 Architecture = "amd64"
	// Arm ...
	Arm Architecture = "arm"
	// X86 ...
	X86 Architecture = "x86"
)

// PossibleArchitectureValues returns an array of possible values for the Architecture const type.
func PossibleArchitectureValues() []Architecture {
	return []Architecture{Amd64, Arm, X86}
}

// BaseImageDependencyType enumerates the values for base image dependency type.
type BaseImageDependencyType string

const (
	// BuildTime ...
	BuildTime BaseImageDependencyType = "BuildTime"
	// RunTime ...
	RunTime BaseImageDependencyType = "RunTime"
)

// PossibleBaseImageDependencyTypeValues returns an array of possible values for the BaseImageDependencyType const type.
func PossibleBaseImageDependencyTypeValues() []BaseImageDependencyType {
	return []BaseImageDependencyType{BuildTime, RunTime}
}

// BaseImageTriggerType enumerates the values for base image trigger type.
type BaseImageTriggerType string

const (
	// All ...
	All BaseImageTriggerType = "All"
	// Runtime ...
	Runtime BaseImageTriggerType = "Runtime"
)

// PossibleBaseImageTriggerTypeValues returns an array of possible values for the BaseImageTriggerType const type.
func PossibleBaseImageTriggerTypeValues() []BaseImageTriggerType {
	return []BaseImageTriggerType{All, Runtime}
}

// DefaultAction enumerates the values for default action.
type DefaultAction string

const (
	// DefaultActionAllow ...
	DefaultActionAllow DefaultAction = "Allow"
	// DefaultActionDeny ...
	DefaultActionDeny DefaultAction = "Deny"
)

// PossibleDefaultActionValues returns an array of possible values for the DefaultAction const type.
func PossibleDefaultActionValues() []DefaultAction {
	return []DefaultAction{DefaultActionAllow, DefaultActionDeny}
}

// ImportMode enumerates the values for import mode.
type ImportMode string

const (
	// Force ...
	Force ImportMode = "Force"
	// NoForce ...
	NoForce ImportMode = "NoForce"
)

// PossibleImportModeValues returns an array of possible values for the ImportMode const type.
func PossibleImportModeValues() []ImportMode {
	return []ImportMode{Force, NoForce}
}

// OS enumerates the values for os.
type OS string

const (
	// Linux ...
	Linux OS = "Linux"
	// Windows ...
	Windows OS = "Windows"
)

// PossibleOSValues returns an array of possible values for the OS const type.
func PossibleOSValues() []OS {
	return []OS{Linux, Windows}
}

// PasswordName enumerates the values for password name.
type PasswordName string

const (
	// Password ...
	Password PasswordName = "password"
	// Password2 ...
	Password2 PasswordName = "password2"
)

// PossiblePasswordNameValues returns an array of possible values for the PasswordName const type.
func PossiblePasswordNameValues() []PasswordName {
	return []PasswordName{Password, Password2}
}

// PolicyStatus enumerates the values for policy status.
type PolicyStatus string

const (
	// Disabled ...
	Disabled PolicyStatus = "disabled"
	// Enabled ...
	Enabled PolicyStatus = "enabled"
)

// PossiblePolicyStatusValues returns an array of possible values for the PolicyStatus const type.
func PossiblePolicyStatusValues() []PolicyStatus {
	return []PolicyStatus{Disabled, Enabled}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Canceled ...
	Canceled ProvisioningState = "Canceled"
	// Creating ...
	Creating ProvisioningState = "Creating"
	// Deleting ...
	Deleting ProvisioningState = "Deleting"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
	// Updating ...
	Updating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Canceled, Creating, Deleting, Failed, Succeeded, Updating}
}

// RegistryUsageUnit enumerates the values for registry usage unit.
type RegistryUsageUnit string

const (
	// Bytes ...
	Bytes RegistryUsageUnit = "Bytes"
	// Count ...
	Count RegistryUsageUnit = "Count"
)

// PossibleRegistryUsageUnitValues returns an array of possible values for the RegistryUsageUnit const type.
func PossibleRegistryUsageUnitValues() []RegistryUsageUnit {
	return []RegistryUsageUnit{Bytes, Count}
}

// ResourceIdentityType enumerates the values for resource identity type.
type ResourceIdentityType string

const (
	// None ...
	None ResourceIdentityType = "None"
	// SystemAssigned ...
	SystemAssigned ResourceIdentityType = "SystemAssigned"
	// SystemAssignedUserAssigned ...
	SystemAssignedUserAssigned ResourceIdentityType = "SystemAssigned, UserAssigned"
	// UserAssigned ...
	UserAssigned ResourceIdentityType = "UserAssigned"
)

// PossibleResourceIdentityTypeValues returns an array of possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{None, SystemAssigned, SystemAssignedUserAssigned, UserAssigned}
}

// RunStatus enumerates the values for run status.
type RunStatus string

const (
	// RunStatusCanceled ...
	RunStatusCanceled RunStatus = "Canceled"
	// RunStatusError ...
	RunStatusError RunStatus = "Error"
	// RunStatusFailed ...
	RunStatusFailed RunStatus = "Failed"
	// RunStatusQueued ...
	RunStatusQueued RunStatus = "Queued"
	// RunStatusRunning ...
	RunStatusRunning RunStatus = "Running"
	// RunStatusStarted ...
	RunStatusStarted RunStatus = "Started"
	// RunStatusSucceeded ...
	RunStatusSucceeded RunStatus = "Succeeded"
	// RunStatusTimeout ...
	RunStatusTimeout RunStatus = "Timeout"
)

// PossibleRunStatusValues returns an array of possible values for the RunStatus const type.
func PossibleRunStatusValues() []RunStatus {
	return []RunStatus{RunStatusCanceled, RunStatusError, RunStatusFailed, RunStatusQueued, RunStatusRunning, RunStatusStarted, RunStatusSucceeded, RunStatusTimeout}
}

// RunType enumerates the values for run type.
type RunType string

const (
	// AutoBuild ...
	AutoBuild RunType = "AutoBuild"
	// AutoRun ...
	AutoRun RunType = "AutoRun"
	// QuickBuild ...
	QuickBuild RunType = "QuickBuild"
	// QuickRun ...
	QuickRun RunType = "QuickRun"
)

// PossibleRunTypeValues returns an array of possible values for the RunType const type.
func PossibleRunTypeValues() []RunType {
	return []RunType{AutoBuild, AutoRun, QuickBuild, QuickRun}
}

// SecretObjectType enumerates the values for secret object type.
type SecretObjectType string

const (
	// Opaque ...
	Opaque SecretObjectType = "Opaque"
	// Vaultsecret ...
	Vaultsecret SecretObjectType = "Vaultsecret"
)

// PossibleSecretObjectTypeValues returns an array of possible values for the SecretObjectType const type.
func PossibleSecretObjectTypeValues() []SecretObjectType {
	return []SecretObjectType{Opaque, Vaultsecret}
}

// SkuName enumerates the values for sku name.
type SkuName string

const (
	// Basic ...
	Basic SkuName = "Basic"
	// Classic ...
	Classic SkuName = "Classic"
	// Premium ...
	Premium SkuName = "Premium"
	// Standard ...
	Standard SkuName = "Standard"
)

// PossibleSkuNameValues returns an array of possible values for the SkuName const type.
func PossibleSkuNameValues() []SkuName {
	return []SkuName{Basic, Classic, Premium, Standard}
}

// SkuTier enumerates the values for sku tier.
type SkuTier string

const (
	// SkuTierBasic ...
	SkuTierBasic SkuTier = "Basic"
	// SkuTierClassic ...
	SkuTierClassic SkuTier = "Classic"
	// SkuTierPremium ...
	SkuTierPremium SkuTier = "Premium"
	// SkuTierStandard ...
	SkuTierStandard SkuTier = "Standard"
)

// PossibleSkuTierValues returns an array of possible values for the SkuTier const type.
func PossibleSkuTierValues() []SkuTier {
	return []SkuTier{SkuTierBasic, SkuTierClassic, SkuTierPremium, SkuTierStandard}
}

// SourceControlType enumerates the values for source control type.
type SourceControlType string

const (
	// Github ...
	Github SourceControlType = "Github"
	// VisualStudioTeamService ...
	VisualStudioTeamService SourceControlType = "VisualStudioTeamService"
)

// PossibleSourceControlTypeValues returns an array of possible values for the SourceControlType const type.
func PossibleSourceControlTypeValues() []SourceControlType {
	return []SourceControlType{Github, VisualStudioTeamService}
}

// SourceRegistryLoginMode enumerates the values for source registry login mode.
type SourceRegistryLoginMode string

const (
	// SourceRegistryLoginModeDefault ...
	SourceRegistryLoginModeDefault SourceRegistryLoginMode = "Default"
	// SourceRegistryLoginModeNone ...
	SourceRegistryLoginModeNone SourceRegistryLoginMode = "None"
)

// PossibleSourceRegistryLoginModeValues returns an array of possible values for the SourceRegistryLoginMode const type.
func PossibleSourceRegistryLoginModeValues() []SourceRegistryLoginMode {
	return []SourceRegistryLoginMode{SourceRegistryLoginModeDefault, SourceRegistryLoginModeNone}
}

// SourceTriggerEvent enumerates the values for source trigger event.
type SourceTriggerEvent string

const (
	// Commit ...
	Commit SourceTriggerEvent = "commit"
	// Pullrequest ...
	Pullrequest SourceTriggerEvent = "pullrequest"
)

// PossibleSourceTriggerEventValues returns an array of possible values for the SourceTriggerEvent const type.
func PossibleSourceTriggerEventValues() []SourceTriggerEvent {
	return []SourceTriggerEvent{Commit, Pullrequest}
}

// TaskStatus enumerates the values for task status.
type TaskStatus string

const (
	// TaskStatusDisabled ...
	TaskStatusDisabled TaskStatus = "Disabled"
	// TaskStatusEnabled ...
	TaskStatusEnabled TaskStatus = "Enabled"
)

// PossibleTaskStatusValues returns an array of possible values for the TaskStatus const type.
func PossibleTaskStatusValues() []TaskStatus {
	return []TaskStatus{TaskStatusDisabled, TaskStatusEnabled}
}

// TokenType enumerates the values for token type.
type TokenType string

const (
	// OAuth ...
	OAuth TokenType = "OAuth"
	// PAT ...
	PAT TokenType = "PAT"
)

// PossibleTokenTypeValues returns an array of possible values for the TokenType const type.
func PossibleTokenTypeValues() []TokenType {
	return []TokenType{OAuth, PAT}
}

// TriggerStatus enumerates the values for trigger status.
type TriggerStatus string

const (
	// TriggerStatusDisabled ...
	TriggerStatusDisabled TriggerStatus = "Disabled"
	// TriggerStatusEnabled ...
	TriggerStatusEnabled TriggerStatus = "Enabled"
)

// PossibleTriggerStatusValues returns an array of possible values for the TriggerStatus const type.
func PossibleTriggerStatusValues() []TriggerStatus {
	return []TriggerStatus{TriggerStatusDisabled, TriggerStatusEnabled}
}

// TrustPolicyType enumerates the values for trust policy type.
type TrustPolicyType string

const (
	// Notary ...
	Notary TrustPolicyType = "Notary"
)

// PossibleTrustPolicyTypeValues returns an array of possible values for the TrustPolicyType const type.
func PossibleTrustPolicyTypeValues() []TrustPolicyType {
	return []TrustPolicyType{Notary}
}

// Type enumerates the values for type.
type Type string

const (
	// TypeDockerBuildRequest ...
	TypeDockerBuildRequest Type = "DockerBuildRequest"
	// TypeEncodedTaskRunRequest ...
	TypeEncodedTaskRunRequest Type = "EncodedTaskRunRequest"
	// TypeFileTaskRunRequest ...
	TypeFileTaskRunRequest Type = "FileTaskRunRequest"
	// TypeRunRequest ...
	TypeRunRequest Type = "RunRequest"
	// TypeTaskRunRequest ...
	TypeTaskRunRequest Type = "TaskRunRequest"
)

// PossibleTypeValues returns an array of possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{TypeDockerBuildRequest, TypeEncodedTaskRunRequest, TypeFileTaskRunRequest, TypeRunRequest, TypeTaskRunRequest}
}

// TypeBasicTaskStepProperties enumerates the values for type basic task step properties.
type TypeBasicTaskStepProperties string

const (
	// TypeDocker ...
	TypeDocker TypeBasicTaskStepProperties = "Docker"
	// TypeEncodedTask ...
	TypeEncodedTask TypeBasicTaskStepProperties = "EncodedTask"
	// TypeFileTask ...
	TypeFileTask TypeBasicTaskStepProperties = "FileTask"
	// TypeTaskStepProperties ...
	TypeTaskStepProperties TypeBasicTaskStepProperties = "TaskStepProperties"
)

// PossibleTypeBasicTaskStepPropertiesValues returns an array of possible values for the TypeBasicTaskStepProperties const type.
func PossibleTypeBasicTaskStepPropertiesValues() []TypeBasicTaskStepProperties {
	return []TypeBasicTaskStepProperties{TypeDocker, TypeEncodedTask, TypeFileTask, TypeTaskStepProperties}
}

// TypeBasicTaskStepUpdateParameters enumerates the values for type basic task step update parameters.
type TypeBasicTaskStepUpdateParameters string

const (
	// TypeBasicTaskStepUpdateParametersTypeDocker ...
	TypeBasicTaskStepUpdateParametersTypeDocker TypeBasicTaskStepUpdateParameters = "Docker"
	// TypeBasicTaskStepUpdateParametersTypeEncodedTask ...
	TypeBasicTaskStepUpdateParametersTypeEncodedTask TypeBasicTaskStepUpdateParameters = "EncodedTask"
	// TypeBasicTaskStepUpdateParametersTypeFileTask ...
	TypeBasicTaskStepUpdateParametersTypeFileTask TypeBasicTaskStepUpdateParameters = "FileTask"
	// TypeBasicTaskStepUpdateParametersTypeTaskStepUpdateParameters ...
	TypeBasicTaskStepUpdateParametersTypeTaskStepUpdateParameters TypeBasicTaskStepUpdateParameters = "TaskStepUpdateParameters"
)

// PossibleTypeBasicTaskStepUpdateParametersValues returns an array of possible values for the TypeBasicTaskStepUpdateParameters const type.
func PossibleTypeBasicTaskStepUpdateParametersValues() []TypeBasicTaskStepUpdateParameters {
	return []TypeBasicTaskStepUpdateParameters{TypeBasicTaskStepUpdateParametersTypeDocker, TypeBasicTaskStepUpdateParametersTypeEncodedTask, TypeBasicTaskStepUpdateParametersTypeFileTask, TypeBasicTaskStepUpdateParametersTypeTaskStepUpdateParameters}
}

// Variant enumerates the values for variant.
type Variant string

const (
	// V6 ...
	V6 Variant = "v6"
	// V7 ...
	V7 Variant = "v7"
	// V8 ...
	V8 Variant = "v8"
)

// PossibleVariantValues returns an array of possible values for the Variant const type.
func PossibleVariantValues() []Variant {
	return []Variant{V6, V7, V8}
}

// WebhookAction enumerates the values for webhook action.
type WebhookAction string

const (
	// ChartDelete ...
	ChartDelete WebhookAction = "chart_delete"
	// ChartPush ...
	ChartPush WebhookAction = "chart_push"
	// Delete ...
	Delete WebhookAction = "delete"
	// Push ...
	Push WebhookAction = "push"
	// Quarantine ...
	Quarantine WebhookAction = "quarantine"
)

// PossibleWebhookActionValues returns an array of possible values for the WebhookAction const type.
func PossibleWebhookActionValues() []WebhookAction {
	return []WebhookAction{ChartDelete, ChartPush, Delete, Push, Quarantine}
}

// WebhookStatus enumerates the values for webhook status.
type WebhookStatus string

const (
	// WebhookStatusDisabled ...
	WebhookStatusDisabled WebhookStatus = "disabled"
	// WebhookStatusEnabled ...
	WebhookStatusEnabled WebhookStatus = "enabled"
)

// PossibleWebhookStatusValues returns an array of possible values for the WebhookStatus const type.
func PossibleWebhookStatusValues() []WebhookStatus {
	return []WebhookStatus{WebhookStatusDisabled, WebhookStatusEnabled}
}
