package db

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	// 在所有测试函数运行前执行
	// 暂时不执行 InitConnection，因为我们想在 TestInitConnection 中测试它
	code := m.Run()

	// 在所有测试函数运行后执行，断开连接
	if Client != nil {
		Client.Disconnect(context.TODO())
	}

	os.Exit(code)
}

func TestInitConnection(t *testing.T) {
	t.Parallel() // 允许并行执行测试

	// 环境变量：可以从环境变量中获取数据库 URI，更安全
	// 这里为了方便测试，我们直接使用常量
	// uri := os.Getenv("MONGO_URI")
	// if uri == "" {
	// 	t.Skip("MONGO_URI 环境变量未设置，跳过测试。")
	// }

	// 调用 InitConnection 函数
	InitConnection()

	// 检查 Client 是否已成功初始化
	if Client == nil {
		t.Fatalf("InitConnection 失败，Client 为 nil")
	}

	// 使用一个超时上下文来ping数据库，确保连接是活跃的
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := Client.Ping(ctx, nil)
	if err != nil {
		t.Fatalf("连接到 MongoDB 失败: %v", err)
	}

	fmt.Println("数据库连接测试成功!")
}

// 示例：测试 GetUsersByName 函数
func TestGetUsersByName(t *testing.T) {
	t.Skip("这是一个示例测试，请根据你的数据创建实际测试")

	// 确保数据库已连接
	if Client == nil {
		t.Fatal("数据库未连接")
	}

	// 假设数据库中存在 '张明' 这个用户
	users, err := GetUsersByName("张明")
	if err != nil {
		t.Fatalf("查询用户失败: %v", err)
	}

	if len(users) == 0 {
		t.Errorf("查询结果为空，预期至少有一个用户")
	}

	// 你可以进一步检查返回的用户数据是否符合预期
	for _, user := range users {
		t.Logf("找到用户: %s", user.DisplayName)
	}
}
