package cfg

import (
	"errors"
	"github.com/buzhiyun/go-utils/file"
	"github.com/buzhiyun/go-utils/log"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type configFile struct {
	fileName string // Default configuration file name.
	fileType string
	//searchPaths   *[]string // Searching path array.
	configMap map[string]interface{} // The pared JSON objects for configuration files.
	available bool

	cacheString  map[string]string
	cacheStrings map[string][]string
	cacheInt     map[string]int
	cacheInt64   map[string]int64
	cacheBool    map[string]bool
}

var (
	supportedFileTypes = []string{"yaml", "yml"}
	searchPath         = []string{
		file.GetAppDir(),
		filepath.Join(file.GetAppDir(), "config"),
		file.GetWorkDir(),
		file.GetHomeDir(),
		os.TempDir(),
	}
	defaultconfigFileName = "config"
)

// 读取新的配置文件
func newConfigFile(file ...string) (*configFile, error) {
	var (
		name = defaultconfigFileName
	)

	if len(file) > 0 {
		name = file[0]
	}

	configPath, configType, ok := configFilePath(name)

	if !ok {
		log.Errorf("没有找到配置文件 %s", name)
		return nil, errors.New("没有找到配置文件")
	}

	var cfgData map[string]interface{}

	// 对 yaml 处理
	log.Debugf("加载配置文件 %s", configPath)

	if configType == "yaml" || configType == "yml" {
		if c, err := os.ReadFile(configPath); err != nil {
			log.Errorf("读取配置文件错误 %s ", err.Error())
			return nil, err
		} else {
			if err = yaml.Unmarshal(c, &cfgData); err != nil {
				log.Errorf("读取配置文件错误 %s , %s", configPath, err.Error())
				return nil, err
			}
		}

	}

	cfg := &configFile{
		fileName:     configPath,
		fileType:     configType,
		configMap:    cfgData,
		available:    true,
		cacheString:  map[string]string{},
		cacheStrings: map[string][]string{},
		cacheInt:     map[string]int{},
		cacheInt64:   map[string]int64{},
		cacheBool:    map[string]bool{},
	}

	return cfg, nil
}

// 搜索配置文件路径
func configFilePath(filename string) (configPath, configType string, exits bool) {
	for _, dir := range searchPath {
		for _, fileType := range supportedFileTypes {
			cPath := filepath.Join(dir, filename+"."+fileType)
			if file.FileExist(cPath) {
				return cPath, fileType, true
			}
		}
	}
	return "", "", false

}

// 获取配置的值
func (c *configFile) Get(pattern string) (value interface{}, success bool) {
	value, _ok := cache.Load(pattern)
	if _ok {
		return value, _ok
	}

	// 没有指定文件的情况下，优先从环境变量里面找
	_v := os.Getenv(pattern)
	if len(_v) > 0 {
		return _v, true
	}

	cfgMap := c.configMap
	var (
		v  interface{}
		ok bool
	)
	for idx, key := range strings.Split(pattern, ".") {

		if idx == 0 {
			v, ok = cfgMap[key]
			if !ok {
				return nil, false
			}
		} else {
			cMap, succ := v.(map[string]interface{})
			if !succ {
				return nil, false
			}
			v, succ = cMap[key]
			if !ok {
				return nil, false
			}
			ok = succ
		}
	}
	if ok {
		cache.Store(pattern, v)
	}

	return v, ok
}

// 扫描字符串值
func (c *configFile) GetString(pattern string) (value string, ok bool) {
	if v, _ok := c.cacheString[pattern]; _ok {
		return v, _ok
	}

	// 没有指定文件的情况下，优先从环境变量里面找
	_v := os.Getenv(pattern)
	if len(_v) > 0 {
		return _v, true
	}

	v, ok := c.Get(pattern)
	if ok {
		value, ok = v.(string)
		c.cacheString[pattern] = value

		return
	}
	return
}

// 扫描字符串值
func (c *configFile) GetStrings(pattern string) (value []string, ok bool) {
	if v, _ok := c.cacheStrings[pattern]; _ok {
		return v, _ok
	}

	v, ok := c.Get(pattern)
	if ok {
		interfaceValues, ok := v.([]interface{})
		if !ok {
			return value, false
		}
		for _, v1 := range interfaceValues {
			stringValue, ok := v1.(string)
			if !ok {
				return value, false
			}
			value = append(value, stringValue)

		}
		c.cacheStrings[pattern] = value
		return value, true
	}
	return
}

// 扫描int64值
func (c *configFile) GetInt64(pattern string) (value int64, ok bool) {
	if v, _ok := c.cacheInt64[pattern]; _ok {
		return v, _ok
	}

	// 没有指定文件的情况下，优先从环境变量里面找
	_v := os.Getenv(pattern)
	if len(_v) > 0 {
		if __v, err := strconv.ParseInt(_v, 10, 64); err == nil {
			return __v, true
		}
	}

	v, ok := c.Get(pattern)
	if ok {
		value, ok = v.(int64)

		if ok {
			c.cacheInt64[pattern] = value
		}
		return
	}
	return
}

// 扫描int值
func (c *configFile) GetInt(pattern string) (value int, ok bool) {
	if v, _ok := c.cacheInt[pattern]; _ok {
		return v, _ok
	}

	// 没有指定文件的情况下，优先从环境变量里面找
	_v := os.Getenv(pattern)
	if len(_v) > 0 {
		if __v, err := strconv.Atoi(_v); err == nil {
			return __v, true
		}
	}

	v, ok := c.Get(pattern)
	if ok {
		value, ok = v.(int)
		if ok {
			c.cacheInt[pattern] = value
		}
		return
	}
	return
}

// 扫描bool值
func (c *configFile) GetBool(pattern string) (value bool, ok bool) {
	if v, _ok := c.cacheBool[pattern]; _ok {
		return v, _ok
	}

	// 没有指定文件的情况下，优先从环境变量里面找
	_v := os.Getenv(pattern)
	if len(_v) > 0 {
		_v = strings.ToLower(_v)
		if _v == "true" {
			return true, true
		} else if _v == "false" {
			return false, true
		}
	}

	v, ok := c.Get(pattern)
	if ok {
		value, ok = v.(bool)
		if ok {
			c.cacheBool[pattern] = value
		}
		return
	}
	return
}

// 扫描bool值
func (c *configFile) Scan(pattern string, out interface{}) (ok bool) {
	v, ok := c.Get(pattern)
	if ok {
		if err := mapstructure.Decode(v, &out); err != nil {
			log.Errorf("加载 %s 出错, %v", pattern, v)
			return false
		}
		return true
	}
	return
}

// 重载文件
func (c *configFile) Reload() (err error) {
	c.available = false
	var cfgData map[string]interface{}
	log.Debugf("加载配置文件 %s", c.fileName)

	// 对 yaml 处理
	if c.fileType == "yaml" {
		if conf, err := os.ReadFile(c.fileName); err != nil {
			log.Errorf("读取配置文件错误 %s ", err.Error())
			return err
		} else {
			if err = yaml.Unmarshal(conf, &cfgData); err != nil {
				log.Errorf("读取配置文件错误 %s , %s", c.fileName, err.Error())
				return err
			}
		}

		c.configMap = cfgData
	}
	c.available = true

	// 清缓存
	c.cacheString = map[string]string{}
	c.cacheStrings = map[string][]string{}
	c.cacheInt = map[string]int{}
	c.cacheInt64 = map[string]int64{}
	c.cacheBool = map[string]bool{}

	return
}

// 配置是否可用
func (c *configFile) Available() (ok bool) {
	return c.available
}
