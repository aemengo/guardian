package bundlerules

import (
	spec "code.cloudfoundry.org/guardian/gardener/container-spec"
	"code.cloudfoundry.org/guardian/rundmc/goci"
	"github.com/opencontainers/runtime-spec/specs-go"
	"math"
)

type Windows struct{}

func (w Windows) Apply(bndl goci.Bndl, spec spec.DesiredContainerSpec) (goci.Bndl, error) {
	if spec.BaseConfig.Windows == nil {
		return bndl, nil
	}

	var (
		limit  = spec.Limits.Memory.LimitInBytes
		shares = shares(limit)
	)

	return bndl.WithWindows(*spec.BaseConfig.Windows).
		WithWindowsMemoryLimit(specs.WindowsMemoryResources{Limit: &limit}).
		WithWindowsCPUShares(specs.WindowsCPUResources{Shares: &shares}), nil
}
// min(10000*(application_memory / 8 GB), 10000)
func shares(memInBytes uint64) uint16 {
	var (
		eightGBInBytes float64 = 8589934592
		memPercent             = float64(memInBytes) / eightGBInBytes
	)

	memPercent = memPercent * 10000
	return uint16(math.Min(10000, memPercent))
}
