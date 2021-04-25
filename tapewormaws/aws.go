package tapewormaws

import "log"
import "fmt"
import "github.com/aws/aws-sdk-go/aws"
import "github.com/aws/aws-sdk-go/aws/credentials"
import "github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
import "github.com/aws/aws-sdk-go/aws/ec2metadata"
import "github.com/aws/aws-sdk-go/aws/session"



func RunAWS() {

    newSession, err := session.NewSession()

    if err != nil {
        log.Fatal(err)
    }

    awsCredentials := credentials.NewChainCredentials(
        []credentials.Provider{
            &ec2rolecreds.EC2RoleProvider{
                Client: ec2metadata.New(newSession, aws.NewConfig()),
            },
            &credentials.SharedCredentialsProvider{},
            &credentials.EnvProvider{},
        })

	fmt.Println(awsCredentials)

}