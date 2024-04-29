package utils

type Model struct {
	ID            string
	Name          string
	ContextWindow int
	Recommended   bool
	Developer     string
}

var GROQ_MODELS = []Model{
	{
		ID:            "llama3-8b-8192",
		Name:          "LLaMA3 8b",
		ContextWindow: 8192,
		Recommended:   false,
		Developer:     "Meta",
	},
	{
		ID:            "llama3-70b-8192",
		Name:          "LLaMA3 70b",
		ContextWindow: 8192,
		Recommended:   true,
		Developer:     "Meta",
	},
	{
		ID:            "mixtral-8x7b-32768",
		Name:          "Mixtral 8x7b",
		ContextWindow: 32768,
		Recommended:   false,
		Developer:     "Mistral",
	},
	{
		ID:            "gemma-7b-it",
		Name:          "Gemma 7b",
		ContextWindow: 8192,
		Recommended:   false,
		Developer:     "Google",
	},
}
