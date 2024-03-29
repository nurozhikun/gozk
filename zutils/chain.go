package zutils

func ChainCall(fns ...func() error) error {
	for _, fn := range fns {
		err := fn()
		if err != nil {
			return err
		}
	}
	return nil
}
