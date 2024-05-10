package app

import (
	"github.com/konimarti/kalman"
	"gonum.org/v1/gonum/mat"
	"math"
	"math/rand"
)

type Coordinate struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func NewCoordinate(x, y float64) *Coordinate {
	return &Coordinate{
		X: x,
		Y: y,
	}
}

func GenerateUniformData() []*Coordinate {
	var s []*Coordinate

	for i := 0; i < 1000; i++ {
		x := rand.NormFloat64() * 10
		y := rand.NormFloat64() * 10

		c := NewCoordinate(x, y)
		s = append(s, c)
	}

	return s

}

/*
	중심 극한 정리

여러 개의 독립적이고 동일한 확률변수를 더 했을 때 그 합의 분포가 정규 분포를 따라감
균등 분포에 확률변수를 더하면 정규분포가 됨
*/
func GeneratePdfData() []*Coordinate {

	var s []*Coordinate
	corr := 0.5
	for i := 0; i < 1000; i++ {
		val1 := rand.NormFloat64()
		//평균 + (결과값 - 결과평균) * 표준 편차 / 결과표준편차
		// 평균 * 난수 + 루트(표준편차 - 평균) * 난수
		val3 := corr*val1 + math.Sqrt(1-corr)*rand.NormFloat64()
		x := 5.0 + val1
		y := 5.0 + val3

		c := NewCoordinate(x, y)
		s = append(s, c)
	}

	return s
}

func ExecuteKalmanFilterByLib(source []Coordinate) {

	ctx := newContext(source[0].X, source[0].Y)
	lti, nse := newSetup()

	// create Kalman filter
	filter := kalman.NewFilter(lti, nse)
	// no control
	control := mat.NewVecDense(4, nil)

	for i := range source {

		measurement := mat.NewVecDense(2, []float64{source[i].X, source[i].Y})

		filtered := filter.Apply(ctx, measurement, control)

		source[i].X = filtered.AtVec(0)
		source[i].Y = filtered.AtVec(1)
	}

}

func ExecuteKalmanFilter() {

}

func GaussianRandom(avg float64, stdev float64) {
	var v1, v2, s, temp float64

Loop:
	for {
		v1 = 2*(rand.Float64()/4) - 1
		v2 = 2*(rand.Float64()/4) - 1
		s = v1*v1 + v2*v2

		if s >= 1 || s == 0 {
			goto Loop
		} else {
			break Loop
		}
	}

	s = math.Sqrt((-2 * math.Log(s)) / s)
	temp = v1 * s
	temp = (stdev * temp) + avg

}

// 로지스틱 분포
func Inv_logistic_cdf(u float64) float64 {
	return math.Log(u / (1 - u))
}

// 레일리 분포
func Inv_rayleigh_cdf(u float64) float64 {
	return math.Sqrt(-2 * math.Log(1-u))
}
