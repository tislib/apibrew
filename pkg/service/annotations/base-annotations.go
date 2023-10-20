package annotations

// values:
// true
const Enabled = "true"

// resource flags
const KeepHistory = "KeepHistory"
const HistoryResource = "HistoryResource"
const NormalizedResource = "NormalizedResource"
const AutoCreated = "AutoCreated"
const EnableAudit = "EnableAudit"
const DisableVersion = "DisableVersion"
const DisableBackup = "DisableBackup"

// property params
const SourceDef = "SourceDef"
const SourceIdentity = "SourceIdentity"
const SourceMatchKey = "SourceMatchKey"
const Identity = "Identity"
const SpecialProperty = "SpecialProperty"
const PrimaryProperty = "PrimaryProperty"

// hcl
const IsHclLabel = "IsHclLabel"
const HclBlock = "HclBlock"

// request annotations
const IgnoreIfExists = "IgnoreIfExists"
const CheckVersion = "CheckVersion"

// security
const AllowPublicAccess = "AllowPublicAccess"

// sql
const SQLType = "SQLType"
const SQLUseTextType = "SQLUseTextType"
const UseJoinTable = "UseJoinTable"

// code generator
const TypeName = "TypeName"
const SelfContainedProperty = "SelfContainedProperty"
const AllowEmptyPrimitive = "AllowEmptyPrimitive"
const CommonType = "CommonType"

// restapi
const RestApiDisabled = "RestApiDisabled"

// openapi
const OpenApiGroup = "OpenApiGroup"
const OpenApiHide = "OpenApiHide"

// service
const ServiceKey = "ServiceKey"

// extensionId
const ExtensionId = "ExtensionId"

// bypass extensions
const BypassExtensions = "BypassExtensions"

var ClientAllowedAnnotations = map[string]bool{
	BypassExtensions: true,
}
