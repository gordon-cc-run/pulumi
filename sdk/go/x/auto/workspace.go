// Copyright 2016-2020, Pulumi Corporation.
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

package auto

import (
	"context"

	"github.com/pulumi/pulumi/sdk/v2/go/common/workspace"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Workspace is the execution context containing a single Pulumi project, a program, and multiple stacks.
// Workspaces are used to manage the execution environment, providing various utilities such as plugin
// installation, environment configuration ($PULUMI_HOME), and creation, deletion, and listing of Stacks.
type Workspace interface {
	// ProjectSettings returns the settings object for the current project if any.
	ProjectSettings(context.Context) (*workspace.Project, error)
	// WriteProjectSettings overwrites the settings object in the current project.
	// There can only be a single project per workspace. Fails is new project name does not match old.
	WriteProjectSettings(context.Context, *workspace.Project) error
	// StackSettings returns the settings object for the stack matching the specified fullyQualifiedStackName if any.
	StackSettings(context.Context, string) (*workspace.ProjectStack, error)
	// WriteStackSettings overwrites the settings object for the stack matching the specified fullyQualifiedStackName.
	WriteStackSettings(context.Context, string, *workspace.ProjectStack) error
	// SerializeArgsForOp is hook to provide additional args to every CLI commands before they are executed.
	// Provided with fullyQualifiedStackName,
	// returns a list of args to append to an invoked command ["--config=...", ].
	SerializeArgsForOp(context.Context, string) ([]string, error)
	// PostOpCallback is a hook executed after every command. Called with the fullyQualifiedStackName.
	// An extensibility point to perform workspace cleanup (CLI operations may create/modify a pulumi.stack.yaml).
	PostOpCallback(context.Context, string) error
	// GetConfig returns the value associated with the specified fullyQualifiedStackName and key,
	// scoped to the current workspace.
	GetConfig(context.Context, string, string) (ConfigValue, error)
	// GetAllConfig returns the config map for the specified fullyQualifiedStackName, scoped to the current workspace.
	GetAllConfig(context.Context, string) (ConfigMap, error)
	// SetConfig sets the specified key-value pair on the provided fullyQualifiedStackName.
	SetConfig(context.Context, string, string, ConfigValue) error
	// SetAllConfig sets all values in the provided config map for the specified fullyQualifiedStackName.
	SetAllConfig(context.Context, string, ConfigMap) error
	// RemoveConfig removes the specified key-value pair on the provided fullyQualifiedStackName.
	RemoveConfig(context.Context, string, string) error
	// RemoveAllConfig removes all values in the provided key list for the specified fullyQualifiedStackName.
	RemoveAllConfig(context.Context, string, []string) error
	// RefreshConfig gets and sets the config map used with the last Update for Stack matching fullyQualifiedStackName.
	RefreshConfig(context.Context, string) (ConfigMap, error)
	// WorkDir returns the working directory to run Pulumi CLI commands.
	WorkDir() string
	// PulumiHome returns the directory override for CLI metadata if set.
	PulumiHome() *string
	// WhoAmI returns the currently authenticated user.
	WhoAmI(context.Context) (string, error)
	// Stack returns a summary of the currently selected stack, if any.
	Stack(context.Context) (*StackSummary, error)
	// CreateStack creates and sets a new stack with the fullyQualifiedStackName, failing if one already exists.
	CreateStack(context.Context, string) error
	// SelectStack selects and sets an existing stack matching the fullyQualifiedStackName, failing if none exists.
	SelectStack(context.Context, string) error
	// RemoveStack deletes the stack and all associated configuration and history.
	RemoveStack(context.Context, string) error
	// ListStacks returns all Stacks created under the current Project.
	// This queries underlying backend and may return stacks not present in the Workspace.
	ListStacks(context.Context) ([]StackSummary, error)
	// InstallPlugin acquires the plugin matching the specified name and version.
	InstallPlugin(context.Context, string, string) error
	// RemovePlugin deletes the plugin matching the specified name and verision.
	RemovePlugin(context.Context, string, string) error
	// ListPlugins lists all installed plugins.
	ListPlugins(context.Context) ([]workspace.PluginInfo, error)
	// Program returns the program `pulumi.RunFunc` to be used for Preview/Update if any.
	// If none is specified, the stack will refer to ProjectSettings for this information.
	Program() pulumi.RunFunc
	// SetProgram sets the program associated with the Workspace to the specified `pulumi.RunFunc`.
	SetProgram(pulumi.RunFunc)
}

// ConfigValue is a configuration value used by a Pulumi program.
// Allows differentiating between secret and plaintext values by setting the `Secret` property.
type ConfigValue struct {
	Value  string
	Secret bool
}

// ConfigMap is a map of ConfigValue used by Pulumi programs.
// Allows differentiating between secret and plaintext values.
type ConfigMap map[string]ConfigValue

// StackSummary is a description of a stack and its current status.
type StackSummary struct {
	Name             string `json:"name"`
	Current          bool   `json:"current"`
	LastUpdate       string `json:"lastUpdate,omitempty"`
	UpdateInProgress bool   `json:"updateInProgress"`
	ResourceCount    *int   `json:"resourceCount,omitempty"`
	URL              string `json:"url,omitempty"`
}
