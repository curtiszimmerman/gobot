/**
 * @project GoBot
 * GoBot IRC channel bot written in Go.
 * @file gobot_test.go
 * Test file for gobot.go
 * @author curtis zimmerman
 * @contact hey@curtisz.com
 * @license MIT
 * @version 0.0.1a
 */

/*START gobot_test.go*/
package main

import (
	"reflect"
	"testing"
)

func T_ReturnVersion() (v *Version) { return }

func TestVersion(t *testing.T) {
	//var versionStock = &Version{}
	var versionStock = T_ReturnVersion()
	v := version()
	if reflect.TypeOf(v) != reflect.TypeOf(versionStock) {
		t.Error("expected Version object returned, got %t", v)
	}
}

/*END gobot_test.go*/
