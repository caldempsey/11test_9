package volume

import (
	"context"
)

// Volume is an interface that must be implemented by a volume.
// The implementation of the volume interface should have the ability/permissions
// to mount the volume onto a container inside a node. For instance, this could
// be an AWS EBS volume mounted onto a container inside an EC2 machine, a FUSE
// volume on top of S3/GCS, etc.
type Volume interface {
	// ID returns the ID of the volume.
	ID() string
	// Mount mounts the volume to the given path.
	Mount(ctx context.Context, path string) error
	// Unmount unmounts the volume from the given path.
	Unmount(ctx context.Context, path string) error
}

type Driver interface {
	Mount(ctx context.Context, volumeID, path string) error
	Unmount(ctx context.Context, volumeID, path string) error
}

// Provider allows management of volumes from different providers (eg AWS, GCP).
type Provider interface {
	// Name returns the name of the provider.
	Name() string
	// VolumeCreate creates a volume. The size is in GB.
	VolumeCreate(ctx context.Context, size int) (Volume, error)
	// VolumeDelete deletes the volume.
	VolumeDelete(ctx context.Context) error
	// VolumeGet gets the volume.
	VolumeGet(ctx context.Context, id string) (Volume, error)
	// VolumeList lists all volumes the provider has.
	VolumeList(ctx context.Context) ([]Volume, error)
	// VolumeResize resizes the volume to the given size in GB.
	VolumeResize(ctx context.Context, size int) error
}
