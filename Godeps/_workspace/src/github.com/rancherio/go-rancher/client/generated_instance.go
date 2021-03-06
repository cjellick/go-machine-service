package client

const (
	INSTANCE_TYPE = "instance"
)

type Instance struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    ExternalId string `json:"externalId,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    State string `json:"state,omitempty"`
    
    Transitioning string `json:"transitioning,omitempty"`
    
    TransitioningMessage string `json:"transitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"transitioningProgress,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
}

type InstanceCollection struct {
	Collection
	Data []Instance `json:"data,omitempty"`
}

type InstanceClient struct {
	rancherClient *RancherClient
}

type InstanceOperations interface {
	List(opts *ListOpts) (*InstanceCollection, error)
	Create(opts *Instance) (*Instance, error)
	Update(existing *Instance, updates interface{}) (*Instance, error)
	ById(id string) (*Instance, error)
	Delete(container *Instance) error
    
    ActionAllocate (*Instance) (*Instance, error)
    
    
    ActionConsole (*Instance, *InstanceConsoleInput) (*InstanceConsole, error)
    
    
    ActionCreate (*Instance) (*Instance, error)
    
    
    ActionDeallocate (*Instance) (*Instance, error)
    
    
    ActionMigrate (*Instance) (*Instance, error)
    
    
    ActionPurge (*Instance) (*Instance, error)
    
    
    ActionRemove (*Instance) (*Instance, error)
    
    
    ActionRestart (*Instance) (*Instance, error)
    
    
    ActionRestore (*Instance) (*Instance, error)
    
    
    ActionStart (*Instance) (*Instance, error)
    
    
    ActionStop (*Instance, *InstanceStop) (*Instance, error)
    
    
    ActionUpdate (*Instance) (*Instance, error)
    
}

func newInstanceClient(rancherClient *RancherClient) *InstanceClient {
	return &InstanceClient{
		rancherClient: rancherClient,
	}
}

func (c *InstanceClient) Create(container *Instance) (*Instance, error) {
	resp := &Instance{}
	err := c.rancherClient.doCreate(INSTANCE_TYPE, container, resp)
	return resp, err
}

func (c *InstanceClient) Update(existing *Instance, updates interface{}) (*Instance, error) {
	resp := &Instance{}
	err := c.rancherClient.doUpdate(INSTANCE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *InstanceClient) List(opts *ListOpts) (*InstanceCollection, error) {
	resp := &InstanceCollection{}
	err := c.rancherClient.doList(INSTANCE_TYPE, opts, resp)
	return resp, err
}

func (c *InstanceClient) ById(id string) (*Instance, error) {
	resp := &Instance{}
	err := c.rancherClient.doById(INSTANCE_TYPE, id, resp)
	return resp, err
}

func (c *InstanceClient) Delete(container *Instance) error {
	return c.rancherClient.doResourceDelete(INSTANCE_TYPE, &container.Resource)
}
    
func (c *InstanceClient) ActionAllocate (resource *Instance) (*Instance, error) {
    
	resp := &Instance{}
    
	err := c.rancherClient.doAction(INSTANCE_TYPE, "allocate", &resource.Resource, nil, resp)
    
	return resp, err
}
    
func (c *InstanceClient) ActionConsole (resource *Instance, input *InstanceConsoleInput) (*InstanceConsole, error) {
    
	resp := &InstanceConsole{}
    
	err := c.rancherClient.doAction(INSTANCE_TYPE, "console", &resource.Resource, input, resp)
    
	return resp, err
}
    
func (c *InstanceClient) ActionCreate (resource *Instance) (*Instance, error) {
    
	resp := &Instance{}
    
	err := c.rancherClient.doAction(INSTANCE_TYPE, "create", &resource.Resource, nil, resp)
    
	return resp, err
}
    
func (c *InstanceClient) ActionDeallocate (resource *Instance) (*Instance, error) {
    
	resp := &Instance{}
    
	err := c.rancherClient.doAction(INSTANCE_TYPE, "deallocate", &resource.Resource, nil, resp)
    
	return resp, err
}
    
func (c *InstanceClient) ActionMigrate (resource *Instance) (*Instance, error) {
    
	resp := &Instance{}
    
	err := c.rancherClient.doAction(INSTANCE_TYPE, "migrate", &resource.Resource, nil, resp)
    
	return resp, err
}
    
func (c *InstanceClient) ActionPurge (resource *Instance) (*Instance, error) {
    
	resp := &Instance{}
    
	err := c.rancherClient.doAction(INSTANCE_TYPE, "purge", &resource.Resource, nil, resp)
    
	return resp, err
}
    
func (c *InstanceClient) ActionRemove (resource *Instance) (*Instance, error) {
    
	resp := &Instance{}
    
	err := c.rancherClient.doAction(INSTANCE_TYPE, "remove", &resource.Resource, nil, resp)
    
	return resp, err
}
    
func (c *InstanceClient) ActionRestart (resource *Instance) (*Instance, error) {
    
	resp := &Instance{}
    
	err := c.rancherClient.doAction(INSTANCE_TYPE, "restart", &resource.Resource, nil, resp)
    
	return resp, err
}
    
func (c *InstanceClient) ActionRestore (resource *Instance) (*Instance, error) {
    
	resp := &Instance{}
    
	err := c.rancherClient.doAction(INSTANCE_TYPE, "restore", &resource.Resource, nil, resp)
    
	return resp, err
}
    
func (c *InstanceClient) ActionStart (resource *Instance) (*Instance, error) {
    
	resp := &Instance{}
    
	err := c.rancherClient.doAction(INSTANCE_TYPE, "start", &resource.Resource, nil, resp)
    
	return resp, err
}
    
func (c *InstanceClient) ActionStop (resource *Instance, input *InstanceStop) (*Instance, error) {
    
	resp := &Instance{}
    
	err := c.rancherClient.doAction(INSTANCE_TYPE, "stop", &resource.Resource, input, resp)
    
	return resp, err
}
    
func (c *InstanceClient) ActionUpdate (resource *Instance) (*Instance, error) {
    
	resp := &Instance{}
    
	err := c.rancherClient.doAction(INSTANCE_TYPE, "update", &resource.Resource, nil, resp)
    
	return resp, err
}
