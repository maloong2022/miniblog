package errorsx_test

import (
	"fmt"
	"github.com/maloong2022/miniblog/internal/pkg/errno"
	"github.com/maloong2022/miniblog/pkg/errorsx"
	"testing"
)

func TestErrorsXMethods(t *testing.T) {

	errx := errorsx.New(500, "InternalError.DBConnection", "Something went wrong: %s", "DB connection failed")

	fmt.Println(errx)

	errx.WithMetadata(map[string]string{

		"user_id": "12345",

		"request_id": "abc-def",
	})

	errx.KV("trace_id", "xyz-789")

	errx.WithMessage("Updated message: %s", "retry failed")

	fmt.Println(errx)

	someerr := doSomething()

	fmt.Println(someerr)

	fmt.Println(errno.ErrUsernameInvalid.Is(someerr))

	fmt.Println(errno.ErrPasswordInvalid.Is(someerr))

}

func doSomething() error {

	return errno.ErrUsernameInvalid.WithMessage("Username is too short")

}
