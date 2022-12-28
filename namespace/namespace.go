package namespace

import "fmt"

const (
	DefaultNamespace = "global"

	HTTPHeaderKey = "X-EDOBTC-NS"
)

type NamespaceContext string

func ValueOrDefault() string {
	return DefaultNamespace
}

func OptionallyPrefix(ns string, id string) string {
	if ns == "" {
		return id
	}

	return fmt.Sprintf("%s|%s", ns, id)
}
