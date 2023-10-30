package configs

const (
	FrontendServerIP     = "http://212.233.89.163"
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
	PaymentTable        = "public.payment"
	SubscriptionTable   = "public.subscription"
	SubscribeLevelTable = "public.subscription_level"
	UserTable           = "public.user"
	SourceDriver        = "file://"
	PaymentAPI          = "https://api.yookassa.ru/v3/"
	PaymentKeyPath      = "API_key"
	ShopId              = "273632"
	ReturnURL           = FrontendServerIP + FrontendServerPort + "/payment"
	FakeRedirectURL     = "https://yoomoney.ru/payments/external/confirmation?orderId=22e12f66-000f-5000-8000-18db351245c7"
)
