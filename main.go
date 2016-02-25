package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/devicefarm"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

var last_seen map[string]string

func main() {
	last_seen = make(map[string]string)
	go updateDevices()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"devices": last_seen})
	})
	r.Run()
}

func repeatedlyUpdateDevices() {
	for range time.Tick(30 * time.Second) {
		updateDevices()
	}
}

func updateDevices() {
	log.Println("Devices [UPDATING] devices...")
	svc := devicefarm.New(session.New(), &aws.Config{Region: aws.String("us-west-2")})
	t := time.Now().String()

	params := &devicefarm.ListDevicesInput{}
	// 	Arn:       aws.String("AmazonResourceName"),
	// 	NextToken: aws.String("PaginationToken"),
	// }
	resp, err := svc.ListDevices(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	for _, device := range resp.Devices {
		last_seen[*device.Arn] = t
	}
	log.Println("Devices updated")
}
