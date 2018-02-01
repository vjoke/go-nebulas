// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: config.proto

/*
Package nebletpb is a generated protocol buffer package.

It is generated from these files:
	config.proto

It has these top-level messages:
	Config
	NetworkConfig
	ChainConfig
	RPCConfig
	AppConfig
	PprofConfig
	MiscConfig
	StatsConfig
	InfluxdbConfig
*/
package nebletpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Reporting modules.
type StatsConfig_ReportingModule int32

const (
	StatsConfig_Influxdb StatsConfig_ReportingModule = 0
)

var StatsConfig_ReportingModule_name = map[int32]string{
	0: "Influxdb",
}
var StatsConfig_ReportingModule_value = map[string]int32{
	"Influxdb": 0,
}

func (x StatsConfig_ReportingModule) String() string {
	return proto.EnumName(StatsConfig_ReportingModule_name, int32(x))
}
func (StatsConfig_ReportingModule) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorConfig, []int{7, 0}
}

// Neblet global configurations.
type Config struct {
	// Network config.
	Network *NetworkConfig `protobuf:"bytes,1,opt,name=network" json:"network,omitempty"`
	// Chain config.
	Chain *ChainConfig `protobuf:"bytes,2,opt,name=chain" json:"chain,omitempty"`
	// RPC config.
	Rpc *RPCConfig `protobuf:"bytes,3,opt,name=rpc" json:"rpc,omitempty"`
	// Stats config.
	Stats *StatsConfig `protobuf:"bytes,100,opt,name=stats" json:"stats,omitempty"`
	// Misc config.
	Misc *MiscConfig `protobuf:"bytes,101,opt,name=misc" json:"misc,omitempty"`
	// App Config.
	App *AppConfig `protobuf:"bytes,102,opt,name=app" json:"app,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

func (m *Config) GetNetwork() *NetworkConfig {
	if m != nil {
		return m.Network
	}
	return nil
}

func (m *Config) GetChain() *ChainConfig {
	if m != nil {
		return m.Chain
	}
	return nil
}

func (m *Config) GetRpc() *RPCConfig {
	if m != nil {
		return m.Rpc
	}
	return nil
}

func (m *Config) GetStats() *StatsConfig {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *Config) GetMisc() *MiscConfig {
	if m != nil {
		return m.Misc
	}
	return nil
}

func (m *Config) GetApp() *AppConfig {
	if m != nil {
		return m.App
	}
	return nil
}

type NetworkConfig struct {
	// Neb seed node address.
	Seed []string `protobuf:"bytes,1,rep,name=seed" json:"seed,omitempty"`
	// Listen addresses.
	Listen []string `protobuf:"bytes,2,rep,name=listen" json:"listen,omitempty"`
	// Network node privateKey address. If nil, generate a new node.
	PrivateKey string `protobuf:"bytes,3,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	// Network ID
	NetworkId uint32 `protobuf:"varint,4,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
}

func (m *NetworkConfig) Reset()                    { *m = NetworkConfig{} }
func (m *NetworkConfig) String() string            { return proto.CompactTextString(m) }
func (*NetworkConfig) ProtoMessage()               {}
func (*NetworkConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{1} }

func (m *NetworkConfig) GetSeed() []string {
	if m != nil {
		return m.Seed
	}
	return nil
}

func (m *NetworkConfig) GetListen() []string {
	if m != nil {
		return m.Listen
	}
	return nil
}

func (m *NetworkConfig) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

func (m *NetworkConfig) GetNetworkId() uint32 {
	if m != nil {
		return m.NetworkId
	}
	return 0
}

type ChainConfig struct {
	// ChainID.
	ChainId uint32 `protobuf:"varint,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// genesis conf file path
	Genesis string `protobuf:"bytes,2,opt,name=genesis,proto3" json:"genesis,omitempty"`
	// Data dir.
	Datadir string `protobuf:"bytes,11,opt,name=datadir,proto3" json:"datadir,omitempty"`
	// Key dir.
	Keydir string `protobuf:"bytes,12,opt,name=keydir,proto3" json:"keydir,omitempty"`
	// storage type mysql leveldb memory
	Storage string `protobuf:"bytes,13,opt,name=storage,proto3" json:"storage,omitempty"`
	// mysql DSN (Data Source Name)
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	MysqlDsn string `protobuf:"bytes,18,opt,name=mysql_dsn,json=mysqlDsn,proto3" json:"mysql_dsn,omitempty"`
	// mysql db
	MysqlDb string `protobuf:"bytes,19,opt,name=mysql_db,json=mysqlDb,proto3" json:"mysql_db,omitempty"`
	// start mine at launch
	StartMine bool `protobuf:"varint,20,opt,name=start_mine,json=startMine,proto3" json:"start_mine,omitempty"`
	// Coinbase.
	Coinbase string `protobuf:"bytes,21,opt,name=coinbase,proto3" json:"coinbase,omitempty"`
	// Miner.
	Miner string `protobuf:"bytes,22,opt,name=miner,proto3" json:"miner,omitempty"`
	// Passphrase.
	Passphrase string `protobuf:"bytes,23,opt,name=passphrase,proto3" json:"passphrase,omitempty"`
	// Lowest GasPrice.
	GasPrice string `protobuf:"bytes,24,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	// Max GasLimit.
	GasLimit string `protobuf:"bytes,25,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	// Supported signature cipher list. ["ECC_SECP256K1"]
	SignatureCiphers []string `protobuf:"bytes,26,rep,name=signature_ciphers,json=signatureCiphers" json:"signature_ciphers,omitempty"`
}

func (m *ChainConfig) Reset()                    { *m = ChainConfig{} }
func (m *ChainConfig) String() string            { return proto.CompactTextString(m) }
func (*ChainConfig) ProtoMessage()               {}
func (*ChainConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{2} }

func (m *ChainConfig) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *ChainConfig) GetGenesis() string {
	if m != nil {
		return m.Genesis
	}
	return ""
}

func (m *ChainConfig) GetDatadir() string {
	if m != nil {
		return m.Datadir
	}
	return ""
}

func (m *ChainConfig) GetKeydir() string {
	if m != nil {
		return m.Keydir
	}
	return ""
}

func (m *ChainConfig) GetStorage() string {
	if m != nil {
		return m.Storage
	}
	return ""
}

func (m *ChainConfig) GetMysqlDsn() string {
	if m != nil {
		return m.MysqlDsn
	}
	return ""
}

func (m *ChainConfig) GetMysqlDb() string {
	if m != nil {
		return m.MysqlDb
	}
	return ""
}

func (m *ChainConfig) GetStartMine() bool {
	if m != nil {
		return m.StartMine
	}
	return false
}

func (m *ChainConfig) GetCoinbase() string {
	if m != nil {
		return m.Coinbase
	}
	return ""
}

func (m *ChainConfig) GetMiner() string {
	if m != nil {
		return m.Miner
	}
	return ""
}

func (m *ChainConfig) GetPassphrase() string {
	if m != nil {
		return m.Passphrase
	}
	return ""
}

func (m *ChainConfig) GetGasPrice() string {
	if m != nil {
		return m.GasPrice
	}
	return ""
}

func (m *ChainConfig) GetGasLimit() string {
	if m != nil {
		return m.GasLimit
	}
	return ""
}

func (m *ChainConfig) GetSignatureCiphers() []string {
	if m != nil {
		return m.SignatureCiphers
	}
	return nil
}

type RPCConfig struct {
	// RPC listen addresses.
	RpcListen []string `protobuf:"bytes,1,rep,name=rpc_listen,json=rpcListen" json:"rpc_listen,omitempty"`
	// HTTP listen addresses.
	HttpListen []string `protobuf:"bytes,2,rep,name=http_listen,json=httpListen" json:"http_listen,omitempty"`
	// Enabled HTTP modules.["api", "admin"]
	HttpModule []string `protobuf:"bytes,3,rep,name=http_module,json=httpModule" json:"http_module,omitempty"`
}

func (m *RPCConfig) Reset()                    { *m = RPCConfig{} }
func (m *RPCConfig) String() string            { return proto.CompactTextString(m) }
func (*RPCConfig) ProtoMessage()               {}
func (*RPCConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{3} }

func (m *RPCConfig) GetRpcListen() []string {
	if m != nil {
		return m.RpcListen
	}
	return nil
}

func (m *RPCConfig) GetHttpListen() []string {
	if m != nil {
		return m.HttpListen
	}
	return nil
}

func (m *RPCConfig) GetHttpModule() []string {
	if m != nil {
		return m.HttpModule
	}
	return nil
}

type AppConfig struct {
	LogLevel string `protobuf:"bytes,1,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
	LogFile  string `protobuf:"bytes,2,opt,name=log_file,json=logFile,proto3" json:"log_file,omitempty"`
	// log file age, unit is s.
	LogAge            uint32 `protobuf:"varint,3,opt,name=log_age,json=logAge,proto3" json:"log_age,omitempty"`
	EnableCrashReport bool   `protobuf:"varint,4,opt,name=enable_crash_report,json=enableCrashReport,proto3" json:"enable_crash_report,omitempty"`
	CrashReportUrl    string `protobuf:"bytes,5,opt,name=crash_report_url,json=crashReportUrl,proto3" json:"crash_report_url,omitempty"`
	// pprof config
	Pprof   *PprofConfig `protobuf:"bytes,6,opt,name=pprof" json:"pprof,omitempty"`
	Version string       `protobuf:"bytes,100,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *AppConfig) Reset()                    { *m = AppConfig{} }
func (m *AppConfig) String() string            { return proto.CompactTextString(m) }
func (*AppConfig) ProtoMessage()               {}
func (*AppConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{4} }

func (m *AppConfig) GetLogLevel() string {
	if m != nil {
		return m.LogLevel
	}
	return ""
}

func (m *AppConfig) GetLogFile() string {
	if m != nil {
		return m.LogFile
	}
	return ""
}

func (m *AppConfig) GetLogAge() uint32 {
	if m != nil {
		return m.LogAge
	}
	return 0
}

func (m *AppConfig) GetEnableCrashReport() bool {
	if m != nil {
		return m.EnableCrashReport
	}
	return false
}

func (m *AppConfig) GetCrashReportUrl() string {
	if m != nil {
		return m.CrashReportUrl
	}
	return ""
}

func (m *AppConfig) GetPprof() *PprofConfig {
	if m != nil {
		return m.Pprof
	}
	return nil
}

func (m *AppConfig) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type PprofConfig struct {
	// pprof listen address, if not configured, the function closes.
	HttpListen string `protobuf:"bytes,1,opt,name=http_listen,json=httpListen,proto3" json:"http_listen,omitempty"`
	// cpu profiling file, if not configured, the profiling not start
	Cpuprofile string `protobuf:"bytes,2,opt,name=cpuprofile,proto3" json:"cpuprofile,omitempty"`
	// memory profiling file, if not configured, the profiling not start
	Memprofile string `protobuf:"bytes,3,opt,name=memprofile,proto3" json:"memprofile,omitempty"`
}

func (m *PprofConfig) Reset()                    { *m = PprofConfig{} }
func (m *PprofConfig) String() string            { return proto.CompactTextString(m) }
func (*PprofConfig) ProtoMessage()               {}
func (*PprofConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{5} }

func (m *PprofConfig) GetHttpListen() string {
	if m != nil {
		return m.HttpListen
	}
	return ""
}

func (m *PprofConfig) GetCpuprofile() string {
	if m != nil {
		return m.Cpuprofile
	}
	return ""
}

func (m *PprofConfig) GetMemprofile() string {
	if m != nil {
		return m.Memprofile
	}
	return ""
}

type MiscConfig struct {
	// Default encryption ciper when create new keystore file.
	DefaultKeystoreFileCiper string `protobuf:"bytes,1,opt,name=default_keystore_file_ciper,json=defaultKeystoreFileCiper,proto3" json:"default_keystore_file_ciper,omitempty"`
}

func (m *MiscConfig) Reset()                    { *m = MiscConfig{} }
func (m *MiscConfig) String() string            { return proto.CompactTextString(m) }
func (*MiscConfig) ProtoMessage()               {}
func (*MiscConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{6} }

func (m *MiscConfig) GetDefaultKeystoreFileCiper() string {
	if m != nil {
		return m.DefaultKeystoreFileCiper
	}
	return ""
}

type StatsConfig struct {
	// Enable metrics or not.
	EnableMetrics   bool                          `protobuf:"varint,1,opt,name=enable_metrics,json=enableMetrics,proto3" json:"enable_metrics,omitempty"`
	ReportingModule []StatsConfig_ReportingModule `protobuf:"varint,2,rep,packed,name=reporting_module,json=reportingModule,enum=nebletpb.StatsConfig_ReportingModule" json:"reporting_module,omitempty"`
	// Influxdb config.
	Influxdb    *InfluxdbConfig `protobuf:"bytes,11,opt,name=influxdb" json:"influxdb,omitempty"`
	MetricsTags []string        `protobuf:"bytes,12,rep,name=metrics_tags,json=metricsTags" json:"metrics_tags,omitempty"`
}

func (m *StatsConfig) Reset()                    { *m = StatsConfig{} }
func (m *StatsConfig) String() string            { return proto.CompactTextString(m) }
func (*StatsConfig) ProtoMessage()               {}
func (*StatsConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{7} }

func (m *StatsConfig) GetEnableMetrics() bool {
	if m != nil {
		return m.EnableMetrics
	}
	return false
}

func (m *StatsConfig) GetReportingModule() []StatsConfig_ReportingModule {
	if m != nil {
		return m.ReportingModule
	}
	return nil
}

func (m *StatsConfig) GetInfluxdb() *InfluxdbConfig {
	if m != nil {
		return m.Influxdb
	}
	return nil
}

func (m *StatsConfig) GetMetricsTags() []string {
	if m != nil {
		return m.MetricsTags
	}
	return nil
}

type InfluxdbConfig struct {
	// Host.
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// Port.
	Port uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	// Database name.
	Db string `protobuf:"bytes,3,opt,name=db,proto3" json:"db,omitempty"`
	// Auth user.
	User string `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	// Auth password.
	Password string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
}

func (m *InfluxdbConfig) Reset()                    { *m = InfluxdbConfig{} }
func (m *InfluxdbConfig) String() string            { return proto.CompactTextString(m) }
func (*InfluxdbConfig) ProtoMessage()               {}
func (*InfluxdbConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{8} }

func (m *InfluxdbConfig) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *InfluxdbConfig) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *InfluxdbConfig) GetDb() string {
	if m != nil {
		return m.Db
	}
	return ""
}

func (m *InfluxdbConfig) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *InfluxdbConfig) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*Config)(nil), "nebletpb.Config")
	proto.RegisterType((*NetworkConfig)(nil), "nebletpb.NetworkConfig")
	proto.RegisterType((*ChainConfig)(nil), "nebletpb.ChainConfig")
	proto.RegisterType((*RPCConfig)(nil), "nebletpb.RPCConfig")
	proto.RegisterType((*AppConfig)(nil), "nebletpb.AppConfig")
	proto.RegisterType((*PprofConfig)(nil), "nebletpb.PprofConfig")
	proto.RegisterType((*MiscConfig)(nil), "nebletpb.MiscConfig")
	proto.RegisterType((*StatsConfig)(nil), "nebletpb.StatsConfig")
	proto.RegisterType((*InfluxdbConfig)(nil), "nebletpb.InfluxdbConfig")
	proto.RegisterEnum("nebletpb.StatsConfig_ReportingModule", StatsConfig_ReportingModule_name, StatsConfig_ReportingModule_value)
}

func init() { proto.RegisterFile("config.proto", fileDescriptorConfig) }

var fileDescriptorConfig = []byte{
	// 870 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x55, 0x51, 0x6f, 0xdb, 0x36,
	0x10, 0x9e, 0xed, 0xc4, 0x91, 0xce, 0x8e, 0x9b, 0x32, 0x69, 0xc3, 0xb6, 0x58, 0x9b, 0x09, 0x08,
	0x60, 0xa0, 0x80, 0x81, 0x65, 0x7b, 0xdd, 0x43, 0xe1, 0x61, 0x40, 0x90, 0x64, 0x08, 0xb4, 0xed,
	0x59, 0xa0, 0x24, 0x5a, 0x26, 0x42, 0x4b, 0x1c, 0x49, 0xa7, 0x0d, 0xf6, 0xb2, 0x3f, 0xb0, 0xbf,
	0xb0, 0x7f, 0x3a, 0x60, 0xb8, 0x13, 0x65, 0x3b, 0x46, 0xdf, 0x78, 0xdf, 0xf7, 0xdd, 0x51, 0x3c,
	0x7e, 0x3c, 0xc1, 0xb8, 0x68, 0xea, 0x85, 0xaa, 0x66, 0xc6, 0x36, 0xbe, 0x61, 0x51, 0x2d, 0x73,
	0x2d, 0xbd, 0xc9, 0x93, 0x7f, 0xfa, 0x30, 0x9c, 0x13, 0xc5, 0xbe, 0x87, 0xa3, 0x5a, 0xfa, 0xcf,
	0x8d, 0x7d, 0xe0, 0xbd, 0x8b, 0xde, 0x74, 0x74, 0x75, 0x3e, 0xeb, 0x64, 0xb3, 0x5f, 0x5b, 0xa2,
	0x55, 0xa6, 0x9d, 0x8e, 0x7d, 0x84, 0xc3, 0x62, 0x29, 0x54, 0xcd, 0xfb, 0x94, 0xf0, 0x6a, 0x9b,
	0x30, 0x47, 0x38, 0xc8, 0x5b, 0x0d, 0xbb, 0x84, 0x81, 0x35, 0x05, 0x1f, 0x90, 0xf4, 0x74, 0x2b,
	0x4d, 0xef, 0xe7, 0x41, 0x88, 0x3c, 0xd6, 0x74, 0x5e, 0x78, 0xc7, 0xcb, 0xfd, 0x9a, 0xbf, 0x21,
	0xdc, 0xd5, 0x24, 0x0d, 0x9b, 0xc2, 0xc1, 0x4a, 0xb9, 0x82, 0x4b, 0xd2, 0x9e, 0x6d, 0xb5, 0x77,
	0xca, 0x15, 0x41, 0x4a, 0x0a, 0xdc, 0x5d, 0x18, 0xc3, 0x17, 0xfb, 0xbb, 0x7f, 0x32, 0xa6, 0xdb,
	0x5d, 0x18, 0x93, 0xfc, 0x05, 0xc7, 0xcf, 0xce, 0xca, 0x18, 0x1c, 0x38, 0x29, 0x4b, 0xde, 0xbb,
	0x18, 0x4c, 0xe3, 0x94, 0xd6, 0xec, 0x35, 0x0c, 0xb5, 0x72, 0x5e, 0xe2, 0xb9, 0x11, 0x0d, 0x11,
	0xfb, 0x00, 0x23, 0x63, 0xd5, 0xa3, 0xf0, 0x32, 0x7b, 0x90, 0x4f, 0x74, 0xd2, 0x38, 0x85, 0x00,
	0xdd, 0xc8, 0x27, 0xf6, 0x2d, 0x40, 0x68, 0x5d, 0xa6, 0x4a, 0x7e, 0x70, 0xd1, 0x9b, 0x1e, 0xa7,
	0x71, 0x40, 0xae, 0xcb, 0xe4, 0xdf, 0x01, 0x8c, 0x76, 0x1a, 0xc7, 0xde, 0x40, 0x44, 0xad, 0x43,
	0x71, 0x8f, 0xc4, 0x47, 0x14, 0x5f, 0x97, 0x8c, 0xc3, 0x51, 0x25, 0x6b, 0xe9, 0x94, 0xa3, 0xde,
	0xc7, 0x69, 0x17, 0x22, 0x53, 0x0a, 0x2f, 0x4a, 0x65, 0xf9, 0xa8, 0x65, 0x42, 0x88, 0x9f, 0xfd,
	0x20, 0x9f, 0x90, 0x18, 0x13, 0x11, 0x22, 0xcc, 0x70, 0xbe, 0xb1, 0xa2, 0x92, 0xfc, 0xb8, 0xcd,
	0x08, 0x21, 0x7b, 0x07, 0xf1, 0xea, 0xc9, 0xfd, 0xa9, 0xb3, 0xd2, 0xd5, 0x9c, 0x11, 0x17, 0x11,
	0xf0, 0xb3, 0xab, 0xf1, 0xeb, 0x02, 0x99, 0xf3, 0xd3, 0x36, 0xaf, 0xe5, 0x72, 0x3c, 0xa7, 0xf3,
	0xc2, 0xfa, 0x6c, 0xa5, 0x6a, 0xc9, 0xcf, 0x2e, 0x7a, 0xd3, 0x28, 0x8d, 0x09, 0xb9, 0x53, 0xb5,
	0x64, 0x6f, 0x21, 0x2a, 0x1a, 0x55, 0xe7, 0xc2, 0x49, 0xfe, 0xaa, 0xad, 0xda, 0xc5, 0xec, 0x0c,
	0x0e, 0x31, 0xc9, 0xf2, 0xd7, 0x44, 0xb4, 0x01, 0x7b, 0x0f, 0x60, 0x84, 0x73, 0x66, 0x69, 0x31,
	0xe7, 0x3c, 0x34, 0x76, 0x83, 0xe0, 0x87, 0x56, 0xc2, 0x65, 0xc6, 0xaa, 0x42, 0x72, 0xde, 0x96,
	0xac, 0x84, 0xbb, 0xc7, 0xb8, 0x23, 0xb5, 0x5a, 0x29, 0xcf, 0xdf, 0x6c, 0xc8, 0x5b, 0x8c, 0xd9,
	0x47, 0x78, 0xe9, 0x54, 0x55, 0x0b, 0xbf, 0xb6, 0x32, 0x2b, 0x94, 0x59, 0x4a, 0xeb, 0xf8, 0x5b,
	0xba, 0xd6, 0x93, 0x0d, 0x31, 0x6f, 0xf1, 0x44, 0x43, 0xbc, 0x71, 0x2b, 0x1e, 0xd2, 0x9a, 0x22,
	0x0b, 0x4e, 0x68, 0xfd, 0x11, 0x5b, 0x53, 0xdc, 0x6e, 0xcc, 0xb0, 0xf4, 0xde, 0x64, 0xcf, 0x9c,
	0x02, 0x08, 0xed, 0x09, 0x56, 0x4d, 0xb9, 0xd6, 0x92, 0x0f, 0xb6, 0x82, 0x3b, 0x42, 0x92, 0xff,
	0x7a, 0x10, 0x6f, 0xec, 0x89, 0xa7, 0xd0, 0x4d, 0x95, 0x69, 0xf9, 0x28, 0x35, 0xb9, 0x21, 0x4e,
	0x23, 0xdd, 0x54, 0xb7, 0x18, 0xe3, 0x5d, 0x20, 0xb9, 0x50, 0x5a, 0x76, 0x7e, 0xd0, 0x4d, 0xf5,
	0x8b, 0xd2, 0x92, 0x9d, 0x03, 0x2e, 0x33, 0xbc, 0xdd, 0x01, 0x79, 0x68, 0xa8, 0x9b, 0xea, 0x53,
	0x25, 0xd9, 0x0c, 0x4e, 0x65, 0x2d, 0x72, 0x2d, 0xb3, 0xc2, 0x0a, 0xb7, 0xcc, 0xac, 0x34, 0x8d,
	0xf5, 0xe4, 0xca, 0x28, 0x7d, 0xd9, 0x52, 0x73, 0x64, 0x52, 0x22, 0xd8, 0x14, 0x4e, 0x76, 0x85,
	0xd9, 0xda, 0x6a, 0x7e, 0x48, 0x7b, 0x4d, 0x8a, 0xad, 0xec, 0x0f, 0xab, 0xf1, 0x09, 0x1b, 0x63,
	0x9b, 0x05, 0x1f, 0xee, 0x3f, 0xe1, 0x7b, 0x84, 0xbb, 0x27, 0x4c, 0x1a, 0x74, 0xdf, 0xa3, 0xb4,
	0x4e, 0x35, 0x35, 0xbd, 0xf8, 0x38, 0xed, 0xc2, 0xa4, 0x86, 0xd1, 0x8e, 0x7e, 0xbf, 0xa1, 0x6d,
	0x0b, 0x76, 0x1b, 0xfa, 0x1e, 0xa0, 0x30, 0x6b, 0xcc, 0xd8, 0xb6, 0x61, 0x07, 0x41, 0x7e, 0x25,
	0x57, 0x1d, 0x1f, 0x5e, 0xe7, 0x16, 0x49, 0x6e, 0x00, 0xb6, 0x63, 0x83, 0xfd, 0x04, 0xef, 0x4a,
	0xb9, 0x10, 0x6b, 0xed, 0xf1, 0x31, 0xe3, 0x8b, 0x90, 0xd4, 0x5f, 0x34, 0x89, 0xb4, 0x61, 0x7b,
	0x1e, 0x24, 0x37, 0x41, 0x81, 0x1d, 0x9f, 0x23, 0x9f, 0xfc, 0xdd, 0x87, 0xd1, 0xce, 0xc0, 0x62,
	0x97, 0x30, 0x09, 0xdd, 0x5e, 0x49, 0x6f, 0x55, 0xe1, 0xa8, 0x42, 0x94, 0x1e, 0xb7, 0xe8, 0x5d,
	0x0b, 0xb2, 0x7b, 0x38, 0x69, 0xdb, 0xab, 0xea, 0xaa, 0x73, 0x06, 0x5a, 0x67, 0x72, 0x75, 0xf9,
	0xd5, 0x41, 0x38, 0x4b, 0x3b, 0x75, 0x6b, 0x9a, 0xf4, 0x85, 0x7d, 0x0e, 0xb0, 0x1f, 0x21, 0x52,
	0xf5, 0x42, 0xaf, 0xbf, 0x94, 0x39, 0x0d, 0x84, 0xd1, 0x15, 0xdf, 0x56, 0xba, 0x0e, 0x4c, 0xb8,
	0x92, 0x8d, 0x92, 0x7d, 0x07, 0xe3, 0xf0, 0x9d, 0x99, 0x17, 0x95, 0xe3, 0x63, 0x72, 0xe7, 0x28,
	0x60, 0xbf, 0x8b, 0xca, 0x25, 0x1f, 0xe0, 0xc5, 0xde, 0xe6, 0x6c, 0x0c, 0x51, 0x57, 0xf1, 0xe4,
	0x9b, 0xe4, 0x0b, 0x4c, 0x9e, 0xd7, 0xc7, 0x61, 0xba, 0x6c, 0x9c, 0x0f, 0xcd, 0xa3, 0x35, 0x62,
	0xe4, 0xbb, 0x3e, 0x99, 0x93, 0xd6, 0x6c, 0x02, 0xfd, 0x32, 0x0f, 0x37, 0xd4, 0x2f, 0x73, 0xd4,
	0xac, 0x9d, 0xb4, 0xe4, 0xcd, 0x38, 0xa5, 0x35, 0x0e, 0x11, 0x1c, 0x00, 0x9f, 0x1b, 0x5b, 0x06,
	0x1b, 0x6e, 0xe2, 0x7c, 0x48, 0xbf, 0xb9, 0x1f, 0xfe, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x29,
	0x8d, 0x23, 0xf6, 0x06, 0x00, 0x00,
}
