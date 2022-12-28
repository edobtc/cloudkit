package agent

// Agent exists as a temporary provisioner
// that can place install itself on nodes during the provisioning
// process  and conduct some actions to ease provisioning,
// such as
// - sending data about the node back to the control plane
// - fetching metadata
// - installing additional items
// - conducting audit/verification
