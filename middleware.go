package warrant

import (
	"log"
	"net/http"

	"github.com/warrant-dev/warrant-go/config"
)

type GetObjectIdFunc func(r *http.Request) string

type GetUserIdFunc func(r *http.Request) string

type NewEnsureIsAuthorizedFunc func(handler http.Handler, options EnsureIsAuthorizedOptions) *EnsureIsAuthorized

type NewEnsureHasPermissionFunc func(handler http.Handler, options EnsureHasPermissionOptions) *EnsureHasPermission

type MiddlewareConfig struct {
	ApiKey         string
	GetObjectId    GetObjectIdFunc
	GetUserId      GetUserIdFunc
	OnAccessDenied http.HandlerFunc
}

type Middleware struct {
	config MiddlewareConfig
	client Client
}

type EnsureIsAuthorizedOptions struct {
	ObjectType string
	ObjectId   string
	Relation   string
	UserId     string
}

type EnsureIsAuthorized struct {
	handler http.Handler
	mw      Middleware
	options EnsureIsAuthorizedOptions
}

type EnsureHasPermissionOptions struct {
	PermissionId string
	UserId       string
}

type EnsureHasPermission struct {
	handler http.Handler
	mw      Middleware
	options EnsureHasPermissionOptions
}

func defaultOnAccessDenied(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
}

func (eia *EnsureIsAuthorized) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	objectId := eia.options.ObjectId
	userId := eia.options.UserId

	if objectId == "" {
		objectId = eia.mw.config.GetObjectId(r)
	}

	if userId == "" {
		userId = eia.mw.config.GetUserId(r)
	}

	isAuthorized, err := eia.mw.client.Check(&WarrantCheckParams{
		WarrantCheck: WarrantCheck{
			Object: Object{
				ObjectType: eia.options.ObjectType,
				ObjectId:   objectId,
			},
			Relation: eia.options.Relation,
			Subject: Subject{
				ObjectType: "user",
				ObjectId:   userId,
			},
		},
	})
	if err != nil {
		log.Println(err)
		eia.mw.config.OnAccessDenied(w, r)
		return
	}

	if !isAuthorized {
		eia.mw.config.OnAccessDenied(w, r)
		return
	}

	eia.handler.ServeHTTP(w, r)
}

func (mw Middleware) NewEnsureIsAuthorized(handler http.Handler, options EnsureIsAuthorizedOptions) *EnsureIsAuthorized {
	if options.ObjectId == "" && mw.config.GetObjectId == nil {
		panic("You must either provide GetObjectId to the Warrant middleware when calling NewMiddleware or provide an ObjectId when calling NewEnsureIsAuthorized.")
	}

	if options.UserId == "" && mw.config.GetUserId == nil {
		panic("You must either provide GetUserId to the Warrant middleware when calling NewMiddleware or provide a UserId when calling NewEnsureIsAuthorized.")
	}

	return &EnsureIsAuthorized{
		handler: handler,
		mw:      mw,
		options: options,
	}
}

func (ehp *EnsureHasPermission) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userId := ehp.options.UserId
	if userId == "" {
		userId = ehp.mw.config.GetUserId(r)
	}

	isAuthorized, err := ehp.mw.client.CheckUserHasPermission(&PermissionCheckParams{
		PermissionId: ehp.options.PermissionId,
		UserId:       userId,
	})
	if err != nil {
		log.Println(err)
		ehp.mw.config.OnAccessDenied(w, r)
		return
	}

	if !isAuthorized {
		ehp.mw.config.OnAccessDenied(w, r)
		return
	}

	ehp.handler.ServeHTTP(w, r)
}

func (mw Middleware) NewEnsureHasPermission(handler http.Handler, options EnsureHasPermissionOptions) *EnsureHasPermission {
	return &EnsureHasPermission{
		handler: handler,
		mw:      mw,
		options: options,
	}
}

func NewMiddleware(middlewareConfig MiddlewareConfig) *Middleware {
	if middlewareConfig.OnAccessDenied == nil {
		middlewareConfig.OnAccessDenied = defaultOnAccessDenied
	}

	return &Middleware{
		config: middlewareConfig,
		client: NewClient(config.ClientConfig{
			ApiKey:                  middlewareConfig.ApiKey,
			ApiEndpoint:             ApiEndpoint,
			AuthorizeEndpoint:       AuthorizeEndpoint,
			SelfServiceDashEndpoint: SelfServiceDashEndpoint,
		}),
	}
}
