// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import (
	fmt "fmt"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	time "time"
)

var ApplicationFieldPathsNested = []string{
	"attributes",
	"contact_info",
	"created_at",
	"description",
	"ids",
	"ids.application_id",
	"name",
	"updated_at",
}

var ApplicationFieldPathsTopLevel = []string{
	"attributes",
	"contact_info",
	"created_at",
	"description",
	"ids",
	"name",
	"updated_at",
}

func (dst *Application) SetFields(src *Application, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "ids":
			if len(subs) > 0 {
				newDst := &dst.ApplicationIdentifiers
				var newSrc *ApplicationIdentifiers
				if src != nil {
					newSrc = &src.ApplicationIdentifiers
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ApplicationIdentifiers = src.ApplicationIdentifiers
				} else {
					var zero ApplicationIdentifiers
					dst.ApplicationIdentifiers = zero
				}
			}
		case "created_at":
			if len(subs) > 0 {
				return fmt.Errorf("'created_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.CreatedAt = src.CreatedAt
			} else {
				var zero time.Time
				dst.CreatedAt = zero
			}
		case "updated_at":
			if len(subs) > 0 {
				return fmt.Errorf("'updated_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.UpdatedAt = src.UpdatedAt
			} else {
				var zero time.Time
				dst.UpdatedAt = zero
			}
		case "name":
			if len(subs) > 0 {
				return fmt.Errorf("'name' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Name = src.Name
			} else {
				var zero string
				dst.Name = zero
			}
		case "description":
			if len(subs) > 0 {
				return fmt.Errorf("'description' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Description = src.Description
			} else {
				var zero string
				dst.Description = zero
			}
		case "attributes":
			if len(subs) > 0 {
				return fmt.Errorf("'attributes' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Attributes = src.Attributes
			} else {
				dst.Attributes = nil
			}
		case "contact_info":
			if len(subs) > 0 {
				return fmt.Errorf("'contact_info' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ContactInfo = src.ContactInfo
			} else {
				dst.ContactInfo = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

var ApplicationsFieldPathsNested = []string{
	"applications",
}

var ApplicationsFieldPathsTopLevel = []string{
	"applications",
}

func (dst *Applications) SetFields(src *Applications, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "applications":
			if len(subs) > 0 {
				return fmt.Errorf("'applications' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Applications = src.Applications
			} else {
				dst.Applications = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

var GetApplicationRequestFieldPathsNested = []string{
	"application_ids",
	"application_ids.application_id",
	"field_mask",
}

var GetApplicationRequestFieldPathsTopLevel = []string{
	"application_ids",
	"field_mask",
}

func (dst *GetApplicationRequest) SetFields(src *GetApplicationRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "application_ids":
			if len(subs) > 0 {
				newDst := &dst.ApplicationIdentifiers
				var newSrc *ApplicationIdentifiers
				if src != nil {
					newSrc = &src.ApplicationIdentifiers
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ApplicationIdentifiers = src.ApplicationIdentifiers
				} else {
					var zero ApplicationIdentifiers
					dst.ApplicationIdentifiers = zero
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				var zero github_com_gogo_protobuf_types.FieldMask
				dst.FieldMask = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

var ListApplicationsRequestFieldPathsNested = []string{
	"collaborator",
	"collaborator.ids.organization_ids",
	"collaborator.ids.organization_ids.organization_id",
	"collaborator.ids.user_ids",
	"collaborator.ids.user_ids.email",
	"collaborator.ids.user_ids.user_id",
	"field_mask",
	"ids",
	"limit",
	"order",
	"page",
}

var ListApplicationsRequestFieldPathsTopLevel = []string{
	"collaborator",
	"field_mask",
	"ids",
	"limit",
	"order",
	"page",
}

func (dst *ListApplicationsRequest) SetFields(src *ListApplicationsRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "collaborator":
			if len(subs) > 0 {
				newDst := dst.Collaborator
				if newDst == nil {
					newDst = &OrganizationOrUserIdentifiers{}
					dst.Collaborator = newDst
				}
				var newSrc *OrganizationOrUserIdentifiers
				if src != nil {
					newSrc = src.Collaborator
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Collaborator = src.Collaborator
				} else {
					dst.Collaborator = nil
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				var zero github_com_gogo_protobuf_types.FieldMask
				dst.FieldMask = zero
			}
		case "order":
			if len(subs) > 0 {
				return fmt.Errorf("'order' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Order = src.Order
			} else {
				var zero string
				dst.Order = zero
			}
		case "limit":
			if len(subs) > 0 {
				return fmt.Errorf("'limit' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Limit = src.Limit
			} else {
				var zero uint32
				dst.Limit = zero
			}
		case "page":
			if len(subs) > 0 {
				return fmt.Errorf("'page' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Page = src.Page
			} else {
				var zero uint32
				dst.Page = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

var CreateApplicationRequestFieldPathsNested = []string{
	"application",
	"application.attributes",
	"application.contact_info",
	"application.created_at",
	"application.description",
	"application.ids",
	"application.ids.application_id",
	"application.name",
	"application.updated_at",
	"collaborator",
	"collaborator.ids.organization_ids",
	"collaborator.ids.organization_ids.organization_id",
	"collaborator.ids.user_ids",
	"collaborator.ids.user_ids.email",
	"collaborator.ids.user_ids.user_id",
	"ids",
}

var CreateApplicationRequestFieldPathsTopLevel = []string{
	"application",
	"collaborator",
	"ids",
}

func (dst *CreateApplicationRequest) SetFields(src *CreateApplicationRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "application":
			if len(subs) > 0 {
				newDst := &dst.Application
				var newSrc *Application
				if src != nil {
					newSrc = &src.Application
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Application = src.Application
				} else {
					var zero Application
					dst.Application = zero
				}
			}
		case "collaborator":
			if len(subs) > 0 {
				newDst := &dst.Collaborator
				var newSrc *OrganizationOrUserIdentifiers
				if src != nil {
					newSrc = &src.Collaborator
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Collaborator = src.Collaborator
				} else {
					var zero OrganizationOrUserIdentifiers
					dst.Collaborator = zero
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

var UpdateApplicationRequestFieldPathsNested = []string{
	"application",
	"application.attributes",
	"application.contact_info",
	"application.created_at",
	"application.description",
	"application.ids",
	"application.ids.application_id",
	"application.name",
	"application.updated_at",
	"field_mask",
}

var UpdateApplicationRequestFieldPathsTopLevel = []string{
	"application",
	"field_mask",
}

func (dst *UpdateApplicationRequest) SetFields(src *UpdateApplicationRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "application":
			if len(subs) > 0 {
				newDst := &dst.Application
				var newSrc *Application
				if src != nil {
					newSrc = &src.Application
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Application = src.Application
				} else {
					var zero Application
					dst.Application = zero
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				var zero github_com_gogo_protobuf_types.FieldMask
				dst.FieldMask = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

var CreateApplicationAPIKeyRequestFieldPathsNested = []string{
	"application_ids",
	"application_ids.application_id",
	"name",
	"rights",
}

var CreateApplicationAPIKeyRequestFieldPathsTopLevel = []string{
	"application_ids",
	"name",
	"rights",
}

func (dst *CreateApplicationAPIKeyRequest) SetFields(src *CreateApplicationAPIKeyRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "application_ids":
			if len(subs) > 0 {
				newDst := &dst.ApplicationIdentifiers
				var newSrc *ApplicationIdentifiers
				if src != nil {
					newSrc = &src.ApplicationIdentifiers
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ApplicationIdentifiers = src.ApplicationIdentifiers
				} else {
					var zero ApplicationIdentifiers
					dst.ApplicationIdentifiers = zero
				}
			}
		case "name":
			if len(subs) > 0 {
				return fmt.Errorf("'name' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Name = src.Name
			} else {
				var zero string
				dst.Name = zero
			}
		case "rights":
			if len(subs) > 0 {
				return fmt.Errorf("'rights' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Rights = src.Rights
			} else {
				dst.Rights = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

var UpdateApplicationAPIKeyRequestFieldPathsNested = []string{
	"api_key",
	"api_key.id",
	"api_key.key",
	"api_key.name",
	"api_key.rights",
	"application_ids",
	"application_ids.application_id",
}

var UpdateApplicationAPIKeyRequestFieldPathsTopLevel = []string{
	"api_key",
	"application_ids",
}

func (dst *UpdateApplicationAPIKeyRequest) SetFields(src *UpdateApplicationAPIKeyRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "application_ids":
			if len(subs) > 0 {
				newDst := &dst.ApplicationIdentifiers
				var newSrc *ApplicationIdentifiers
				if src != nil {
					newSrc = &src.ApplicationIdentifiers
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ApplicationIdentifiers = src.ApplicationIdentifiers
				} else {
					var zero ApplicationIdentifiers
					dst.ApplicationIdentifiers = zero
				}
			}
		case "api_key":
			if len(subs) > 0 {
				newDst := &dst.APIKey
				var newSrc *APIKey
				if src != nil {
					newSrc = &src.APIKey
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.APIKey = src.APIKey
				} else {
					var zero APIKey
					dst.APIKey = zero
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

var SetApplicationCollaboratorRequestFieldPathsNested = []string{
	"application_ids",
	"application_ids.application_id",
	"collaborator",
	"collaborator.ids",
	"collaborator.ids.ids.organization_ids",
	"collaborator.ids.ids.organization_ids.organization_id",
	"collaborator.ids.ids.user_ids",
	"collaborator.ids.ids.user_ids.email",
	"collaborator.ids.ids.user_ids.user_id",
	"collaborator.rights",
	"ids",
}

var SetApplicationCollaboratorRequestFieldPathsTopLevel = []string{
	"application_ids",
	"collaborator",
	"ids",
}

func (dst *SetApplicationCollaboratorRequest) SetFields(src *SetApplicationCollaboratorRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "application_ids":
			if len(subs) > 0 {
				newDst := &dst.ApplicationIdentifiers
				var newSrc *ApplicationIdentifiers
				if src != nil {
					newSrc = &src.ApplicationIdentifiers
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ApplicationIdentifiers = src.ApplicationIdentifiers
				} else {
					var zero ApplicationIdentifiers
					dst.ApplicationIdentifiers = zero
				}
			}
		case "collaborator":
			if len(subs) > 0 {
				newDst := &dst.Collaborator
				var newSrc *Collaborator
				if src != nil {
					newSrc = &src.Collaborator
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Collaborator = src.Collaborator
				} else {
					var zero Collaborator
					dst.Collaborator = zero
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}