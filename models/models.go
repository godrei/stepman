package models

// -------------------
// --- Common models

// EnvironmentItemOptionsModel ...
type EnvironmentItemOptionsModel struct {
	Title             *string  `json:"title,omitempty" yaml:"title,omitempty"`
	Description       *string  `json:"description,omitempty" yaml:"description,omitempty"`
	ValueOptions      []string `json:"value_options,omitempty" yaml:"value_options,omitempty"`
	IsRequired        *bool    `json:"is_required,omitempty" yaml:"is_required,omitempty"`
	IsExpand          *bool    `json:"is_expand,omitempty" yaml:"is_expand,omitempty"`
	IsDontChangeValue *bool    `json:"is_dont_change_value,omitempty" yaml:"is_dont_change_value,omitempty"`
}

// EnvironmentItemModel ...
type EnvironmentItemModel map[string]interface{}

// StepSourceModel ...
type StepSourceModel struct {
	Git *string `json:"git,omitempty" yaml:"git,omitempty"`
}

// StepModel ...
type StepModel struct {
	Title               *string                `json:"title,omitempty" yaml:"title,omitempty"`
	Description         *string                `json:"description,omitempty" yaml:"description,omitempty"`
	Summary             *string                `json:"summary,omitempty" yaml:"summary,omitempty"`
	Website             *string                `json:"website,omitempty" yaml:"website,omitempty"`
	SourceCodeURL       *string                `json:"source_code_url,omitempty" yaml:"source_code_url,omitempty"`
	SupportURL          *string                `json:"support_url,omitempty" yaml:"support_url,omitempty"`
	Source              StepSourceModel        `json:"source,omitempty" yaml:"source,omitempty"`
	HostOsTags          []string               `json:"host_os_tags,omitempty" yaml:"host_os_tags,omitempty"`
	ProjectTypeTags     []string               `json:"project_type_tags,omitempty" yaml:"project_type_tags,omitempty"`
	TypeTags            []string               `json:"type_tags,omitempty" yaml:"type_tags,omitempty"`
	IsRequiresAdminUser *bool                  `json:"is_requires_admin_user,omitempty" yaml:"is_requires_admin_user,omitempty"`
	IsAlwaysRun         *bool                  `json:"is_always_run,omitempty" yaml:"is_always_run,omitempty"`
	IsNotImportant      *bool                  `json:"is_not_important,omitempty" yaml:"is_not_important,omitempty"`
	Inputs              []EnvironmentItemModel `json:"inputs,omitempty" yaml:"inputs,omitempty"`
	Outputs             []EnvironmentItemModel `json:"outputs,omitempty" yaml:"outputs,omitempty"`
}

// -------------------
// --- Steplib models

// StepGroupModel ...
type StepGroupModel struct {
	Versions            map[string]StepModel `json:"versions"`
	LatestVersionNumber string               `json:"latest_version_number"`
}

// StepHash ...
type StepHash map[string]StepGroupModel

// DownloadLocationModel ...
type DownloadLocationModel struct {
	Type string `json:"type"`
	Src  string `json:"src"`
}

// StepCollectionModel ...
type StepCollectionModel struct {
	FormatVersion        string                  `json:"format_version" yaml:"format_version"`
	GeneratedAtTimeStamp int64                   `json:"generated_at_timestamp" yaml:"generated_at_timestamp"`
	Steps                StepHash                `json:"steps" yaml:"steps"`
	SteplibSource        string                  `json:"steplib_source" yaml:"steplib_source"`
	DownloadLocations    []DownloadLocationModel `json:"download_locations" yaml:"download_locations"`
}

// WorkFlowModel ...
type WorkFlowModel struct {
	FormatVersion string      `json:"format_version"`
	Environments  []string    `json:"environments"`
	Steps         []StepModel `json:"steps"`
}

// -------------------
// --- Bitrise-cli models

// StepListItemModel ...
type StepListItemModel map[string]StepModel

// WorkflowModel ...
type WorkflowModel struct {
	Environments []EnvironmentItemModel `json:"environments"`
	Steps        []StepListItemModel    `json:"steps"`
}

// AppModel ...
type AppModel struct {
	Environments []EnvironmentItemModel `json:"environments" yaml:"environments"`
}

// BitriseConfigModel ...
type BitriseConfigModel struct {
	FormatVersion string                   `json:"format_version" yaml:"format_version"`
	App           AppModel                 `json:"app" yaml:"app"`
	Workflows     map[string]WorkflowModel `json:"workflows" yaml:"workflows"`
}
