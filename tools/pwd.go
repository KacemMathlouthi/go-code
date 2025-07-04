package tools

func Pwd() (string, error) {
	result, err := Shell("pwd")
	if err != nil {
		return "", err
	}
	return result, nil
}
