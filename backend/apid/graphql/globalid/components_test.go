package globalid

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	type wants struct {
		cluster string
		res     string
		nsp     string
		resType string
		id      string
		err     bool
	}
	testCases := []struct {
		gid  string
		want wants
	}{
		{"users", wants{err: true}},
		{"users:cat/123", wants{err: true}},
		{"users:ns:cat/123", wants{err: true}},
		{"srn:users", wants{err: true}},
		{"srn:users:123", wants{res: "users", id: "123"}},
		{"srn:u:ns:1", wants{res: "u", nsp: "ns", id: "1"}},
		{"srn:users:cat/123", wants{res: "users", resType: "cat", id: "123"}},
		{"srn:x:y:*:z", wants{res: "x", nsp: "y", id: "*:z"}},
		{"srn:@us-west:x:y:z", wants{cluster: "us-west", res: "x", nsp: "y", id: "z"}},
		{"srn:@us-west:users:bob", wants{cluster: "us-west", res: "users", id: "bob"}},
		{"srn:@us-west:mutators", wants{err: true}},
		{"srn:@us-west:", wants{err: true}},
	}
	for _, tc := range testCases {
		t.Run(tc.gid, func(t *testing.T) {
			components, err := Parse(tc.gid)
			if tc.want.err {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tc.want.cluster, components.Cluster())
			assert.Equal(t, tc.want.res, components.Resource())
			assert.Equal(t, tc.want.nsp, components.Namespace())
			assert.Equal(t, tc.want.resType, components.ResourceType())
			assert.Equal(t, tc.want.id, components.UniqueComponent())
		})
	}
}

func TestStandardComponentsString(t *testing.T) {
	assert := assert.New(t)
	testCases := []struct {
		components StandardComponents
		want       string
	}{
		{
			components: StandardComponents{
				resource:        "users",
				uniqueComponent: "123",
			},
			want: "srn:users:123",
		},
		{
			StandardComponents{
				resource:        "users",
				resourceType:    "cat",
				uniqueComponent: "123",
			},
			"srn:users:cat/123",
		},
		{
			StandardComponents{
				resource:        "users",
				namespace:       "default",
				uniqueComponent: "123",
			},
			"srn:users:default:123",
		},
		{
			StandardComponents{
				resource:        "users",
				namespace:       "default",
				uniqueComponent: "123",
			},
			"srn:users:default:123",
		},
		{
			StandardComponents{
				resource:        "users",
				resourceType:    "cat",
				namespace:       "default",
				uniqueComponent: "123",
			},
			"srn:users:default:cat/123",
		},
		{
			StandardComponents{
				resource:        "silences",
				namespace:       "default",
				uniqueComponent: "123:456",
			},
			"srn:silences:default:123:456",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.want, func(t *testing.T) {
			gid := tc.components.String()
			assert.Equal(tc.want, gid)
		})
	}
}
