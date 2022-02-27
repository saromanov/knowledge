package validation

import (
	"errors"

	restModel "github.com/saromanov/knowledge/internal/models/rest"
)

var (
	errNoTitle = errors.New("title is not defined")
	errNoAuthor = errors.New("author is not defined")
)

// PostPafe provides validation of the create page request
func PostPage(m *restModel.Page) error {
	if m.Title == "" {
		return errNoTitle
	}
	if m.AuthorID == "" {
		return errNoAuthor
	}
	return nil
}
