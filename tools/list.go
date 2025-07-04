package tools

func List() (string, error) {
	result, err := Shell("ls")
	if err != nil {
		return "", err
	}
	return result, nil
}
