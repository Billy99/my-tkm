//go:build linux
// +build linux

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
*/

const (
	// DefaultPath = "/run/tkm/tkm.sock"
	CSI_SOCKET = "/run/tkm/csi/csi.sock"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	//ctx := context.Background()

	// connect to the TKM server - NOOP for now
	/*
		conn, err := createConnection(ctx)
		if err != nil {
			log.Printf("failed to create client connection: %v", err)
			return
		}
	*/

	defer func() {
		log.Printf("Closing Connection\n")
		//conn.Close()
	}()

	ticker := time.NewTicker(10 * time.Second)
	go func() {
		var count uint64
		for range ticker.C {
			count = count + 1
			log.Printf("Still Alive: %d\n", count)
		}
	}()

	<-stop

	log.Printf("Exiting...\n")
}

/*
func createConnection(ctx context.Context) (*grpc.ClientConn, error) {
	var (
		addr        string
		local_creds credentials.TransportCredentials
	)

	addr = fmt.Sprintf("unix://%s", DefaultPath)
	local_creds = insecure.NewCredentials()

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(local_creds))
	if err == nil {
		return conn, nil
	}
	log.Printf("did not connect: %v", err)

	return nil, fmt.Errorf("unable to establish connection")
}
*/
