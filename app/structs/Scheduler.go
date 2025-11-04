package structs

import "context"
import "git-evac-app/actions"
import "git-evac/schemas"
import "sync"

type Action struct {
	Type       string              `json:"type"`
	Owner      string              `json:"owner"`
	Repository string              `json:"repository"`
	Response   *schemas.Repository `json:"response"`
	Error      error               `json:"error"`
}

type Scheduler struct {
	Queue      []Action
	Results    chan Action
	cancelfunc context.CancelFunc
	mutex      *sync.Mutex
	waitgroup  *sync.WaitGroup
}

func NewScheduler() *Scheduler {

	var scheduler Scheduler

	scheduler.Queue = make([]Action, 0)
	scheduler.Results = make(chan Action, 0)
	scheduler.cancelfunc = nil
	scheduler.mutex = &sync.Mutex{}
	scheduler.waitgroup = &sync.WaitGroup{}

	return &scheduler

}

func (scheduler *Scheduler) Add(typ string, owner string, repository string) bool {

	found := false

	for _, action := range scheduler.Queue {

		if action.Type == typ && action.Owner == owner && action.Repository == repository {
			found = true
			break
		}

	}

	if found == false {

		scheduler.mutex.Lock()
		scheduler.Queue = append(scheduler.Queue, Action{
			Type:       typ,
			Owner:      owner,
			Repository: repository,
			Response:   nil,
			Error:      nil,
		})
		scheduler.mutex.Unlock()

		return true

	}

	return false

}

func (scheduler *Scheduler) Reset() {

	scheduler.mutex.Lock()

	if scheduler.cancelfunc != nil {
		scheduler.cancelfunc()
		scheduler.cancelfunc = nil
	}

	scheduler.mutex.Unlock()

	go func() {

		scheduler.mutex.Lock()

		scheduler.waitgroup.Wait()
		close(scheduler.Results)

		scheduler.waitgroup = &sync.WaitGroup{}
		scheduler.Queue = make([]Action, 0)
		scheduler.Results = make(chan Action, 0)

		scheduler.mutex.Unlock()

	}()

}

func (scheduler *Scheduler) Start() {

	scheduler.mutex.Lock()

	batch := make([]Action, len(scheduler.Queue))
	copy(batch, scheduler.Queue)

	scheduler.Queue = make([]Action, 0)
	scheduler.Results = make(chan Action, len(batch))

	ctx, cancel := context.WithCancel(context.Background())
	scheduler.cancelfunc = cancel

	scheduler.mutex.Unlock()

	go func() {

		for _, action := range batch {

			scheduler.waitgroup.Add(1)

			select {
			case <-ctx.Done():
				return
			default:
				// Continue
			}

			var response *schemas.Repository = nil
			var err error = nil

			switch action.Type {
			case "clone":
				response, err = actions.Clone(action.Owner, action.Repository)
			case "fix":
				response, err = actions.Fix(action.Owner, action.Repository)
			case "commit":
				response, err = actions.Commit(action.Owner, action.Repository)
			case "pull":
				response, err = actions.Pull(action.Owner, action.Repository)
			case "push":
				response, err = actions.Push(action.Owner, action.Repository)
			case "backup":
				response, err = actions.Backup(action.Owner, action.Repository)
			case "restore":
				response, err = actions.Restore(action.Owner, action.Repository)
			}

			action.Response = response
			action.Error = err

			select {
			case <-ctx.Done():
				return
			default:
				scheduler.Results <- action
			}

			scheduler.waitgroup.Done()

		}

		close(scheduler.Results)

	}()

}

func (scheduler *Scheduler) Stop() {

	scheduler.mutex.Lock()

	if scheduler.cancelfunc != nil {
		scheduler.cancelfunc()
		scheduler.cancelfunc = nil
	}

	go func() {
		scheduler.waitgroup.Wait()
		close(scheduler.Results)
	}()

	scheduler.mutex.Unlock()

}

func (scheduler *Scheduler) Wait() {
	scheduler.waitgroup.Wait()
}
