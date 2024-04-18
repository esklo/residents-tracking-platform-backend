package app

import (
	"context"
	protoAuth "github.com/esklo/residents-tracking-platform/gen/proto/auth"
	protoContact "github.com/esklo/residents-tracking-platform/gen/proto/contact"
	protoDepartment "github.com/esklo/residents-tracking-platform/gen/proto/department"
	protoDistrict "github.com/esklo/residents-tracking-platform/gen/proto/district"
	protoFile "github.com/esklo/residents-tracking-platform/gen/proto/file"
	protoGeo "github.com/esklo/residents-tracking-platform/gen/proto/geo"
	protoRequest "github.com/esklo/residents-tracking-platform/gen/proto/request"
	protoTheme "github.com/esklo/residents-tracking-platform/gen/proto/theme"
	protoUser "github.com/esklo/residents-tracking-platform/gen/proto/user"
	"github.com/esklo/residents-tracking-platform/internal/app/provider"
	"github.com/esklo/residents-tracking-platform/internal/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

type App struct {
	serviceProvider *provider.ServiceProvider
	grpcServer      *grpc.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	go func() {
		err := a.runHTTPServer()
		if err != nil {
			log.Fatal(errors.Wrap(err, "http server error"))
		}
	}()
	return a.runGRPCServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCServer,
		a.migrateDatabase,
		a.bootstrapDatabase,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = provider.NewServiceProvider()
	return nil
}

func (a *App) initGRPCServer(_ context.Context) error {
	a.grpcServer = grpc.NewServer(
		grpc.MaxRecvMsgSize(50*1024*1024),
		grpc.MaxSendMsgSize(50*1024*1024),
	)
	reflection.Register(a.grpcServer)
	protoUser.RegisterUserServiceServer(a.grpcServer, a.serviceProvider.UserImpl())
	protoAuth.RegisterAuthServiceServer(a.grpcServer, a.serviceProvider.AuthImpl())
	protoGeo.RegisterGeoServiceServer(a.grpcServer, a.serviceProvider.GeoImpl())
	protoDistrict.RegisterDistrictServiceServer(a.grpcServer, a.serviceProvider.DistrictImpl())
	protoFile.RegisterFileServiceServer(a.grpcServer, a.serviceProvider.FileImpl())
	protoDepartment.RegisterDepartmentServiceServer(a.grpcServer, a.serviceProvider.DepartmentImpl())
	protoTheme.RegisterThemeServiceServer(a.grpcServer, a.serviceProvider.ThemeImpl())
	protoRequest.RegisterRequestServiceServer(a.grpcServer, a.serviceProvider.RequestImpl())
	protoContact.RegisterContactServiceServer(a.grpcServer, a.serviceProvider.ContactImpl())
	return nil
}

func (a *App) runHTTPServer() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	if err := protoUser.RegisterUserServiceHandlerFromEndpoint(ctx, mux, a.serviceProvider.GRPCConfig().Address(), opts); err != nil {
		return errors.Wrap(err, "failed to register user HTTP server")
	}

	if err := protoAuth.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, a.serviceProvider.GRPCConfig().Address(), opts); err != nil {
		return errors.Wrap(err, "failed to register auth HTTP server")
	}

	if err := protoGeo.RegisterGeoServiceHandlerFromEndpoint(ctx, mux, a.serviceProvider.GRPCConfig().Address(), opts); err != nil {
		return errors.Wrap(err, "failed to register geo HTTP server")
	}

	if err := protoDistrict.RegisterDistrictServiceHandlerFromEndpoint(ctx, mux, a.serviceProvider.GRPCConfig().Address(), opts); err != nil {
		return errors.Wrap(err, "failed to register district HTTP server")
	}

	if err := protoFile.RegisterFileServiceHandlerFromEndpoint(ctx, mux, a.serviceProvider.GRPCConfig().Address(), opts); err != nil {
		return errors.Wrap(err, "failed to register file HTTP server")
	}

	if err := protoDepartment.RegisterDepartmentServiceHandlerFromEndpoint(ctx, mux, a.serviceProvider.GRPCConfig().Address(), opts); err != nil {
		return errors.Wrap(err, "failed to register department HTTP server")
	}

	if err := protoTheme.RegisterThemeServiceHandlerFromEndpoint(ctx, mux, a.serviceProvider.GRPCConfig().Address(), opts); err != nil {
		return errors.Wrap(err, "failed to register theme HTTP server")
	}

	if err := protoRequest.RegisterRequestServiceHandlerFromEndpoint(ctx, mux, a.serviceProvider.GRPCConfig().Address(), opts); err != nil {
		return errors.Wrap(err, "failed to register request HTTP server")
	}

	if err := protoContact.RegisterContactServiceHandlerFromEndpoint(ctx, mux, a.serviceProvider.GRPCConfig().Address(), opts); err != nil {
		return errors.Wrap(err, "failed to register contact HTTP server")
	}

	withCors := cors.AllowAll().Handler(mux)

	log.Printf("HTTP server is running on %s", a.serviceProvider.HTTPConfig().Address())
	if err := http.ListenAndServe(a.serviceProvider.HTTPConfig().Address(), withCors); err != nil {
		return errors.Wrap(err, "failed to start HTTP server")
	}

	return nil
}

func (a *App) runGRPCServer() error {
	log.Printf("GRPC server is running on %s", a.serviceProvider.GRPCConfig().Address())

	list, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) migrateDatabase(_ context.Context) error {
	db, err := a.serviceProvider.DatabaseConnection()
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return errors.Wrap(err, "can not init postgres with instance")
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///app/migrations",
		"postgres", driver)
	if err != nil {
		return errors.Wrap(err, "can not create migrate db instance")
	}

	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return errors.Wrap(err, "can not complete migrations")
		}
	}
	return nil
}