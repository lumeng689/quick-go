package mock

func ReadFromDB() (string, error) {
	return "hello", nil
}

func HelloWorld() (string, error) {
	data, err := ReadFromDB() // external dependency
	if err != nil {
		return "", err
	}

	return data, nil
}
