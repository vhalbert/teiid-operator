package constants

import "github.com/teiid/teiid-operator/pkg/apis/vdb/v1alpha1"

const (
	// ImageStreamNamespace default namespace for the ImageStreams
	ImageStreamNamespace = "openshift"
	// ImageRegistry ...
	ImageRegistry = "docker.io"
	// ImageRepo ...
	ImageRepo = "fabric8"
	// ImageStreamTag default tag name for the ImageStreams
	ImageStreamTag = "latest" // "latest-java11"
)

// RuntimeImageDefaults ...
var RuntimeImageDefaults = map[v1alpha1.RuntimeType][]v1alpha1.Image{
	v1alpha1.SpringbootRuntimeType: {
		{
			ImageStreamName:      "s2i-java",
			ImageStreamNamespace: ImageStreamNamespace,
			ImageStreamTag:       "latest-java11",
			ImageRegistry:        ImageRegistry,
			ImageRepository:      ImageRepo,
			BuilderImage:         true,
		},
		{
			BuilderImage: false,
		},
	},
}
