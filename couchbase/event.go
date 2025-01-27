package couchbase

import "time"

type Event struct {
	CollectionName string
	EventTime      time.Time
	Key            []byte
	Value          []byte
	IsDeleted      bool
	IsExpired      bool
	IsMutated      bool
}

func NewDeleteEvent(key []byte, value []byte, collectionName string, eventTime time.Time) Event {
	return Event{
		Key:            key,
		Value:          value,
		IsDeleted:      true,
		CollectionName: collectionName,
		EventTime:      eventTime,
	}
}

func NewExpireEvent(key []byte, value []byte, collectionName string, eventTime time.Time) Event {
	return Event{
		Key:            key,
		Value:          value,
		IsExpired:      true,
		CollectionName: collectionName,
		EventTime:      eventTime,
	}
}

func NewMutateEvent(key []byte, value []byte, collectionName string, eventTime time.Time) Event {
	return Event{
		Key:            key,
		Value:          value,
		IsMutated:      true,
		CollectionName: collectionName,
		EventTime:      eventTime,
	}
}
