package playwright

func SnapshotForAI(page Page, options ...LocatorAriaSnapshotOptions) (string, error) {
	var option LocatorAriaSnapshotOptions
	if len(options) == 1 {
		option = options[0]
	}
	p := page.(*pageImpl)
	ret, err := p.channel.Send("snapshotForAI", option)
	if err != nil {
		return "", err
	}
	return ret.(string), nil
}
