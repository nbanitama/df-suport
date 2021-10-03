package df

type Intent struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Responses []Response `json:"responses"`
	Priority  int        `json:"priority"`
	Events    []Event    `json:"events"`
}

type Event struct {
	Name string `json:"name"`
}

type Response struct {
	ResetContents    bool              `json:"resetContexts"`
	Action           string            `json:"action"`
	AffectedContexts []AffectedContext `json:"affectedContexts"`
}

type AffectedContext struct {
	Name     string `json:"name"`
	Lifespan int    `json:"lifespan"`
}
