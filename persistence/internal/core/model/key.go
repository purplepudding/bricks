package model

import (
	"fmt"
)

type StorageKey interface {
	Category() string
	Key() string
}

type PlayerStorageKey struct {
	TitleID  string
	PlayerID string
	Datatype string
}

func (psk PlayerStorageKey) Category() string {
	return fmt.Sprintf("player.%s.%s", psk.TitleID, psk.Datatype)
}

func (psk PlayerStorageKey) Key() string {
	return psk.PlayerID
}

type TitleStorageKey struct {
	TitleID  string
	Datatype string
}

func (tsk TitleStorageKey) Category() string {
	return fmt.Sprintf("title.%s", tsk.TitleID)
}

func (tsk TitleStorageKey) Key() string {
	return tsk.Datatype
}
