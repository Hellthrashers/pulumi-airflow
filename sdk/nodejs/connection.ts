// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides an Airflow connection.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airflow from "pulumi-airflow";
 *
 * const example = new airflow.Connection("example", {
 *     connType: "example",
 *     connectionId: "example",
 * });
 * ```
 *
 * ## Import
 *
 * Connections can be imported using the connection key. terraform
 *
 * ```sh
 *  $ pulumi import airflow:index/connection:Connection default example
 * ```
 */
export class Connection extends pulumi.CustomResource {
    /**
     * Get an existing Connection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConnectionState, opts?: pulumi.CustomResourceOptions): Connection {
        return new Connection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'airflow:index/connection:Connection';

    /**
     * Returns true if the given object is an instance of Connection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Connection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Connection.__pulumiType;
    }

    /**
     * The connection type.
     */
    public readonly connType!: pulumi.Output<string>;
    /**
     * The connection ID.
     */
    public readonly connectionId!: pulumi.Output<string>;
    /**
     * The description of the connection.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Other values that cannot be put into another field, e.g. RSA keys.
     */
    public readonly extra!: pulumi.Output<string | undefined>;
    /**
     * The host of the connection.
     */
    public readonly host!: pulumi.Output<string | undefined>;
    /**
     * The login of the connection.
     */
    public readonly login!: pulumi.Output<string | undefined>;
    /**
     * The paasword of the connection.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * The port of the connection.
     */
    public readonly port!: pulumi.Output<number | undefined>;
    /**
     * The schema of the connection.
     */
    public readonly schema!: pulumi.Output<string | undefined>;

    /**
     * Create a Connection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConnectionArgs | ConnectionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ConnectionState | undefined;
            resourceInputs["connType"] = state ? state.connType : undefined;
            resourceInputs["connectionId"] = state ? state.connectionId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["extra"] = state ? state.extra : undefined;
            resourceInputs["host"] = state ? state.host : undefined;
            resourceInputs["login"] = state ? state.login : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["schema"] = state ? state.schema : undefined;
        } else {
            const args = argsOrState as ConnectionArgs | undefined;
            if ((!args || args.connType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connType'");
            }
            if ((!args || args.connectionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectionId'");
            }
            resourceInputs["connType"] = args ? args.connType : undefined;
            resourceInputs["connectionId"] = args ? args.connectionId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["extra"] = args ? args.extra : undefined;
            resourceInputs["host"] = args ? args.host : undefined;
            resourceInputs["login"] = args ? args.login : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["schema"] = args ? args.schema : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Connection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Connection resources.
 */
export interface ConnectionState {
    /**
     * The connection type.
     */
    connType?: pulumi.Input<string>;
    /**
     * The connection ID.
     */
    connectionId?: pulumi.Input<string>;
    /**
     * The description of the connection.
     */
    description?: pulumi.Input<string>;
    /**
     * Other values that cannot be put into another field, e.g. RSA keys.
     */
    extra?: pulumi.Input<string>;
    /**
     * The host of the connection.
     */
    host?: pulumi.Input<string>;
    /**
     * The login of the connection.
     */
    login?: pulumi.Input<string>;
    /**
     * The paasword of the connection.
     */
    password?: pulumi.Input<string>;
    /**
     * The port of the connection.
     */
    port?: pulumi.Input<number>;
    /**
     * The schema of the connection.
     */
    schema?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Connection resource.
 */
export interface ConnectionArgs {
    /**
     * The connection type.
     */
    connType: pulumi.Input<string>;
    /**
     * The connection ID.
     */
    connectionId: pulumi.Input<string>;
    /**
     * The description of the connection.
     */
    description?: pulumi.Input<string>;
    /**
     * Other values that cannot be put into another field, e.g. RSA keys.
     */
    extra?: pulumi.Input<string>;
    /**
     * The host of the connection.
     */
    host?: pulumi.Input<string>;
    /**
     * The login of the connection.
     */
    login?: pulumi.Input<string>;
    /**
     * The paasword of the connection.
     */
    password?: pulumi.Input<string>;
    /**
     * The port of the connection.
     */
    port?: pulumi.Input<number>;
    /**
     * The schema of the connection.
     */
    schema?: pulumi.Input<string>;
}
