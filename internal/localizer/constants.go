package localizer

const (
	UseContextCommand          = "sqlcmd config use-context mssql"
	PasswordEnvVar             = "SQLCMD_PASSWORD"
	PasswordEnvVar2            = "SQLCMDPASSWORD"
	EndpointFlag               = "--endpoint"
	FeedbackUrl                = "https://aka.ms/sqlcmd-feedback"
	PasswordEncryptFlag        = "--password-encryption"
	AuthTypeFlag               = "--auth-type"
	ModernAuthTypeBasic        = "basic"
	ModernAuthTypeOther        = "other"
	UserNameFlag               = "--username"
	NameFlag                   = "--name"
	GetContextCommand          = "sqlcmd config get-contexts"
	GetEndpointsCommand        = "sqlcmd config get-endpoints"
	GetUsersCommand            = "sqlcmd config get-users"
	RunQueryExample            = "sqlcmd query \"SELECT @@SERVERNAME\""
	UninstallCommand           = "sqlcmd uninstall"
	AcceptEulaFlag             = "--accept-eula"
	AcceptEulaEnvVar           = "SQLCMD_ACCEPT_EULA"
	PodmanPsCommand            = "podman ps"
	DockerPsCommand            = "docker ps"
	HelpFlag                   = "--help"
	QueryAndExitFlag           = "-Q"
	QueryFlag                  = "-q"
	DbNameVar                  = "SQLCMDDBNAME"
	BatchTerminatorGo          = "GO"
	ConnStrPattern             = "[[tcp:]|[lpc:]|[np:]]server[\\instance_name][,port]"
	ServerEnvVar               = "SQLCMDSERVER"
	InsertKeyword              = "INSERT"
	PacketSizeVar              = "SQLCMDPACKETSIZE"
	LoginTimeOutVar            = "SQLCMDLOGINTIMEOUT"
	WorkstationVar             = "SQLCMDWORKSTATION"
	ApplicationIntentFlagShort = "-K"
	DosErrorLevel              = "DOS ERRORLEVEL"
	StdoutName                 = "stdout"
	ColSeparatorVar            = "SQLCMDCOLSEP"
	ErrorLevel                 = "ERRORLEVEL"
	AppIntentValues            = `"readonly"`
	FormatValues               = "\"horiz\",\"horizontal\",\"vert\",\"vertical\""
	ErrToStderrValues          = "\"-1\",\"0\",\"1\""
)