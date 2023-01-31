package links

import (
	"errors"
	"math/rand"
	"time"

	"github.com/martin2844/link-shorten/pkg/db"
	"gorm.io/gorm"
)

func CreateLink(original string) (string, error) {
	// First search db to check if link exists, this should search for original. If original exists, return existing short.
	var link db.Link
	result := db.Instance.First(&link, "original = ?", original)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return link.Short, nil
	}
	// If original does not exist, create new link and return short.
	short := randString(8)
	newLink := db.Link{Original: original, Short: short}
	result = db.Instance.Create(&newLink)
	if result.Error != nil {
		return "", result.Error
	}
	return short, nil
}

func GetLink(short string) (string, error) {
	var link db.Link
	result := db.Instance.First(&link, "short = ?", short)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return "", result.Error
	}
	return link.Original, nil
}

func GetAllLinks() ([]db.Link, error) {
	var links []db.Link
	result := db.Instance.Find(&links)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return links, nil
}

// The randString function takes an integer parameter n, which represents the length of the random string to generate.
// The function uses the math/rand package to generate a random number and the time package to seed the random number generator.
// The function then creates a slice of runes with the length of n, and fills it with random characters from the set of letters and numbers.
// Finally, the function returns the string version of the slice of runes.
func randString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
