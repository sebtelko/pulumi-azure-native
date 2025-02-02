// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./environment";
export * from "./environmentSetting";
export * from "./galleryImage";
export * from "./getEnvironment";
export * from "./getEnvironmentSetting";
export * from "./getGalleryImage";
export * from "./getGlobalUserEnvironment";
export * from "./getGlobalUserOperationBatchStatus";
export * from "./getGlobalUserOperationStatus";
export * from "./getGlobalUserPersonalPreferences";
export * from "./getLab";
export * from "./getLabAccount";
export * from "./getLabAccountRegionalAvailability";
export * from "./getUser";
export * from "./lab";
export * from "./labAccount";
export * from "./listGlobalUserEnvironments";
export * from "./listGlobalUserLabs";
export * from "./user";

// Export enums:
export * from "../types/enums/labservices";

// Export sub-modules:
import * as v20181015 from "./v20181015";

export {
    v20181015,
};

// Import resources to register:
import { Environment } from "./environment";
import { EnvironmentSetting } from "./environmentSetting";
import { GalleryImage } from "./galleryImage";
import { Lab } from "./lab";
import { LabAccount } from "./labAccount";
import { User } from "./user";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:labservices:Environment":
                return new Environment(name, <any>undefined, { urn })
            case "azure-native:labservices:EnvironmentSetting":
                return new EnvironmentSetting(name, <any>undefined, { urn })
            case "azure-native:labservices:GalleryImage":
                return new GalleryImage(name, <any>undefined, { urn })
            case "azure-native:labservices:Lab":
                return new Lab(name, <any>undefined, { urn })
            case "azure-native:labservices:LabAccount":
                return new LabAccount(name, <any>undefined, { urn })
            case "azure-native:labservices:User":
                return new User(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "labservices", _module)
