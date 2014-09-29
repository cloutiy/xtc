package asm

import (
  "bitbucket.org/yyuu/bs/core"
)

type SuffixedSymbol struct {
  ClassName string
  Base core.ISymbol
  Suffix string
}

func NewSuffixedSymbol(base core.ISymbol, suffix string) *SuffixedSymbol {
  return &SuffixedSymbol { "asm.SuffixedSymbol", base, suffix }
}

func (self SuffixedSymbol) IsZero() bool {
  return false
}

func (self SuffixedSymbol) GetName() string {
  return self.Base.GetName()
}

func (self SuffixedSymbol) String() string {
  return self.Base.String() + self.Suffix
}
