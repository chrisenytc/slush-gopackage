/*
 * <%= appNameSlug %>
 * https://github.com/<%= userName %>/<%= appNameSlug %>
 *
 * Copyright (c) <%= year %>, <%= authorName %>
 * Licensed under the <%= license %> license.
 */

package main

import (
	"testing"
	"github.com/<%= userName %>/<%= appNameSlug %>"
)

func TestReturnMsg(t *testing.T) {
	msg := "Testing Message func"
	testMsg := <%= appNameSlug %>.ReturnMsg(msg)
	if msg != testMsg {
		t.Errorf("'%s' is different of '%s'", msg, testMsg)
	}
}
