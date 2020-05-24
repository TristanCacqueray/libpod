package entities

import "github.com/containers/image/v5/types"

// PlayKubeOptions controls playing kube YAML files.
type PlayKubeOptions struct {
	// Authfile - path to an authentication file.
	Authfile string
	// CertDir - to a directory containing TLS certifications and keys.
	CertDir string
	// Username for authenticating against the registry.
	Username string
	// Password for authenticating against the registry.
	Password string
	// Network - name of the CNI network to connect to.
	Network string
	// Quiet - suppress output when pulling images.
	Quiet bool
	// SignaturePolicy - path to a signature-policy file.
	SignaturePolicy string
	// SkipTLSVerify - skip https and certificate validation when
	// contacting container registries.
	SkipTLSVerify types.OptionalBool
	// SeccompProfileRoot - path to a directory containing seccomp
	// profiles.
	SeccompProfileRoot string
}

// PlayKubePods represents a single pod and associated containers created by play kube
type PlayKubePod struct {
	ID string
	// Containers - the IDs of the containers running in the created pod.
	Containers []string
	// Logs - non-fatal erros and log messages while processing.
	Logs []string
}

// PlayKubeReport contains the results of running play kube.
type PlayKubeReport struct {
	// Pods - pods created by play kube.
	Pods []PlayKubePod
}
