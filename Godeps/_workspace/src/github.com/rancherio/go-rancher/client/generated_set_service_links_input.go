package client

const (
	SET_SERVICE_LINKS_INPUT_TYPE = "setServiceLinksInput"
)

type SetServiceLinksInput struct {
	Resource
    
    ServiceIds []string `json:"serviceIds,omitempty"`
    
}

type SetServiceLinksInputCollection struct {
	Collection
	Data []SetServiceLinksInput `json:"data,omitempty"`
}

type SetServiceLinksInputClient struct {
	rancherClient *RancherClient
}

type SetServiceLinksInputOperations interface {
	List(opts *ListOpts) (*SetServiceLinksInputCollection, error)
	Create(opts *SetServiceLinksInput) (*SetServiceLinksInput, error)
	Update(existing *SetServiceLinksInput, updates interface{}) (*SetServiceLinksInput, error)
	ById(id string) (*SetServiceLinksInput, error)
	Delete(container *SetServiceLinksInput) error
}

func newSetServiceLinksInputClient(rancherClient *RancherClient) *SetServiceLinksInputClient {
	return &SetServiceLinksInputClient{
		rancherClient: rancherClient,
	}
}

func (c *SetServiceLinksInputClient) Create(container *SetServiceLinksInput) (*SetServiceLinksInput, error) {
	resp := &SetServiceLinksInput{}
	err := c.rancherClient.doCreate(SET_SERVICE_LINKS_INPUT_TYPE, container, resp)
	return resp, err
}

func (c *SetServiceLinksInputClient) Update(existing *SetServiceLinksInput, updates interface{}) (*SetServiceLinksInput, error) {
	resp := &SetServiceLinksInput{}
	err := c.rancherClient.doUpdate(SET_SERVICE_LINKS_INPUT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *SetServiceLinksInputClient) List(opts *ListOpts) (*SetServiceLinksInputCollection, error) {
	resp := &SetServiceLinksInputCollection{}
	err := c.rancherClient.doList(SET_SERVICE_LINKS_INPUT_TYPE, opts, resp)
	return resp, err
}

func (c *SetServiceLinksInputClient) ById(id string) (*SetServiceLinksInput, error) {
	resp := &SetServiceLinksInput{}
	err := c.rancherClient.doById(SET_SERVICE_LINKS_INPUT_TYPE, id, resp)
	return resp, err
}

func (c *SetServiceLinksInputClient) Delete(container *SetServiceLinksInput) error {
	return c.rancherClient.doResourceDelete(SET_SERVICE_LINKS_INPUT_TYPE, &container.Resource)
}
