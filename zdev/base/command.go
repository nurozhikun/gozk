package base

import "github.com/nurozhikun/gozk/zreflect"

func (c *Command) Make() *Command {
	if c.BodyMap == nil {
		c.BodyMap = NewMap()
	}
	return c
}

func (c *Command) StructToMap() *Command {
	// c.Make()
	c.BodyMap.InsertMap(zreflect.StructFieldsByTag(c.BodyStruct, StructTag))
	return c
}

func (c *Command) SetField(name string, value interface{}) *Command {
	c.BodyMap.Insert(name, value)
	return c
}

func (c *Command) SetFieldStream(stream IStream) *Command {
	c.BodyMap.Insert(FieldStream, stream)
	return c
}

func (c *Command) SetFieldCustom(custom ICustom) *Command {
	c.BodyMap.Insert(FieldCustom, custom)
	return c
}
