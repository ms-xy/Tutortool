package errorutils

func GetString(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
