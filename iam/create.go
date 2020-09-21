package iam

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

func iamSession() *iam.IAM {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return iam.New(sess)
}

// CreateUser creates a new IAM user
func CreateUser(name string) error {

	svc := iamSession()

	_, err := svc.GetUser(&iam.GetUserInput{
		UserName: &name,
	})

	if awserr, ok := err.(awserr.Error); ok && awserr.Code() == iam.ErrCodeNoSuchEntityException {
		result, err := svc.CreateUser(&iam.CreateUserInput{
			UserName: &name,
		})

		if err != nil {
			fmt.Println("CreateUser Error", err)
			return err
		}

		fmt.Println("Success", result)
	} else {
		fmt.Println("GetUser Error", err)
	}
	return nil

}

// ListUsers returns a list of users in an account
func ListUsers() error {
	svc := iamSession()
	result, err := svc.ListUsers(&iam.ListUsersInput{
		MaxItems: aws.Int64(10),
	})

	if err != nil {
		fmt.Println("Error", err)
		return err
	}

	for i, user := range result.Users {
		if user == nil {
			continue
		}
		fmt.Printf("%d user %s created %v\n", i, *user.UserName, user.CreateDate)
	}
	return nil
}
