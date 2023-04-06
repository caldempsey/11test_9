// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type UnweaveBuildStatus string

const (
	UnweaveBuildStatusInitializing UnweaveBuildStatus = "initializing"
	UnweaveBuildStatusBuilding     UnweaveBuildStatus = "building"
	UnweaveBuildStatusSuccess      UnweaveBuildStatus = "success"
	UnweaveBuildStatusFailed       UnweaveBuildStatus = "failed"
	UnweaveBuildStatusError        UnweaveBuildStatus = "error"
	UnweaveBuildStatusCanceled     UnweaveBuildStatus = "canceled"
)

func (e *UnweaveBuildStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UnweaveBuildStatus(s)
	case string:
		*e = UnweaveBuildStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for UnweaveBuildStatus: %T", src)
	}
	return nil
}

type NullUnweaveBuildStatus struct {
	UnweaveBuildStatus UnweaveBuildStatus
	Valid              bool // Valid is true if String is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUnweaveBuildStatus) Scan(value interface{}) error {
	if value == nil {
		ns.UnweaveBuildStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UnweaveBuildStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUnweaveBuildStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.UnweaveBuildStatus, nil
}

type UnweaveSessionStatus string

const (
	UnweaveSessionStatusInitializing UnweaveSessionStatus = "initializing"
	UnweaveSessionStatusRunning      UnweaveSessionStatus = "running"
	UnweaveSessionStatusTerminated   UnweaveSessionStatus = "terminated"
	UnweaveSessionStatusError        UnweaveSessionStatus = "error"
)

func (e *UnweaveSessionStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UnweaveSessionStatus(s)
	case string:
		*e = UnweaveSessionStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for UnweaveSessionStatus: %T", src)
	}
	return nil
}

type NullUnweaveSessionStatus struct {
	UnweaveSessionStatus UnweaveSessionStatus
	Valid                bool // Valid is true if String is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUnweaveSessionStatus) Scan(value interface{}) error {
	if value == nil {
		ns.UnweaveSessionStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UnweaveSessionStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUnweaveSessionStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.UnweaveSessionStatus, nil
}

type UnweaveAccount struct {
	ID string `json:"id"`
}

type UnweaveBuild struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	ProjectID   string             `json:"projectID"`
	BuilderType string             `json:"builderType"`
	Status      UnweaveBuildStatus `json:"status"`
	CreatedBy   string             `json:"createdBy"`
	CreatedAt   time.Time          `json:"createdAt"`
	StartedAt   sql.NullTime       `json:"startedAt"`
	FinishedAt  sql.NullTime       `json:"finishedAt"`
	UpdatedAt   time.Time          `json:"updatedAt"`
	MetaData    json.RawMessage    `json:"metaData"`
}

type UnweaveNode struct {
	ID        string          `json:"id"`
	Provider  string          `json:"provider"`
	Region    string          `json:"region"`
	Metadata  json.RawMessage `json:"metadata"`
	Status    string          `json:"status"`
	CreatedAt time.Time       `json:"createdAt"`
	ReadyAt   sql.NullTime    `json:"readyAt"`
	OwnerID   string          `json:"ownerID"`
}

type UnweaveNodeSshKey struct {
	NodeID    string    `json:"nodeID"`
	SshKeyID  string    `json:"sshKeyID"`
	CreatedAt time.Time `json:"createdAt"`
}

type UnweaveProject struct {
	ID             string         `json:"id"`
	DefaultBuildID sql.NullString `json:"defaultBuildID"`
}

type UnweaveSession struct {
	ID             string               `json:"id"`
	Name           string               `json:"name"`
	NodeID         string               `json:"nodeID"`
	Region         string               `json:"region"`
	CreatedBy      string               `json:"createdBy"`
	CreatedAt      time.Time            `json:"createdAt"`
	ReadyAt        sql.NullTime         `json:"readyAt"`
	ExitedAt       sql.NullTime         `json:"exitedAt"`
	Status         UnweaveSessionStatus `json:"status"`
	ProjectID      string               `json:"projectID"`
	SshKeyID       sql.NullString       `json:"sshKeyID"`
	ConnectionInfo json.RawMessage      `json:"connectionInfo"`
	Error          sql.NullString       `json:"error"`
	BuildID        sql.NullString       `json:"buildID"`
	Spec           json.RawMessage      `json:"spec"`
	CommitID       sql.NullString       `json:"commitID"`
	GitRemoteUrl   sql.NullString       `json:"gitRemoteUrl"`
	Command        []string             `json:"command"`
}

type UnweaveSshKey struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	OwnerID   string    `json:"ownerID"`
	CreatedAt time.Time `json:"createdAt"`
	PublicKey string    `json:"publicKey"`
	IsActive  bool      `json:"isActive"`
}
