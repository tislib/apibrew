package impl

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/helper"
	"github.com/apibrew/apibrew/pkg/logging"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/service"
	jwt2 "github.com/apibrew/apibrew/pkg/service/jwt"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
	"os"
	"time"
)

type authenticationService struct {
	recordService         service.RecordService
	privateKey            *rsa.PrivateKey
	publicKey             *rsa.PublicKey
	DisableAuthentication bool
}

func (s *authenticationService) AuthenticationDisabled() bool {
	return s.DisableAuthentication
}

func (s *authenticationService) Authenticate(ctx context.Context, username string, password string, term model.TokenTerm) (*model.Token, errors.ServiceError) {
	if s.DisableAuthentication {
		return &model.Token{
			Term:       term,
			Content:    "",
			Expiration: timestamppb.New(time.Now().Add(time.Minute).UTC()),
		}, nil
	}
	logger := log.WithFields(logging.CtxFields(ctx))

	logger.Debug("Begin Authenticate")

	defer logger.Debug("End Authenticate")

	systemCtx := util.WithSystemContext(ctx)

	// locate user
	user, err := s.LocateUser(systemCtx, username, password)

	if err != nil {
		return nil, errors.AuthenticationFailedError
	}

	return s.prepareToken(systemCtx, term, user)
}

func (s *authenticationService) prepareToken(ctx context.Context, term model.TokenTerm, user *resource_model.User) (*model.Token, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))
	// Prepare token
	expiration := s.ExpirationFromTerm(term)
	//roles, err := s.LocateRoles(ctx, user.Roles)

	//if err != nil {
	//	return nil, err
	//}

	//var constraints []abs.SecurityConstraint
	//
	//for _, role := range roles {
	//	for _, constraint := range role.SecurityConstraints {
	//		constraint.Role = &role.Name
	//		constraints = append(constraints, &constraint)
	//	}
	//}
	//
	//for _, constraint := range user.SecurityConstraints {
	//	constraints = append(constraints, &constraint)
	//}

	token, err := jwt2.JwtUserDetailsSign(jwt2.JwtUserDetailsSignParams{
		Key: *s.privateKey,
		UserDetails: abs.UserDetails{
			UserId:   user.Id.String(),
			Username: user.Username,
			Roles:    user.Roles,
		},
		ExpiresAt: expiration,
		Issuer:    "github.com/apibrew/apibrew",
	})

	logger.Tracef("Token prepared: %s", token)

	if err != nil {
		logger.Warning("Token preparation error", err)
		return nil, err
	}

	return &model.Token{
		Term:       term,
		Content:    token,
		Expiration: timestamppb.New(expiration.UTC()),
	}, nil
}

func (s *authenticationService) RenewToken(ctx context.Context, oldToken string, term model.TokenTerm) (*model.Token, errors.ServiceError) {
	userDetails, err := jwt2.JwtVerifyAndUnpackUserDetails(*s.publicKey, oldToken)

	if err != nil {
		return nil, err
	}

	systemCtx := util.WithSystemContext(ctx)

	user, err := s.FindUser(systemCtx, userDetails.Username)

	if err != nil {
		return nil, err
	}

	return s.prepareToken(systemCtx, term, user)
}

func (s *authenticationService) ParseAndVerifyToken(token string) (*abs.UserDetails, errors.ServiceError) {
	return jwt2.JwtVerifyAndUnpackUserDetails(*s.publicKey, token)
}

type RequestWithToken interface {
	GetToken() string
}

func (s *authenticationService) LocateUser(ctx context.Context, username, password string) (*resource_model.User, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))

	logger.Debugf("Locating user: %s", username)

	user, err := s.FindUser(ctx, username)
	if err != nil {
		logger.Debugf("Could not find user: %s", username)
		return nil, err
	}

	logger.Debugf("Checking password: %s", username)
	if user.Password == nil || util.VerifyKey(*user.Password, password) != nil {
		logger.Debugf("Password is wrong: %s", username)
		return nil, errors.AuthenticationFailedError
	}

	return user, nil
}

func (s *authenticationService) FindUser(ctx context.Context, username string) (*resource_model.User, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))

	logger.Debug("FindUser with username: ", username)

	res, err := s.recordService.FindBy(ctx, resources.UserResource.Namespace, resources.UserResource.Name, "username", username)

	if err != nil {
		return nil, err
	}

	return resource_model.UserMapperInstance.FromRecord(res), nil
}

func (s *authenticationService) LocateRoles(ctx context.Context, names []string) ([]*resource_model.Role, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))

	logger.Debug("Locating role by names: ", names)

	if len(names) == 0 {
		return []*resource_model.Role{}, nil
	}

	res, _, err := s.recordService.List(ctx, service.RecordListParams{
		Namespace: resources.RoleResource.Namespace,
		Resource:  resources.RoleResource.Name,
		Query:     helper.NewQueryBuilder().In("name", util.ArrayMap(names, func(s string) interface{} { return s })),
	})

	if err != nil {
		return nil, err
	}

	return util.ArrayMap(res, resource_model.RoleMapperInstance.FromRecord), nil
}

func (s *authenticationService) Init(config *model.AppConfig) {
	s.DisableAuthentication = config.GetDisableAuthentication()

	if config.DisableAuthentication {
		return
	}

	if config.JwtPrivateKey == "" {
		priv, err := rsa.GenerateKey(rand.Reader, 2048)

		if err != nil {
			panic(err)
		}

		s.privateKey = priv
		s.publicKey = &priv.PublicKey
	} else {
		privateKeyContent, err := os.ReadFile(config.JwtPrivateKey)
		if err != nil {
			panic(err)
		}
		publicKeyContent, err := os.ReadFile(config.JwtPublicKey)
		if err != nil {
			panic(err)
		}

		privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyContent)
		if err != nil {
			panic(err)
		}
		publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyContent)
		if err != nil {
			panic(err)
		}

		s.privateKey = privateKey
		s.publicKey = publicKey
	}
}

func (s *authenticationService) ExpirationFromTerm(term model.TokenTerm) time.Time {
	switch term {
	case model.TokenTerm_VERY_SHORT:
		return time.Now().Add(time.Minute)
	case model.TokenTerm_SHORT:
		return time.Now().Add(20 * time.Minute)
	case model.TokenTerm_MIDDLE:
		return time.Now().Add(2 * 24 * time.Hour)
	case model.TokenTerm_LONG:
		return time.Now().Add(2 * 30 * 24 * time.Hour)
	case model.TokenTerm_VERY_LONG:
		return time.Now().Add(2 * 365 * 24 * time.Hour)
	default:
		panic("unknown token term:" + term.String())
	}
}

func NewAuthenticationService(recordService service.RecordService, service service.AuthorizationService) service.AuthenticationService {
	return &authenticationService{
		recordService: recordService,
	}
}