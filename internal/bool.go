package internal

import "github.com/open-compiler-project/vm/internal/reflect"

type Bool bool

func (*Bool) Type() reflect.Type { return reflect.Bool }
