package bizrule

import "reflect"

func And(channels ...<-chan struct{}) <-chan struct{} {
	ch := make(chan struct{})
	if len(channels) == 0 {
		close(ch)
		return ch
	}
	go func() {
		cases := make([]reflect.SelectCase, 0, len(channels))
		for _, ch := range channels {
			cases = append(cases, reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)})
		}
		for i, trigger := 0, len(channels); i < trigger; i++ {
			reflect.Select(cases)
		}
		close(ch)
	}()
	return ch
}

func Or(channels ...<-chan struct{}) <-chan struct{} {
	ch := make(chan struct{})
	if len(channels) == 0 {
		close(ch)
		return ch
	}
	go func() {
		cases := make([]reflect.SelectCase, 0, len(channels))
		for _, ch := range channels {
			cases = append(cases, reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)})
		}
		reflect.Select(cases)
		close(ch)
	}()
	return ch
}
