package rolele

import "log"

type Action struct {
	ID      int
	Name    string
	Comment string
}

func ApplyActions(handlerAllowActions []Action, actions []string) bool {
	for _, handlerAllowAction := range handlerAllowActions {
		for _, action := range actions {
			if handlerAllowAction.Name == action {
				log.Printf("allow with action: %s\n", action)
				return true
			}
		}
	}

	log.Printf("action is not allow\n")
	return false
}
