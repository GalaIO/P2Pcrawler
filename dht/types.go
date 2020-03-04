package dht

type Dict map[string]interface{}

func (d *Dict) benEncode() (string, error) {
	return encodeDict(*d)
}

func (d *Dict) benDecode(src string) error {
	tmp, err := decodeDict(src)
	if err != nil {
		return err
	}
	*d = tmp
	return nil
}

type List []interface{}

func (l *List) benEncode() (string, error) {
	return encodeSlice(*l)
}

func (l *List) benDecode(src string) error {
	tmp, err := decodeSlice(src)
	if err != nil {
		return err
	}
	*l = tmp
	return nil
}
