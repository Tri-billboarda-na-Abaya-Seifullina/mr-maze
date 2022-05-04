package service

import "fmt"

func (s *service) ExampleServiceMethod(name string) (string, error) {
	return fmt.Sprintf("Hello, %s", name), nil
}
