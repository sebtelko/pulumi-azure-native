// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./getIntegrationAccount";
export * from "./getIntegrationAccountAgreement";
export * from "./getIntegrationAccountAssembly";
export * from "./getIntegrationAccountBatchConfiguration";
export * from "./getIntegrationAccountCertificate";
export * from "./getIntegrationAccountMap";
export * from "./getIntegrationAccountPartner";
export * from "./getIntegrationAccountSchema";
export * from "./getIntegrationAccountSession";
export * from "./getIntegrationServiceEnvironment";
export * from "./getIntegrationServiceEnvironmentManagedApi";
export * from "./getRosettaNetProcessConfiguration";
export * from "./getWorkflow";
export * from "./getWorkflowAccessKey";
export * from "./integrationAccount";
export * from "./integrationAccountAgreement";
export * from "./integrationAccountAssembly";
export * from "./integrationAccountBatchConfiguration";
export * from "./integrationAccountCertificate";
export * from "./integrationAccountMap";
export * from "./integrationAccountPartner";
export * from "./integrationAccountSchema";
export * from "./integrationAccountSession";
export * from "./integrationServiceEnvironment";
export * from "./integrationServiceEnvironmentManagedApi";
export * from "./listIntegrationAccountAgreementContentCallbackUrl";
export * from "./listIntegrationAccountAssemblyContentCallbackUrl";
export * from "./listIntegrationAccountCallbackUrl";
export * from "./listIntegrationAccountKeyVaultKeys";
export * from "./listIntegrationAccountMapContentCallbackUrl";
export * from "./listIntegrationAccountPartnerContentCallbackUrl";
export * from "./listIntegrationAccountSchemaContentCallbackUrl";
export * from "./listWorkflowAccessKeySecretKeys";
export * from "./listWorkflowCallbackUrl";
export * from "./listWorkflowRunActionExpressionTraces";
export * from "./listWorkflowRunActionRepetitionExpressionTraces";
export * from "./listWorkflowTriggerCallbackUrl";
export * from "./listWorkflowVersionTriggerCallbackUrl";
export * from "./rosettaNetProcessConfiguration";
export * from "./workflow";
export * from "./workflowAccessKey";

// Export enums:
export * from "../types/enums/logic";

// Export sub-modules:
import * as v20150201preview from "./v20150201preview";
import * as v20150801preview from "./v20150801preview";
import * as v20160601 from "./v20160601";
import * as v20180701preview from "./v20180701preview";
import * as v20190501 from "./v20190501";

export {
    v20150201preview,
    v20150801preview,
    v20160601,
    v20180701preview,
    v20190501,
};

// Import resources to register:
import { IntegrationAccount } from "./integrationAccount";
import { IntegrationAccountAgreement } from "./integrationAccountAgreement";
import { IntegrationAccountAssembly } from "./integrationAccountAssembly";
import { IntegrationAccountBatchConfiguration } from "./integrationAccountBatchConfiguration";
import { IntegrationAccountCertificate } from "./integrationAccountCertificate";
import { IntegrationAccountMap } from "./integrationAccountMap";
import { IntegrationAccountPartner } from "./integrationAccountPartner";
import { IntegrationAccountSchema } from "./integrationAccountSchema";
import { IntegrationAccountSession } from "./integrationAccountSession";
import { IntegrationServiceEnvironment } from "./integrationServiceEnvironment";
import { IntegrationServiceEnvironmentManagedApi } from "./integrationServiceEnvironmentManagedApi";
import { RosettaNetProcessConfiguration } from "./rosettaNetProcessConfiguration";
import { Workflow } from "./workflow";
import { WorkflowAccessKey } from "./workflowAccessKey";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:logic:IntegrationAccount":
                return new IntegrationAccount(name, <any>undefined, { urn })
            case "azure-native:logic:IntegrationAccountAgreement":
                return new IntegrationAccountAgreement(name, <any>undefined, { urn })
            case "azure-native:logic:IntegrationAccountAssembly":
                return new IntegrationAccountAssembly(name, <any>undefined, { urn })
            case "azure-native:logic:IntegrationAccountBatchConfiguration":
                return new IntegrationAccountBatchConfiguration(name, <any>undefined, { urn })
            case "azure-native:logic:IntegrationAccountCertificate":
                return new IntegrationAccountCertificate(name, <any>undefined, { urn })
            case "azure-native:logic:IntegrationAccountMap":
                return new IntegrationAccountMap(name, <any>undefined, { urn })
            case "azure-native:logic:IntegrationAccountPartner":
                return new IntegrationAccountPartner(name, <any>undefined, { urn })
            case "azure-native:logic:IntegrationAccountSchema":
                return new IntegrationAccountSchema(name, <any>undefined, { urn })
            case "azure-native:logic:IntegrationAccountSession":
                return new IntegrationAccountSession(name, <any>undefined, { urn })
            case "azure-native:logic:IntegrationServiceEnvironment":
                return new IntegrationServiceEnvironment(name, <any>undefined, { urn })
            case "azure-native:logic:IntegrationServiceEnvironmentManagedApi":
                return new IntegrationServiceEnvironmentManagedApi(name, <any>undefined, { urn })
            case "azure-native:logic:RosettaNetProcessConfiguration":
                return new RosettaNetProcessConfiguration(name, <any>undefined, { urn })
            case "azure-native:logic:Workflow":
                return new Workflow(name, <any>undefined, { urn })
            case "azure-native:logic:WorkflowAccessKey":
                return new WorkflowAccessKey(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "logic", _module)
