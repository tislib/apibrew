import { PermissionChecks, AccessMap } from "./model.ts";

import { namespacePermissions } from "./helper.ts";

import { Namespace, Resource, SecurityConstraint } from "@apibrew/client";
import { BaseLogger } from "../../logging.ts";

const logger = BaseLogger.child({ component: 'AccessMapComputer' })

function mapConstraintTo(constraint: SecurityConstraint, permissionChecks?: PermissionChecks): boolean {
    if (!permissionChecks) {
        logger.warn('permissionChecks is undefined')
        return false
    }

    if (constraint.operation === 'FULL') {
        permissionChecks.full = true
    } else if (constraint.operation === 'READ') {
        permissionChecks.read = true
    } else if (constraint.operation === 'UPDATE') {
        permissionChecks.update = true
    } else if (constraint.operation === 'CREATE') {
        permissionChecks.create = true
    } else if (constraint.operation === 'DELETE') {
        permissionChecks.delete = true
    } else {
        throw new Error(`Unknown operation ${constraint.operation}`)
    }

    // if (constraint.property == 'owner' && constraint.propertyMode == 'PROPERTY_MATCH_ANY' && constraint.propertyValue == '$username') {
    //     permissionChecks.allowOwnedOnly = true
    // }

    return true
}

export function prepareAccessMap(accessMap: AccessMap, namespaces: Namespace[], resources: Resource[], constraints: SecurityConstraint[]) {
    let updatedAccessMap = { ...accessMap }

    logger.debug('prepareAccessMap', { namespaces, resources, constraints })

    if (!namespaces.some(item => item.name === 'system')) {
        namespaces.push({
            name: 'system',
        } as Namespace)
    }

    updatedAccessMap = {
        ...updatedAccessMap,
        "system": {
            full: false,
            read: false,
            create: false,
            update: false,
            delete: false
        },
    }

    for (const namespace of namespaces) {
        updatedAccessMap = {
            ...updatedAccessMap,
            [`namespace-${namespace.name}`]: {
                full: false,
                read: false,
                create: false,
                update: false,
                delete: false
            },
        }
    }
    for (const resource of resources) {
        updatedAccessMap = {
            ...updatedAccessMap,
            [`resource-${resource.namespace.name}/${resource.name}`]: {
                full: false,
                read: false,
                create: false,
                update: false,
                delete: false
            },
        }

        for (const property of resource.properties) {
            updatedAccessMap = {
                ...updatedAccessMap,
                [`resource-${resource.namespace.name}/${resource.name}-${property.name}`]: {
                    full: false,
                    read: false,
                    create: false,
                    update: false,
                    delete: false
                },
            }
        }
    }

    for (const constraint of constraints) {
        if (constraint.propertyMode || constraint.permit == 'REJECT' || constraint.propertyValue || (constraint.recordIds ?? []).length > 0) {
            logger.debug('Skipping constraint', { constraint })
            continue
        }

        if (!constraint.property) { // resource or namespace or system level
            if (!constraint.resource) { // namespace or system level
                if (!constraint.namespace) { // system level
                    if (mapConstraintTo(constraint, updatedAccessMap['system'])) {
                        constraint.localFlags = {
                            imported: true,
                        }
                    }
                } else { // namespace level
                    if (mapConstraintTo(constraint, updatedAccessMap[`namespace-${constraint.namespace.name}`])) {
                        constraint.localFlags = {
                            imported: true,
                        }
                    }
                }
            } else { // resource level
                if (mapConstraintTo(constraint, updatedAccessMap[`resource-${constraint.namespace?.name}/${constraint.resource.name}`])) {
                    constraint.localFlags = {
                        imported: true,
                    }
                }
            }
        } else { // property level constraint
            if (mapConstraintTo(constraint, updatedAccessMap[`resource-${constraint.namespace?.name}/${constraint.resource}-${constraint.property}`])) {
                constraint.localFlags = {
                    imported: true,
                }
            }
        }
    }

    return updatedAccessMap;
}

export function prepareConstraintsFromAccessMap(constraints: SecurityConstraint[], accessMap: AccessMap, namespaces: Namespace[], resources: Resource[]): SecurityConstraint[] {
    const updatedConstraints: SecurityConstraint[] = []

    const systemPermission = accessMap.system

    const systemConstraint = {
        permit: 'ALLOW',
        localFlags: {
            imported: true,
        }
    } as SecurityConstraint

    if (systemPermission.full) {
        updatedConstraints.push({
            ...systemConstraint,
            operation: 'FULL',
        })
    }
    if (systemPermission.read) {
        updatedConstraints.push({
            ...systemConstraint,
            operation: 'READ',
        })
    }
    if (systemPermission.create) {
        updatedConstraints.push({
            ...systemConstraint,
            operation: 'CREATE',
        })
    }
    if (systemPermission.update) {
        updatedConstraints.push({
            ...systemConstraint,
            operation: 'UPDATE',
        })
    }
    if (systemPermission.delete) {
        updatedConstraints.push({
            ...systemConstraint,
            operation: 'DELETE',
        })
    }

    for (const namespace of namespaces) {
        const namespacePermission = namespacePermissions(accessMap, namespace.name)

        const namespaceConstraint = {
            namespace: namespace,
            permit: 'ALLOW',
            localFlags: {
                imported: true,
            }
        } as SecurityConstraint

        if (namespacePermission.full) {
            updatedConstraints.push({
                ...namespaceConstraint,
                operation: 'FULL',
            })
        }
        if (namespacePermission.read) {
            updatedConstraints.push({
                ...namespaceConstraint,
                operation: 'READ',
            })
        }
        if (namespacePermission.create) {
            updatedConstraints.push({
                ...namespaceConstraint,
                operation: 'CREATE',
            })
        }
        if (namespacePermission.update) {
            updatedConstraints.push({
                ...namespaceConstraint,
                operation: 'UPDATE',
            })
        }
        if (namespacePermission.delete) {
            updatedConstraints.push({
                ...namespaceConstraint,
                operation: 'DELETE',
            })
        }
    }

    for (const resource of resources) {
        const resourcePermission = accessMap[`resource-${resource.namespace.name}/${resource.name}`]

        const resourceConstraint = {
            namespace: resource.namespace,
            resource: resource,
            permit: 'ALLOW',
            localFlags: {
                imported: true,
            }
        } as SecurityConstraint

        if (resourcePermission.full) {
            updatedConstraints.push({
                ...resourceConstraint,
                operation: 'FULL',
            })
        }
        if (resourcePermission.read) {
            updatedConstraints.push({
                ...resourceConstraint,
                operation: 'READ',
            })
        }
        if (resourcePermission.create) {
            updatedConstraints.push({
                ...resourceConstraint,
                operation: 'CREATE',
            })
        }
        if (resourcePermission.update) {
            updatedConstraints.push({
                ...resourceConstraint,
                operation: 'UPDATE',
            })
        }
        if (resourcePermission.delete) {
            updatedConstraints.push({
                ...resourceConstraint,
                operation: 'DELETE',
            })
        }

        for (const property of resource.properties) {
            const propertyPermission = accessMap[`resource-${resource.namespace.name}/${resource.name}-${property.name}`]

            const propertyConstraint = {
                namespace: resource.namespace,
                resource: resource,
                property: property.name,
                permit: 'ALLOW',
                localFlags: {
                    imported: true,
                }
            } as SecurityConstraint

            if (propertyPermission.full) {
                updatedConstraints.push({
                    ...propertyConstraint,
                    operation: 'FULL',
                })
            }
            if (propertyPermission.read) {
                updatedConstraints.push({
                    ...propertyConstraint,
                    operation: 'READ',
                })
            }
            if (propertyPermission.create) {
                updatedConstraints.push({
                    ...propertyConstraint,
                    operation: 'CREATE',
                })
            }
            if (propertyPermission.update) {
                updatedConstraints.push({
                    ...propertyConstraint,
                    operation: 'UPDATE',
                })
            }
            if (propertyPermission.delete) {
                updatedConstraints.push({
                    ...propertyConstraint,
                    operation: 'DELETE',
                })
            }
        }
    }

    // keep constraints which is not related to access map
    for (const constraint of constraints) {
        if ((constraint.localFlags as any)?.imported) {
            continue
        }

        updatedConstraints.push(constraint)
    }

    return updatedConstraints;
}
