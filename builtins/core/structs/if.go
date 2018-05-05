package structs

import (
	"errors"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	proc.GoFunctions["if"] = cmdIf
	proc.GoFunctions["!if"] = cmdIf
}

const (
	fIf = iota
	fThen
	fElse
	fDone
)

func cmdIf(p *proc.Process) error {
	p.Stdout.SetDataType(types.Generic)

	if p.Parameters.Len() == 0 {
		return errors.New("No arguments made. `if` requires parameters.")
	}

	var (
		blocks [3][]rune
		flag   int
	)

	if p.IsMethod {
		// We derive the conditional state from stdin
		flag++
	}

	for i := 0; i < p.Parameters.Len(); i++ {
		switch {
		case i == 0 && !p.IsMethod:
			r, err := p.Parameters.Block(0)
			if err != nil {
				return err
			}

			blocks[0] = r
			flag++

		default:
			if flag == fDone {
				return errors.New("Parameters past end of `then` block.")
			}

			s, err := p.Parameters.String(i)
			if err != nil {
				return err
			}

			matched, err := setFlag(&s, &flag)
			if err != nil {
				return err
			}

			if matched == true {
				continue
			}

			r, err := p.Parameters.Block(i)
			if err != nil {
				return err
			}

			blocks[flag] = r
			flag++
		}
	}

	//debug.Log("if  :", string(blocks[fIf]))
	//debug.Log("then:", string(blocks[fThen]))
	//debug.Log("else:", string(blocks[fElse]))

	var conditional bool

	if len(blocks[fIf]) > 0 {
		// --- IF --- (function)
		stdout := streams.NewStdin()
		stderr := new(streams.Null)
		i, err := lang.RunBlockExistingConfigSpace(blocks[fIf], nil, stdout, stderr, p)
		if err != nil {
			return err
		}

		b, err := stdout.ReadAll()
		if err != nil {
			return err
		}
		conditional = types.IsTrue(b, i)

	} else {
		// --- IF --- (method)
		b, err := p.Stdin.ReadAll()
		if err != nil {
			return err
		}
		conditional = types.IsTrue(b, p.Previous.ExitNum)
	}

	if (conditional && !p.IsNot) || (!conditional && p.IsNot) {
		// --- THEN ---
		if len(blocks[fThen]) > 0 {
			_, err := lang.RunBlockExistingConfigSpace(blocks[fThen], nil, p.Stdout, p.Stderr, p)
			if err != nil {
				return err
			}
		}

	} else {
		// --- ELSE ---
		if len(blocks[fElse]) > 0 {
			_, err := lang.RunBlockExistingConfigSpace(blocks[fElse], nil, p.Stdout, p.Stderr, p)
			if err != nil {
				return err
			}
		} else {
			p.ExitNum = 1
		}
	}

	return nil
}

func setFlag(s *string, flag *int) (bool, error) {
	switch *s {
	case "then":
		if *flag > fThen {
			return false, errors.New("`then` appears too late in parameters.")
		}
		*flag = fThen
		return true, nil

	case "else":
		if *flag > fElse {
			return false, errors.New("`else` appears too late in parameters.")
		}
		*flag = fElse
		return true, nil

	default:
		return false, nil

	}
}