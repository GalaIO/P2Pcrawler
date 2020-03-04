package dht

type Dict map[string]interface{}

func (d *Dict) GetInteger(key string) int {
	dst, ok := d.GetVal(key).(int)
	if !ok {
		panicBizErr(key + "'s val cannot match type of int")
	}
	return dst
}

func (d *Dict) GetString(key string) string {
	dst, ok := d.GetVal(key).(string)
	if !ok {
		panicBizErr(key + "'s val cannot match type of string")
	}
	return dst
}

func (d *Dict) GetList(key string) List {
	dst, ok := d.GetVal(key).(List)
	if !ok {
		panicBizErr(key + "'s val cannot match type of List")
	}
	return dst
}

func (d *Dict) GetDict(key string) Dict {
	dst, ok := d.GetVal(key).(Dict)
	if !ok {
		panicBizErr(key + "'s val cannot match type of Dict")
	}
	return dst
}

func (d *Dict) GetVal(key string) interface{} {
	val, exist := (*d)[key]
	if !exist {
		panicBizErr("cannot find " + key + "'s val")
	}
	return val
}

type List []interface{}

func (l *List) GetVal(index int) interface{} {
	return (*l)[index]
}

func (l *List) GetInteger(index int) int {
	val, ok := l.GetVal(index).(int)
	if !ok {
		panicBizErr("val cannot match type of int")
	}
	return val
}

func (l *List) GetString(index int) string {
	val, ok := l.GetVal(index).(string)
	if !ok {
		panicBizErr("val cannot match type of string")
	}
	return val
}

func (l *List) GetList(index int) List {
	val, ok := l.GetVal(index).(List)
	if !ok {
		panicBizErr("val cannot match type of List")
	}
	return val
}

func (l *List) GetDict(index int) Dict {
	val, ok := l.GetVal(index).(Dict)
	if !ok {
		panicBizErr("val cannot match type of Dict")
	}
	return val
}
