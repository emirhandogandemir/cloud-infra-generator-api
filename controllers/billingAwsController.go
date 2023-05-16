package controllers

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func GetBillingAwsHandler(c *gin.Context) {
    var params models.BillingParamsAws

	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if params.Start == "" {
		params.Start = time.Now().AddDate(-1, 0, 0).Format("2006-01-02")
	}

	if params.End == "" {
		params.End = time.Now().Format("2006-01-02")
	}

	if params.Region ==""{
		params.Region= "us-east-1"
	}

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	cfg.Region="us-east-1"
	client := costexplorer.NewFromConfig(cfg)

	resp, err := client.GetCostAndUsage(context.TODO(), &costexplorer.GetCostAndUsageInput{
		Granularity: types.GranularityMonthly,
		Metrics:     []string{"UnblendedCost"},
		TimePeriod: &types.DateInterval{
			Start: aws.String(params.Start),
			End:   aws.String(params.End),
		},
	})
	if err != nil {
		log.Fatalf("unable to get cost and usage, %v", err)
	}

	c.JSON(http.StatusOK, resp.ResultsByTime)
}