package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mine/templ"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type DBSecret struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Engine   string `json:"engine"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DBName   string `json:"dbname"`
}

func getSecret() (*DBSecret, error) {
	secretName := os.Getenv("SECRET_NAME")
	if secretName == "" {
		secretName = "dbtestsecret" // fallback to hardcoded value if env var not set
	}
	region := os.Getenv("AWS_REGION")
	if region == "" {
		region = "ap-southeast-1" // fallback to hardcoded value if env var not set
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %v", err)
	}

	svc := secretsmanager.NewFromConfig(cfg)
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"),
	}

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to get secret value: %v", err)
	}

	var dbSecret DBSecret
	err = json.Unmarshal([]byte(*result.SecretString), &dbSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal secret: %v", err)
	}

	return &dbSecret, nil
}

func main() {
	dbSecret, err := getSecret()
	if err != nil {
		log.Fatalf("Error getting secret: %v", err)
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
		dbSecret.Host, dbSecret.Port, dbSecret.Username, dbSecret.Password, "postgres")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ.HelloWorld("connstr", connStr).Render(r.Context(), w)
	})

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
