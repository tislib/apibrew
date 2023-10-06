

import { Function } from "./function";


export const FunctionTriggerResource = {
    resource: "FunctionTrigger",
    namespace: "logic",
};

// Sub Types

// Resource Type
export interface FunctionTrigger {
    id: string;
name: string;
resource: string;
namespace: string;
action: 'create' | 'update' | 'delete' | 'list' | 'get';
order?: 'before' | 'after' | 'instead';
async: boolean;
function: Function;
annotations?: object;
createdBy?: string;
updatedBy?: string;
createdOn?: string;
updatedOn?: string;
version: number;

}
// Resource and Property Names
export const FunctionTriggerName = "FunctionTrigger";

export const FunctionTriggerIdName = "Id";

export const FunctionTriggerNameName = "Name";

export const FunctionTriggerResourceName = "Resource";

export const FunctionTriggerNamespaceName = "Namespace";

export const FunctionTriggerActionName = "Action";

export const FunctionTriggerOrderName = "Order";

export const FunctionTriggerAsyncName = "Async";

export const FunctionTriggerFunctionName = "Function";

export const FunctionTriggerAnnotationsName = "Annotations";

export const FunctionTriggerCreatedByName = "CreatedBy";

export const FunctionTriggerUpdatedByName = "UpdatedBy";

export const FunctionTriggerCreatedOnName = "CreatedOn";

export const FunctionTriggerUpdatedOnName = "UpdatedOn";

export const FunctionTriggerVersionName = "Version";


