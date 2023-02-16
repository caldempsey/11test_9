// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	//-----------------------------------------------------------------
	// The queries below return data in the format expected by the API.
	//-----------------------------------------------------------------
	MxSessionGet(ctx context.Context, id uuid.UUID) (MxSessionGetRow, error)
	MxSessionsGet(ctx context.Context, projectID uuid.UUID) ([]MxSessionsGetRow, error)
	ProjectGet(ctx context.Context, id uuid.UUID) (UnweaveProject, error)
	SSHKeyAdd(ctx context.Context, arg SSHKeyAddParams) error
	SSHKeyGetByName(ctx context.Context, arg SSHKeyGetByNameParams) (UnweaveSshKey, error)
	SSHKeyGetByPublicKey(ctx context.Context, arg SSHKeyGetByPublicKeyParams) (UnweaveSshKey, error)
	SSHKeysGet(ctx context.Context, ownerID uuid.UUID) ([]UnweaveSshKey, error)
	SessionCreate(ctx context.Context, arg SessionCreateParams) (uuid.UUID, error)
	SessionGet(ctx context.Context, id uuid.UUID) (UnweaveSession, error)
	SessionGetAllActive(ctx context.Context) ([]UnweaveSession, error)
	SessionSetError(ctx context.Context, arg SessionSetErrorParams) error
	SessionStatusUpdate(ctx context.Context, arg SessionStatusUpdateParams) error
	SessionUpdateConnectionInfo(ctx context.Context, arg SessionUpdateConnectionInfoParams) error
	SessionsGet(ctx context.Context, arg SessionsGetParams) ([]SessionsGetRow, error)
}

var _ Querier = (*Queries)(nil)
