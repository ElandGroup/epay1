package common

func ResourceMessage(resourceKey string, messages ...string) (message string) {
	message = resourceKey
	for _, v := range messages {
		message += "||" + v
	}
	return
}
