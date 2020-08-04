package processes

import (
	"os"

	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/guardian/rundmc/goci"
	"code.cloudfoundry.org/guardian/rundmc/users"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

//go:generate counterfeiter . EnvDeterminer
type EnvDeterminer interface {
	EnvFor(bndl goci.Bndl, spec garden.ProcessSpec, containerUID int) []string
}

type EnvFunc func(bndl goci.Bndl, spec garden.ProcessSpec, containerUID int) []string

func (fn EnvFunc) EnvFor(bndl goci.Bndl, spec garden.ProcessSpec, containerUID int) []string {
	return fn(bndl, spec, containerUID)
}

type ProcBuilder struct {
	envDeterminer  EnvDeterminer
	nonRootMaxCaps []string
}

func NewBuilder(envDeterminer EnvDeterminer, nonRootMaxCaps []string) *ProcBuilder {
	return &ProcBuilder{
		envDeterminer:  envDeterminer,
		nonRootMaxCaps: nonRootMaxCaps,
	}
}

func (p *ProcBuilder) BuildProcess(bndl goci.Bndl, spec garden.ProcessSpec, user *users.ExecUser) *specs.Process {
	additionalGIDs := []uint32{}
	if os.Geteuid() == 0 {
		additionalGIDs = toUint32Slice(user.Sgids)
	}
	return &specs.Process{
		Args:        append([]string{spec.Path}, spec.Args...),
		ConsoleSize: console(spec),
		Env:         p.envDeterminer.EnvFor(bndl, spec, user.Uid),
		User: specs.User{
			UID:            uint32(user.Uid),
			GID:            uint32(user.Gid),
			AdditionalGids: additionalGIDs,
			Username:       spec.User,
		},
		Cwd:             spec.Dir,
		Capabilities:    p.capabilities(bndl, user.Gid),
		Rlimits:         toRlimits(spec.Limits),
		Terminal:        spec.TTY != nil,
		ApparmorProfile: bndl.Process().ApparmorProfile,
	}
}

func (p *ProcBuilder) capabilities(bndl goci.Bndl, containerUID int) *specs.LinuxCapabilities {
	capsToSet := bndl.Capabilities()
	if containerUID != 0 {
		capsToSet = intersect(capsToSet, p.nonRootMaxCaps)
	}

	// TODO centralize knowledge of garden -> runc capability schema translation
	if len(capsToSet) > 0 {
		return &specs.LinuxCapabilities{
			Bounding:    capsToSet,
			Inheritable: capsToSet,
			Permitted:   capsToSet,
		}
	}

	return nil
}

func console(spec garden.ProcessSpec) *specs.Box {
	consoleBox := &specs.Box{
		Width:  80,
		Height: 24,
	}
	if spec.TTY != nil && spec.TTY.WindowSize != nil {
		consoleBox.Width = uint(spec.TTY.WindowSize.Columns)
		consoleBox.Height = uint(spec.TTY.WindowSize.Rows)
	}
	return consoleBox
}

func intersect(l1 []string, l2 []string) (result []string) {
	for _, a := range l1 {
		for _, b := range l2 {
			if a == b {
				result = append(result, a)
			}
		}
	}

	return result
}

func toUint32Slice(slice []int) []uint32 {
	result := []uint32{}
	for _, i := range slice {
		result = append(result, uint32(i))
	}
	return result
}
