package cmd

import "flag"

type Configs struct {
	task int
	text string
	key  string
}

func NewConfigs() *Configs {
	task := flag.Int("task", 1, "Number of the task to process")
	text := flag.String("text", "255", "Input text for ciphering")
	key := flag.String("key", "250", "Key to cipher text with")
	flag.Parse()
	return &Configs{
		task: *task,
		text: *text,
		key:  *key,
	}
}

func (c Configs) Task() int {
	return c.task
}

func (c Configs) Text() string {
	return c.text
}

func (c Configs) Key() string {
	return c.key
}
