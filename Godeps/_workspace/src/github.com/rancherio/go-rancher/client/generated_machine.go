package client

const (
	MACHINE_TYPE = "machine"
)

type Machine struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    Amazonec2Config *Amazonec2Config `json:"amazonec2Config,omitempty"`
    
    AuthCertificateAuthority string `json:"authCertificateAuthority,omitempty"`
    
    AuthKey string `json:"authKey,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    DigitaloceanConfig *DigitaloceanConfig `json:"digitaloceanConfig,omitempty"`
    
    Driver string `json:"driver,omitempty"`
    
    ExoscaleConfig *ExoscaleConfig `json:"exoscaleConfig,omitempty"`
    
    ExternalId string `json:"externalId,omitempty"`
    
    ExtractedConfig string `json:"extractedConfig,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    OpenstackConfig *OpenstackConfig `json:"openstackConfig,omitempty"`
    
    PacketConfig *PacketConfig `json:"packetConfig,omitempty"`
    
    RackspaceConfig *RackspaceConfig `json:"rackspaceConfig,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    SoftlayerConfig *SoftlayerConfig `json:"softlayerConfig,omitempty"`
    
    State string `json:"state,omitempty"`
    
    Transitioning string `json:"transitioning,omitempty"`
    
    TransitioningMessage string `json:"transitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"transitioningProgress,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
    VirtualboxConfig *VirtualboxConfig `json:"virtualboxConfig,omitempty"`
    
    VmwarevcloudairConfig *VmwarevcloudairConfig `json:"vmwarevcloudairConfig,omitempty"`
    
    VmwarevsphereConfig *VmwarevsphereConfig `json:"vmwarevsphereConfig,omitempty"`
    
}

type MachineCollection struct {
	Collection
	Data []Machine `json:"data,omitempty"`
}

type MachineClient struct {
	rancherClient *RancherClient
}

type MachineOperations interface {
	List(opts *ListOpts) (*MachineCollection, error)
	Create(opts *Machine) (*Machine, error)
	Update(existing *Machine, updates interface{}) (*Machine, error)
	ById(id string) (*Machine, error)
	Delete(container *Machine) error
    
    ActionBootstrap (*Machine) (*PhysicalHost, error)
    
    
    ActionCreate (*Machine) (*PhysicalHost, error)
    
    
    ActionError (*Machine) (*PhysicalHost, error)
    
    
    ActionRemove (*Machine) (*PhysicalHost, error)
    
    
    ActionUpdate (*Machine) (*PhysicalHost, error)
    
}

func newMachineClient(rancherClient *RancherClient) *MachineClient {
	return &MachineClient{
		rancherClient: rancherClient,
	}
}

func (c *MachineClient) Create(container *Machine) (*Machine, error) {
	resp := &Machine{}
	err := c.rancherClient.doCreate(MACHINE_TYPE, container, resp)
	return resp, err
}

func (c *MachineClient) Update(existing *Machine, updates interface{}) (*Machine, error) {
	resp := &Machine{}
	err := c.rancherClient.doUpdate(MACHINE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *MachineClient) List(opts *ListOpts) (*MachineCollection, error) {
	resp := &MachineCollection{}
	err := c.rancherClient.doList(MACHINE_TYPE, opts, resp)
	return resp, err
}

func (c *MachineClient) ById(id string) (*Machine, error) {
	resp := &Machine{}
	err := c.rancherClient.doById(MACHINE_TYPE, id, resp)
	return resp, err
}

func (c *MachineClient) Delete(container *Machine) error {
	return c.rancherClient.doResourceDelete(MACHINE_TYPE, &container.Resource)
}
    
func (c *MachineClient) ActionBootstrap (resource *Machine) (*PhysicalHost, error) {
    
	resp := &PhysicalHost{}
    
	err := c.rancherClient.doAction(MACHINE_TYPE, "bootstrap", &resource.Resource, nil, resp)
    
	return resp, err
}
    
func (c *MachineClient) ActionCreate (resource *Machine) (*PhysicalHost, error) {
    
	resp := &PhysicalHost{}
    
	err := c.rancherClient.doAction(MACHINE_TYPE, "create", &resource.Resource, nil, resp)
    
	return resp, err
}
    
func (c *MachineClient) ActionError (resource *Machine) (*PhysicalHost, error) {
    
	resp := &PhysicalHost{}
    
	err := c.rancherClient.doAction(MACHINE_TYPE, "error", &resource.Resource, nil, resp)
    
	return resp, err
}
    
func (c *MachineClient) ActionRemove (resource *Machine) (*PhysicalHost, error) {
    
	resp := &PhysicalHost{}
    
	err := c.rancherClient.doAction(MACHINE_TYPE, "remove", &resource.Resource, nil, resp)
    
	return resp, err
}
    
func (c *MachineClient) ActionUpdate (resource *Machine) (*PhysicalHost, error) {
    
	resp := &PhysicalHost{}
    
	err := c.rancherClient.doAction(MACHINE_TYPE, "update", &resource.Resource, nil, resp)
    
	return resp, err
}
