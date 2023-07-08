package connect

import (
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"
)

func Connect() {
	fmt.Println("Starting connection...");
	// Create a cluster configuration
	cluster := gocql.NewCluster("3.94.80.110") // Replace with your Cassandra cluster IP addresses
	
	cluster.ProtoVersion = 4
	// Set the keyspace you want to connect to
	// cluster.Keyspace = "test1"
	// cluster.SslOpts = &gocql.SslOptions {
	// 	CaPath: "password-manager-db.pem",
	// }

	
    cluster.DisableInitialHostLookup = true
    // cluster.Authenticator = gocql.PasswordAuthenticator{Username: "test", Password: "testpwd"}
    cluster.Consistency = gocql.Quorum
    cluster.CQLVersion = "3.4.5"
    cluster.IgnorePeerAddr = true
    cluster.DefaultIdempotence = true
    cluster.Timeout = time.Second * 30
    cluster.ConnectTimeout = time.Second * 30

	// Set other optional configuration options

	// Create a session, which represents the connection to the cluster
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	// Use the session to execute queries and perform operations

	fmt.Println("Query executed successfully.")
}
