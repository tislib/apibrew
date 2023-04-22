package test

import (
	log "github.com/sirupsen/logrus"
	"github.com/tislib/apibrew/pkg/model"
	"github.com/tislib/apibrew/pkg/stub"
	"github.com/tislib/apibrew/pkg/test/setup"
	"github.com/tislib/apibrew/pkg/util"
	"testing"
)

func TestDhTestUserCannotCreateUser(t *testing.T) {
	userDhTestCtx := setup.WithUserAuthenticationContext(setup.Ctx, "dh_test", "dh_test")

	_, err := userClient.Create(userDhTestCtx, &stub.CreateUserRequest{
		Users: []*model.User{
			{
				Username: "test123",
			},
		},
	})

	if err == nil {
		t.Error("error expected ErrorCode_ACCESS_DENIED but it succeeded")
		return
	}

	if util.GetErrorCode(err) != model.ErrorCode_ACCESS_DENIED {
		log.Print(err)
		t.Error("error expected ErrorCode_ACCESS_DENIED")
	}
}

func TestDhTestUserCanReadUser(t *testing.T) {
	userDhTestCtx := setup.WithUserAuthenticationContext(setup.Ctx, "dh_test", "dh_test")

	_, err := userClient.List(userDhTestCtx, &stub.ListUserRequest{})

	if err != nil {
		log.Print(err)
		t.Error("read operation should be successful")
		return
	}
}
