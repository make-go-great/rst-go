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

func parseTitle(line string) (title, bool) {
	lines := strings.Split(line, "\n")
	if len(lines) != 3 {
		return title{}, false
	}

	text := strings.TrimSpace(lines[1])

	return title{
		text: text,
	}, true
}

func parseListItem(line string) listItem {
	for listTok := range listTokens {
		line = strings.TrimLeft(line, string(listTok))
	}

	line = strings.TrimSpace(line)

	return listItem{
		text: line,
	}
}
