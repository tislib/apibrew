import { Namespace } from "./namespace";
import { DataSource } from "./data-source";
export declare const ResourceResource: {
    resource: string;
    namespace: string;
};
export interface Resource {
    id: string;
    version: number;
    createdBy: string;
    updatedBy?: string;
    createdOn: string;
    updatedOn?: string;
    name: string;
    namespace: Namespace;
    virtual: boolean;
    types?: object;
    immutable: boolean;
    abstract: boolean;
    dataSource?: DataSource;
    entity?: string;
    catalog?: string;
    annotations?: object;
    indexes?: object;
    securityConstraints?: object[];
    title?: string;
    description?: string;
}
export declare const ResourceName = "Resource";
export declare const ResourceIdName = "Id";
export declare const ResourceVersionName = "Version";
export declare const ResourceCreatedByName = "CreatedBy";
export declare const ResourceUpdatedByName = "UpdatedBy";
export declare const ResourceCreatedOnName = "CreatedOn";
export declare const ResourceUpdatedOnName = "UpdatedOn";
export declare const ResourceNameName = "Name";
export declare const ResourceNamespaceName = "Namespace";
export declare const ResourceVirtualName = "Virtual";
export declare const ResourceTypesName = "Types";
export declare const ResourceImmutableName = "Immutable";
export declare const ResourceAbstractName = "Abstract";
export declare const ResourceDataSourceName = "DataSource";
export declare const ResourceEntityName = "Entity";
export declare const ResourceCatalogName = "Catalog";
export declare const ResourceAnnotationsName = "Annotations";
export declare const ResourceIndexesName = "Indexes";
export declare const ResourceSecurityConstraintsName = "SecurityConstraints";
export declare const ResourceTitleName = "Title";
export declare const ResourceDescriptionName = "Description";