/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 5.3.0
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type MultiEnvironmentDependentFlag struct {
	Name string `json:"name,omitempty"`
	Key string `json:"key,omitempty"`
	Environments []DependentFlagEnvironment `json:"environments,omitempty"`
	Links *DependentFlagsLinks `json:"_links,omitempty"`
	Site *Site `json:"_site,omitempty"`
}
