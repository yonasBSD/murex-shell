package textmanip

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types/define"
)

func init() {
	proc.GoFunctions["left"] = cmdLeft
	proc.GoFunctions["right"] = cmdRight
	proc.GoFunctions["append"] = cmdAppend
	proc.GoFunctions["prepend"] = cmdPrepend
}

func cmdLeft(p *proc.Process) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	left, err := p.Parameters.Int(0)
	if err != nil {
		return err
	}

	var output []string
	p.Stdin.ReadArray(func(b []byte) {
		if len(b) < left {
			output = append(output, string(b))
		} else {
			output = append(output, string(b[:left]))
		}
	})

	b, err := define.MarshalData(p, dt, output)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}

func cmdRight(p *proc.Process) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	right, err := p.Parameters.Int(0)
	if err != nil {
		return err
	}

	var output []string
	p.Stdin.ReadArray(func(b []byte) {
		if len(b) < right {
			output = append(output, string(b))
		} else {
			output = append(output, string(b[len(b)-right:]))
		}
	})

	b, err := define.MarshalData(p, dt, output)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}

func cmdPrepend(p *proc.Process) (err error) {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	prepend := p.Parameters.StringAll()

	var output []string
	p.Stdin.ReadArray(func(b []byte) {
		output = append(output, prepend+string(b))
	})

	b, err := define.MarshalData(p, dt, output)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}

func cmdAppend(p *proc.Process) (err error) {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	s := p.Parameters.StringAll()

	var output []string
	p.Stdin.ReadArray(func(b []byte) {
		output = append(output, string(b)+s)
	})

	b, err := define.MarshalData(p, dt, output)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}
