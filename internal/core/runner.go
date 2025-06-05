package core

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Plugin interface {
	Name() string
	Execute(params map[string]interface{}) error
}

type Runner struct {
	plugins map[string]Plugin
}

func NewRunner() *Runner {
	return &Runner{plugins: make(map[string]Plugin)}
}

func (r *Runner) RegisterPlugin(p Plugin) {
	r.plugins[p.Name()] = p
}

func (r *Runner) Run(path string) error {
	routine, err := LoadRoutine(path)
	if err != nil {
		return err
	}
	for _, step := range routine.Steps {
		p, ok := r.plugins[step.Device]
		if !ok {
			return fmt.Errorf("plugin '%s' not found", step.Device)
		}
		if err := p.Execute(step.Params); err != nil {
			return err
		}
		time.Sleep(time.Duration(step.DelayMS) * time.Millisecond)
	}
	return nil
}

type Routine struct {
	Steps []Step `yaml:"steps" json:"steps"`
}

type Step struct {
	Device  string                 `yaml:"device" json:"device"`
	Params  map[string]interface{} `yaml:"params" json:"params"`
	DelayMS int                    `yaml:"delay_ms" json:"delay_ms"`
}

func LoadRoutine(path string) (*Routine, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var routine Routine
	if err := yaml.Unmarshal(data, &routine); err != nil {
		return nil, err
	}
	return &routine, nil
}
