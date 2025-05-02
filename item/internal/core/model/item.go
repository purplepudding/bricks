package model

import (
	itemv1 "github.com/purplepudding/bricks/api/pkg/pb/bricks/v1/item"
	persistencev1 "github.com/purplepudding/bricks/item/pkg/pb/bricks/item/v1/persistence"
	"google.golang.org/protobuf/types/known/structpb"
)

type Item struct {
	ID                   string
	Name                 string
	Labels               []string
	Flags                uint64
	AvailabilitySchedule TimeRangeList
	Version              uint64
	Assets               map[string]*structpb.Value
	Parameters           map[string]*structpb.Value
}

func ItemFromAPIPB(item *itemv1.Item) *Item {
	if item == nil {
		return nil
	}

	return &Item{
		ID:                   item.Id,
		Name:                 item.Name,
		Labels:               item.Labels,
		Flags:                item.Flags,
		AvailabilitySchedule: TimeRangeListFromAPIPB(item.AvailabilitySchedule),
		Version:              item.Version,
	}
}

func ItemFromPersistencePB(item *persistencev1.Item, id string, version uint64) *Item {
	if item == nil {
		return nil
	}

	return &Item{
		ID:                   id,
		Name:                 item.Name,
		Labels:               item.Labels,
		Flags:                item.Flags,
		AvailabilitySchedule: TimeRangeListFromPersistencePB(item.AvailabilitySchedule),
		Version:              version,
	}
}

func (item Item) IntoAPIPB() *itemv1.Item {
	return &itemv1.Item{
		Id:                   item.ID,
		Name:                 item.Name,
		Labels:               item.Labels,
		Flags:                item.Flags,
		AvailabilitySchedule: item.AvailabilitySchedule.IntoAPIPB(),
		Version:              item.Version,
	}
}

func (item Item) IntoPersistencePB() *persistencev1.Item {
	return &persistencev1.Item{
		Name:                 item.Name,
		Labels:               item.Labels,
		Flags:                item.Flags,
		AvailabilitySchedule: item.AvailabilitySchedule.IntoPersistencePB(),
	}
}
