package configs

const (
	// FrontendServerIP   = "http://212.233.89.163"
	FrontendServerIP = "https://my-kopilka.ru"
	FrontendServerPort = ""
	// FrontendServerIP = "http://localhost"
	// FrontendServerPort = ":8000"
	BackendServerPort  = ":8001"

	DriverSQL            = "pgx"
	DatabaseDMS          = "postgres"
	DatabaseUser         = "kopilka"
	DatabaseUserPassword = "12345"
	// DatabaseServerIP     = "localhost"
	DatabaseServerIP = "2023_2_chaihona_no1-postgres-1"
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
	PostTable           = "public.post"
	AttachTable         = "public.post_attach"
	LikeTable           = "public.post_like"
	PaymentTable        = "public.payment"
	CommentTable = "public.post_comment"
	AnalitycsTable = "public.analytics"

	RedisServerIP   = "127.0.0.1"
	RedisServerPort = "6379"
	RedisServerURL = "2023_2_chaihona_no1-redis-1"

	PaymentAPI      = "https://api.yookassa.ru/v3/"
	PaymentKeyPath  = "API_key"
	ShopId          = "294126"
	ReturnURL       = FrontendServerIP + FrontendServerPort + "/payment"
	FakeRedirectURL = "https://yoomoney.ru/payments/external/confirmation?orderId=22e12f66-000f-5000-8000-18db351245c7"

	BasePath = "~/go/src/github.com/M0rdovorot/kopilka"
)