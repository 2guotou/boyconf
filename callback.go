package boyconf

import (
	"fmt"
	"time"

	"github.com/go-fsnotify/fsnotify"
)

//发起监控
func (b *Boy) watch() (err error) {
	var watcher *fsnotify.Watcher
	if watcher, err = fsnotify.NewWatcher(); err == nil {
		go b.watchAndReload(watcher)
		err = watcher.Add(b.File)
	}
	return
}

//监控变更并重载, 以及调用 Trigger
func (b *Boy) watchAndReload(w *fsnotify.Watcher) {
	defer w.Close()
	for {
		select {
		case event := <-w.Events:
			op := event.Op.String()
			fmt.Println("Boy: Nsnotify Config File Event: " + op)
			if op == "WRITE" {
				time.Sleep(5 * time.Second) //休眠5秒钟
				//释放掉多余的触发事件
				for l := len(w.Events); l > 0; l-- {
					<-w.Events
				}
				//再次加载配置, 然后完成回调操作
				if er := b.load(); er == nil {
					b.trigger()
				} else {
					fmt.Println("Boy: Nsnotify Load Config Error: " + er.Error())
				}
			}
		case err := <-w.Errors:
			fmt.Println("Boy: Nsnotify Exception: " + err.Error())
		}
	}
}

//重载触发
func (b *Boy) trigger() {
	for _, f := range b.reloadTrigger {
		f()
	}
}
