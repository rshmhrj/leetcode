package easy

func isValid(s string) bool {
	if len(s) <= 1 || s == "" || isClose(string(s[0])) {
		return false
	}

	parenStack := &stack{}
	for _, v := range s {
		if isOpen(string(v)) {
			n := &node{
				value: string(v),
				next:  nil,
			}
			parenStack.push(n)
			continue
		}
		if isClose(string(v)) {
			if parenStack.len == 0 || !parenStack.performMatch(string(v)) {
				return false
			}
		}
	}
	if parenStack.len != 0 {
		return false
	}
	return true
}

func isOpen(symbol string) bool {
	if symbol == "(" || symbol == "{" || symbol == "[" {
		return true
	}
	return false
}

func isClose(symbol string) bool {
	if symbol == ")" || symbol == "}" || symbol == "]" {
		return true
	}
	return false
}

func (s *stack) performMatch(symbol string) bool {
	match := s.pop()
	return isMatch(match.value, symbol)
}

func isMatch(o, c string) bool {
	if o == "(" && c == ")" {
		return true
	}
	if o == "{" && c == "}" {
		return true
	}
	if o == "[" && c == "]" {
		return true
	}
	return false
}

type node struct {
	value string
	next  *node
}

type stack struct {
	head *node
	len  int
}

func (s *stack) push(n *node) {
	n.next = s.head
	s.head = n
	s.len++
}

func (s *stack) pop() *node {
	n := s.head
	s.head = s.head.next
	s.len--
	return n
}
