package fizz_buzz

import (
	"testing"
	"reflect"
)

func TestCall1Say1(t *testing.T) {
	acture := CallFizzBuzz(2)
	expected := "2"



	if !reflect.DeepEqual(acture,expected) {
		t.Error(acture, expected)
	}
}

func TestCall3Say3(t *testing.T) {
	acture := CallFizzBuzz(3)
	expected := "fizz"

	if !reflect.DeepEqual(acture,expected) {
		t.Error(acture, expected)
	}
}