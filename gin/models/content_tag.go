package models

type ContentTag struct {
	ID             int32   `json:id`
	OrganizationId int32   `json:organization_id`
	Name           string  `json:name`
	BrandIds       []int32 `json:brand_ids`
	CreatedAt      string  `json:created_at`
	UpdatedAt      string  `json:updated_at`
}
