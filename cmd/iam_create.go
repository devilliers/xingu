package cmd

import (
	"github.com/dooferlad/xingu/iam"
	"github.com/spf13/cobra"
)

var name string

var iamCreateUserCmd = &cobra.Command{
	Use:   "create_user",
	Short: "Create IAM user",
	Args: func(cmd *cobra.Command, args []string) error {
		return defaultArgFromPositionS(args, &name)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return iam.CreateUser(name)
	},
}

func init() {
	iamCmd.AddCommand(iamCreateUserCmd)
	iamCreateUserCmd.Flags().StringToStringVar(&filter, "filters", nil, "ec2 instance filters")
	iamCreateUserCmd.Flags().StringVarP(&name, "name", "n", "", "ec2 instance name")
}
