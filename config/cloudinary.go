package config

import "github.com/cloudinary/cloudinary-go/v2"

type CloudinaryConfig struct {
	Name      string
	APIKey    string
	APISecret string
}

func ConnectCloudinary(cfg *Config) (*cloudinary.Cloudinary, error) {
	client, err := cloudinary.NewFromParams(cfg.Cloudinary.Name, cfg.Cloudinary.APIKey, cfg.Cloudinary.APISecret)
	if err != nil {
		return client, err
	}

	return client, nil
}
