package channel_by_channel

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isChannelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestChannel(t *testing.T) {
	//获取当前节点
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isChannelled(ctx) {
					break
				}
			}
			time.Sleep(time.Millisecond * 5)
			fmt.Println(i, "Done")
		}(i, ctx)

	}
	//执行次方法会让所有的cancel都取消
	cancel()
	time.Sleep(time.Second * 1)
}
