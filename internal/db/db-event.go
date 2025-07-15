package db

type DBEventType int64

const (
    DBCreate DBEventType = iota
    DBRead
    DBUpdate
    DBDelete
)

type DBEvent struct {
    Table string
    Verb DBEventType
    Data any
}
