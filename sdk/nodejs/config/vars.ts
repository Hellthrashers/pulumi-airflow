// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("airflow");

export declare const baseEndpoint: string | undefined;
Object.defineProperty(exports, "baseEndpoint", {
    get() {
        return __config.get("baseEndpoint") ?? utilities.getEnv("AIRFLOW_BASE_ENDPOINT");
    },
    enumerable: true,
});

/**
 * Disable SSL verification
 */
export declare const disableSslVerification: boolean | undefined;
Object.defineProperty(exports, "disableSslVerification", {
    get() {
        return __config.getObject<boolean>("disableSslVerification");
    },
    enumerable: true,
});

/**
 * The oauth to use for API authentication
 */
export declare const oauth2Token: string | undefined;
Object.defineProperty(exports, "oauth2Token", {
    get() {
        return __config.get("oauth2Token");
    },
    enumerable: true,
});

/**
 * The password to use for API basic authentication
 */
export declare const password: string | undefined;
Object.defineProperty(exports, "password", {
    get() {
        return __config.get("password");
    },
    enumerable: true,
});

/**
 * The username to use for API basic authentication
 */
export declare const username: string | undefined;
Object.defineProperty(exports, "username", {
    get() {
        return __config.get("username");
    },
    enumerable: true,
});

