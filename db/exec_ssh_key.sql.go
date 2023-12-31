// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: exec_ssh_key.sql

package db

import (
	"context"
)

const ExecSSHKeyDelete = `-- name: ExecSSHKeyDelete :exec
DELETE FROM unweave.exec_ssh_key
WHERE exec_id = $1
  AND ssh_key_id = $2
`

type ExecSSHKeyDeleteParams struct {
	ExecID   string `json:"execID"`
	SshKeyID string `json:"sshKeyID"`
}

func (q *Queries) ExecSSHKeyDelete(ctx context.Context, arg ExecSSHKeyDeleteParams) error {
	_, err := q.db.ExecContext(ctx, ExecSSHKeyDelete, arg.ExecID, arg.SshKeyID)
	return err
}

const ExecSSHKeyGet = `-- name: ExecSSHKeyGet :one
SELECT exec_id, ssh_key_id
FROM unweave.exec_ssh_key
WHERE exec_id = $1
  AND ssh_key_id = $2
`

type ExecSSHKeyGetParams struct {
	ExecID   string `json:"execID"`
	SshKeyID string `json:"sshKeyID"`
}

func (q *Queries) ExecSSHKeyGet(ctx context.Context, arg ExecSSHKeyGetParams) (UnweaveExecSshKey, error) {
	row := q.db.QueryRowContext(ctx, ExecSSHKeyGet, arg.ExecID, arg.SshKeyID)
	var i UnweaveExecSshKey
	err := row.Scan(&i.ExecID, &i.SshKeyID)
	return i, err
}

const ExecSSHKeyInsert = `-- name: ExecSSHKeyInsert :exec
INSERT INTO unweave.exec_ssh_key (exec_id, ssh_key_id)
VALUES ($1, $2)
`

type ExecSSHKeyInsertParams struct {
	ExecID   string `json:"execID"`
	SshKeyID string `json:"sshKeyID"`
}

func (q *Queries) ExecSSHKeyInsert(ctx context.Context, arg ExecSSHKeyInsertParams) error {
	_, err := q.db.ExecContext(ctx, ExecSSHKeyInsert, arg.ExecID, arg.SshKeyID)
	return err
}

const ExecSSHKeysGetByExecID = `-- name: ExecSSHKeysGetByExecID :many
SELECT exec_id, ssh_key_id
FROM unweave.exec_ssh_key
WHERE exec_id = $1
`

func (q *Queries) ExecSSHKeysGetByExecID(ctx context.Context, execID string) ([]UnweaveExecSshKey, error) {
	rows, err := q.db.QueryContext(ctx, ExecSSHKeysGetByExecID, execID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UnweaveExecSshKey
	for rows.Next() {
		var i UnweaveExecSshKey
		if err := rows.Scan(&i.ExecID, &i.SshKeyID); err != nil {
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
