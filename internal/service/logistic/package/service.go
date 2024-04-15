package logisticPackage

import (
	"github.com/ozonmp/omp-bot/internal/model/logistic"
)


type PackageService interface {
  Describe(packageID uint64) (*logistic.Package, error)
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
  logistic.AllEntities = append(logistic.AllEntities, &pckg)

  return uint64(len(logistic.AllEntities)), nil
}