package etcdv3

import (
	"google.golang.org/grpc/naming"
	"errors"
	etcd3 "github.com/coreos/etcd/clientv3"

	"strings"
	"fmt"
)

type resolver struct {
	serviceName string
}

func NewResolver(serviceName string) *resolver {
	return &resolver{serviceName: serviceName}
}

func (re *resolver) Resolve(target string) (naming.Watcher, error) {
	if re.serviceName == "" {
		return nil, errors.New("grpclb: no service name provided")
	}

	//etcd client
	client, err := etcd3.New(etcd3.Config{
		Endpoints: strings.Split(target, ","),
	})

	if err != nil {
		return nil, fmt.Errorf("grpclb: create etcd3 client failed: %s", err.Error())
	}
	return &watcher{re: re, client: *client}, nil
}
