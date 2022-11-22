package gcnf

import "github.com/layasugar/laya/core/constants"

// AppName 返回当前app名称
func AppName() string {
	if IsSet(constants.KEY_APPNAME) {
		return GetString(constants.KEY_APPNAME)
	}
	return constants.DEFAULT_APPNAME
}

// AppMode 返回当前的环境
func AppMode() string {
	if IsSet(constants.KEY_APPMODE) {
		return GetString(constants.KEY_APPMODE)
	}
	return constants.DEFAULT_APPMODE
}

// AppVersion 返回app的版本号
func AppVersion() string {
	if IsSet(constants.KEY_APPVERSION) {
		return GetString(constants.KEY_APPVERSION)
	}
	return constants.DEFAULT_APPVERSION
}

// Listen 监听地址
func Listen() string {
	if IsSet(constants.KEY_APPLISTEN) {
		return GetString(constants.KEY_APPLISTEN)
	}
	return constants.DEFAULT_LISTEN
}
