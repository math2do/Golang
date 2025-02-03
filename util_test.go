package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	var num int = 15251
	res := isPalindrome(num)
	assert.False(t, res, "expected:%v, found:%v", false, res)
}
