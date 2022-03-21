package convert

import (
	restModel "github.com/saromanov/knowledge/internal/models/rest"
	"github.com/saromanov/knowledge/internal/models/storage"
	storageModel "github.com/saromanov/knowledge/internal/models/storage"
)

// RestPageToStoragePage converts rest model to storage model
func RestPageToStoragePage(rm *restModel.Page)*storageModel.Page{
	return &storageModel.Page{
		Title: rm.Title,
		Body: rm.Body,
		AuthorID: rm.AuthorID,
		CreatedAt: rm.CreatedAt,
	}
}

func RestAuthorToStorageAuthor(rm *restModel.Author)*storageModel.Author {
	return &storage.Author{
		
	}
}