package main

import (
    "fmt"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

func main() {
    // Create a session object to talk to S3 service
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-west-2"),
    })

    if err != nil {
        fmt.Println("Error creating session:", err)
        return
    }

    // Create a new instance of the service's client with a Session.
    svc := s3.New(sess)

    // Call S3 ListBuckets
    result, err := svc.ListBuckets(nil)
    if err != nil {
        fmt.Println("Error listing buckets:", err)
        return
    }

    fmt.Println("Buckets:")
    for _, b := range result.Buckets {
        fmt.Println(*b.Name)
    }
}
