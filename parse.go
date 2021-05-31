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
			//
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
			//
		}
	}

	return nodes
}

func parseTitle(line string) (Title, bool) {
	lines := strings.Split(line, "\n")
	if len(lines) != 3 {
		return Title{}, false
	}

	text := strings.TrimSpace(lines[1])

	return Title{
		text: text,
	}, true
}

func parseListItem(line string) ListItem {
	for listTok := range listTokens {
		line = strings.TrimLeft(line, string(listTok))
	}

	line = strings.TrimSpace(line)

	return ListItem{
		text: line,
	}
}
