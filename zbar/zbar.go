package zbar

import (
	"encoding/xml"
	"errors"
	"os/exec"
	"strings"
)

var execPath string

type Zbar struct {
	Source []string
}

func init() {
	cmd := exec.Command("which", "zbarimg")
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	execPath = strings.TrimSuffix(string(out), "\n")
}

func (z *Zbar) Decode() (*Barcodes, error) {
	if len(z.Source) == 0 {
		return nil, errors.New("未设定文件路径")
	}

	args := make([]string, 0)
	args = append(args, "-q")
	args = append(args, "--xml")
	args = append(args, z.Source...)
	cmd := exec.Command(execPath, args...)

	out, err := cmd.CombinedOutput()
	if err != nil && cmd.ProcessState.ExitCode() != 4 {
		return nil, errors.New(string(out))
	}

	return z.toEntity(out)
}

func (z *Zbar) SetPath(zbarPath string) {
	execPath = zbarPath
}

func (z *Zbar) toEntity(data []byte) (*Barcodes, error) {
	res := &Barcodes{}
	err := xml.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
