package zupgrd

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	"gitee.com/sienectagv/gozk/zarchive"
	"gitee.com/sienectagv/gozk/zlogger"
	"gitee.com/sienectagv/gozk/znet"
	"gitee.com/sienectagv/gozk/zutils"
	"gopkg.in/ini.v1"
)

type IniUpgrade struct {
	Url            string `ini:"url"`
	InstallPath    string `ini:"install_path"`
	cachePath      string
	remoteVerFile  string
	currentVerFile string
}

type IniVersion struct {
	Version   string `ini:"version"`
	Package   string `ini:"package"`
	Installed string `ini:"installed"` //installed version
}

func (cfg *IniUpgrade) remoteName() string {
	return path.Join(cfg.cachePath, cfg.remoteVerFile)
}

func (cfg *IniUpgrade) currentName() string {
	return path.Join(cfg.cachePath, cfg.currentVerFile)
}

//fileName don't include ".ini"
func LoadCfgUpgrade() *IniUpgrade {
	rootPath := zutils.RootPath()
	cfg := &IniUpgrade{
		cachePath:      path.Join(rootPath, "cache"),
		remoteVerFile:  "remote.ini",
		currentVerFile: "current.ini",
	}
	fileName := path.Join(rootPath, "ini/root.ini")
	sections, err := ini.Load(fileName)
	if nil == err {
		sections.MapTo(cfg)
	} else {
		zlogger.Error("failed to load .ini", fileName, err)
	}
	os.MkdirAll(cfg.cachePath, os.ModeType)
	if !path.IsAbs(cfg.InstallPath) {
		// cfg.InstallPath = path.Join(rootPath, cfg.InstallPath)
		cfg.InstallPath, _ = filepath.Abs(cfg.InstallPath)
		zlogger.Info(cfg.InstallPath)
	}
	return cfg
}

type UpgradeApp struct {
	cfg *IniUpgrade
}

func NewUpgradeApp() *UpgradeApp {
	ua := &UpgradeApp{}
	ua.cfg = LoadCfgUpgrade()
	return ua
}

// func (ua *UpgradeApp) AddToLoopGroup(loopGroup *zutils.LoopGroup) {
// 	loopGroup.GoLoop("_upgrade_",
// 		func() int {
// 			ua.Run()
// 			return 1
// 		},
// 		time.Minute,
// 		nil)
// }

func (ua *UpgradeApp) UpdateRemoteVersion(fileName string) error {
	if nil == ua.cfg {
		return zutils.ErrNullParam
	}
	err := znet.HttpGetFile(ua.cfg.Url+fileName,
		func(body io.Reader) {
			file, err := os.Create(ua.cfg.remoteName())
			if err != nil {
				return
			}
			defer file.Close()
			io.Copy(file, body)
		})
	if nil != err {
		zlogger.Error(err)
	}
	return err
}

func (ua *UpgradeApp) TryDownPackage() error {
	remote := &IniVersion{}
	zlogger.Info(ua.cfg.remoteName())
	file, err := ini.Load(ua.cfg.remoteName())
	if err != nil {
		zlogger.Error(err)
		return err
	}
	err = file.MapTo(remote)
	// err := ini.StrictMapTo(remote, ua.cfg.remoteName())
	zlogger.Info(*remote)
	if nil != err {
		zlogger.Error(err)
		return err
	}
	current := &IniVersion{}
	ini.StrictMapTo(current, ua.cfg.currentName())
	if remote.Version <= current.Version || len(remote.Package) == 0 {
		zlogger.Info("No latest version!")
		return zutils.ErrNotFound
	}
	//download package
	var errcopy error = zutils.ErrCmdUnexist
	err = znet.HttpGetFile(ua.cfg.Url+remote.Package,
		func(body io.Reader) {
			file, err := os.Create(filepath.Join(ua.cfg.cachePath, remote.Package))
			if err != nil {
				return
			}
			defer file.Close()
			_, errcopy = io.Copy(file, body)
		})
	if nil != err {
		zlogger.Error(err)
		return err
	}
	if nil != errcopy {
		zlogger.Error(errcopy)
		return errcopy
	}
	//write current
	file = ini.Empty()
	file.ReflectFrom(remote)
	file.SaveTo(ua.cfg.currentName())
	return nil
}

func (ua *UpgradeApp) Install() error {
	current := &IniVersion{}
	file, _ := ini.Load(ua.cfg.currentName())
	file.MapTo(current)
	// ini.StrictMapTo(current, ua.cfg.currentName())
	if current.Installed >= current.Version {
		zlogger.Info("It's latest version, don't need to install!")
		return nil
	}
	err := zarchive.UnzipToFolder(filepath.Join(ua.cfg.cachePath, current.Package), ua.cfg.InstallPath)
	if err != nil {
		zlogger.Info("Failed to install the latest version package")
		return err
	}
	current.Installed = current.Version
	file.ReflectFrom(current)
	file.SaveTo(ua.cfg.currentName())
	//write current
	fmt.Printf("Upgrade to version:%s\n", current.Version)
	return nil
}

func (ua *UpgradeApp) Run() {
	if nil == ua.cfg {
		ua.cfg = LoadCfgUpgrade()
		zlogger.Info(*ua.cfg)
	}
}
