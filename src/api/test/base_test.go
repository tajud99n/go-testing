package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/tajud99n/go-testing/src/api/app"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	fmt.Println("about to start application...")
	go app.StartApp()
	fmt.Println("about to start functional tests...")

	os.Exit(m.Run())
}
