package configs

const (
	FrontendServerIP     = "http://localhost"
	FrontendServerPort   = ":8000"
	BackendServerPort    = ":8001"
	DriverSQL            = "pgx"
	DatabaseDMS          = "postgres"
	DatabaseUser         = "kopilka"
	DatabaseUserPassword = "12345"
	DatabaseServerIP     = "localhost"
	DatabaseServerPort   = "5432"
	DatabaseName         = "kopilka"
	DatabaseURL          = DatabaseDMS + "://" + DatabaseUser +
		":" + DatabaseUserPassword + "@" + DatabaseServerIP +
		":" + DatabaseServerPort + "/" + DatabaseName
	MigrationsPath      = "db/migrations"
	SourceDriver        = "file://"
	UserTable           = "public.user"
	SubscribeLevelTable = "public.subscription_level"
	SubscriptionTable   = "public.subscription"
)
