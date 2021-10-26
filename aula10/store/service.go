package store

import "sync"

type Service struct {
	Name, ContainerID  string
	Image, IPAddress   string
	Endpoint, Hostname string
	Labels             map[string]string
	Ports              []string
}

func NewService(name, cid, img, ip, endpoint, hostname string,
	labels map[string]string, ports []string) Service {
	return Service{
		name, cid, img, ip, endpoint, hostname, labels, ports,
	}
}

type ServiceStore struct {
	services map[string]Service
	sync.RWMutex
}

func (ss *ServiceStore) Add(name string, svc Service) {
	ss.Lock()
	defer ss.Unlock()

	ss.services[name] = svc
}

func (ss *ServiceStore) Del(name string) {
	ss.Lock()
	defer ss.Unlock()

	delete(ss.services, name)
}

func (ss *ServiceStore) List() []Service {
	ss.RLock()
	defer ss.RUnlock()

	svcs := make([]Service, 0)

	for _, svc := range ss.services {
		svcs = append(svcs, svc)
	}

	return svcs
}

func NewStore() ServiceStore {
	return ServiceStore{
		services: make(map[string]Service),
	}
}
