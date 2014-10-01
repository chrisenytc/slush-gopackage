/*
 * <%= appNameSlug %>
 * https://github.com/<%= userName %>/<%= appNameSlug %>
 *
 * Copyright (c) <%= year %>, <%= authorName %>
 * Licensed under the <%= license %> license.
 */

package main

/*
 * Dependencies
 */

import (
	"fmt"
	"github.com/<%= userName %>/<%= appNameSlug %>/lib"
)

/*
 * Main application
 */

func main() {
	msg := lib.ReturnMsg("Welcome to Go Package")
	fmt.Println(msg)
}
