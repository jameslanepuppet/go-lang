package main

import (
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"

    "fmt"
)

func main() {
// Create a Session with a custom region
session, err := session.NewSession(&aws.Config{ Region: aws.String("us-west-2"),})

    // Create new EC2 client
    ec2Svc := ec2.New(session)

    // Call to get detailed information on each instance
    result, err := ec2Svc.DescribeInstances(nil)
    if err != nil {
        fmt.Println("Error", err)
    } else {
        fmt.Println("Success", result)
	}

}