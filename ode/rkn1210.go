// Copyright ©2022 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ode

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

// RKN1210 Runge-Kutta-Nyström integrator of order 12(10).
type RKN1210 struct {
	dom        float64
	y, dy, aux *mat.VecDense
	// Low and high order terms from integration.
	hFDb, hFb, hFDbhat, hFbhat *mat.VecDense
	// second order differential equations.
	diffs [rk1210Len]*mat.VecDense
	// Step control.
	atol, minStep, maxStep float64
	fx                     func(y *mat.VecDense, t float64, x mat.Vector)
}

// Number of Runge-Kutta iteration steps to calculate both high and low order solutions.
const rk1210Len = 17

// Init initializes the Runge-Kutta-Nyström integrator with a second order
// ODE initial value problem.
func (rk *RKN1210) Init(ivp IVP2) {
	rk.fx = ivp.Func
	rk.dom, rk.y = ivp.T0, mat.VecDenseCopyOf(ivp.Y0)
	rk.dy = mat.VecDenseCopyOf(ivp.DY0)
	ny := rk.y.Len()
	rk.aux = mat.NewVecDense(ny, nil)
	rk.hFDbhat = mat.NewVecDense(ny, nil)
	rk.hFbhat = mat.NewVecDense(ny, nil)
	rk.hFDb = mat.NewVecDense(ny, nil)
	rk.hFb = mat.NewVecDense(ny, nil)
	for j := range rk.diffs {
		rk.diffs[j] = mat.NewVecDense(ny, nil)
	}
}

// State stores the current values of the solved ODE as calculated by the
// RKN integrator into s.
func (rk *RKN1210) State(s *State2) {
	s.T = rk.dom
	s.Y.CloneFromVec(rk.y)
	s.DY.CloneFromVec(rk.dy)
}

// SetState sets the current values of the solved ODE to s.
func (rk *RKN1210) SetState(s State2) {
	rk.dom = s.T
	rk.y.CloneFromVec(s.Y)
	rk.dy.CloneFromVec(s.DY)
}

// Step implements Integrator interface. Advances solution by step h. If algorithm
// is set to adaptive then h is just a suggestion.
func (rk *RKN1210) Step(h float64) (float64, error) {
	adaptive := rk.atol > 0
	// Declare shorthand variable names.
	y := rk.y
	dy := rk.dy
	aux := rk.aux
	// high order B's.
	hFDbH := rk.hFDbhat
	hFbH := rk.hFbhat
	// Low order B's for error estimation.
	hFDb := rk.hFDb
	hFb := rk.hFb

	fun := rk.fx
	F := rk.diffs
	t := rk.dom

solve:
	hFDbH.Zero()
	hFbH.Zero()
	hFb.Zero()
	hFDb.Zero()
	h2 := h * h
	for j, Fj := range F {
		// aux = y + h*c[j] + F*h*h*a[j]
		hc := h * rkn12c[j]
		aux.AddScaledVec(y, hc, dy)
		for i, Fi := range F[:j] {
			aux.AddScaledVec(aux, h2*rkn12A[j][i], Fi)
		}
		// finally F[:,j] = Func( aux ) @ t+h*c[j]
		fun(Fj, t+hc, aux) // 17 function evaluations.
		// Calculate high order h*F*b
		hFDbH.AddScaledVec(hFDbH, h*rkn12bphat[j], Fj)
		hFbH.AddScaledVec(hFbH, h*rkn12bhat[j], Fj)
		if adaptive {
			// Low order h*F*b for error estimation if user requested tolerance.
			hFb.AddScaledVec(hFb, h*rkn12b[j], Fj)
			hFDb.AddScaledVec(hFDb, h*rkn12bp[j], Fj)
		}
	}
	if adaptive {
		const (
			preCond = 11.0
			relax   = 0.9
		)
		// Calculate the difference between high and low order terms.
		aux.SubVec(hFb, hFbH)
		errMax := h * mat.Norm(aux, math.Inf(1)) // error ~ h*| y_l- y_h |
		aux.SubVec(hFDb, hFDbH)
		// In taking the Max we use worst case error.
		errMax = math.Max(mat.Norm(aux, math.Inf(1)), errMax) // error ~ | y'_l- y'_h |
		errRatio := rk.atol / (errMax * h * preCond)
		hnew := relax * math.Pow(errRatio, 1./11.)
		hnew = math.Min(math.Max(hnew, rk.minStep), rk.maxStep)
		if errMax > rk.atol && h > rk.minStep {
			// Error is not permissible and we may redo the step.
			h = hnew
			goto solve
		}
		// The error is within tolerance and we may suggest the user use a larger step.
		// Modify return value to suggest new step.
		defer func() { h = hnew }()
	}
	// Calculate next step solutions with high order B's:
	//  y[i+1] = y[i] + h*(dy[i] + hFbhat)
	//  dy[i+1] = dy[i] + hFDbhat
	aux.AddVec(dy, hFbH)
	y.AddScaledVec(y, h, aux)
	dy.AddVec(dy, hFDbH)
	rk.dom += h
	return h, nil
}

// NewRKN1210 configures a new RKN1210 instance.
func NewRKN1210(cfg Parameters) *RKN1210 {
	if (cfg.AbsTolerance != 0 && cfg.MaxStep <= 0) || cfg.MaxStep < cfg.MinStep || cfg.MinStep < 0 {
		panic("invalid parameter supplied")
	}
	return &RKN1210{
		atol:    cfg.AbsTolerance,
		minStep: cfg.MinStep,
		maxStep: cfg.MaxStep,
	}
}

var (
	rkn12c = [rk1210Len]float64{0.0e0, 2.0e-2, 4.0e-2, 1.0e-1, 1.33333333333333333333333333333e-1, 1.6e-1, 5.0e-2, 2.0e-1, 2.5e-1, 3.33333333333333333333333333333e-1, 5.0e-1, 5.55555555555555555555555555556e-1, 7.5e-1, 8.57142857142857142857142857143e-1, 9.45216222272014340129957427739e-1, 1.0e0, 1.0e0}
	rkn12A = [rk1210Len][rk1210Len]float64{
		1:  {0: 2e-4},
		2:  {2.66666666666666666666666666667e-4, 5.33333333333333333333333333333e-4},
		3:  {2.91666666666666666666666666667e-3, -4.16666666666666666666666666667e-3, 6.25e-3},
		4:  {1.64609053497942386831275720165e-3, 0, 5.48696844993141289437585733882e-3, 1.75582990397805212620027434842e-3},
		5:  {1.9456e-3, 0, 7.15174603174603174603174603175e-3, 2.91271111111111111111111111111e-3, 7.89942857142857142857142857143e-4},
		6:  {5.6640625e-4, 0, 8.80973048941798941798941798942e-4, -4.36921296296296296296296296296e-4, 3.39006696428571428571428571429e-4, -9.94646990740740740740740740741e-5},
		7:  {3.08333333333333333333333333333e-3, 0, 0, 1.77777777777777777777777777778e-3, 2.7e-3, 1.57828282828282828282828282828e-3, 1.08606060606060606060606060606e-2},
		8:  {3.65183937480112971375119150338e-3, 0, 3.96517171407234306617557289807e-3, 3.19725826293062822350093426091e-3, 8.22146730685543536968701883401e-3, -1.31309269595723798362013884863e-3, 9.77158696806486781562609494147e-3, 3.75576906923283379487932641079e-3},
		9:  {3.70724106871850081019565530521e-3, 0, 5.08204585455528598076108163479e-3, 1.17470800217541204473569104943e-3, -2.11476299151269914996229766362e-2, 6.01046369810788081222573525136e-2, 2.01057347685061881846748708777e-2, -2.83507501229335808430366774368e-2, 1.48795689185819327555905582479e-2},
		10: {3.51253765607334415311308293052e-2, 0, -8.61574919513847910340576078545e-3, -5.79144805100791652167632252471e-3, 1.94555482378261584239438810411e0, -3.43512386745651359636787167574e0, -1.09307011074752217583892572001e-1, 2.3496383118995166394320161088e0, -7.56009408687022978027190729778e-1, 1.09528972221569264246502018618e-1},
		11: {2.05277925374824966509720571672e-2, 0, -7.28644676448017991778247943149e-3, -2.11535560796184024069259562549e-3, 9.27580796872352224256768033235e-1, -1.65228248442573667907302673325e0, -2.10795630056865698191914366913e-2, 1.20653643262078715447708832536e0, -4.13714477001066141324662463645e-1, 9.07987398280965375956795739516e-2, 5.35555260053398504916870658215e-3},
		12: {-1.43240788755455150458921091632e-1, 0, 1.25287037730918172778464480231e-2, 6.82601916396982712868112411737e-3, -4.79955539557438726550216254291e0, 5.69862504395194143379169794156e0, 7.55343036952364522249444028716e-1, -1.27554878582810837175400796542e-1, -1.96059260511173843289133255423e0, 9.18560905663526240976234285341e-1, -2.38800855052844310534827013402e-1, 1.59110813572342155138740170963e-1},
		13: {8.04501920552048948697230778134e-1, 0, -1.66585270670112451778516268261e-2, -2.1415834042629734811731437191e-2, 1.68272359289624658702009353564e1, -1.11728353571760979267882984241e1, -3.37715929722632374148856475521e0, -1.52433266553608456461817682939e1, 1.71798357382154165620247684026e1, -5.43771923982399464535413738556e0, 1.38786716183646557551256778839e0, -5.92582773265281165347677029181e-1, 2.96038731712973527961592794552e-2},
		14: {-9.13296766697358082096250482648e-1, 0, 2.41127257578051783924489946102e-3, 1.76581226938617419820698839226e-2, -1.48516497797203838246128557088e1, 2.15897086700457560030782161561e0, 3.99791558311787990115282754337e0, 2.84341518002322318984542514988e1, -2.52593643549415984378843352235e1, 7.7338785423622373655340014114e0, -1.8913028948478674610382580129e0, 1.00148450702247178036685959248e0, 4.64119959910905190510518247052e-3, 1.12187550221489570339750499063e-2},
		15: {-2.75196297205593938206065227039e-1, 0, 3.66118887791549201342293285553e-2, 9.7895196882315626246509967162e-3, -1.2293062345886210304214726509e1, 1.42072264539379026942929665966e1, 1.58664769067895368322481964272e0, 2.45777353275959454390324346975e0, -8.93519369440327190552259086374e0, 4.37367273161340694839327077512e0, -1.83471817654494916304344410264e0, 1.15920852890614912078083198373e0, -1.72902531653839221518003422953e-2, 1.93259779044607666727649875324e-2, 5.20444293755499311184926401526e-3},
		16: {1.30763918474040575879994562983e0, 0, 1.73641091897458418670879991296e-2, -1.8544456454265795024362115588e-2, 1.48115220328677268968478356223e1, 9.38317630848247090787922177126e0, -5.2284261999445422541474024553e0, -4.89512805258476508040093482743e1, 3.82970960343379225625583875836e1, -1.05873813369759797091619037505e1, 2.43323043762262763585119618787e0, -1.04534060425754442848652456513e0, 7.17732095086725945198184857508e-2, 2.16221097080827826905505320027e-3, 7.00959575960251423699282781988e-3, 0},
	}
	// low order b.
	rkn12b = [rk1210Len]float64{1.70087019070069917527544646189e-2, 0.0e0, 0.0e0, 0.0e0, 0.0e0, 0.0e0, 7.22593359308314069488600038463e-2, 3.72026177326753045388210502067e-1, -4.01821145009303521439340233863e-1, 3.35455068301351666696584034896e-1, -1.31306501075331808430281840783e-1, 1.89431906616048652722659836455e-1, 2.68408020400290479053691655806e-2, 1.63056656059179238935180933102e-2, 3.79998835669659456166597387323e-3, 0.0e0, 0.0e0}
	// low order b'.
	rkn12bp = [rk1210Len]float64{1.70087019070069917527544646189e-2, 0.0e0, 0.0e0, 0.0e0, 0.0e0, 0.0e0, 7.60624588745593757356421093119e-2, 4.65032721658441306735263127583e-1, -5.35761526679071361919120311817e-1, 5.03182602452027500044876052344e-1, -2.62613002150663616860563681567e-1, 4.26221789886109468625984632024e-1, 1.07363208160116191621476662322e-1, 1.14139659241425467254626653171e-1, 6.93633866500486770090602920091e-2, 2.0e-2, 0.0e0}
	// high order b.
	rkn12bhat = [rk1210Len]float64{1.21278685171854149768890395495e-2, 0.0e0, 0.0e0, 0.0e0, 0.0e0, 0.0e0, 8.62974625156887444363792274411e-2, 2.52546958118714719432343449316e-1, -1.97418679932682303358307954886e-1, 2.03186919078972590809261561009e-1, -2.07758080777149166121933554691e-2, 1.09678048745020136250111237823e-1, 3.80651325264665057344878719105e-2, 1.16340688043242296440927709215e-2, 4.65802970402487868693615238455e-3, 0.0e0, 0.0e0}
	// high order b'.
	rkn12bphat = [rk1210Len]float64{1.21278685171854149768890395495e-2, 0.0e0, 0.0e0, 0.0e0, 0.0e0, 0.0e0, 9.08394342270407836172412920433e-2, 3.15683697648393399290429311645e-1, -2.63224906576909737811077273181e-1, 3.04780378618458886213892341513e-1, -4.15516161554298332243867109382e-2, 2.46775609676295306562750285101e-1, 1.52260530105866022937951487642e-1, 8.14384816302696075086493964505e-2, 8.50257119389081128008018326881e-2, -9.15518963007796287314100251351e-3, 2.5e-2}
)
