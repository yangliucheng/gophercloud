package users

import (
	"github.com/rackspace/gophercloud"
	os "github.com/rackspace/gophercloud/openstack/db/v1/users"
	"github.com/rackspace/gophercloud/pagination"
)

// Create will create a new database user for the specified database instance.
func Create(client *gophercloud.ServiceClient, instanceID string, opts os.CreateOptsBuilder) os.CreateResult {
	return os.Create(client, instanceID, opts)
}

// List will list all available users for a specified database instance.
func List(client *gophercloud.ServiceClient, instanceID string) pagination.Pager {
	return os.List(client, instanceID)
}

// Delete will permanently remove a user from a specified database instance.
func Delete(client *gophercloud.ServiceClient, instanceID, userName string) os.DeleteResult {
	return os.Delete(client, instanceID, userName)
}
