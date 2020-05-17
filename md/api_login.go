package md

import (
	"errors"

	"github.com/qin-nz/goctp/signal"
	"github.com/qin-nz/libctp"
	"github.com/sirupsen/logrus"
)

// 行情用户登录
func (p *apispi) ReqUserLogin(reqf libctp.CThostFtdcReqUserLoginField) error {
	reqid := p.newRequestId()

	ret := p.api.ReqUserLogin(reqf, reqid)

	if ret != 0 {
		// TODO: 获取具体的 return msg
		return errors.New("发送用户登录请求失败！")
	}

	logrus.WithField("RequestID", reqid).Println("行情系统账号登陆中...")
	err := p.s.Wait(signal.Login)
	return err
}

// 登录请求响应
func (p *apispi) OnRspUserLogin(pRspUserLogin libctp.CThostFtdcRspUserLoginField, pRspInfo libctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	p.s.Trigger(signal.Login, nil) // nRequestID 返回恒为0，不能用

	if pRspInfo.GetErrorID() != 0 {
		logrus.WithFields(logrus.Fields{
			"errmsg": pRspInfo.GetErrorMsg(),
			"errid":  pRspInfo.GetErrorID(),
		}).Warn("行情系统登录异常")
	}

	if bIsLast {
		logrus.Printf("行情系统登陆成功，当前交易日： %v %v %v %d\n", p.api.GetTradingDay(), bIsLast, pRspInfo, nRequestID)
	} else {
		logrus.Debugf("行情系统登录异常")

	}
}
