package vyze

// IDs

type IDsIn chan<- ID
type IDsOut <-chan ID

func (s IDsOut) Unique() IDsOut {
	idsOut := make(chan ID)
	go func(idsOut chan ID) {
		idsEmitted := IDSet{}
		for id := range s {
			if !idsEmitted[id] {
				idsOut <- id
				idsEmitted[id] = true
			}
		}
		close(idsOut)
	}(idsOut)
	return idsOut
}

func (s IDsOut) Set() IDSet {
	if s == nil {
		return nil
	}
	vals := IDSet{}
	for val := range s {
		vals[val] = true
	}
	return vals
}

func (s IDsOut) List() IDList {
	if s == nil {
		return nil
	}
	vals := IDList{}
	for val := range s {
		vals = append(vals, val)
	}
	return vals
}

func (s IDsOut) Length() int64 {
	if s == nil {
		return 0
	}
	count := int64(0)
	for range s {
		count++
	}
	return count
}

func (s IDsOut) Pipe(s2 IDsIn) {
	go func() {
		for i := range s {
			s2 <- i
		}
		close(s2)
	}()
}

// ID Pairs

type IDPairsIn chan<- IDPair
type IDPairsOut <-chan IDPair

func (s IDPairsOut) Map() IDMap {
	if s == nil {
		return nil
	}
	vals := IDMap{}
	for val := range s {
		vals[val[0]] = val[1]
	}
	return vals
}

func (s IDPairsOut) List() []IDPair {
	idPairs := []IDPair{}
	for idp := range s {
		idPairs = append(idPairs, idp)
	}
	return idPairs
}

func (s IDPairsOut) Keys() IDsOut {
	if s == nil {
		return nil
	}
	vals := make(chan ID)
	go func() {
		for val := range s {
			vals <- val[0]
		}
		close(vals)
	}()
	return vals
}

func (s IDPairsOut) Values() IDsOut {
	if s == nil {
		return nil
	}
	vals := make(chan ID)
	go func() {
		for val := range s {
			vals <- val[1]
		}
		close(vals)
	}()
	return vals
}

func (s IDPairsOut) Length() int64 {
	if s == nil {
		return 0
	}
	count := int64(0)
	for range s {
		count++
	}
	return count
}

// Keyed IDPairs

type KeyedIDPairsIn chan<- KeyedIDPair
type KeyedIDPairsOut <-chan KeyedIDPair

func (s KeyedIDPairsOut) Pairs() IDPairsOut {
	if s == nil {
		return nil
	}
	vals := make(chan IDPair)
	go func() {
		for idp := range s {
			vals <- idp.Pair
		}
		close(vals)
	}()
	return vals
}

func (s KeyedIDPairsOut) List() []KeyedIDPair {
	idPairs := []KeyedIDPair{}
	for idp := range s {
		idPairs = append(idPairs, idp)
	}
	return idPairs
}

func (s KeyedIDPairsOut) PairMap() IDMap {
	if s == nil {
		return nil
	}
	vals := IDMap{}
	for val := range s {
		vals[val.Pair[0]] = val.Pair[1]
	}
	return vals
}

func (s KeyedIDPairsOut) PairList() []IDPair {
	idPairs := []IDPair{}
	for idp := range s {
		idPairs = append(idPairs, idp.Pair)
	}
	return idPairs
}

func (s KeyedIDPairsOut) PairValueSet() IDSet {
	if s == nil {
		return nil
	}
	vals := IDSet{}
	for val := range s {
		vals[val.Pair[1]] = true
	}
	return vals
}

func (s KeyedIDPairsOut) PairValueList() IDList {
	if s == nil {
		return nil
	}
	vals := IDList{}
	for val := range s {
		vals = append(vals, val.Pair[1])
	}
	return vals
}

func (s KeyedIDPairsOut) PairKeys() IDsOut {
	if s == nil {
		return nil
	}
	vals := make(chan ID)
	go func() {
		for val := range s {
			vals <- val.Pair[0]
		}
		close(vals)
	}()
	return vals
}

func (s KeyedIDPairsOut) PairValues() IDsOut {
	if s == nil {
		return nil
	}
	vals := make(chan ID)
	go func() {
		for val := range s {
			vals <- val.Pair[1]
		}
		close(vals)
	}()
	return vals
}

func (s KeyedIDPairsOut) Length() int64 {
	if s == nil {
		return 0
	}
	count := int64(0)
	for range s {
		count++
	}
	return count
}

// Binary

type BinaryOut <-chan Binary

func (s BinaryOut) List() []Binary {
	if s == nil {
		return nil
	}
	vals := []Binary{}
	for val := range s {
		vals = append(vals, val)
	}
	return vals
}

type BinaryPairOut <-chan [2]Binary

func (s BinaryPairOut) List() [][2]Binary {
	if s == nil {
		return nil
	}
	vals := [][2]Binary{}
	for val := range s {
		vals = append(vals, val)
	}
	return vals
}

func (s BinaryOut) IDs() IDsOut {
	if s == nil {
		return nil
	}
	ids := make(chan ID)
	go func(ids chan ID) {
		for b := range s {
			ids <- B2ID(b)
		}
		close(ids)
	}(ids)
	return ids
}
