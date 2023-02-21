package image

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"time"

	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/service"
	_ "golang.org/x/image/webp"
)

const (
	expiredAfter = time.Hour
)

func IDs(is []people.Image) []uint {
	ids := make([]uint, len(is))
	for i, p := range is {
		ids[i] = p.ID
	}
	return ids
}

func isExpired(img people.Image) bool {
	return time.Now().Add(-expiredAfter).After(img.CreatedAt)
}

func validate(img people.Image, userID uint) error {
	if img.UserID != userID {
		return service.NewError(people.AuthError, fmt.Sprintf("You do not have permission to use this image: %v", img.ID))
	}
	if isExpired(img) {
		return service.NewError(people.ResourceError, fmt.Sprintf("Image is expired: %v", img.ID))
	}
	return nil
}

func getDimensions(img people.Image) (int, int, error) {
	// TODO: remove trailing slash
	f, err := os.Open(path(img.CreatedAt, img.Name))
	if err != nil {
		return 0, 0, err
	}
	defer f.Close()
	cfg, _, err := image.DecodeConfig(f)
	if err != nil {
		return 0, 0, err
	}
	return cfg.Width, cfg.Height, nil
}
