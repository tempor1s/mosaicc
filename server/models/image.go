package models

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Image represents an uploaded image in the database
type Image struct {
	UserID       string    `json:"-" bson:"user_id"`                   // the ID of the auth0 user who owns the image
	FullURL      string    `json:"full_url" bson:"full_url"`           // the full URL of the image (stored on google cloud storage)
	ShortURL     string    `json:"short_url" bson:"short_url"`         // the CDN url of the iamge (stored on google cloud cdn)
	Name         string    `json:"img_name" bson:"img_name"`           // the name of the image (what the image was named when it was uploaded)
	UploadedName string    `json:"uploaded_name" bson:"uploaded_name"` // the name of the image when it gets uploaded to google cloud (different than given name)
	UploadDate   time.Time `json:"upload_date" bson:"upload_date"`     // when the image was uploaded (for sorting)
}

// CreateImage will create a new image object in the database
func (s *Store) CreateImage(img Image) error {
	// get the images collection
	collection := s.Database.Collection("images")
	// create the image in the database
	_, err := collection.InsertOne(context.Background(), img)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// GetImagesByUser will return all the images for a a given auth0 user id
func (s *Store) GetImagesByUser(id string) ([]Image, error) {
	// get the images collection
	collection := s.Database.Collection("images")

	ctx := context.Background()

	// empty image array to decode into
	cursor, err := collection.Find(ctx, bson.M{"user_id": id})
	if err != nil {
		return []Image{}, err
	}

	var images []Image
	if err = cursor.All(ctx, &images); err != nil {
		return []Image{}, err
	}

	return images, nil
}
