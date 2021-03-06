package client

const (
	PROCESS_DEFINITION_TYPE = "processDefinition"
)

type ProcessDefinition struct {
	Resource
    
    ExtensionBased bool `json:"extensionBased,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    PostProcessListeners interface{} `json:"postProcessListeners,omitempty"`
    
    PreProcessListeners interface{} `json:"preProcessListeners,omitempty"`
    
    ProcessHandlers interface{} `json:"processHandlers,omitempty"`
    
    ResourceType string `json:"resourceType,omitempty"`
    
    StateTransitions []interface{} `json:"stateTransitions,omitempty"`
    
}

type ProcessDefinitionCollection struct {
	Collection
	Data []ProcessDefinition `json:"data,omitempty"`
}

type ProcessDefinitionClient struct {
	rancherClient *RancherClient
}

type ProcessDefinitionOperations interface {
	List(opts *ListOpts) (*ProcessDefinitionCollection, error)
	Create(opts *ProcessDefinition) (*ProcessDefinition, error)
	Update(existing *ProcessDefinition, updates interface{}) (*ProcessDefinition, error)
	ById(id string) (*ProcessDefinition, error)
	Delete(container *ProcessDefinition) error
}

func newProcessDefinitionClient(rancherClient *RancherClient) *ProcessDefinitionClient {
	return &ProcessDefinitionClient{
		rancherClient: rancherClient,
	}
}

func (c *ProcessDefinitionClient) Create(container *ProcessDefinition) (*ProcessDefinition, error) {
	resp := &ProcessDefinition{}
	err := c.rancherClient.doCreate(PROCESS_DEFINITION_TYPE, container, resp)
	return resp, err
}

func (c *ProcessDefinitionClient) Update(existing *ProcessDefinition, updates interface{}) (*ProcessDefinition, error) {
	resp := &ProcessDefinition{}
	err := c.rancherClient.doUpdate(PROCESS_DEFINITION_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ProcessDefinitionClient) List(opts *ListOpts) (*ProcessDefinitionCollection, error) {
	resp := &ProcessDefinitionCollection{}
	err := c.rancherClient.doList(PROCESS_DEFINITION_TYPE, opts, resp)
	return resp, err
}

func (c *ProcessDefinitionClient) ById(id string) (*ProcessDefinition, error) {
	resp := &ProcessDefinition{}
	err := c.rancherClient.doById(PROCESS_DEFINITION_TYPE, id, resp)
	return resp, err
}

func (c *ProcessDefinitionClient) Delete(container *ProcessDefinition) error {
	return c.rancherClient.doResourceDelete(PROCESS_DEFINITION_TYPE, &container.Resource)
}
