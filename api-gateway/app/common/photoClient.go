package common

import (
	"log"
	"fmt"
	"google.golang.org/grpc"
	"github.com/ktamashun/flickrphotosearch/api-gateway/apis"
)

var PhotoClient apis.PhotoClient

func init() {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(getConnectionString(), opts...)
	if err != nil {
		log.Fatal(err)
	}
	//defer conn.Close()

	PhotoClient = apis.NewPhotoClient(conn)
}

func getConnectionString() string {
	var connstr = fmt.Sprintf("%s:%s",
		AppConfig.BackendHost,
		AppConfig.BackendPort,
	)

	log.Println(connstr)

	return connstr
}
