package jstart

type RuleRegistry struct {
	rules map[string]Rule
}

func NewRuleRegistry() *RuleRegistry {
	return &RuleRegistry{
		rules: make(map[string]Rule),
	}
}

func (r *RuleRegistry) Register(rule Rule) {
	r.rules[rule.Name()] = rule
}

func (r *RuleRegistry) Find(name string) (Rule, bool) {
	rule, exist := r.rules[name]
	return rule, exist
}
