import {Button} from "@mui/material"
import {PageLayout} from "../../layout/PageLayout"
import {Api, Delete, Edit, PlusOneOutlined, Search} from "@mui/icons-material"
import {Resource, ResourceProperty} from "../../model"
import {useNavigate} from "react-router-dom"
import {DataGrid, GridActionsCellItem, GridColDef, GridRowParams, GridValueGetterParams} from '@mui/x-data-grid';
import {Record, RecordService} from "@apibrew/core-lib"
import React, {useContext, useEffect, useState} from "react"
// import {SdkDrawer} from "../sdk/SdkDrawer"
import {Crud} from "../../model/ui/crud.ts";
import {Icon} from "../Icon.tsx";
import {ModuleService} from "@apibrew/core-lib";
import {LayoutContext, useBreadCramps} from "../../context/layout-context.ts";
import {useErrorHandler} from "../../hooks/error-handler.tsx";
import Box from "@mui/material/Box";

export interface ListProps {
    resource: Resource
    crudConfig: Crud
}

export function List(props: ListProps) {
    const navigate = useNavigate()
    const [list, setList] = useState<Record[]>([])
    const [showSdk, setShowSdk] = useState(false)
    const errorHandler = useErrorHandler()
    const layoutContext = useContext(LayoutContext)

    const load = () => {
        RecordService.list<Record>(props.resource.namespace ?? 'default', props.resource.name).then((data) => {
            setList(data)
        }, errorHandler)
    }

    useBreadCramps()

    const resourcePropertyMap = new Map<string, ResourceProperty>()

    props.resource.properties.forEach((property) => {
        resourcePropertyMap.set(property.name, property)
    })

    useEffect(() => {
        load();
    }, [props.resource])

    let columns: GridColDef[] = props.resource.properties.map((property) => {
        return {
            field: property.name,
            type: 'string',
            headerName: property.name,
            width: 150,
            editable: false,
            valueGetter: (params: GridValueGetterParams<any, any>) => {
                const prop = props.resource.properties.find((p) => p.name === property.name)!

                if (prop.type === 'REFERENCE' && params.row[property.name]) {
                    return params.row[property.name].name
                }

                return params.row[property.name]
            }
        } as GridColDef
    })

    if (props.crudConfig?.gridConfig?.columns) {
        columns = props.crudConfig.gridConfig.columns.filter(item => !item.disabled).map((column) => {
            return {
                field: column.name,
                type: column.type ?? 'string',
                headerName: column.title,
                width: (((column.width ?? 0) > 0) ? column.width : 150),
                flex: column.flex,
                sortable: column.sortable,
                filterable: column.filterable,
                editable: false,
                valueGetter: (params: GridValueGetterParams<any, any>) => {
                    const prop = resourcePropertyMap.get(column.name)

                    if (prop === undefined) {
                        return 'Unknown column!'
                    }

                    if (prop.type === 'REFERENCE' && params.row[column.name]) {
                        return params.row[column.name].name
                    }

                    return params.row[column.name]
                }
            } as GridColDef
        })
    }

    const actionColumn: GridColDef = {
        field: 'actions',
        type: 'actions',
        maxWidth: 200,
        minWidth: 200,
        align: 'right',
        headerName: 'Actions',
        getActions: (params: GridRowParams) => {
            const actions = [
                <GridActionsCellItem label='Edit' icon={<Edit/>} onClick={() => {
                    navigate(`${params.id}/edit`)
                }}/>,
                <GridActionsCellItem label='Delete' icon={<Delete/>} onClick={() => {
                    RecordService.remove(props.resource.namespace ?? 'default', props.resource.name, params.id as string).then(() => {
                        load()
                    })
                }}/>,
                <GridActionsCellItem label='Delete' icon={<Search/>} onClick={() => {
                    navigate(`${params.id}/view`)
                }}/>,
            ]

            if (props.crudConfig?.gridConfig?.disableDefaultActions) {
                actions.length = 0
            }

            if (props.crudConfig?.gridConfig?.actions) {
                props.crudConfig.gridConfig.actions.forEach((action) => {
                    actions.push(<GridActionsCellItem
                        label={action.title}
                        icon={<Icon name={action.icon}/>}
                        onClick={() => {
                            ModuleService.executeActionComponent(action.component, params.id as string, layoutContext).then()
                        }}/>)
                })
            }

            return actions
        }
    }

    if (!columns.some(item => item.flex)) {
        actionColumn.maxWidth = undefined
        actionColumn.flex = 1
    }

    columns.push(actionColumn)


    return (
        <PageLayout>
            <Box sx={{display: 'flex', paddingBottom: '10px', width: '100%'}}>
                <Button variant={'contained'} color='success' onClick={() => {
                    navigate('new')
                }} startIcon={<PlusOneOutlined/>}>New {props.resource.name}</Button>
                <Box flexGrow={1}/>
                <Button variant={'contained'} color='primary' onClick={() => {
                    setShowSdk(true)
                }} startIcon={<Api/>}>sdk</Button>
                <Button variant={'contained'} color='secondary' onClick={() => {
                    navigate('settings')
                }} startIcon={<Api/>}>Crud Settings</Button>
            </Box>

            {/*<SdkDrawer resource={props.resource} open={showSdk} onClose={() => {*/}
            {/*    setShowSdk(false)*/}
            {/*}}/>*/}
            <DataGrid
                rows={list}
                columns={columns}
                density='compact'
                initialState={{
                    pagination: {
                        paginationModel: {
                            pageSize: 100,
                        },
                    },
                }}
                disableRowSelectionOnClick
                autoHeight={true}
                // rowHeight={35}
                disableColumnMenu={true}
            />

        </PageLayout>
    )
}
