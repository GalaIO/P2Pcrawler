package misc

import "strings"

type Dict map[string]interface{}

func (d *Dict) GetInteger(key string) int {
	dst, ok := d.GetVal(key).(int)
	if !ok {
		PanicBizErr(key + "'s val cannot match type of int")
	}
	return dst
}

func (d *Dict) GetString(key string) string {
	dst, ok := d.GetVal(key).(string)
	if !ok {
		PanicBizErr(key + "'s val cannot match type of string")
	}
	return dst
}

func (d *Dict) GetIntegerOrDefault(key string, val int) int {
	dst, ok := d.GetValOrDefault(key, val).(int)
	if !ok {
		return val
	}
	return dst
}

func (d *Dict) GetStringOrDefault(key string, val string) string {
	dst, ok := d.GetValOrDefault(key, val).(string)
	if !ok {
		return val
	}
	return dst
}

func (d *Dict) GetList(key string) List {
	dst, ok := d.GetVal(key).(List)
	if !ok {
		PanicBizErr(key + "'s val cannot match type of List")
	}
	return dst
}

func (d *Dict) GetDict(key string) Dict {
	dst, ok := d.GetVal(key).(Dict)
	if !ok {
		PanicBizErr(key + "'s val cannot match type of Dict")
	}
	return dst
}

func (d *Dict) GetVal(key string) interface{} {
	val, exist := (*d)[key]
	if !exist {
		PanicBizErr("cannot find " + key + "'s val")
	}
	return val
}

func (d *Dict) GetValOrDefault(key string, val interface{}) interface{} {
	val, exist := (*d)[key]
	if !exist {
		return val
	}
	return val
}

func (d *Dict) Exist(key string) bool {
	_, exist := (*d)[key]
	return exist
}

type List []interface{}

func (l *List) GetVal(index int) interface{} {
	return (*l)[index]
}

func (l *List) GetInteger(index int) int {
	val, ok := l.GetVal(index).(int)
	if !ok {
		PanicBizErr("val cannot match type of int")
	}
	return val
}

func (l *List) GetString(index int) string {
	val, ok := l.GetVal(index).(string)
	if !ok {
		PanicBizErr("val cannot match type of string")
	}
	return val
}

func (l *List) GetList(index int) List {
	val, ok := l.GetVal(index).(List)
	if !ok {
		PanicBizErr("val cannot match type of List")
	}
	return val
}

func (l *List) GetDict(index int) Dict {
	val, ok := l.GetVal(index).(Dict)
	if !ok {
		PanicBizErr("val cannot match type of Dict")
	}
	return val
}

func (l *List) Exist(index int) bool {
	return len(*l) > index
}

func (l *List) ContainsString(target string) bool {
	for i := range *l {
		if strings.EqualFold(l.GetString(i), target) {
			return true
		}
	}
	return false
}
