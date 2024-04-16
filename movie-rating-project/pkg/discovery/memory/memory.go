package memory

import (
	"context"
	"errors"
	"movieexample/pkg/discovery"
	"sync"
	"time"
)

// Registry defines an in-memory service registry
type Registry struct {
	sync.RWMutex
	serviceAddrs map[string]map[string]*serviceInstance
}

type serviceInstance struct {
	hostPort   string
	lastActive time.Time
}

// NewRegistry creates a new in-memory service registry instance
func NewRegistry() *Registry {
	return &Registry{serviceAddrs: map[string]map[string]*serviceInstance{}}
}

// Register creates a service record in the registry
func (r *Registry) Register(ctx context.Context, instanceID string, serviceName string, hostPort string) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.serviceAddrs[serviceName]; !ok {
		r.serviceAddrs[serviceName] = map[string]*serviceInstance{}
	}
	r.serviceAddrs[serviceName][instanceID] = &serviceInstance{hostPort: hostPort, lastActive: time.Now()}
	return nil
}

// Deregister removes a service record from the registry
func (r *Registry) Deregister(ctx context.Context, instanceID string, serviceName string) error {

	r.Lock()
	defer r.Unlock()

	if _, ok := r.serviceAddrs[serviceName]; !ok {
		return nil
	}
	delete(r.serviceAddrs[serviceName], instanceID)
	return nil
}

// ReportHealthyState is a push mechanism for reporting healthy state to registry
func (r *Registry) ReportHealthyState(instaceID string, serviceName string) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.serviceAddrs[serviceName]; !ok {
		return errors.New("service is not registered yet")
	}
	if _, ok := r.serviceAddrs[serviceName][instaceID]; !ok {
		return errors.New("service instance is not registered yet")
	}
	r.serviceAddrs[serviceName][instaceID].lastActive = time.Now()
	return nil
}

// ServiceAddresses return the list fo addresses of active instances of the given service
func (r *Registry) ServiceAddresses(ctx context.Context, serviceName string) ([]string, error) {
	r.Lock()
	defer r.Unlock()
	if len(r.serviceAddrs[serviceName]) == 0{
		return nil, discovery.ErrNotFound
	}
	var res []string
	for _, i := range r.serviceAddrs[serviceName] {
		if i.lastActive.Before(time.Now().Add(-5 * time.Second)){
			continue
		}
		res = append(res, i.hostPort)
	}
	return res, nil
}
