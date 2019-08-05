package machine

// Machine defines a machine struct from a provider.
type Machine struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Status   string   `json:"status"`
	SizeSlug string   `json:"size_slug"`
	Created  string   `json:"created_at"`
	Region   string   `json:"region"`
	Tags     []string `json:"tags"`

	Provider string `json:"provider"`
}

// MachineCreated defines the response for creating a machine.
type MachineCreated struct {
	ID       int    `json:"id"`
	Provider string `json:"provider"`
}

// MachineDeleted defines the response for deleted a machine.
type MachineDeleted struct {
	ID       int    `json:"id"`
	Provider string `json:"provider"`
}
