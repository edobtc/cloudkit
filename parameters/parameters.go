package parameters

import (
	"fmt"
	"sort"
	"strings"
)

// Parameter is a simple name value pair representing parameters
// that can act as declared inputs or units when designing an experiment.
// Either in defining the whitelist for an resource, or collecting input values
// of a single experiment request in flight that needs assignment
type Parameter struct {
	Name  string
	Value interface{}
}

// Parameters is a collection of parameters
type Parameters []Parameter

// String returns the sorted, name.value pair of all strings in the parameters
// list for inclusion as a seed string for hashing
func (p Parameters) String() string {
	var params []string

	sort.Slice(p, func(i, j int) bool { return p[i].Name < p[j].Name })

	for _, param := range p {
		paramString := fmt.Sprintf("%v.%v", param.Name, param.Value)

		params = append(params, paramString)
	}

	return strings.Join(params[:], ".")
}

// WhitelistedString returns a String() but with parameters not in the supplied whitelist
// excluded
func (p Parameters) WhitelistedString(whitelist Parameters) string {
	var params Parameters
	ws := map[string]int{}

	for _, wp := range whitelist {
		ws[wp.Name] = 0
	}

	for _, param := range p {
		if _, ok := ws[param.Name]; ok {
			params = append(params, param)
		}
	}

	return params.String()
}
