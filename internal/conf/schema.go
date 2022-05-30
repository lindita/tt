package conf

import (
	"fmt"
	"time"
)

type Config struct {
	Runtime        Runtime          `yaml:"runtime"`
	Database       DataBase         `yaml:"database"`
	HTTPServerAddr string           `yaml:"http_server_addr" mapstructure:"http_server_addr"`
	Redis          RedisConfig      `yaml:"redis"`
	Alipay         AlipayConfig     `yaml:"alipay"`
	OpenSearch     OpenSearchConfig `yaml:"open_search" mapstructure:"open_search"`
	SLS            AliyunSLS        `yaml:"aliyun_sls" mapstructure:"aliyun_sls"`
	Qiniu          Qiniu            `yaml:"qiniu"`
	Xsj            Xsj              `yaml:"xsj_web" mapstructure:"xsj_web"`
}

// Runtime 运行环境变量
type Runtime string

// DataBase 数据库配置
type DataBase struct {
	Default MySQLHAConfig `yaml:"default"`
}

// MySQLHAConfig 主从配置
type MySQLHAConfig struct {
	Master  MySQLConfig  `yaml:"master"`
	Slave   *MySQLConfig `yaml:"slave"`
	Slave01 *MySQLConfig `yaml:"slave01"`
}

type Xsj struct {
	App    string `yaml:"app"`
	Secret string `yaml:"secret"`
	Host   string `yaml:"host"`
}

type Qiniu struct {
	AccessKey     string `yaml:"access_key" mapstructure:"access_key"`
	SecretKey     string `yaml:"secret_key" mapstructure:"secret_key"`
	Bucket        string `yaml:"bucket"`
	Domain        string `yaml:"domain"`
	DomainApp     string `yaml:"domain_app" mapstructure:"domain_app"`
	SourceDomain  string `yaml:"source_domain" mapstructure:"source_domain"`
	DomainComment string `yaml:"domain_comment" mapstructure:"domain_comment"`
}

// MySQLConfig 数据库配置
type MySQLConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
	Dbname   string `yaml:"charset"`
	ConnMax  int    `yaml:"connMax"`
	ConnMin  int    `yaml:"connMin"`
	Timeout  string `yaml:"timeout"`
}

type RedisConfig struct {
	Default   ServerInfo   `yaml:"default"`
	HsDefault ServerInfo   `yaml:"hs_default" mapstructure:"hs_default"`
	Policy    ServerPolicy `yaml:"policy"`
}

type AlipayConfig struct {
	Pay  AlipayInfo `yaml:"pay"`
	Auth AlipayInfo `yaml:"auth"`
}

type OpenSearchConfig struct {
	EndPoint        string `yaml:"endpoint"`
	AccessKeyId     string `yaml:"access_key_id" mapstructure:"access_key_id"`
	AccessKeySecret string `yaml:"access_key_secret" mapstructure:"access_key_secret"`
}

type QiNiu struct {
	AccessKey     string `yaml:"access_key"  mapstructure:"access_key"`
	SecretKey     string `yaml:"secret_key"  mapstructure:"secret_key"`
	Bucket        string `yaml:"bucket"`
	Domain        string `yaml:"domain"`
	DomainApp     string `yaml:"domain_app"  mapstructure:"domain_app"`
	SourceDomain  string `yaml:"source_domain"  mapstructure:"source_domain"`
	DomainComment string `yaml:"domain_comment"  mapstructure:"domain_comment"`

	EndPoint        string `yaml:"endpoint"`
	AccessKeyId     string `yaml:"access_key_id" mapstructure:"access_key_id"`
	AccessKeySecret string `yaml:"access_key_secret" mapstructure:"access_key_secret"`
}

type ServerPolicy struct {
	// Maximum number of retries before giving up.
	// Default is to not retry failed commands.
	MaxRetries int
	// Minimum backoff between each retry.
	// Default is 8 milliseconds; -1 disables backoff.
	MinRetryBackoff time.Duration
	// Maximum backoff between each retry.
	// Default is 512 milliseconds; -1 disables backoff.
	MaxRetryBackoff time.Duration

	// Dial timeout for establishing new connections.
	// Default is 5 seconds.
	DialTimeout time.Duration `yaml:"timeout_connect"`
	// Timeout for socket reads. If reached, commands will fail
	// with a timeout instead of blocking. Use value -1 for no timeout and 0 for default.
	// Default is 3 seconds.
	ReadTimeout time.Duration `yaml:"timeout_read"`
	// Timeout for socket writes. If reached, commands will fail
	// with a timeout instead of blocking.
	// Default is ReadTimeout.
	WriteTimeout time.Duration `yaml:"timeout_write"`

	// Maximum number of socket connections.
	// Default is 10 connections per every CPU as reported by runtime.NumCPU.
	PoolSize int
	// Minimum number of idle connections which is useful when establishing
	// new connection is slow.
	MinIdleConns int
	// Connection age at which client retires (closes) the connection.
	// Default is to not close aged connections.
	MaxConnAge time.Duration
	// Amount of time client waits for connection if all connections
	// are busy before returning an error.
	// Default is ReadTimeout + 1 second.
	PoolTimeout time.Duration
	// Amount of time after which client closes idle connections.
	// Should be less than server's timeout.
	// Default is 5 minutes. -1 disables idle timeout check.
	IdleTimeout time.Duration
	// Frequency of idle checks made by idle connections reaper.
	// Default is 1 minute. -1 disables idle connections reaper,
	// but idle connections are still discarded by the client
	// if IdleTimeout is set.
	IdleCheckFrequency time.Duration
}

type ServerInfo struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Timeout  int64  `yaml:"timeout"`
	PoolSize int    `yaml:"pool_size"`
}

type AlipayInfo struct {
	AppId                   string `yaml:"app_id" mapstructure:"app_id"`
	PrivateKey              string `yaml:"private_key" mapstructure:"private_key"`
	AppCertPublicKeyPath    string `yaml:"app_cert_public_key_path" mapstructure:"app_cert_public_key_path"`
	AlipayRootCertPath      string `yaml:"alipay_root_cert_path" mapstructure:"alipay_root_cert_path"`
	AlipayCertPublicKeyRSA2 string `yaml:"alipay_cert_public_key_rsa_2" mapstructure:"alipay_cert_public_key_rsa_2"`
	AlipayUserId            string `yaml:"alipay_user_id" mapstructure:"alipay_user_id"`
}

func (t ServerInfo) GetAddr() string {
	return t.Host + ":" + fmt.Sprintf("%d", t.Port)
}

// AliyunSLS AliyunSls aliyun sls
type AliyunSLS struct {
	Key      string `yaml:"key"`
	Secret   string `yaml:"secret"`
	EndPoint string `yaml:"endpoint"`
	Project  string `yaml:"project"`
	Store    string `yaml:"store"`
}
