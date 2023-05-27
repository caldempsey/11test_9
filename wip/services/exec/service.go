package exec

import (
	"context"
	"fmt"
	"time"

	"github.com/unweave/unweave/api/types"
)

type Service struct {
	store         Store
	driver        Driver
	stateInformer informer
}

func NewService(store Store, driver Driver) (*Service, error) {
	s := &Service{
		store:         store,
		driver:        driver,
		stateInformer: NewStateInformer(store, driver),
	}

	s.stateInformer.watch()

	execs, err := store.ListAll()
	if err != nil {
		return nil, fmt.Errorf("failed to init informer, failed list all execs: %w", err)
	}

	for _, e := range execs {
		e := e
		o := s.newStateObserver(e)
		s.stateInformer.register(o)
	}

	return s, nil
}

func (s *Service) newStateObserver(exec types.Exec) observer {
	return &stateObserver{exec: exec, srv: s}
}

func (s *Service) Create(ctx context.Context, project string, params types.ExecCreateParams) (types.Exec, error) {
	// TODO:
	// 	- Parse image and buildID
	//  - Parse network
	// 	- Parse volumes

	image := ""
	execID, err := s.driver.Create(ctx, project, image, params.Spec)
	if err != nil {
		return types.Exec{}, err
	}

	exec := types.Exec{
		ID:        execID,
		Name:      "",
		CreatedAt: time.Now(),
		CreatedBy: "",
		Image:     image,
		BuildID:   nil,
		Status:    types.StatusInitializing,
		Command:   params.Command,
		Keys:      nil,
		Volumes:   nil,
		Network:   types.ExecNetwork{},
		Spec:      types.HardwareSpec{},
		CommitID:  params.CommitID,
		GitURL:    params.GitURL,
		Region:    "", // Set later once
		Provider:  params.Provider,
	}
	if err = s.store.Create(project, exec); err != nil {
		return types.Exec{}, fmt.Errorf("failed to add exec to store: %w", err)
	}

	o := s.newStateObserver(exec)
	s.stateInformer.register(o)

	return exec, nil
}

func (s *Service) Get(ctx context.Context, id string) (types.Exec, error) {
	return types.Exec{}, nil
}

func (s *Service) List(ctx context.Context, project string) ([]types.Exec, error) {
	return nil, nil
}

func (s *Service) Terminate(ctx context.Context, id string) error {

	return nil
}
