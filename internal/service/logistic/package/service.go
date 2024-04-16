package logisticPackage

import (
	"github.com/ozonmp/omp-bot/internal/model/logistic"
)


type PackageService interface {
  Get(packageID uint64) (*logistic.Package, error)
  List(cursor uint64, limit uint64) ([]logistic.Package, error)
  Create(logistic.Package) (uint64, error)
  Update(packageID uint64, packageItem logistic.Package) error
  Remove(packageID uint64) (bool, error)
}

type DummyPackageService struct {}

func NewDummyPackageService() *DummyPackageService {
  return &DummyPackageService{}
}

func (ps *DummyPackageService) Create(pckg logistic.Package) (uint64, error) {
  logistic.AllEntities = append(logistic.AllEntities, pckg)

  return uint64(len(logistic.AllEntities)), nil
}

func (ps *DummyPackageService) Get(id uint64) (logistic.Package, error) {
  item := logistic.AllEntities[id]

  return item, nil
}

func (ps *DummyPackageService) List(cursor uint64, limit uint64) ([]logistic.Package, error) {
  return logistic.AllEntities[cursor:cursor + limit], nil
}

func (ps *DummyPackageService) Count() (int, error) {
  return len(logistic.AllEntities), nil
}


func (ps *DummyPackageService) Remove(packageID uint64) (bool, error) {
  lastElemIndex := len(logistic.AllEntities) - 1

  logistic.AllEntities[packageID] = logistic.AllEntities[lastElemIndex]
  logistic.AllEntities = logistic.AllEntities[:lastElemIndex]

  return true, nil
}
