package execx

import "fmt"

// ExitStatus stores exit information of the command
type ExitStatus struct {
	Code     int
	Signaled bool
	Killed   bool
	Timeout  bool
	Canceled bool
	Err      error
}

func (es *ExitStatus) Error() string {
	if es.Err != nil {
		return es.Err.Error()
	}
	return fmt.Sprintf("exit command with %d", es.Code)
}
