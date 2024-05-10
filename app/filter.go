package app

import (
	"github.com/konimarti/kalman"
	"github.com/konimarti/lti"
	"gonum.org/v1/gonum/mat"
)

func newContext(v1 float64, v2 float64) *kalman.Context {
	// define current context
	ctx := kalman.Context{
		// init state: pos_x = 0, pox_y = 0, v_x = 30 km/h, v_y = 10 km/h
		X: mat.NewVecDense(4, []float64{0, 0, v1, v2}),
		// initial covariance matrix
		P: mat.NewDense(4, 4, []float64{
			1, 0, 0, 0,
			0, 1, 0, 0,
			0, 0, 1, 0,
			0, 0, 0, 1,
		}),
	}
	return &ctx
}

func newSetup() (lti.Discrete, kalman.Noise) {
	// define LTI system
	dt := 0.1
	lti := lti.Discrete{
		Ad: mat.NewDense(4, 4, []float64{
			1, 0, dt, 0,
			0, 1, 0, dt,
			0, 0, 1, 0,
			0, 0, 0, 1,
		}),
		Bd: mat.NewDense(4, 4, nil),
		C: mat.NewDense(2, 4, []float64{
			0, 0, 1, 0,
			0, 0, 0, 1,
		}),
		D: mat.NewDense(2, 4, nil),
	}

	// G
	G := mat.NewDense(4, 2, []float64{
		0, 0,
		0, 0,
		1, 0,
		0, 1,
	})
	var Gd mat.Dense
	Gd.Mul(lti.Ad, G)

	// process model covariance matrix
	qk := mat.NewDense(2, 2, []float64{
		0.01, 0,
		0, 0.01,
	})
	var Q mat.Dense
	Q.Product(&Gd, qk, Gd.T())
	// define system and measurement noise
	// measurement errors
	corr := 0.5
	R := mat.NewDense(2, 2, []float64{1, corr, corr, 1})

	// create noise struct
	nse := kalman.Noise{&Q, R}

	return lti, nse
}
