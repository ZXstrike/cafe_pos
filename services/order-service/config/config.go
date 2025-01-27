package config

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServePort  string
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
	MySQL      MySQLConfig
}

type MySQLConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

var Config *AppConfig

func LoadConfig() (*AppConfig, error) {
	_, err := os.Stat(".env")
	if err == nil {
		err := godotenv.Load()
		if err != nil {
			return nil, fmt.Errorf("error loading .env file: %w", err)
		}
	}

	privateKey, err := loadECDSAPrivateKey(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		return nil, fmt.Errorf("error loading private key: %w", err)
	}

	publicKey, err := loadECDSAPublicKey(os.Getenv("PUBLIC_KEY"))
	if err != nil {
		return nil, fmt.Errorf("error loading public key: %w", err)
	}

	Config = &AppConfig{
		ServePort:  os.Getenv("SERVE_PORT"),
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		MySQL: MySQLConfig{
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Database: os.Getenv("DB_DATABASE"),
		},
	}

	return Config, nil
}

func loadECDSAPrivateKey(key string) (*ecdsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(key))
	if block == nil {
		return nil, fmt.Errorf("error decoding private key")
	}

	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("error parsing private key: %w", err)
	}

	return privateKey, nil
}

func loadECDSAPublicKey(key string) (*ecdsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(key))
	if block == nil {
		return nil, fmt.Errorf("error decoding public key")
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("error parsing public key: %w", err)
	}

	return publicKey.(*ecdsa.PublicKey), nil
}
