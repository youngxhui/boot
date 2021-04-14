package srv

import (
	"boot/protos"
	"context"
	"google.golang.org/protobuf/types/known/anypb"

	"github.com/youngxhui/power/log"
)

type UserServer struct {
}

func (u UserServer) Login(ctx context.Context, info *protos.UserInfo) (*protos.Result, error) {
	if info.Username == "123" && info.Password == "1234" {
		return &protos.Result{
			Code:    0,
			Message: "登录成功",
			Data:    nil,
		}, nil
	}
	//client := http.Client{}
	//url := "http://localhost:8501/v1/models/fib:predict"

	//resp, err := client.Post(url, "application/json", strings.NewReader("{\"instances\": [[[2.0],[3.0],[5.0]]]}"))
	//if err != nil {
	//	log.Error(err.Error())
	//}
	//body, err := ioutil.ReadAll(resp.Body)
	//log.Info(body)

	//b := string(body)
	//log.Debug(b)
	token := protos.Success{
		Token: "Java",
	}
	any, _ := anypb.New(&token)
	log.Info("any:", any)
	newToken := protos.Success{}

	log.Info("new Token", newToken.Token)
	return &protos.Result{Code: -1,
		Message: "账号或密码错误",
		Data:    any}, nil
}
