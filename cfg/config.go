package cfg

type Conf interface {
	Available() (ok bool)

	// Get 获取配置的内容 ， ok 的话就说明是读取到了
	// pattern 格式，用 . 来做分级 例如： a.b.c
	Get(pattern string) (value interface{}, ok bool)

	// 重新加载配置文件
	Reload() (err error)

	// 扫描值
	GetString(pattern string) (string, bool)
	GetInt(pattern string) (int, bool)
	GetInt64(pattern string) (int64, bool)
	GetBool(pattern string) (bool, bool)
}

var config map[string]Conf

func Config(filename ...string) Conf {
	var fName = ""
	if len(filename) > 0 {
		fName = filename[0]
	}
	if cfg, ok := config[fName]; ok {
		if cfg.Available() {
			return cfg
		}
	}

	cfgFile, err := NewConfigFile(filename...)
	if err != nil {
		return &ConfigFile{available: false}
	}
	if config == nil {
		config = make(map[string]Conf)
	}
	config[fName] = cfgFile
	return cfgFile
}
