package main

// 実装の詳細
type ServiceImpl struct{}

func (s *ServiceImpl) Apply(id int) error { return nil }

// 上位階層が定義する抽象
type OrderService interface {
	Apply(int) error
}

// 上位階層の利用者側の型
type Application struct {
	os OrderService
}

// 多言語のコンストラクタインジェクションに相当する実装
func NewApplication(os OrderService) *Application {
	return &Application{os: os}
}

func (app *Application) Apply(id int) error {
	return app.os.Apply(id)
}

func main() {
	app := NewApplication(&ServiceImpl{})
	app.Apply(19)
}
