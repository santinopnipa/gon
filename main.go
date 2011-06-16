package main

import "web"
import "framework/starter"

func main() {
    web.Get("/(.*)", starter.Get)
    web.Run("0.0.0.0:8080")
}
