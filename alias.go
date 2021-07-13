package alias

import "fmt"

type alias struct {
	Key     string
	Command string
}

type List []alias

func (l *List) Add(key string, command string) {
	a := alias{
		Key:     key,
		Command: command,
	}
	*l = append(*l, a)
}

func (l *List) Delete(key string) error {
	ls := *l
	if key == "" {
		return fmt.Errorf("alias %s does not exist", key)
	}
	for i, a := range ls {
		if a.Key == key {
			*l = append(ls[:i-1], ls[i:]...)
			return nil
		}
	}
	return nil
}
