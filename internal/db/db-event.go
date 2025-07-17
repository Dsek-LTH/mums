package db

type DBEventType string

const (
    DBCreate DBEventType = "create"
    DBRead DBEventType = "read"
    DBUpdate DBEventType = "update"
    DBDelete DBEventType = "delete"
)

type DBEvent struct {
    Table string
    Type DBEventType
    Data any
}
