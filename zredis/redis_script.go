package zredis

import "github.com/gomodule/redigo/redis"

const (
	luaResetSet = `redis.call('DEL', KEYS[1])
		return redis.call('SADD', KEYS[1], unpack(ARGV))
		`
)

func (p *Pool) ResetSet(key interface{}, values ...interface{}) (int, error) {
	script := redis.NewScript(1, luaResetSet)
	conn := p.Get()
	defer conn.Close()
	para := make([]interface{}, 0, 1+len(values))
	para = append(para, key)
	for _, v := range values {
		para = append(para, v)
	}
	ret, err := redis.Int(script.Do(conn, para...))
	if nil != err {
		return 0, err
	}
	return ret, nil
}
