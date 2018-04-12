package archive

import (
	"github.com/forfd8960/go-archive/model"
)

type Itemer interface {
	GetItemList(offset, limit int32) ([]*model.ArchiveItem, error)
	GetItemCount() (int64, error)
}
