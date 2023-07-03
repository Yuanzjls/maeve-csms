package ocpp201

type GenericStatusEnumType string

const GenericStatusEnumTypeAccepted GenericStatusEnumType = "Accepted"
const GenericStatusEnumTypeRejected GenericStatusEnumType = "Rejected"

type SignCertificateResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status GenericStatusEnumType `json:"status" yaml:"status" mapstructure:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty" yaml:"statusInfo,omitempty" mapstructure:"statusInfo,omitempty"`
}

func (*SignCertificateResponseJson) IsResponse() {}
