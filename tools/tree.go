package tools

func Tree(path string) (string, error) {
	result, err := Shell("tree " + path)
	if err != nil {
		return "", err
	}
	return result, nil
}
