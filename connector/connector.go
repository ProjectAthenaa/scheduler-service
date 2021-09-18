package connector

import "context"

type Server struct {
	UnimplementedSchedulerServer
}

func (s Server) GetScheduledTasks(ctx context.Context, user *User) (*Tasks, error) {
	panic("implement me")
}
