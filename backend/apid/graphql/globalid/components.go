package globalid

import (
	"errors"
	"strings"
)

//
// IDs
//

// Components describes the components of a global identifier.
//
// When represented as a string the ID appears in the follwoing format, parens
// denote optional components.
//
//   srn:(?@cluster:)resource:(?namespace:)(?resourceType/)uniqueComponents
//
// Example global IDs
//
//   srn:entities:default:selene.local
//   srn:events:sales:check/aG93ZHkgYnVkCg==
//   srn:checks:auto:disk-full
//   srn:users:deanlearner
//   srn:@us-west-2b:checks:parts:disk-chk
//   srn:@local:mutators:sales:jsonify
//
type Components interface {
	// Resource definition associated with this ID.
	Resource() string

	// Namespace is the name of the namespace the resource belongs to.
	Namespace() string

	// ResourceType is a optional element that describes any sort of sub-type
	// of the resource.
	ResourceType() string

	// Cluster is an optional element that describes which cluster the resource
	// may belong to.
	Cluster() string

	// UniqueComponent is a string that uniquely identify a resource; often
	// times this is the resource's name.
	UniqueComponent() string

	// String return string representation of ID
	String() string
}

// StandardComponents describes the standard components of a global identifier.
type StandardComponents struct {
	cluster         string
	resource        string
	namespace       string
	resourceType    string
	uniqueComponent string
}

// String returns the string representation of the global ID.
func (id StandardComponents) String() string {
	cluster := id.cluster
	if cluster != "" {
		cluster = "@" + cluster
	}

	nameComponents := append([]string{id.resourceType}, id.uniqueComponent)
	nameComponents = omitEmpty(nameComponents)
	pathComponents := omitEmpty([]string{
		cluster,
		id.resource,
		id.namespace,
	})

	// srn:{pathComponents}:{nameComponents}
	return "srn:" + strings.Join(pathComponents, ":") +
		":" + strings.Join(nameComponents, "/")
}

// Cluster associated with the resource.
func (id StandardComponents) Cluster() string {
	return id.cluster
}

// Resource definition associated with this ID.
func (id StandardComponents) Resource() string {
	return id.resource
}

// Namespace is the name of the namespace the resource belongs to.
func (id StandardComponents) Namespace() string {
	return id.namespace
}

// ResourceType is a optional element that describes any sort of sub-type of
// the resource.
func (id StandardComponents) ResourceType() string {
	return id.resourceType
}

// UniqueComponent is a string that uniquely identify a resource; often times
// this is the resource's name.
func (id StandardComponents) UniqueComponent() string {
	return id.uniqueComponent
}

// Parse takes a global ID string, decodes it and returns it's components.
func Parse(gid string) (StandardComponents, error) {
	id := StandardComponents{}
	pathComponents := strings.SplitN(gid, ":", 4)

	// Should be at least srn:resource:name
	if len(pathComponents) < 3 {
		return id, errors.New("given global ID does not appear valid")
	}

	if pathComponents[0] != "srn" {
		return id, errors.New("given string does not appear to be a Sensu global ID")
	}
	pathComponents = pathComponents[1:]

	// If present, pop the cluster from the path components, eg. @cluster:resource:ns:type/name
	//                                                           ^^^^^^^^
	cluster := pathComponents[0]
	if cluster[:1] == "@" && len(cluster) > 1 {
		id.cluster = cluster[1:]
		pathComponents = pathComponents[1:]
		if len(pathComponents) > 1 {
			nameComponents := strings.SplitN(pathComponents[1], ":", 2)
			pathComponents = append(pathComponents[:1], nameComponents...)
		}
		if len(pathComponents) < 2 {
			return id, errors.New("given global ID does not appear to contain a name component")
		}
	}

	// Pop the resource from the path components, eg. resource:ns:type/name
	//                                                ^^^^^^^^
	id.resource = pathComponents[0]
	pathComponents = pathComponents[1:]

	// Pop the name components from the path components, eg. ns:type/name
	//                                                          ^^^^^^^^^
	nameComponents := strings.Split(pathComponents[len(pathComponents)-1], "/")
	pathComponents = pathComponents[0 : len(pathComponents)-1]

	// If present pop the ns from the path components, eg. ns
	//                                                     ^^
	if len(pathComponents) > 0 {
		id.namespace = pathComponents[0]
	}

	// If present pop the type from the name components, eg. type/my-great-check
	//                                                       ^^^^
	if len(nameComponents) > 1 {
		id.resourceType = nameComponents[0]
		nameComponents = nameComponents[1:]
	}

	// Pop the remaining element from the name components, eg. my-great-check
	//                                                         ^^^^^^^^^^^^^^
	id.uniqueComponent = nameComponents[0]

	return id, nil
}

func omitEmpty(in []string) (out []string) {
	for _, n := range in {
		if n != "" {
			out = append(out, n)
		}
	}

	return
}
