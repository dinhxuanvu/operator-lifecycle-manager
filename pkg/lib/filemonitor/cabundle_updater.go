package filemonitor

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"path/filepath"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
)

type cabundlestore struct {
	mutex        sync.RWMutex
	certpool     *x509.CertPool
	clientCAPath string
}

type getConfigFn = func(*tls.ClientHelloInfo) (*tls.Config, error)

func NewCABundleStore(clientCAPath string) *cabundlestore {
	pem, err := ioutil.ReadFile(clientCAPath)
	if err != nil {
		panic(err)
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(pem)

	return &cabundlestore{
		mutex:        sync.RWMutex{},
		certpool:     pool,
		clientCAPath: clientCAPath,
	}
}

func (c *cabundlestore) GetConfig(h *tls.ClientHelloInfo) (*tls.Config, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	cfg := tls.Config{
		ClientCAs:  c.certpool,
		ClientAuth: tls.VerifyClientCertIfGiven,
	}
	return &cfg, nil
}

func (c *cabundlestore) storeCABundle(clientCAPath string) error {
	pem, err := ioutil.ReadFile(clientCAPath)
	if err == nil {
		c.mutex.Lock()
		defer c.mutex.Unlock()
		pool := x509.NewCertPool()
		pool.AppendCertsFromPEM(pem)
		c.certpool = pool
	}
	return err
}

func (c *cabundlestore) HandleCABundleUpdate(logger logrus.FieldLogger, event fsnotify.Event) {
	switch op := event.Op; op {
	case fsnotify.Create:
		logger.Debugf("got fs event for %v", event.Name)

		if err := c.storeCABundle(c.clientCAPath); err != nil {
			logger.Debugf("unable to reload ca bundle: %v", err)
		} else {
			logger.Debugf("successfully reload ca bundle: %v", err)
		}
	}
}

func GetCABundleReloadFn(logger *logrus.Logger, clientCAPath string) (getConfigFn, error) {
	cabundlestore := NewCABundleStore(clientCAPath)
	watcher, err := NewWatch(logger, []string{filepath.Dir(clientCAPath)}, cabundlestore.HandleCABundleUpdate)
	if err != nil {
		return nil, err
	}
	watcher.Run(context.Background())

	return cabundlestore.GetConfig, nil
}
