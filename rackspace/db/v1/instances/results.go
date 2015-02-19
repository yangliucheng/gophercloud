package instances

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
	os "github.com/rackspace/gophercloud/openstack/db/v1/instances"
	"github.com/rackspace/gophercloud/pagination"
	"github.com/rackspace/gophercloud/rackspace/db/v1/datastores"
)

// Instance represents a remote MySQL instance.
type Instance struct {
	// Indicates the datetime that the instance was created
	Created string //time.Time

	// Indicates the most recent datetime that the instance was updated.
	Updated string //time.Time

	// Indicates how the instance stores data.
	Datastore datastores.DatastorePartial

	// Indicates the hardware flavor the instance uses.
	Flavor os.Flavor

	// A DNS-resolvable hostname associated with the database instance (rather
	// than an IPv4 address). Since the hostname always resolves to the correct
	// IP address of the database instance, this relieves the user from the task
	// of maintaining the mapping. Note that although the IP address may likely
	// change on resizing, migrating, and so forth, the hostname always resolves
	// to the correct database instance.
	Hostname string

	// Indicates the unique identifier for the instance resource.
	ID string

	// Exposes various links that reference the instance resource.
	Links []gophercloud.Link

	// The human-readable name of the instance.
	Name string

	// The build status of the instance.
	Status string

	// Information about the attached volume of the instance.
	Volume os.Volume

	// IP indicates the various IP addresses which allow access.
	IP []string

	// Indicates whether this instance is a replica of another source instance.
	ReplicaOf *Instance `mapstructure:"replica_of" json:"replica_of"`

	// Indicates whether this instance is the source of other replica instances.
	Replicas []Instance
}

func commonExtract(err error, body interface{}) (*Instance, error) {
	if err != nil {
		return nil, err
	}

	var response struct {
		Instance Instance `mapstructure:"instance"`
	}

	err = mapstructure.Decode(body, &response)
	return &response.Instance, err
}

// CreateResult represents the result of a Create operation.
type CreateResult struct {
	os.CreateResult
}

// Extract will retrieve an instance from a create result.
func (r CreateResult) Extract() (*Instance, error) {
	return commonExtract(r.Err, r.Body)
}

// GetResult represents the result of a Get operation.
type GetResult struct {
	os.GetResult
}

// Extract will extract an Instance from a GetResult.
func (r GetResult) Extract() (*Instance, error) {
	return commonExtract(r.Err, r.Body)
}

// ConfigResult represents the result of getting default configuration for an
// instance.
type ConfigResult struct {
	gophercloud.Result
}

// DetachResult represents the result of detaching a replica from its source.
type DetachResult struct {
	gophercloud.ErrResult
}

// Extract will extract the configuration information (in the form of a map)
// about a particular instance.
func (r ConfigResult) Extract() (map[string]string, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var response struct {
		Instance struct {
			Config map[string]string `mapstructure:"configuration"`
		} `mapstructure:"instance"`
	}

	err := mapstructure.Decode(r.Body, &response)
	return response.Instance.Config, err
}

// UpdateResult represents the result of an Update operation.
type UpdateResult struct {
	gophercloud.ErrResult
}

// ExtractInstances retrieves a slice of instances from a paginated collection.
func ExtractInstances(page pagination.Page) ([]Instance, error) {
	casted := page.(os.InstancePage).Body

	var response struct {
		Instances []Instance `mapstructure:"instances"`
	}

	err := mapstructure.Decode(casted, &response)
	return response.Instances, err
}
