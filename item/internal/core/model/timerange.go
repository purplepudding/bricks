package model

import (
	"time"

	"github.com/purplepudding/bricks/api/pkg/pb/bricks/v1/common"
	"github.com/purplepudding/bricks/item/pkg/pb/bricks/item/v1/persistence"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TimeRange struct {
	From *time.Time
	To   *time.Time
}

type TimeRangeList []TimeRange

func TimeRangeListFromAPIPB(atrl []*common.TimeRange) TimeRangeList {
	var trl TimeRangeList

	for _, tr := range atrl {
		var from, to *time.Time

		if tr.From != nil {
			f := (*tr.From).AsTime()
			from = &f
		}
		if tr.To != nil {
			t := (*tr.To).AsTime()
			to = &t
		}

		trl = append(trl, TimeRange{From: from, To: to})
	}

	return trl
}

func TimeRangeListFromPersistencePB(ptrl []*persistence.TimeRange) TimeRangeList {
	var trl TimeRangeList

	for _, tr := range ptrl {
		var from, to *time.Time

		if tr.From != nil {
			f := (*tr.From).AsTime()
			from = &f
		}
		if tr.To != nil {
			t := (*tr.To).AsTime()
			to = &t
		}

		trl = append(trl, TimeRange{From: from, To: to})
	}

	return trl
}

func (trl TimeRangeList) IntoAPIPB() []*common.TimeRange {
	var ctrl []*common.TimeRange

	for _, tr := range trl {
		var from, to *timestamppb.Timestamp

		if tr.From != nil {
			from = timestamppb.New(*tr.From)
		}
		if tr.To != nil {
			to = timestamppb.New(*tr.To)
		}

		ctrl = append(ctrl, &common.TimeRange{From: from, To: to})
	}

	return ctrl
}

func (trl TimeRangeList) IntoPersistencePB() []*persistence.TimeRange {
	var ctrl []*persistence.TimeRange

	for _, tr := range trl {
		var from, to *timestamppb.Timestamp

		if tr.From != nil {
			from = timestamppb.New(*tr.From)
		}
		if tr.To != nil {
			to = timestamppb.New(*tr.To)
		}

		ctrl = append(ctrl, &persistence.TimeRange{From: from, To: to})
	}

	return ctrl
}
