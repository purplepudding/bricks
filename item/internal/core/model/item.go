package model

import (
	"log/slog"
	"time"

	commonv1 "github.com/purplepudding/bricks/api/pkg/pb/bricks/v1/common"
	itemv1 "github.com/purplepudding/bricks/api/pkg/pb/bricks/v1/item"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Item struct {
	ID           string
	Name         string
	Labels       []string
	Flags        uint64
	Availability TimeRangeList
	Version      uint64
	Assets       StringAnyMap
	Parameters   StringAnyMap
}

func ItemFromPB(item *itemv1.Item) *Item {
	if item == nil {
		return nil
	}

	return &Item{
		ID:     item.Id,
		Name:   item.Name,
		Labels: item.Labels,
		Flags:  item.Flags,
	}
}

func (item Item) IntoPB() *itemv1.Item {
	return &itemv1.Item{
		Id:                   item.ID,
		Name:                 item.Name,
		Labels:               item.Labels,
		Flags:                item.Flags,
		AvailabilitySchedule: item.Availability.IntoPB(),
		Version:              item.Version,
	}
}

type StringAnyMap map[string]any

func (sam StringAnyMap) IntoPB() map[string]*structpb.Value {
	var m map[string]*structpb.Value
	for k, v := range sam {
		sv, err := structpb.NewValue(v)
		if err != nil {
			slog.Error("unexpected value in StringAnyMap, continuing", "err", err, "key", k, "value", v)
			continue
		}

		m[k] = sv
	}

	return m
}

type TimeRange struct {
	From *time.Time
	To   *time.Time
}

type TimeRangeList []TimeRange

func (trl TimeRangeList) IntoPB() []*commonv1.TimeRange {
	var ctrl []*commonv1.TimeRange

	for _, tr := range trl {
		var from, to *timestamppb.Timestamp

		if tr.From != nil {
			from = timestamppb.New(*tr.From)
		}
		if tr.To != nil {
			to = timestamppb.New(*tr.To)
		}

		ctrl = append(ctrl, &commonv1.TimeRange{
			From: from,
			To:   to,
		})
	}

	return ctrl
}
