package logging

import "maps"

type Label struct {
	key   string
	value string
}

func NewLabel(key, value string) *Label {
	return &Label{
		key:   key,
		value: value,
	}
}

func (label *Label) Key() string {
	return label.key
}

func (label *Label) Value() string {
	return label.value
}

type Labels struct {
	labels []*Label
	index  map[string][]*Label
}

func NewLabels(labels ...*Label) *Labels {
	index := make(map[string][]*Label)

	for _, label := range labels {
		key := label.Key()
		index[key] = append(index[key], label)
	}

	return &Labels{
		labels: labels,
		index:  index,
	}
}

func (labels *Labels) Get(key string) []*Label {
	return labels.index[key]
}

func (labels *Labels) All() []string {
	var keys []string

	for key := range labels.index {
		keys = append(keys, key)
	}

	return keys
}

func (labels *Labels) Clone() *Labels {
	newLabels := make([]*Label, len(labels.labels))
	copy(newLabels, labels.labels)

	newIndex := maps.Clone(labels.index)

	return &Labels{
		labels: newLabels,
		index:  newIndex,
	}
}
