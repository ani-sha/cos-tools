/*
Connector Service Fleet Manager Admin APIs

Connector Service Fleet Manager Admin is a Rest API to manage connector clusters.

API version: 0.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

// ConnectorType Represents a connector type supported by the API
type ConnectorType struct {
	Id   string `json:"id,omitempty"`
	Kind string `json:"kind,omitempty"`
	Href string `json:"href,omitempty"`
	// Name of the connector type.
	Name string `json:"name"`
	// Version of the connector type.
	Version string `json:"version"`
	// Channels of the connector type.
	Channels []Channel `json:"channels,omitempty"`
	// A description of the connector.
	Description string `json:"description,omitempty"`
	// Connector type is deprecated and removed from the catalog.
	Deprecated bool `json:"deprecated,omitempty"`
	// URL to an icon of the connector.
	IconHref string `json:"icon_href,omitempty"`
	// Labels used to categorize the connector
	Labels []string `json:"labels,omitempty"`
	// Name-value string annotations for resource
	Annotations map[string]string `json:"annotations,omitempty"`
	// Ranking for featured connectors
	FeaturedRank int32 `json:"featured_rank,omitempty"`
	// The capabilities supported by the connector
	Capabilities []string `json:"capabilities,omitempty"`
	// A json schema that can be used to validate a ConnectorRequest connector field.
	Schema map[string]interface{} `json:"schema"`
}
