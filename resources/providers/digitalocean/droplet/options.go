package droplet

// region mappings
var Regions = map[string]string{
	"nyc1": "New York 1",
	"nyc2": "New York 2",
	"nyc3": "New York 3",
	"ams1": "Amsterdam 1",
	"sfo1": "San Francisco 1",
	"sfo2": "San Francisco 2",
	"sfo3": "San Francisco 3",
}

// droplet size mappings
var DropletSizes = map[string]string{
	"default": "s-1vcpu-2gb",
	"small":   "s-1vcpu-2gb",
	"medium":  "s-1vcpu-2gb",
	"large":   "s-1vcpu-2gb",
}
