package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "example.com/go-usermgmt-grpc/usermanagement"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var new_users = make(map[string]int32)
	new_users["Toni"] = 22
	new_users["Nikola"] = 25
	new_users["Sefer"] = 23

	for name, age := range new_users {
		r, err := c.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})
		if err != nil {
			log.Fatalf("Could not create: %v", err)
		}
		log.Printf(`User details:
		Name: %s
		Age: %d
		ID: %d`, r.GetName(), r.GetAge(), r.GetId())
	}
	params := &pb.GetUserParams{}
	r, err := c.GetUsers(ctx, params)
	if err != nil {
		log.Fatalf("Could not retrieve users: %v", err)
	}

	log.Print("\nUSER LIST:\n")
	fmt.Printf("r.GetUsers(): %v\n", r.GetUsers())
}
