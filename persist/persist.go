package persist

import "dicuz-crawler/model"

type Saver interface {
	Init()
	Save(item model.Item) error
	Close()
}
