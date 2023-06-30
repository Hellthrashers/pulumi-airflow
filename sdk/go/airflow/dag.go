// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airflow

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Airflow DAG.
//
// > Note this resource adpots an existing DAG and does not create a one, Also on delete the resource by default. A DAG is only deleted from state and not acutally deleted.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/Hellthrashers/pulumi-airflow/sdk/go/airflow"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := airflow.NewDag(ctx, "example", &airflow.DagArgs{
//				DagId:    pulumi.String("example"),
//				IsPaused: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// DAGs can be imported using the DAG Id. terraform
//
// ```sh
//
//	$ pulumi import airflow:index/dag:Dag default example
//
// ```
type Dag struct {
	pulumi.CustomResourceState

	// The ID of the DAG.
	DagId     pulumi.StringOutput  `pulumi:"dagId"`
	DeleteDag pulumi.BoolPtrOutput `pulumi:"deleteDag"`
	// User-provided DAG description, which can consist of several sentences or paragraphs that describe DAG contents.
	Description pulumi.StringOutput `pulumi:"description"`
	// The key containing the encrypted path to the file. Encryption and decryption take place only on the server. This prevents the client from reading an non-DAG file.
	FileToken pulumi.StringOutput `pulumi:"fileToken"`
	// The absolute path to the file.
	Fileloc pulumi.StringOutput `pulumi:"fileloc"`
	// Whether the DAG is currently seen by the scheduler(s).
	IsActive pulumi.BoolOutput `pulumi:"isActive"`
	// Whether the DAG is paused.
	IsPaused pulumi.BoolOutput `pulumi:"isPaused"`
	// Whether the DAG is SubDAG.
	IsSubdag pulumi.BoolOutput `pulumi:"isSubdag"`
	// If the DAG is SubDAG then it is the top level DAG identifier. Otherwise, null.
	RootDagId pulumi.StringOutput `pulumi:"rootDagId"`
}

// NewDag registers a new resource with the given unique name, arguments, and options.
func NewDag(ctx *pulumi.Context,
	name string, args *DagArgs, opts ...pulumi.ResourceOption) (*Dag, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DagId == nil {
		return nil, errors.New("invalid value for required argument 'DagId'")
	}
	if args.IsPaused == nil {
		return nil, errors.New("invalid value for required argument 'IsPaused'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Dag
	err := ctx.RegisterResource("airflow:index/dag:Dag", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDag gets an existing Dag resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDag(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DagState, opts ...pulumi.ResourceOption) (*Dag, error) {
	var resource Dag
	err := ctx.ReadResource("airflow:index/dag:Dag", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Dag resources.
type dagState struct {
	// The ID of the DAG.
	DagId     *string `pulumi:"dagId"`
	DeleteDag *bool   `pulumi:"deleteDag"`
	// User-provided DAG description, which can consist of several sentences or paragraphs that describe DAG contents.
	Description *string `pulumi:"description"`
	// The key containing the encrypted path to the file. Encryption and decryption take place only on the server. This prevents the client from reading an non-DAG file.
	FileToken *string `pulumi:"fileToken"`
	// The absolute path to the file.
	Fileloc *string `pulumi:"fileloc"`
	// Whether the DAG is currently seen by the scheduler(s).
	IsActive *bool `pulumi:"isActive"`
	// Whether the DAG is paused.
	IsPaused *bool `pulumi:"isPaused"`
	// Whether the DAG is SubDAG.
	IsSubdag *bool `pulumi:"isSubdag"`
	// If the DAG is SubDAG then it is the top level DAG identifier. Otherwise, null.
	RootDagId *string `pulumi:"rootDagId"`
}

type DagState struct {
	// The ID of the DAG.
	DagId     pulumi.StringPtrInput
	DeleteDag pulumi.BoolPtrInput
	// User-provided DAG description, which can consist of several sentences or paragraphs that describe DAG contents.
	Description pulumi.StringPtrInput
	// The key containing the encrypted path to the file. Encryption and decryption take place only on the server. This prevents the client from reading an non-DAG file.
	FileToken pulumi.StringPtrInput
	// The absolute path to the file.
	Fileloc pulumi.StringPtrInput
	// Whether the DAG is currently seen by the scheduler(s).
	IsActive pulumi.BoolPtrInput
	// Whether the DAG is paused.
	IsPaused pulumi.BoolPtrInput
	// Whether the DAG is SubDAG.
	IsSubdag pulumi.BoolPtrInput
	// If the DAG is SubDAG then it is the top level DAG identifier. Otherwise, null.
	RootDagId pulumi.StringPtrInput
}

func (DagState) ElementType() reflect.Type {
	return reflect.TypeOf((*dagState)(nil)).Elem()
}

type dagArgs struct {
	// The ID of the DAG.
	DagId     string `pulumi:"dagId"`
	DeleteDag *bool  `pulumi:"deleteDag"`
	// Whether the DAG is paused.
	IsPaused bool `pulumi:"isPaused"`
}

// The set of arguments for constructing a Dag resource.
type DagArgs struct {
	// The ID of the DAG.
	DagId     pulumi.StringInput
	DeleteDag pulumi.BoolPtrInput
	// Whether the DAG is paused.
	IsPaused pulumi.BoolInput
}

func (DagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dagArgs)(nil)).Elem()
}

type DagInput interface {
	pulumi.Input

	ToDagOutput() DagOutput
	ToDagOutputWithContext(ctx context.Context) DagOutput
}

func (*Dag) ElementType() reflect.Type {
	return reflect.TypeOf((**Dag)(nil)).Elem()
}

func (i *Dag) ToDagOutput() DagOutput {
	return i.ToDagOutputWithContext(context.Background())
}

func (i *Dag) ToDagOutputWithContext(ctx context.Context) DagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DagOutput)
}

// DagArrayInput is an input type that accepts DagArray and DagArrayOutput values.
// You can construct a concrete instance of `DagArrayInput` via:
//
//	DagArray{ DagArgs{...} }
type DagArrayInput interface {
	pulumi.Input

	ToDagArrayOutput() DagArrayOutput
	ToDagArrayOutputWithContext(context.Context) DagArrayOutput
}

type DagArray []DagInput

func (DagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Dag)(nil)).Elem()
}

func (i DagArray) ToDagArrayOutput() DagArrayOutput {
	return i.ToDagArrayOutputWithContext(context.Background())
}

func (i DagArray) ToDagArrayOutputWithContext(ctx context.Context) DagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DagArrayOutput)
}

// DagMapInput is an input type that accepts DagMap and DagMapOutput values.
// You can construct a concrete instance of `DagMapInput` via:
//
//	DagMap{ "key": DagArgs{...} }
type DagMapInput interface {
	pulumi.Input

	ToDagMapOutput() DagMapOutput
	ToDagMapOutputWithContext(context.Context) DagMapOutput
}

type DagMap map[string]DagInput

func (DagMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Dag)(nil)).Elem()
}

func (i DagMap) ToDagMapOutput() DagMapOutput {
	return i.ToDagMapOutputWithContext(context.Background())
}

func (i DagMap) ToDagMapOutputWithContext(ctx context.Context) DagMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DagMapOutput)
}

type DagOutput struct{ *pulumi.OutputState }

func (DagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dag)(nil)).Elem()
}

func (o DagOutput) ToDagOutput() DagOutput {
	return o
}

func (o DagOutput) ToDagOutputWithContext(ctx context.Context) DagOutput {
	return o
}

// The ID of the DAG.
func (o DagOutput) DagId() pulumi.StringOutput {
	return o.ApplyT(func(v *Dag) pulumi.StringOutput { return v.DagId }).(pulumi.StringOutput)
}

func (o DagOutput) DeleteDag() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Dag) pulumi.BoolPtrOutput { return v.DeleteDag }).(pulumi.BoolPtrOutput)
}

// User-provided DAG description, which can consist of several sentences or paragraphs that describe DAG contents.
func (o DagOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Dag) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The key containing the encrypted path to the file. Encryption and decryption take place only on the server. This prevents the client from reading an non-DAG file.
func (o DagOutput) FileToken() pulumi.StringOutput {
	return o.ApplyT(func(v *Dag) pulumi.StringOutput { return v.FileToken }).(pulumi.StringOutput)
}

// The absolute path to the file.
func (o DagOutput) Fileloc() pulumi.StringOutput {
	return o.ApplyT(func(v *Dag) pulumi.StringOutput { return v.Fileloc }).(pulumi.StringOutput)
}

// Whether the DAG is currently seen by the scheduler(s).
func (o DagOutput) IsActive() pulumi.BoolOutput {
	return o.ApplyT(func(v *Dag) pulumi.BoolOutput { return v.IsActive }).(pulumi.BoolOutput)
}

// Whether the DAG is paused.
func (o DagOutput) IsPaused() pulumi.BoolOutput {
	return o.ApplyT(func(v *Dag) pulumi.BoolOutput { return v.IsPaused }).(pulumi.BoolOutput)
}

// Whether the DAG is SubDAG.
func (o DagOutput) IsSubdag() pulumi.BoolOutput {
	return o.ApplyT(func(v *Dag) pulumi.BoolOutput { return v.IsSubdag }).(pulumi.BoolOutput)
}

// If the DAG is SubDAG then it is the top level DAG identifier. Otherwise, null.
func (o DagOutput) RootDagId() pulumi.StringOutput {
	return o.ApplyT(func(v *Dag) pulumi.StringOutput { return v.RootDagId }).(pulumi.StringOutput)
}

type DagArrayOutput struct{ *pulumi.OutputState }

func (DagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Dag)(nil)).Elem()
}

func (o DagArrayOutput) ToDagArrayOutput() DagArrayOutput {
	return o
}

func (o DagArrayOutput) ToDagArrayOutputWithContext(ctx context.Context) DagArrayOutput {
	return o
}

func (o DagArrayOutput) Index(i pulumi.IntInput) DagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Dag {
		return vs[0].([]*Dag)[vs[1].(int)]
	}).(DagOutput)
}

type DagMapOutput struct{ *pulumi.OutputState }

func (DagMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Dag)(nil)).Elem()
}

func (o DagMapOutput) ToDagMapOutput() DagMapOutput {
	return o
}

func (o DagMapOutput) ToDagMapOutputWithContext(ctx context.Context) DagMapOutput {
	return o
}

func (o DagMapOutput) MapIndex(k pulumi.StringInput) DagOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Dag {
		return vs[0].(map[string]*Dag)[vs[1].(string)]
	}).(DagOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DagInput)(nil)).Elem(), &Dag{})
	pulumi.RegisterInputType(reflect.TypeOf((*DagArrayInput)(nil)).Elem(), DagArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DagMapInput)(nil)).Elem(), DagMap{})
	pulumi.RegisterOutputType(DagOutput{})
	pulumi.RegisterOutputType(DagArrayOutput{})
	pulumi.RegisterOutputType(DagMapOutput{})
}
