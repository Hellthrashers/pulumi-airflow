// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package airflow

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Airflow connection.
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
//			_, err := airflow.NewConnection(ctx, "example", &airflow.ConnectionArgs{
//				ConnType:     pulumi.String("example"),
//				ConnectionId: pulumi.String("example"),
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
// Connections can be imported using the connection key. terraform
//
// ```sh
//
//	$ pulumi import airflow:index/connection:Connection default example
//
// ```
type Connection struct {
	pulumi.CustomResourceState

	// The connection type.
	ConnType pulumi.StringOutput `pulumi:"connType"`
	// The connection ID.
	ConnectionId pulumi.StringOutput `pulumi:"connectionId"`
	// The description of the connection.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Other values that cannot be put into another field, e.g. RSA keys.
	Extra pulumi.StringPtrOutput `pulumi:"extra"`
	// The host of the connection.
	Host pulumi.StringPtrOutput `pulumi:"host"`
	// The login of the connection.
	Login pulumi.StringPtrOutput `pulumi:"login"`
	// The paasword of the connection.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The port of the connection.
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// The schema of the connection.
	Schema pulumi.StringPtrOutput `pulumi:"schema"`
}

// NewConnection registers a new resource with the given unique name, arguments, and options.
func NewConnection(ctx *pulumi.Context,
	name string, args *ConnectionArgs, opts ...pulumi.ResourceOption) (*Connection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnType == nil {
		return nil, errors.New("invalid value for required argument 'ConnType'")
	}
	if args.ConnectionId == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Connection
	err := ctx.RegisterResource("airflow:index/connection:Connection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnection gets an existing Connection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionState, opts ...pulumi.ResourceOption) (*Connection, error) {
	var resource Connection
	err := ctx.ReadResource("airflow:index/connection:Connection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Connection resources.
type connectionState struct {
	// The connection type.
	ConnType *string `pulumi:"connType"`
	// The connection ID.
	ConnectionId *string `pulumi:"connectionId"`
	// The description of the connection.
	Description *string `pulumi:"description"`
	// Other values that cannot be put into another field, e.g. RSA keys.
	Extra *string `pulumi:"extra"`
	// The host of the connection.
	Host *string `pulumi:"host"`
	// The login of the connection.
	Login *string `pulumi:"login"`
	// The paasword of the connection.
	Password *string `pulumi:"password"`
	// The port of the connection.
	Port *int `pulumi:"port"`
	// The schema of the connection.
	Schema *string `pulumi:"schema"`
}

type ConnectionState struct {
	// The connection type.
	ConnType pulumi.StringPtrInput
	// The connection ID.
	ConnectionId pulumi.StringPtrInput
	// The description of the connection.
	Description pulumi.StringPtrInput
	// Other values that cannot be put into another field, e.g. RSA keys.
	Extra pulumi.StringPtrInput
	// The host of the connection.
	Host pulumi.StringPtrInput
	// The login of the connection.
	Login pulumi.StringPtrInput
	// The paasword of the connection.
	Password pulumi.StringPtrInput
	// The port of the connection.
	Port pulumi.IntPtrInput
	// The schema of the connection.
	Schema pulumi.StringPtrInput
}

func (ConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionState)(nil)).Elem()
}

type connectionArgs struct {
	// The connection type.
	ConnType string `pulumi:"connType"`
	// The connection ID.
	ConnectionId string `pulumi:"connectionId"`
	// The description of the connection.
	Description *string `pulumi:"description"`
	// Other values that cannot be put into another field, e.g. RSA keys.
	Extra *string `pulumi:"extra"`
	// The host of the connection.
	Host *string `pulumi:"host"`
	// The login of the connection.
	Login *string `pulumi:"login"`
	// The paasword of the connection.
	Password *string `pulumi:"password"`
	// The port of the connection.
	Port *int `pulumi:"port"`
	// The schema of the connection.
	Schema *string `pulumi:"schema"`
}

// The set of arguments for constructing a Connection resource.
type ConnectionArgs struct {
	// The connection type.
	ConnType pulumi.StringInput
	// The connection ID.
	ConnectionId pulumi.StringInput
	// The description of the connection.
	Description pulumi.StringPtrInput
	// Other values that cannot be put into another field, e.g. RSA keys.
	Extra pulumi.StringPtrInput
	// The host of the connection.
	Host pulumi.StringPtrInput
	// The login of the connection.
	Login pulumi.StringPtrInput
	// The paasword of the connection.
	Password pulumi.StringPtrInput
	// The port of the connection.
	Port pulumi.IntPtrInput
	// The schema of the connection.
	Schema pulumi.StringPtrInput
}

func (ConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionArgs)(nil)).Elem()
}

type ConnectionInput interface {
	pulumi.Input

	ToConnectionOutput() ConnectionOutput
	ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput
}

func (*Connection) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil)).Elem()
}

func (i *Connection) ToConnectionOutput() ConnectionOutput {
	return i.ToConnectionOutputWithContext(context.Background())
}

func (i *Connection) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionOutput)
}

// ConnectionArrayInput is an input type that accepts ConnectionArray and ConnectionArrayOutput values.
// You can construct a concrete instance of `ConnectionArrayInput` via:
//
//	ConnectionArray{ ConnectionArgs{...} }
type ConnectionArrayInput interface {
	pulumi.Input

	ToConnectionArrayOutput() ConnectionArrayOutput
	ToConnectionArrayOutputWithContext(context.Context) ConnectionArrayOutput
}

type ConnectionArray []ConnectionInput

func (ConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Connection)(nil)).Elem()
}

func (i ConnectionArray) ToConnectionArrayOutput() ConnectionArrayOutput {
	return i.ToConnectionArrayOutputWithContext(context.Background())
}

func (i ConnectionArray) ToConnectionArrayOutputWithContext(ctx context.Context) ConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionArrayOutput)
}

// ConnectionMapInput is an input type that accepts ConnectionMap and ConnectionMapOutput values.
// You can construct a concrete instance of `ConnectionMapInput` via:
//
//	ConnectionMap{ "key": ConnectionArgs{...} }
type ConnectionMapInput interface {
	pulumi.Input

	ToConnectionMapOutput() ConnectionMapOutput
	ToConnectionMapOutputWithContext(context.Context) ConnectionMapOutput
}

type ConnectionMap map[string]ConnectionInput

func (ConnectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Connection)(nil)).Elem()
}

func (i ConnectionMap) ToConnectionMapOutput() ConnectionMapOutput {
	return i.ToConnectionMapOutputWithContext(context.Background())
}

func (i ConnectionMap) ToConnectionMapOutputWithContext(ctx context.Context) ConnectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionMapOutput)
}

type ConnectionOutput struct{ *pulumi.OutputState }

func (ConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil)).Elem()
}

func (o ConnectionOutput) ToConnectionOutput() ConnectionOutput {
	return o
}

func (o ConnectionOutput) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return o
}

// The connection type.
func (o ConnectionOutput) ConnType() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.ConnType }).(pulumi.StringOutput)
}

// The connection ID.
func (o ConnectionOutput) ConnectionId() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.ConnectionId }).(pulumi.StringOutput)
}

// The description of the connection.
func (o ConnectionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Other values that cannot be put into another field, e.g. RSA keys.
func (o ConnectionOutput) Extra() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringPtrOutput { return v.Extra }).(pulumi.StringPtrOutput)
}

// The host of the connection.
func (o ConnectionOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringPtrOutput { return v.Host }).(pulumi.StringPtrOutput)
}

// The login of the connection.
func (o ConnectionOutput) Login() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringPtrOutput { return v.Login }).(pulumi.StringPtrOutput)
}

// The paasword of the connection.
func (o ConnectionOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// The port of the connection.
func (o ConnectionOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Connection) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

// The schema of the connection.
func (o ConnectionOutput) Schema() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringPtrOutput { return v.Schema }).(pulumi.StringPtrOutput)
}

type ConnectionArrayOutput struct{ *pulumi.OutputState }

func (ConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Connection)(nil)).Elem()
}

func (o ConnectionArrayOutput) ToConnectionArrayOutput() ConnectionArrayOutput {
	return o
}

func (o ConnectionArrayOutput) ToConnectionArrayOutputWithContext(ctx context.Context) ConnectionArrayOutput {
	return o
}

func (o ConnectionArrayOutput) Index(i pulumi.IntInput) ConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Connection {
		return vs[0].([]*Connection)[vs[1].(int)]
	}).(ConnectionOutput)
}

type ConnectionMapOutput struct{ *pulumi.OutputState }

func (ConnectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Connection)(nil)).Elem()
}

func (o ConnectionMapOutput) ToConnectionMapOutput() ConnectionMapOutput {
	return o
}

func (o ConnectionMapOutput) ToConnectionMapOutputWithContext(ctx context.Context) ConnectionMapOutput {
	return o
}

func (o ConnectionMapOutput) MapIndex(k pulumi.StringInput) ConnectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Connection {
		return vs[0].(map[string]*Connection)[vs[1].(string)]
	}).(ConnectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionInput)(nil)).Elem(), &Connection{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionArrayInput)(nil)).Elem(), ConnectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionMapInput)(nil)).Elem(), ConnectionMap{})
	pulumi.RegisterOutputType(ConnectionOutput{})
	pulumi.RegisterOutputType(ConnectionArrayOutput{})
	pulumi.RegisterOutputType(ConnectionMapOutput{})
}
