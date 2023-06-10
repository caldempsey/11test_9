// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: exec.sql

package db

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/lib/pq"
)

const ExecCreate = `-- name: ExecCreate :exec
insert into unweave.exec (id, created_by, project_id,
                          region, name, spec, metadata, commit_id, git_remote_url,
                          command,
                          build_id, image, provider)
values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
`

type ExecCreateParams struct {
	ID           string          `json:"id"`
	CreatedBy    string          `json:"createdBy"`
	ProjectID    string          `json:"projectID"`
	Region       string          `json:"region"`
	Name         string          `json:"name"`
	Spec         json.RawMessage `json:"spec"`
	Metadata     json.RawMessage `json:"metadata"`
	CommitID     sql.NullString  `json:"commitID"`
	GitRemoteUrl sql.NullString  `json:"gitRemoteUrl"`
	Command      []string        `json:"command"`
	BuildID      sql.NullString  `json:"buildID"`
	Image        string          `json:"image"`
	Provider     string          `json:"provider"`
}

func (q *Queries) ExecCreate(ctx context.Context, arg ExecCreateParams) error {
	_, err := q.db.ExecContext(ctx, ExecCreate,
		arg.ID,
		arg.CreatedBy,
		arg.ProjectID,
		arg.Region,
		arg.Name,
		arg.Spec,
		arg.Metadata,
		arg.CommitID,
		arg.GitRemoteUrl,
		pq.Array(arg.Command),
		arg.BuildID,
		arg.Image,
		arg.Provider,
	)
	return err
}

const ExecGet = `-- name: ExecGet :one
select id, name, region, created_by, created_at, ready_at, exited_at, status, project_id, error, build_id, spec, commit_id, git_remote_url, command, metadata, image, provider
from unweave.exec
where id = $1
   or name = $1
`

func (q *Queries) ExecGet(ctx context.Context, idOrName string) (UnweaveExec, error) {
	row := q.db.QueryRowContext(ctx, ExecGet, idOrName)
	var i UnweaveExec
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Region,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.ReadyAt,
		&i.ExitedAt,
		&i.Status,
		&i.ProjectID,
		&i.Error,
		&i.BuildID,
		&i.Spec,
		&i.CommitID,
		&i.GitRemoteUrl,
		pq.Array(&i.Command),
		&i.Metadata,
		&i.Image,
		&i.Provider,
	)
	return i, err
}

const ExecGetAllActive = `-- name: ExecGetAllActive :many
select id, name, region, created_by, created_at, ready_at, exited_at, status, project_id, error, build_id, spec, commit_id, git_remote_url, command, metadata, image, provider
from unweave.exec
where status = 'initializing'
   or status = 'running'
`

func (q *Queries) ExecGetAllActive(ctx context.Context) ([]UnweaveExec, error) {
	rows, err := q.db.QueryContext(ctx, ExecGetAllActive)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UnweaveExec
	for rows.Next() {
		var i UnweaveExec
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Region,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.ReadyAt,
			&i.ExitedAt,
			&i.Status,
			&i.ProjectID,
			&i.Error,
			&i.BuildID,
			&i.Spec,
			&i.CommitID,
			&i.GitRemoteUrl,
			pq.Array(&i.Command),
			&i.Metadata,
			&i.Image,
			&i.Provider,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const ExecList = `-- name: ExecList :many
select id, name, region, created_by, created_at, ready_at, exited_at, status, project_id, error, build_id, spec, commit_id, git_remote_url, command, metadata, image, provider
from unweave.exec as e
where (e.provider = coalesce($1, e.provider))
  and project_id = coalesce($2, project_id)
  and (($3 = true and (status = 'initializing' or status = 'running'))
    or $3 = false)
`

type ExecListParams struct {
	FilterProvider  sql.NullString `json:"filterProvider"`
	FilterProjectID sql.NullString `json:"filterProjectID"`
	FilterActive    interface{}    `json:"filterActive"`
}

func (q *Queries) ExecList(ctx context.Context, arg ExecListParams) ([]UnweaveExec, error) {
	rows, err := q.db.QueryContext(ctx, ExecList, arg.FilterProvider, arg.FilterProjectID, arg.FilterActive)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UnweaveExec
	for rows.Next() {
		var i UnweaveExec
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Region,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.ReadyAt,
			&i.ExitedAt,
			&i.Status,
			&i.ProjectID,
			&i.Error,
			&i.BuildID,
			&i.Spec,
			&i.CommitID,
			&i.GitRemoteUrl,
			pq.Array(&i.Command),
			&i.Metadata,
			&i.Image,
			&i.Provider,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const ExecListActiveByProvider = `-- name: ExecListActiveByProvider :many
select id, name, region, created_by, created_at, ready_at, exited_at, status, project_id, error, build_id, spec, commit_id, git_remote_url, command, metadata, image, provider
from unweave.exec as e
where provider = $1
  and (status = 'initializing'
    or status = 'running')
`

func (q *Queries) ExecListActiveByProvider(ctx context.Context, provider string) ([]UnweaveExec, error) {
	rows, err := q.db.QueryContext(ctx, ExecListActiveByProvider, provider)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UnweaveExec
	for rows.Next() {
		var i UnweaveExec
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Region,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.ReadyAt,
			&i.ExitedAt,
			&i.Status,
			&i.ProjectID,
			&i.Error,
			&i.BuildID,
			&i.Spec,
			&i.CommitID,
			&i.GitRemoteUrl,
			pq.Array(&i.Command),
			&i.Metadata,
			&i.Image,
			&i.Provider,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const ExecListByProvider = `-- name: ExecListByProvider :many
select id, name, region, created_by, created_at, ready_at, exited_at, status, project_id, error, build_id, spec, commit_id, git_remote_url, command, metadata, image, provider
from unweave.exec as e
where e.provider = $1
`

func (q *Queries) ExecListByProvider(ctx context.Context, provider string) ([]UnweaveExec, error) {
	rows, err := q.db.QueryContext(ctx, ExecListByProvider, provider)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UnweaveExec
	for rows.Next() {
		var i UnweaveExec
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Region,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.ReadyAt,
			&i.ExitedAt,
			&i.Status,
			&i.ProjectID,
			&i.Error,
			&i.BuildID,
			&i.Spec,
			&i.CommitID,
			&i.GitRemoteUrl,
			pq.Array(&i.Command),
			&i.Metadata,
			&i.Image,
			&i.Provider,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const ExecSetError = `-- name: ExecSetError :exec
update unweave.exec
set status = 'error'::unweave.exec_status,
    error  = $2
where id = $1
`

type ExecSetErrorParams struct {
	ID    string         `json:"id"`
	Error sql.NullString `json:"error"`
}

func (q *Queries) ExecSetError(ctx context.Context, arg ExecSetErrorParams) error {
	_, err := q.db.ExecContext(ctx, ExecSetError, arg.ID, arg.Error)
	return err
}

const ExecSetFailed = `-- name: ExecSetFailed :exec
update unweave.exec
set status = 'failed'::unweave.exec_status,
    error  = $2
where id = $1
`

type ExecSetFailedParams struct {
	ID    string         `json:"id"`
	Error sql.NullString `json:"error"`
}

func (q *Queries) ExecSetFailed(ctx context.Context, arg ExecSetFailedParams) error {
	_, err := q.db.ExecContext(ctx, ExecSetFailed, arg.ID, arg.Error)
	return err
}

const ExecStatusUpdate = `-- name: ExecStatusUpdate :exec
update unweave.exec
set status    = $2,
    ready_at  = coalesce($3, ready_at),
    exited_at = coalesce($4, exited_at)
where id = $1
`

type ExecStatusUpdateParams struct {
	ID       string            `json:"id"`
	Status   UnweaveExecStatus `json:"status"`
	ReadyAt  sql.NullTime      `json:"readyAt"`
	ExitedAt sql.NullTime      `json:"exitedAt"`
}

func (q *Queries) ExecStatusUpdate(ctx context.Context, arg ExecStatusUpdateParams) error {
	_, err := q.db.ExecContext(ctx, ExecStatusUpdate,
		arg.ID,
		arg.Status,
		arg.ReadyAt,
		arg.ExitedAt,
	)
	return err
}

const ExecUpdateConnectionInfo = `-- name: ExecUpdateConnectionInfo :exec
update unweave.exec
set metadata = jsonb_set(metadata, '{connection_info}', $2::jsonb)
where id = $1
`

type ExecUpdateConnectionInfoParams struct {
	ID             string          `json:"id"`
	ConnectionInfo json.RawMessage `json:"connectionInfo"`
}

func (q *Queries) ExecUpdateConnectionInfo(ctx context.Context, arg ExecUpdateConnectionInfoParams) error {
	_, err := q.db.ExecContext(ctx, ExecUpdateConnectionInfo, arg.ID, arg.ConnectionInfo)
	return err
}

const ExecsGet = `-- name: ExecsGet :many
select id, name, region, created_by, created_at, ready_at, exited_at, status, project_id, error, build_id, spec, commit_id, git_remote_url, command, metadata, image, provider
from unweave.exec
where project_id = $1
order by unweave.exec.created_at desc
limit $2 offset $3
`

type ExecsGetParams struct {
	ProjectID string `json:"projectID"`
	Limit     int32  `json:"limit"`
	Offset    int32  `json:"offset"`
}

func (q *Queries) ExecsGet(ctx context.Context, arg ExecsGetParams) ([]UnweaveExec, error) {
	rows, err := q.db.QueryContext(ctx, ExecsGet, arg.ProjectID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UnweaveExec
	for rows.Next() {
		var i UnweaveExec
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Region,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.ReadyAt,
			&i.ExitedAt,
			&i.Status,
			&i.ProjectID,
			&i.Error,
			&i.BuildID,
			&i.Spec,
			&i.CommitID,
			&i.GitRemoteUrl,
			pq.Array(&i.Command),
			&i.Metadata,
			&i.Image,
			&i.Provider,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
