package conf

import (
	"flag"
	"github.com/pelletier/go-toml/v2"
	"github.com/shopspring/decimal"
	"github.com/spf13/cast"
	"math"
	"os"
	"strings"
	"time"
)

const (
	defaultExpireTime          = 600      // 订单默认有效期 10分钟
	DefaultUsdtCnyRate         = 6.4      // 默认USDT汇率
	DefaultTrxCnyRate          = 0.95     // 默认TRX汇率
	defaultAuthToken           = "123234" // 默认授权码
	defaultListen              = ":8080"  // 默认监听地址
	defaultPaymentMinAmount    = 0.01
	defaultPaymentMaxAmount    = 99999
	defaultUsdtAtomicity       = 0.01 // 原子精度
	defaultTrxAtomicity        = 0.01
	defaultTronGrpcNode        = "18.141.79.38:50051"               // 默认GRPC节点
	defaultPolygonRpcEndpoint  = "https://polygon-rpc.com/"         // 默认Polygon RPC节点
	defaultEthereumRpcEndpoint = "https://ethereum.publicnode.com/" // 默认Ethereum RPC节点
	defaultOutputLog           = "/var/log/bepusdt.log"             // 默认日志输出文件
	defaultSqlitePath          = "/var/lib/bepusdt/sqlite.db"       // 默认数据库文件
)

const (
	Tron     = "tron"
	Polygon  = "polygon"
	Ethereum = "ethereum"
)

var (
	cfg  Conf
	path string
)

func Init() {
	flag.StringVar(&path, "conf", "./conf.toml", "config file path")
	flag.Parse()

	data, err := os.ReadFile(path)
	if err != nil {

		panic("配置文件加载失败：" + err.Error())
	}

	if err = toml.Unmarshal(data, &cfg); err != nil {

		panic("配置数据解析失败：" + err.Error())
	}
}

func GetUsdtRate() string {
	if cfg.Pay.UsdtRate != "" {

		return cfg.Pay.UsdtRate
	}

	return cast.ToString(DefaultUsdtCnyRate)
}

func GetTrxRate() string {
	if cfg.Pay.TrxRate != "" {

		return cfg.Pay.TrxRate
	}

	return cast.ToString(DefaultTrxCnyRate)
}

func GetUsdtAtomicity() (decimal.Decimal, int) {
	var val = defaultUsdtAtomicity
	if cfg.Pay.UsdtAtom != 0 {

		val = cfg.Pay.UsdtAtom
	}

	var atom = decimal.NewFromFloat(val)

	return atom, cast.ToInt(math.Abs(float64(atom.Exponent())))
}

func GetTrxAtomicity() (decimal.Decimal, int) {
	var val = defaultTrxAtomicity
	if cfg.Pay.TrxAtom != 0 {

		val = cfg.Pay.TrxAtom
	}

	var atom = decimal.NewFromFloat(val)

	return atom, cast.ToInt(math.Abs(float64(atom.Exponent())))
}

func GetExpireTime() time.Duration {
	if cfg.Pay.ExpireTime == 0 {

		return time.Duration(defaultExpireTime)
	}

	return time.Duration(cfg.Pay.ExpireTime)
}

func GetAuthToken() string {
	if cfg.AuthToken == "" {

		return defaultAuthToken
	}

	return cfg.AuthToken
}

func GetAppUri(host string) string {
	if cfg.AppUri != "" {

		return cfg.AppUri
	}

	return host
}

func GetStaticPath() string {

	return cfg.StaticPath
}

func GetSqlitePath() string {
	if cfg.SqlitePath != "" {

		return cfg.SqlitePath
	}

	return defaultSqlitePath
}

func GetOutputLog() string {
	if cfg.OutputLog != "" {

		return cfg.OutputLog
	}

	return defaultOutputLog
}

func GetListen() string {
	if cfg.Listen != "" {

		return cfg.Listen
	}

	return defaultListen
}

func GetTronGrpcNode() string {
	if cfg.TronGrpcNode != "" {

		return cfg.TronGrpcNode
	}

	return defaultTronGrpcNode
}

func GetPolygonRpcEndpoint() string {
	if cfg.PolygonRpcEndpoint != "" {

		return cfg.PolygonRpcEndpoint
	}

	return defaultPolygonRpcEndpoint
}

func GetEthereumRpcEndpoint() string {
	if cfg.EthereumRpcEndpoint != "" {

		return cfg.EthereumRpcEndpoint
	}

	return defaultEthereumRpcEndpoint
}

func BotToken() string {
	var token = strings.TrimSpace(os.Getenv("BOT_TOKEN"))
	if token != "" {

		return token
	}

	return cfg.Bot.Token
}

func BotAdminID() int64 {
	var id = strings.TrimSpace(os.Getenv("BOT_ADMIN_ID"))
	if id != "" {

		return cast.ToInt64(id)
	}

	return cfg.Bot.AdminID
}

func BotNotifyTarget() string {
	if cfg.Bot.GroupID != "" {

		return cfg.Bot.GroupID
	}

	return cast.ToString(cfg.Bot.AdminID)
}

func GetWalletAddress() []string {

	return cfg.Pay.WalletAddress
}

func GetTradeIsConfirmed() bool {

	return cfg.Pay.TradeIsConfirmed
}

func GetPaymentAmountMin() decimal.Decimal {
	var val = defaultPaymentMinAmount
	if cfg.Pay.PaymentAmountMin != 0 {

		val = cfg.Pay.PaymentAmountMin
	}

	return decimal.NewFromFloat(val)
}

func GetPaymentAmountMax() decimal.Decimal {
	var val float64 = defaultPaymentMaxAmount
	if cfg.Pay.PaymentAmountMax != 0 {

		val = cfg.Pay.PaymentAmountMax
	}

	return decimal.NewFromFloat(val)
}
