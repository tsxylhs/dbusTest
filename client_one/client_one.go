package main

import (
	"encoding/json"
	"github.com/godbus/dbus"
	"github.com/godbus/dbus/introspect"
	"os"
)

func main() {
	conn, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}
	node, err2 := introspect.Call(conn.Object("org.freedesktop.hello", "/org/freedesktop/hello"))
	if err2 != nil {
		panic(err2)
	}
	data, _ := json.MarshalIndent(node, "", "   ")
	os.Stdout.Write(data)
}
