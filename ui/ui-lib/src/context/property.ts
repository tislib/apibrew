import { ResourceProperty } from "@apibrew/client";
import React from "react";

export const ResourcePropertyContext = React.createContext<ResourceProperty | undefined>(undefined)

export function useResourceProperty(required: boolean) {
    const ResourceProperty = React.useContext(ResourcePropertyContext)

    if (required && !ResourceProperty) {
        throw new Error("useResourceProperty must be used within a ResourcePropertyProvider")
    }
    return ResourceProperty
}