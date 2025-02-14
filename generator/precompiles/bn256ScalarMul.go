// Copyright 2020 Marius van der Wijden
// This file is part of the fuzzy-vm library.
//
// The fuzzy-vm library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The fuzzy-vm library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the fuzzy-vm library. If not, see <http://www.gnu.org/licenses/>.

package precompiles

import (
	"github.com/rgeraldes24/goevmlab/program"
	"github.com/theQRL/FuzzyVM/filler"
	"github.com/theQRL/go-zond/common"
	"github.com/theQRL/go-zond/crypto/bn256"
)

var bn256mulAddr, _ = common.NewAddressFromString("Z0000000000000000000000000000000000000007")

type bn256MulCaller struct{}

func (*bn256MulCaller) call(p *program.Program, f *filler.Filler) error {
	k := f.BigInt32()
	point := new(bn256.G1).ScalarBaseMult(k)
	scalar := f.BigInt32()
	c := CallObj{
		Gas:       f.GasInt(),
		Address:   bn256mulAddr,
		InOffset:  0,
		InSize:    96,
		OutOffset: 0,
		OutSize:   64,
		Value:     f.BigInt32(),
	}
	// 64 bytes curve point
	p.Mstore(point.Marshal(), 0)
	// 32 bytes scalar
	bytes := make([]byte, 32)
	copy(bytes, scalar.Bytes())
	p.Mstore(bytes[:], 64)
	CallRandomizer(p, f, c)
	return nil
}
