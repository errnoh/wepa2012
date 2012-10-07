package main

type Counter int

func (c *Counter) Get() int {
	*c++
	return int(*c)
}
