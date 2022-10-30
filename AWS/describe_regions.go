package main

import (
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/aws/awserr"

    "fmt"
)

func main() {
// Create a Session with a custom region
session, err := session.NewSession(&aws.Config{Region: aws.String("us-west-2"),})

    // Create new EC2 client
    Ec2Svc := ec2.New(session)
	input := &ec2.DescribeRegionsInput{}
	
	result, err := Ec2Svc.DescribeRegions(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	
	fmt.Println(result)

}