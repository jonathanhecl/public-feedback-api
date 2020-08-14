package extras

import (
	"crypto/sha1"
	"fmt"
	"io"
	"regexp"
	"time"
)

func ValidateEmail(email string) bool {

	return regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`).MatchString(email)

}

func GenerateModeratorLink(messageID string, createdAt time.Time, emailModerator string) string {

	h := sha1.New()
	io.WriteString(h, ex.ps)
	io.WriteString(h, messageID)
	io.WriteString(h, createdAt.String())
	io.WriteString(h, emailModerator)
	return fmt.Sprintf("%x", h.Sum(nil))

}
