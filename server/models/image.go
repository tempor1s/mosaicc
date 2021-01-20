package models

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

// GetImageByID will get the image by its object id
// the (bool) represents if the image exists or not
func (s *Store) GetImageByID(id string) (Image, bool, error) {
	// get the images collection
	collection := s.Database.Collection("images")
	ctx := context.Background()

	// get the database image
	var image Image
	err := collection.FindOne(ctx, bson.M{"img_name": id}).Decode(&image)

	if err != nil {
		if errors.Is(err, mongo.ErrNilDocument) {
			// does not exist because document not found
			return Image{}, false, errors.New("image not found")
		}
		// does not exist and other err
		return Image{}, false, err
	}

	// image exists, and return it with no err
	return image, true, nil
}

// DeleteImageByID will delete the image with the given ID
func (s *Store) DeleteImageByID(id string) error {
	// get the images collection
	collection := s.Database.Collection("images")
	ctx := context.Background()

	// get the database image
	_, err := collection.DeleteOne(ctx, bson.M{"img_name": id})
	if err != nil {
		return err
	}

	return nil
}
