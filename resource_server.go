package main

import (
    "github.com/hashicorp/terraform/helper/schema"
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/ecs"
    "github.com/aws/aws-sdk-go/aws/session"
    "strconv"
)

func resourceServer() *schema.Resource {
    return &schema.Resource{
        Create: resourceServerCreate,
        Read:   resourceServerRead,
        Update: resourceServerUpdate,
        Delete: resourceServerDelete,

        Schema: map[string]*schema.Schema{
            "cluster": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "service": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "taskDefinition": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "min": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "50",
            },

        },
    
    }
}

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
    //sess := session.Must(session.NewSession())
    sess := session.Must(session.NewSession(&aws.Config{
    Region: aws.String("us-west-2"),
    }))
    svc := ecs.New(sess)
   
   //Get variables
    service := d.Get("service").(string)
    cluster := d.Get("cluster").(string)
    min := d.Get("min").(string)
    i64, err:= strconv.ParseInt(min, 10, 64)


    params := &ecs.UpdateServiceInput{
        Service: aws.String(service), // Required
        Cluster: aws.String(cluster),
        DeploymentConfiguration: &ecs.DeploymentConfiguration{
            MaximumPercent:        aws.Int64(200),
            MinimumHealthyPercent: aws.Int64(i64),
        },
        //DesiredCount:   aws.Int64(1),
        //TaskDefinition: aws.String("String"),
    }
    resp, err := svc.UpdateService(params)

    if err != nil {
        // Print the error, cast err to awserr.Error to get the Code and
        // Message from an error.
        fmt.Println(err.Error())
        return nil
    }

    // Pretty-print the response data.
    fmt.Println(resp)   
    return nil
}
func resourceServerRead(d *schema.ResourceData, m interface{}) error {
    return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
    return nil
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
    return nil
}

