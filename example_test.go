package swagger

import (
	"github.com/sohaha/zlsgo/znet"
)

func ExampleSwagger() {
	r := znet.New()

	// import _ "app/docs"

	RegisterDocs(r)

	znet.Run()
}
