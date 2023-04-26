// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides an Airflow pool.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airflow from "@pulumi/airflow";
 *
 * const example = new airflow.Pool("example", {
 *     slots: 2,
 * });
 * ```
 *
 * ## Import
 *
 * Pools can be imported using the pool name. terraform
 *
 * ```sh
 *  $ pulumi import airflow:index/pool:Pool default example
 * ```
 */
export class Pool extends pulumi.CustomResource {
    /**
     * Get an existing Pool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PoolState, opts?: pulumi.CustomResourceOptions): Pool {
        return new Pool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'airflow:index/pool:Pool';

    /**
     * Returns true if the given object is an instance of Pool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Pool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Pool.__pulumiType;
    }

    /**
     * The name of pool.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The number of slots used by running/queued tasks at the moment.
     */
    public /*out*/ readonly occupiedSlots!: pulumi.Output<number>;
    /**
     * The number of free slots at the moment.
     */
    public /*out*/ readonly openSlots!: pulumi.Output<number>;
    /**
     * The number of slots used by queued tasks at the moment.
     */
    public /*out*/ readonly queuedSlots!: pulumi.Output<number>;
    /**
     * The maximum number of slots that can be assigned to tasks. One job may occupy one or more slots.
     */
    public readonly slots!: pulumi.Output<number>;
    /**
     * The number of slots used by running tasks at the moment.
     */
    public /*out*/ readonly usedSlots!: pulumi.Output<number>;

    /**
     * Create a Pool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PoolArgs | PoolState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PoolState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["occupiedSlots"] = state ? state.occupiedSlots : undefined;
            resourceInputs["openSlots"] = state ? state.openSlots : undefined;
            resourceInputs["queuedSlots"] = state ? state.queuedSlots : undefined;
            resourceInputs["slots"] = state ? state.slots : undefined;
            resourceInputs["usedSlots"] = state ? state.usedSlots : undefined;
        } else {
            const args = argsOrState as PoolArgs | undefined;
            if ((!args || args.slots === undefined) && !opts.urn) {
                throw new Error("Missing required property 'slots'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["slots"] = args ? args.slots : undefined;
            resourceInputs["occupiedSlots"] = undefined /*out*/;
            resourceInputs["openSlots"] = undefined /*out*/;
            resourceInputs["queuedSlots"] = undefined /*out*/;
            resourceInputs["usedSlots"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Pool.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Pool resources.
 */
export interface PoolState {
    /**
     * The name of pool.
     */
    name?: pulumi.Input<string>;
    /**
     * The number of slots used by running/queued tasks at the moment.
     */
    occupiedSlots?: pulumi.Input<number>;
    /**
     * The number of free slots at the moment.
     */
    openSlots?: pulumi.Input<number>;
    /**
     * The number of slots used by queued tasks at the moment.
     */
    queuedSlots?: pulumi.Input<number>;
    /**
     * The maximum number of slots that can be assigned to tasks. One job may occupy one or more slots.
     */
    slots?: pulumi.Input<number>;
    /**
     * The number of slots used by running tasks at the moment.
     */
    usedSlots?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Pool resource.
 */
export interface PoolArgs {
    /**
     * The name of pool.
     */
    name?: pulumi.Input<string>;
    /**
     * The maximum number of slots that can be assigned to tasks. One job may occupy one or more slots.
     */
    slots: pulumi.Input<number>;
}