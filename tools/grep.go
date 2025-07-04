package tools

func Grep(pattern string, path string) (string, error) {
	result, err := Shell("grep " + pattern + " " + path)
	if err != nil {
		return "", err
	}
	return result, nil
}
