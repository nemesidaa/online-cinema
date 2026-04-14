package comparations

import (
	"s3receiver/internal/boilerplate"
	"s3receiver/internal/filestorage"
)

var (
	_ boilerplate.FilestorageImplementation = (*filestorage.Filestorage)(nil)
)
