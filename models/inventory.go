package models

type Inventory struct {
	Model

	// AuthID is the foreign key to the Auth model.
	AuthID int

	// ResourceID is the ID of the resource given by its platform, such as
	// an EC2 ID.
	ResourceID string `json:"resource_id" binding:"required"`

	// ResourceCreateOn is the date the resource was created.
	ResourceCreatedOn string `json:"resource_create_time" binding:"required"`
}

func (i *Inventory) Add(authID int) error {
	i.AuthID = authID
	if err := db.Create(i).Error; err != nil {
		return err
	}
	return nil
}
