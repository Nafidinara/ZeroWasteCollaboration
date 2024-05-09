package types

type OrganizationType string

const (
	Community   OrganizationType = "community"
	Company     OrganizationType = "company"
	Institution OrganizationType = "institution"
	NGO         OrganizationType = "ngo"
	Agency      OrganizationType = "agency"
)