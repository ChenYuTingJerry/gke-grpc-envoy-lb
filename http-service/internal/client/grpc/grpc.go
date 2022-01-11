package grpc

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"

	"github.com/jerry-yt-chen/gke-grpc-envoy-lb/echo-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/jerry-yt-chen/gke-grpc-envoy-lb/http-service/configs"
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile("cert/cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Create the credentials and return it
	config := &tls.Config{
		ClientAuth:         tls.RequireAndVerifyClientCert,
		ClientCAs:          certPool,
		InsecureSkipVerify: true,
	}

	return credentials.NewTLS(config), nil
}

func NewEchoGrpcClient() (client proto.EchoServiceClient, close func(), err error) {
	//tlsCredentials, err := loadTLSCredentials()
	//if err != nil {
	//	logrus.Fatal("cannot load TLS credentials: ", err)
	//}
	//opts := grpc.WithTransportCredentials(tlsCredentials)
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	grpcAddr := configs.C.Component.Echo.Host
	grpcAddr = fmt.Sprintf("%s:%s", configs.C.Component.Echo.Host, configs.C.Component.Echo.Port)
	conn, err := grpc.Dial(grpcAddr, opts)
	if err != nil {
		return nil, func() {}, err
	}

	return proto.NewEchoServiceClient(conn), func() {
		conn.Close()
	}, nil
}
