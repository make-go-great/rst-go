package rst

import "strings"

const defaultNodeLen = 10

func Parse(lines []string) []Node {
	if len(lines) == 0 {
		return nil
	}

	nodes := make([]Node, 0, defaultNodeLen)

	for _, line := range lines {
		if strings.HasPrefix(line, string(titleToken)) && strings.HasSuffix(line, string(titleToken)) {
			nodes = append(nodes, parseTitle(line))
			continue
		}

		if strings.HasSuffix(line, string(sectionToken)) {
			nodes = append(nodes, parseSection(line))
			continue
		}

		isListToken := false
		for listTok := range listTokens {
			if strings.HasPrefix(line, string(listTok)) {
				isListToken = true
				break
			}
		}

		if isListToken {
			nodes = append(nodes, parseListItem(line))
		}
	}

	return nodes
}

func parseTitle(line string) (Title, bool) {
	lines := strings.Split(line, "\n")
	if len(lines) != 3 {
		return Title{}, false
	}

	return Title{
		text: lines[1],
	}, true
}

func parseSection(line string) (Section, bool) {
	lines := strings.Split(line, "\n")
	if len(lines) != 2 {
		return Section{}, false
	}

	return NewSection(lines[0]), true
}

func parseListItem(line string) ListItem {
	for listTok := range listTokens {
		line = strings.TrimLeft(line, string(listTok))
	}

	return ListItem{
		text: line,
	}
}
