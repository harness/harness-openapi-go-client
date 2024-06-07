/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub
 *
 * API version: 1.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// This is the favorite entity defined in Harness
type FavoriteDto struct {
	// Organization Identifier for the Entity.
	Org string `json:"org,omitempty"`
	// Project Identifier for the Entity.
	Project string `json:"project,omitempty"`
	// Identifier of the user.
	UserId string `json:"user_id,omitempty"`
	Module *ModuleType `json:"module,omitempty"`
	ResourceType *FavoritesResourceType `json:"resource_type,omitempty"`
	// Resource Identifier of the Entity.
	ResourceId string `json:"resource_id,omitempty"`
}
