import {Role} from './role';
import {Permission} from './permission';

export interface User {
    id: string
    version: number
    auditData?: AuditData
    username: string
    password?: string
    roles?: Role[]
    permissions?: Permission[]
    details?: object
}

export const UserEntityInfo = {
    namespace: "system",
    resource: "User",
    restPath: "system-user",
}

export interface AuditData {
    createdBy: string
    updatedBy: string
    createdOn: string | Date
    updatedOn: string | Date
}

export const UserResource = {
  "name": "User",
  "namespace": {
    "name": "system"
  },
  "properties": [
    {
      "name": "id",
      "type": "UUID",
      "required": true,
      "immutable": true,
      "exampleValue": "a39621a4-6d48-11ee-b962-0242ac120002",
      "description": "The unique identifier of the resource. It is randomly generated and immutable.",
      "annotations": {
        "PrimaryProperty": "true",
        "SpecialProperty": "true"
      }
    },
    {
      "name": "version",
      "type": "INT32",
      "required": true,
      "defaultValue": 1,
      "exampleValue": 1,
      "title": "Version",
      "description": "The version of the resource/record. It is incremented on every update.",
      "annotations": {
        "AllowEmptyPrimitive": "true",
        "SpecialProperty": "true"
      }
    },
    {
      "name": "auditData",
      "type": "STRUCT",
      "typeRef": "AuditData",
      "exampleValue": {
        "createdBy": "admin",
        "createdOn": "2023-11-04T03:41:34+04:00",
        "updatedBy": "admin",
        "updatedOn": "2023-11-04T03:41:34+04:00"
      },
      "title": "Audit Data",
      "description": "The audit data of the resource/record. \nIt contains information about who created the resource/record, when it was created, who last updated the resource/record and when it was last updated.",
      "annotations": {
        "SpecialProperty": "true"
      }
    },
    {
      "name": "username",
      "type": "STRING",
      "required": true,
      "unique": true,
      "length": 256,
      "title": "Username",
      "annotations": {
        "IsHclLabel": "true"
      }
    },
    {
      "name": "password",
      "type": "STRING",
      "length": 256,
      "title": "Password"
    },
    {
      "name": "roles",
      "type": "LIST",
      "item": {
        "name": "roles",
        "type": "REFERENCE",
        "reference": {
          "resource": {
            "name": "Role",
            "namespace": {
              "name": "system"
            }
          },
          "cascade": false
        }
      },
      "title": "Roles"
    },
    {
      "name": "permissions",
      "type": "LIST",
      "item": {
        "name": "permissions",
        "type": "REFERENCE",
        "reference": {
          "resource": {
            "name": "Permission",
            "namespace": {
              "name": "system"
            }
          },
          "cascade": false,
          "backReference": "user"
        }
      },
      "title": "Permissions",
      "description": "The permissions of the user. It is used to define the access control rules for resources for users. When you set permissions it is automatically created though Permission Resource. No need to manage it manually"
    },
    {
      "name": "details",
      "type": "OBJECT",
      "title": "Details",
      "description": "The details of the user. It is used to store additional information about the user."
    }
  ],
  "types": [
    {
      "name": "AuditData",
      "title": "Audit Data",
      "description": "Audit Data is a type that represents the audit data of a resource/record. ",
      "properties": [
        {
          "name": "createdBy",
          "type": "STRING",
          "immutable": true,
          "length": 256,
          "exampleValue": "admin",
          "title": "Created By",
          "description": "The user who created the resource/record.",
          "annotations": {
            "SpecialProperty": "true"
          }
        },
        {
          "name": "updatedBy",
          "type": "STRING",
          "length": 256,
          "exampleValue": "admin",
          "title": "Updated By",
          "description": "The user who last updated the resource/record.",
          "annotations": {
            "SpecialProperty": "true"
          }
        },
        {
          "name": "createdOn",
          "type": "TIMESTAMP",
          "immutable": true,
          "exampleValue": "2023-11-04T03:41:34+04:00",
          "title": "Created On",
          "description": "The timestamp when the resource/record was created.",
          "annotations": {
            "SpecialProperty": "true"
          }
        },
        {
          "name": "updatedOn",
          "type": "TIMESTAMP",
          "exampleValue": "2023-11-04T03:41:34+04:00",
          "title": "Updated On",
          "description": "The timestamp when the resource/record was last updated.",
          "annotations": {
            "SpecialProperty": "true"
          }
        }
      ]
    }
  ],
  "title": "User",
  "description": "User is a resource that defines the access control model. It is used to authenticate and authorize users.",
  "annotations": {
    "EnableAudit": "true",
    "OpenApiGroup": "meta"
  }
} as unknown

