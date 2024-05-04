package bootstrap

import "gorm.io/gorm"

type Application struct {
	Env *Env
	Mysql *gorm.DB
}

func App() Application {
	app := &Application{}

	app.Env = NewEnv()
	app.Mysql = NewMysqlDatabase(app.Env)

	return *app
}

func (app *Application) CloseDBConnection() {
	// app.Mysql.
}