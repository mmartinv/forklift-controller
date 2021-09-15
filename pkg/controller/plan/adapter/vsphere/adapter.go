package vsphere

import (
	api "github.com/konveyor/forklift-controller/pkg/apis/forklift/v1beta1"
	"github.com/konveyor/forklift-controller/pkg/controller/plan/adapter/base"
	plancontext "github.com/konveyor/forklift-controller/pkg/controller/plan/context"
)

//
// vSphere adapter.
type Adapter struct{}

//
// Constructs a vSphere builder.
func (r *Adapter) Builder(ctx *plancontext.Context) (builder base.Builder, err error) {
	b := &Builder{Context: ctx}
	err = b.Load()
	if err != nil {
		return
	}
	builder = b
	return
}

//
// Constructs a vSphere validator.
func (r *Adapter) Validator(plan *api.Plan) (validator base.Validator, err error) {
	v := &Validator{plan: plan}
	err = v.Load()
	if err != nil {
		return
	}
	validator = v
	return
}

//
// Constructs a vSphere client.
func (r *Adapter) Client(ctx *plancontext.Context) (client base.Client, err error) {
	c := &Client{Context: ctx}
	err = c.connect()
	if err != nil {
		return
	}
	client = c
	return
}
