package pathing

import "errors"

func Search(start, end string) ([]string, error) {
	return bfs(start, end, getLinks)
}

func bfs(start, end string, getLinksFunc func(string) ([]string, error)) ([]string, error) {
	if start == end {
		return []string{}, nil
	}

	var frontier []string
	discovered := make(map[string]struct{})
	cameFrom := make(map[string]string)

	frontier = append(frontier, start)
	for len(frontier) > 0 {
		current := frontier[0]
		frontier = frontier[1:]

		links, err := getLinksFunc(current)
		if err != nil {
			return nil, err
		}

		for _, link := range links {
			if link == end {
				cameFrom[link] = current
				return getPath(link, cameFrom), nil
			}

			_, exists := discovered[link]
			if !exists && !contains(frontier, link) {
				cameFrom[link] = current
				frontier = append(frontier, link)
			}
		}

		discovered[current] = struct{}{}
	}

	return []string{}, errors.New("no path found")
}

func getPath(article string, cameFrom map[string]string) []string {
	path := []string{article}
	current := article
	for {
		origin, exists := cameFrom[current]
		if !exists {
			break
		}
		path = append(path, origin)
		current = origin
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

func contains(stringArray []string, value string) bool {
	for _, val := range stringArray {
		if val == value {
			return true
		}
	}
	return false
}
