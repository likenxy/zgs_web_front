// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
    h := server.Default(
        server.WithHostPorts(":9191"),
    )
	register(h)
	h.Spin()
}
