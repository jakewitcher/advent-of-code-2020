package program

type Mask struct {
	Bits []Bit
}

type Bit struct {
	Value int
	Shift int
}

type Decoder interface {
	UpdateMask(mask *Mask)
	Decode(memory Memory, mw *MemWrite)
}

type VersionOneDecoder struct {
	*Mask
}

func (v *VersionOneDecoder) UpdateMask(mask *Mask) {
	v.Mask = mask
}

func (v *VersionOneDecoder) Decode(memory Memory, mw *MemWrite) {
	memory[mw.Index] = v.Apply(mw.Value)
}

func (v *VersionOneDecoder) Apply(value int) int {
	for _, bit := range v.Bits {
		if bit.Value == 0 {
			value = value &^ (1 << bit.Shift)
		} else {
			value = value | (1 << bit.Shift)
		}
	}
	return value
}

type VersionTwoDecoder struct {
	*Mask
}

func (v *VersionTwoDecoder) UpdateMask(mask *Mask) {
	v.Mask = mask
}

func (v *VersionTwoDecoder) Decode(memory Memory, mw *MemWrite) {
	floating := make([]int, 36)
	addresses := make(map[int]interface{})
	address := mw.Index

	for _, bit := range v.Bits {
		floating[bit.Shift] = 1
		if bit.Value == 1 {
			address = address | (1 << bit.Shift)
		}
	}

	addresses[address] = struct{}{}
	for i, n := range floating {
		if n == 1 {
			continue
		}

		for address, _ := range addresses {
			a := address | (1 << i)
			b := address &^ (1 << i)
			addresses[a] = struct{}{}
			addresses[b] = struct{}{}
		}
	}

	for address, _ := range addresses {
		memory[address] = mw.Value
	}
}
