package config

type LoG struct {
	AgencyLinkkey   string `mapstructure:"agencyLinkkey" json:"agencyLinkkey" yaml:"agencyLinkkey"`       // 机构linkkey
	EncryptLinkkey  string `mapstructure:"encryptLinkkey" json:"encryptLinkkey" yaml:"encryptLinkkey"`    // 加解密linkkey
	ApiUrl          string `mapstructure:"apiUrl" json:"apiUrl" yaml:"apiUrl"`                            // 接口地址
	AgencyPublicKey string `mapstructure:"agencyPublicKey" json:"agencyPublicKey" yaml:"agencyPublicKey"` // 公钥
	PublicKey       string `mapstructure:"publicKey" json:"publicKey" yaml:"publicKey"`                   //验签公钥
	PrivateKey      string `mapstructure:"privateKey" json:"privateKey" yaml:"privateKey"`                //验签私钥
}
