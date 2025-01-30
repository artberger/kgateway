//go:build ignore

package listener_options

import (
	"net/http"
	"path/filepath"

	"github.com/kgateway-dev/kgateway/pkg/utils/fsutils"
	"github.com/kgateway-dev/kgateway/test/gomega/matchers"
	e2edefaults "github.com/kgateway-dev/kgateway/test/kubernetes/e2e/defaults"
	"github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	setupManifests = []string{
		filepath.Join(fsutils.MustGetThisDir(), "testdata", "setup.yaml"),
		e2edefaults.CurlPodManifest,
	}
	basicLisOptManifest = filepath.Join(fsutils.MustGetThisDir(), "testdata", "basic-lisopt.yaml")

	// When we apply the setup file, we expect resources to be created with this metadata
	glooProxyObjectMeta = metav1.ObjectMeta{
		Name:      "gloo-proxy-gw",
		Namespace: "default",
	}
	proxyService    = &corev1.Service{ObjectMeta: glooProxyObjectMeta}
	proxyDeployment = &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "gloo-proxy-gw",
			Namespace: "default",
		},
	}
	nginxPod = &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "nginx",
			Namespace: "default",
		},
	}
	exampleSvc = &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "example-svc",
			Namespace: "default",
		},
	}

	expectedHealthyResponse = &matchers.HttpResponse{
		StatusCode: http.StatusOK,
		Body:       gomega.ContainSubstring("Welcome to nginx!"),
	}
)
