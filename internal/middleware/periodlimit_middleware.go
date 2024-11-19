package middleware

import (
	"net/http"

	"github.com/wlevene/loginservice/internal/config"
	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

const (
	seconds = 1
	quota   = 10

	key = "insou_periodlimit"
)

type PeriodLimitMiddleware struct {
	limit_period *limit.PeriodLimit //令牌桶
	c            config.Config
}

func NewPeriodLimitMiddleware(c config.Config) *PeriodLimitMiddleware {
	obj := &PeriodLimitMiddleware{
		c: c,
		limit_period: limit.NewPeriodLimit(
			seconds,
			quota,
			redis.New(c.Redis.Addr),
			key),
	}

	return obj
}

func (m *PeriodLimitMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// pass, err := m.Check("id")
		// if err != nil || !pass {
		// 	res, _ := json.Marshal(map[string]interface{}{
		// 		"code": http.StatusTooManyRequests,
		// 		"msg":  "TooManyRequests OverQuota",
		// 	})
		// 	w.WriteHeader(http.StatusTooManyRequests)
		// 	_, _ = w.Write(res)
		// 	return
		// }

		next(w, r)
	}
}

func (m *PeriodLimitMiddleware) Check(key string) (bool, error) {

	code, err := m.limit_period.Take(key)
	if err != nil {
		logx.Errorf("take out key failure[key:%s,err:%s]", key, err.Error())
		return false, err
	}

	logx.Info("limit:", code)

	// switch val =&gt; process request
	switch code {
	case limit.OverQuota:
		logx.Errorf("OverQuota key: %v", key)
		return false, err
	case limit.Allowed:
		logx.Infof("AllowedQuota key: %v", key)
		return true, nil
	case limit.HitQuota: // hit
		logx.Errorf("HitQuota key: %v", key)
		return true, err
	default:
		logx.Errorf("DefaultQuota key: %v", key)
		// unknown response, we just let the sms go
		return true, nil
	}

}
