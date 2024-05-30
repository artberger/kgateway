package bootstrap

import (
	"context"
	"net"

	"github.com/solo-io/gloo/projects/gloo/pkg/debug"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/durationpb"
	"k8s.io/client-go/kubernetes"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	corecache "github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/cache"
	"github.com/solo-io/solo-kit/pkg/api/v1/control-plane/cache"
	"github.com/solo-io/solo-kit/pkg/api/v1/control-plane/server"
	skkube "github.com/solo-io/solo-kit/pkg/api/v1/resources/common/kubernetes"

	"github.com/solo-io/gloo/pkg/bootstrap/leaderelector"
	gwtranslator "github.com/solo-io/gloo/projects/gateway/pkg/translator"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/upstreams/consul"
	"github.com/solo-io/gloo/projects/gloo/pkg/validation"
)

type Opts struct {
	WriteNamespace               string
	StatusReporterNamespace      string
	WatchNamespaces              []string
	Upstreams                    factory.ResourceClientFactory
	KubeServiceClient            skkube.ServiceClient
	UpstreamGroups               factory.ResourceClientFactory
	Proxies                      factory.ResourceClientFactory
	Secrets                      factory.ResourceClientFactory
	Artifacts                    factory.ResourceClientFactory
	AuthConfigs                  factory.ResourceClientFactory
	RateLimitConfigs             factory.ResourceClientFactory
	GraphQLApis                  factory.ResourceClientFactory
	VirtualServices              factory.ResourceClientFactory
	RouteTables                  factory.ResourceClientFactory
	Gateways                     factory.ResourceClientFactory
	MatchableHttpGateways        factory.ResourceClientFactory
	MatchableTcpGateways         factory.ResourceClientFactory
	VirtualHostOptions           factory.ResourceClientFactory
	RouteOptions                 factory.ResourceClientFactory
	KubeClient                   kubernetes.Interface
	Consul                       Consul
	WatchOpts                    clients.WatchOpts
	DevMode                      bool
	ControlPlane                 ControlPlane
	ValidationServer             ValidationServer
	ProxyDebugServer             ProxyDebugServer
	Settings                     *v1.Settings
	KubeCoreCache                corecache.KubeCoreCache
	ValidationOpts               *gwtranslator.ValidationOpts
	ReadGatwaysFromAllNamespaces bool
	GatewayControllerEnabled     bool
	ProxyCleanup                 func()

	Identity leaderelector.Identity

	GlooGateway GlooGateway
}

type IstioValues struct {
	// Based on istioSDS helm values. Config used for installing Gloo Edge with Istio SDS cert rotation features for Istio mTLS
	SDSEnabled bool
	// Based on enableIstioSidecarOnGateway helm value. If true, the gateway will have the istio sidecar injected.
	SidecarOnGatewayEnabled bool
}

type GlooGateway struct {
	EnableK8sGatewayController bool
	IstioValues                IstioValues
}

type Consul struct {
	ConsulWatcher      consul.ConsulWatcher
	DnsServer          string
	DnsPollingInterval *durationpb.Duration
}

type ControlPlane struct {
	*GrpcService
	SnapshotCache cache.SnapshotCache
	XDSServer     server.Server

	Kube KubernetesControlPlaneConfig
}

// KubernetesControlPlaneConfig contains information about the control plane when
// running on Kubernetes.
type KubernetesControlPlaneConfig struct {
	// the address and port of the socket address that envoy will connect to
	// to receive xDS config from the gloo control plane
	XdsHost string
	XdsPort int32
}

// ValidationServer validates proxies generated by controllors outside the gloo pod
type ValidationServer struct {
	*GrpcService
	Server validation.ValidationServer
}

// ProxyDebugServer returns proxies to callers outside the gloo pod - this is only necessary for UI/debugging purposes.
type ProxyDebugServer struct {
	*GrpcService
	Server debug.ProxyEndpointServer
}

type GrpcService struct {
	Ctx             context.Context
	BindAddr        net.Addr
	GrpcServer      *grpc.Server
	StartGrpcServer bool
}

// GetBindAddress returns the string form of the BindAddr (for example, "192.0.2.1:25", "[2001:db8::1]:80")
func (g *GrpcService) GetBindAddress() string {
	if g == nil {
		return ""
	}
	return g.BindAddr.String()
}

// GetBindPort returns the port if the GrpcService relies on a TCPAddr, 0 otherwise
func (g *GrpcService) GetBindPort() int {
	if g == nil {
		return 0
	}
	tcpAddr, ok := g.BindAddr.(*net.TCPAddr)
	if !ok {
		return 0
	}
	return tcpAddr.Port
}
