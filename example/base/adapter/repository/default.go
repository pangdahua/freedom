package repository

import (
	"github.com/8treenet/freedom"
)

func init() {
	freedom.Prepare(func(initiator freedom.Initiator) {
		initiator.BindRepository(func() *Default {
			return &Default{}
		})
	})
}

// Default .
type Default struct {
	freedom.Repository
}

// GetIP .
func (repo *Default) GetIP() string {
	//repo.DB().Find()
	repo.Worker.Logger().Infof("我是Repository GetIP")
	return repo.Worker.IrisContext().RemoteAddr()
}

// GetUA - implment DefaultRepoInterface interface
func (repo *Default) GetUA() string {
	repo.Worker.Logger().Infof("我是Repository GetUA")
	return repo.Worker.IrisContext().Request().UserAgent()
}