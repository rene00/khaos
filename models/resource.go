package models

import (
	"errors"
)

// Resource is the struct for a resource. An example of a resource is an AWS
// EC2 instance.
type Resource struct {
	Model

	// AuthID is the foreign key to the Auth model.
	AuthID uint `sql:"type:integer REFERENCES auth(id)" gorm:"not null"`

	// ResourceTypeID is the foreign key to the ResourceType model.
	ResourceTypeID uint `sql:"type:integer REFERENCES resource_type(id)" gorm:"not null"`

	// ResourceType is the struct to ResourceType.
	ResourceType ResourceType

	// PlatformResourceID is the ID of the resource given by its platform,
	// such as an EC2 ID.
	PlatformResourceID string `json:"platform_resource_id" binding:"required"`
}

// ResourceType is the struct for a resource type. An example of a resource
// type is AWSInstance.
type ResourceType struct {
	Model

	// Name is the name of the Resource Type.
	Name string `gorm:"unique;not null"`
}

func (r *Resource) Add(authID uint) error {
	r.AuthID = authID
	if err := db.Create(r).Error; err != nil {
		return err
	}
	return nil
}

func GetResource(resourceID uint) (Resource, error) {
	var resource Resource
	err := db.Where("id = ?", resourceID).First(&resource).Error
	if err != nil {
		return resource, errors.New("resource not found")
	}
	err = db.Where("id = ?", resource.ResourceTypeID).First(&resource.ResourceType).Error
	if err != nil {
		return resource, errors.New("resource type not found")
	}
	return resource, nil
}

func GetResourceTypes() ([]*ResourceType, error) {
	var resourceTypes []*ResourceType
	err := db.Find(&resourceTypes).Error
	if err != nil {
		return nil, err
	}
	return resourceTypes, nil

}

func GetResourceType(name string) (ResourceType, error) {
	var resourceType ResourceType
	err := db.Where("name = ?", name).First(&resourceType).Error
	if err != nil {
		return resourceType, errors.New("resource type not found")
	}
	return resourceType, nil
}
