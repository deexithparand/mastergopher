package array

type OneDResponse struct {
	Compute int32
	Size    int32
}

type TwoDResponse struct {
	DiagonalCompute int32
	Compute         int32
}

// one d and two d methods
func Array_Method_OneD(oned [5]int32) OneDResponse {

	// array means size is fixed, i.e its can't append values throws index out of range error

	// find length
	var size int32 = int32(len(oned)) // typecasting the default return type of len function which is int

	// compute sum of elements
	var compute int32 = 0
	for _, v := range oned {
		compute += int32(v) // typecasting default type of the for loop variable
	}

	return OneDResponse{
		Compute: compute,
		Size:    size,
	}
}

func Array_Method_TwoD(twod [5][5]int32) TwoDResponse {

	compute := 0
	diagonalcompute := 0

	// find compute O(N2)
	for i := range twod {
		for j := range twod[i] {
			compute += int(twod[i][j])
		}
	}

	// find diagonal compute values
	for i := range twod {
		for j := range twod[i] {
			if i == j {
				diagonalcompute += int(twod[i][j])
			}
		}
	}

	return TwoDResponse{
		DiagonalCompute: int32(diagonalcompute),
		Compute:         int32(compute),
	}

}
