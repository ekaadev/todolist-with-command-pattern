package src

// TODOLIST MANAGER: HISTORY COMMAND

type HistoryCommand struct {
	History []string
}

func (h *HistoryCommand) Add(history string) {
	isLimit := h.checkLimit()

	if isLimit {
		h.History = append(h.History[:0], h.History[1:]...)
	}

	h.History = append(h.History, history)
}

func (h *HistoryCommand) checkLimit() bool {
	return len(h.History) >= 5
}
