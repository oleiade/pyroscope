// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64

package python

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type PerfPyEvent struct {
	StackStatus uint8
	Err         uint8
	Reserved2   uint8
	Reserved3   uint8
	Pid         uint32
	KernStack   int64
	StackLen    uint32
	Stack       [75]uint32
}

type PerfPyOffsetConfig struct {
	PyThreadStateFrame            int16
	PyThreadStateCframe           int16
	PyCFrameCurrentFrame          int16
	PyCodeObjectCoFilename        int16
	PyCodeObjectCoName            int16
	PyCodeObjectCoVarnames        int16
	PyCodeObjectCoLocalsplusnames int16
	PyTupleObjectObItem           int16
	VFrameCode                    int16
	VFramePrevious                int16
	VFrameLocalsplus              int16
}

type PerfPyPidData struct {
	Offsets PerfPyOffsetConfig
	_       [2]byte
	Version struct {
		Major uint32
		Minor uint32
		Patch uint32
	}
	Musl   uint8
	_      [3]byte
	TssKey int32
}

type PerfPySampleStateT struct {
	SymbolCounter          int64
	Offsets                PerfPyOffsetConfig
	_                      [2]byte
	CurCpu                 uint32
	_                      [4]byte
	FramePtr               uint64
	PythonStackProgCallCnt int64
	Event                  PerfPyEvent
}

type PerfPySymbol struct {
	Classname [32]int8
	Name      [64]int8
	File      [128]int8
}

// LoadPerf returns the embedded CollectionSpec for Perf.
func LoadPerf() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_PerfBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load Perf: %w", err)
	}

	return spec, err
}

// LoadPerfObjects loads Perf and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*PerfObjects
//	*PerfPrograms
//	*PerfMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadPerfObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadPerf()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// PerfSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type PerfSpecs struct {
	PerfProgramSpecs
	PerfMapSpecs
}

// PerfSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type PerfProgramSpecs struct {
	PyperfCollect   *ebpf.ProgramSpec `ebpf:"pyperf_collect"`
	ReadPythonStack *ebpf.ProgramSpec `ebpf:"read_python_stack"`
}

// PerfMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type PerfMapSpecs struct {
	PyEvents    *ebpf.MapSpec `ebpf:"py_events"`
	PyPidConfig *ebpf.MapSpec `ebpf:"py_pid_config"`
	PyProgs     *ebpf.MapSpec `ebpf:"py_progs"`
	PyStateHeap *ebpf.MapSpec `ebpf:"py_state_heap"`
	PySymbols   *ebpf.MapSpec `ebpf:"py_symbols"`
	Stacks      *ebpf.MapSpec `ebpf:"stacks"`
}

// PerfObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadPerfObjects or ebpf.CollectionSpec.LoadAndAssign.
type PerfObjects struct {
	PerfPrograms
	PerfMaps
}

func (o *PerfObjects) Close() error {
	return _PerfClose(
		&o.PerfPrograms,
		&o.PerfMaps,
	)
}

// PerfMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadPerfObjects or ebpf.CollectionSpec.LoadAndAssign.
type PerfMaps struct {
	PyEvents    *ebpf.Map `ebpf:"py_events"`
	PyPidConfig *ebpf.Map `ebpf:"py_pid_config"`
	PyProgs     *ebpf.Map `ebpf:"py_progs"`
	PyStateHeap *ebpf.Map `ebpf:"py_state_heap"`
	PySymbols   *ebpf.Map `ebpf:"py_symbols"`
	Stacks      *ebpf.Map `ebpf:"stacks"`
}

func (m *PerfMaps) Close() error {
	return _PerfClose(
		m.PyEvents,
		m.PyPidConfig,
		m.PyProgs,
		m.PyStateHeap,
		m.PySymbols,
		m.Stacks,
	)
}

// PerfPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadPerfObjects or ebpf.CollectionSpec.LoadAndAssign.
type PerfPrograms struct {
	PyperfCollect   *ebpf.Program `ebpf:"pyperf_collect"`
	ReadPythonStack *ebpf.Program `ebpf:"read_python_stack"`
}

func (p *PerfPrograms) Close() error {
	return _PerfClose(
		p.PyperfCollect,
		p.ReadPythonStack,
	)
}

func _PerfClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed perf_bpfel_arm64.o
var _PerfBytes []byte