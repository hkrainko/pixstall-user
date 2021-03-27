package file

import (
	"context"
	"pixstall-user/domain/file/model"
)

type Repo interface {
	SaveFile(ctx context.Context, file model.File, fileType model.FileType, ownerID string, acl []string) (*string, error)
}
