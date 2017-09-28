package boyconf

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

//Boy is a instance of Config File
type Boy struct {
	File          string      //配置文件
	Config        interface{} //由使用程序制定一个指针
	Env           []string    //环境变量, 据此按次序加载配置, 最后会使用 env="default" 的配置做补充
	AutoReload    bool        //自动重载配置
	reloadTrigger []func()    //添加一组回调函数
}

// Init ...
// 拉取配置期间发生任何错误都会 return 一个 Error, 调用层请做好应对策略.
func Init(b *Boy, fs ...func()) (err error) {
	if b.File == "" {
		return errors.New("Config File Not Specified")
	}
	if b.Env == nil {
		b.Env = []string{}
	}
	b.reloadTrigger = fs
	if err = b.load(); err != nil {
		return err
	}
	if b.AutoReload {
		err = b.watch()
	}
	return
}

// map key 必须是小写, 所以需要一次 loop 将 key 转换为 lower
// 配置的覆盖逻辑为:
// key > value  覆盖
// key > []     覆盖
// key > struct merge & 覆盖
// key > map    merge & 覆盖
func (b *Boy) load() (err error) {
	content, err := ioutil.ReadFile(b.File)
	if err != nil {
		return errors.New("Boy: Config File Read Error: " + err.Error())
	}
	rawCnfs := map[string]json.RawMessage{}
	err = json.Unmarshal(content, &rawCnfs)
	if err != nil {
		return errors.New("Boy: Config First Parse Error: " + err.Error())
	}
	config := b.Config
	for _, env := range b.Env {
		if raw, ok := rawCnfs[env]; ok {
			if err = json.Unmarshal(raw, &config); err != nil {
				return fmt.Errorf("Boy: <%s> config unmarshal error: %s", env, err.Error())
			}
		} else {
			return fmt.Errorf("Boy: <%s> config not found", env)
		}
	}
	b.Config = config
	return
}
