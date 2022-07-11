package cfg

import (
	"errors"
	"github.com/buzhiyun/go-utils/file"
	"github.com/kataras/golog"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type ConfigFile struct {
	fileName string // Default configuration file name.
	fileType string
	//searchPaths   *[]string // Searching path array.
	configMap map[string]interface{} // The pared JSON objects for configuration files.
	available bool
}

var (
	supportedFileTypes = []string{"yaml"}
	SearchPath         = []string{
		file.GetAppDir(),
		file.GetWorkDir(),
		file.GetHomeDir(),
		os.TempDir(),
	}
	DefaultConfigFileName = "config"
)

// 读取新的配置文件
func NewConfigFile(file ...string) (*ConfigFile, error) {
	var (
		name = DefaultConfigFileName
	)

	if len(file) > 0 {
		name = file[0]
	}

	configPath, configType, ok := configFilePath(name)

	if !ok {
		golog.Errorf("没有找到配置文件 %s", name)
		return nil, errors.New("没有找到配置文件")
	}

	var cfgData map[string]interface{}

	// 对 yaml 处理
	golog.Debugf("加载配置文件 %s", configPath)

	if configType == "yaml" {
		if c, err := ioutil.ReadFile(configPath); err != nil {
			golog.Errorf("读取配置文件错误 %s ", err.Error())
			return nil, err
		} else {
			if err = yaml.Unmarshal(c, &cfgData); err != nil {
				golog.Errorf("读取配置文件错误 %s , %s", configPath, err.Error())
				return nil, err
			}
		}

	}

	cfg := &ConfigFile{
		fileName:  configPath,
		fileType:  configType,
		configMap: cfgData,
		available: true,
	}

	return cfg, nil
}

// 搜索配置文件路径
func configFilePath(filename string) (configPath, configType string, exits bool) {
	for _, dir := range SearchPath {
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
func (c *ConfigFile) Get(pattern string) (value interface{}, success bool) {
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
	return v, ok
}

// 扫描字符串值
func (c *ConfigFile) GetString(pattern string) (value string, ok bool) {
	v, ok := c.Get(pattern)
	if ok {
		value, ok = v.(string)
		return
	}
	return
}

// 扫描int64值
func (c *ConfigFile) GetInt64(pattern string) (value int64, ok bool) {
	v, ok := c.Get(pattern)
	if ok {
		value, ok = v.(int64)
		return
	}
	return
}

// 扫描int值
func (c *ConfigFile) GetInt(pattern string) (value int, ok bool) {
	v, ok := c.Get(pattern)
	if ok {
		value, ok = v.(int)
		return
	}
	return
}

// 扫描bool值
func (c *ConfigFile) GetBool(pattern string) (value bool, ok bool) {
	v, ok := c.Get(pattern)
	if ok {
		value, ok = v.(bool)
		return
	}
	return
}

// 重载文件
func (c *ConfigFile) Reload() (err error) {
	c.available = false
	var cfgData map[string]interface{}
	golog.Debugf("加载配置文件 %s", c.fileName)

	// 对 yaml 处理
	if c.fileType == "yaml" {
		if conf, err := ioutil.ReadFile(c.fileName); err != nil {
			golog.Errorf("读取配置文件错误 %s ", err.Error())
			return err
		} else {
			if err = yaml.Unmarshal(conf, &cfgData); err != nil {
				golog.Errorf("读取配置文件错误 %s , %s", c.fileName, err.Error())
				return err
			}
		}

		c.configMap = cfgData
	}
	c.available = true

	return
}

// 配置是否可用
func (c *ConfigFile) Available() (ok bool) {
	return c.available
}
