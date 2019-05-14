package query

import "github.com/getfider/fider/app/models/dto"

type ListBlobs struct {
	Result []string
}

type GetBlobByKey struct {
	Key string

	Result *dto.Blob
}
