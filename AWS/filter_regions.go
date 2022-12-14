package main

import (

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"

	"fmt"
)
// describes existing instances that are running in EC2. 
func GetRunningInstances(client *ec2.EC2) (*ec2.DescribeInstancesOutput, error) {
	result, err := client.DescribeInstances(&ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("instance-state-name"),
				Values: []*string{
					aws.String("running"),
				},
			},
		},
	})

	if err != nil {
		return nil, err
	}

	return result, err
}


func main() {
	// Create a Session with a custom region
	session, err := session.NewSession(&aws.Config{Region: aws.String("us-west-2"),})

	if err != nil {
		fmt.Printf("Failed to initialize new session: %v", err)
		return
	}

	ec2Client := ec2.New(session)

	runningInstances, err := GetRunningInstances(ec2Client)
	if err != nil {
		fmt.Printf("Couldn't retrieve running instances: %v", err)
		return
	}

	for _, reservation := range runningInstances.Reservations {
		for _, instance := range reservation.Instances {
			fmt.Printf("Found running instance: %s\n", *instance.InstanceId)
		}
	}
}