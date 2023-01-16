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

var AllDropletSizes = map[string]bool{
	"s-1vcpu-512mb-10gb": true,
	"s-1vcpu-1gb":        true,
	"s-1vcpu-1gb-amd":    true,
	"s-1vcpu-1gb-intel":  true,
	"s-1vcpu-2gb":        true,
	"s-1vcpu-2gb-amd":    true,
	"s-1vcpu-2gb-intel":  true,
	"s-2vcpu-2gb":        true,
	"s-2vcpu-2gb-amd":    true,
	"s-2vcpu-2gb-intel":  true,
	"s-2vcpu-4gb":        true,
	"s-2vcpu-4gb-amd":    true,
	"s-2vcpu-4gb-intel":  true,
	"c-2":                true,
	"c2-2vcpu-4gb":       true,
	"s-4vcpu-8gb":        true,
	"s-4vcpu-8gb-amd":    true,
	"s-4vcpu-8gb-intel":  true,
	"g-2vcpu-8gb":        true,
	"gd-2vcpu-8gb":       true,
	"m-2vcpu-16gb":       true,
	"c-4":                true,
	"c2-4vcpu-8gb":       true,
	"s-8vcpu-16gb":       true,
	"m3-2vcpu-16gb":      true,
	"s-8vcpu-16gb-amd":   true,
	"s-8vcpu-16gb-intel": true,
	"g-4vcpu-16gb":       true,
	"so-2vcpu-16gb":      true,
	"m6-2vcpu-16gb":      true,
	"gd-4vcpu-16gb":      true,
	"so1_5-2vcpu-16gb":   true,
	"m-4vcpu-32gb":       true,
	"c-8":                true,
	"c2-8vcpu-16gb":      true,
	"m3-4vcpu-32gb":      true,
	"g-8vcpu-32gb":       true,
	"so-4vcpu-32gb":      true,
	"m6-4vcpu-32gb":      true,
	"gd-8vcpu-32gb":      true,
	"so1_5-4vcpu-32gb":   true,
	"m-8vcpu-64gb":       true,
	"c-16":               true,
	"c2-16vcpu-32gb":     true,
	"m3-8vcpu-64gb":      true,
	"g-16vcpu-64gb":      true,
	"so-8vcpu-64gb":      true,
	"m6-8vcpu-64gb":      true,
	"gd-16vcpu-64gb":     true,
	"so1_5-8vcpu-64gb":   true,
	"m-16vcpu-128gb":     true,
	"c-32":               true,
	"c2-32vcpu-64gb":     true,
	"m3-16vcpu-128gb":    true,
	"c-48":               true,
	"m-24vcpu-192gb":     true,
	"g-32vcpu-128gb":     true,
	"so-16vcpu-128gb":    true,
	"m6-16vcpu-128gb":    true,
	"gd-32vcpu-128gb":    true,
	"c2-48vcpu-96gb":     true,
	"m3-24vcpu-192gb":    true,
	"g-40vcpu-160gb":     true,
	"so1_5-16vcpu-128gb": true,
	"m-32vcpu-256gb":     true,
	"gd-40vcpu-160gb":    true,
	"so-24vcpu-192gb":    true,
	"m6-24vcpu-192gb":    true,
	"m3-32vcpu-256gb":    true,
	"so1_5-24vcpu-192gb": true,
	"so-32vcpu-256gb":    true,
	"m6-32vcpu-256gb":    true,
	"so1_5-32vcpu-256gb": true,
}

// droplet size mappings
var DropletSizes = map[string]string{
	"default": "s-1vcpu-2gb",
	"small":   "s-1vcpu-2gb",
	"medium":  "s-1vcpu-2gb",
	"large":   "s-1vcpu-2gb",
}

func GetRegionName(region string) string {
	return Regions[region]
}

func ValidRegion(region string) bool {
	_, ok := Regions[region]
	return ok
}

func GetDropletSize(size string) string {
	return DropletSizes[size]
}

func ValidDropletSize(size string) bool {
	_, ok := DropletSizes[size]
	return ok
}
