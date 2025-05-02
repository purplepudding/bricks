package common

import (
	commonv1 "github.com/purplepudding/bricks/api/pkg/pb/bricks/v1/common"
)

type Page struct {
	LastID string
	Count  uint32
}

func PageFromCommonPB(page *commonv1.Pagination) *Page {
	if page == nil {
		return nil
	}

	var lastID string
	if page.LastId != nil {
		lastID = *page.LastId
	}

	return &Page{
		LastID: lastID,
		Count:  page.Count,
	}
}
