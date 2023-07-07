import { Function, FunctionName } from "./model/function";
import { FunctionTrigger, FunctionTriggerName } from "./model/function-trigger";
import { ResourceRule, ResourceRuleName } from "./model/resource-rule";
import { ENGINE_REMOTE_ADDR, EXTENSION_NAME, FN_DIR } from "./config";
import { registerExtensions } from "./registrator";
import { Extension } from "./model";
import { components } from './model/base-schema'
import { Module, ModuleName } from "./model/module";
import { scriptFunctionTemplate, moduleFunctionTemplate, moduleInitTemplate } from "./function-template";
import * as fs from 'fs'
import path from "path";
import { PassThrough } from "stream";
import { Extract, extract } from 'tar-fs'
import { initModule } from "./module-init";
import { Lambda, LambdaNameName, LambdaResource } from "./model/lambda";
import { functionRepository, lambdaRepository, moduleRepository } from "./client";
import { FunctionExecution } from "@apibrew/client";
var mkdirp = require('mkdirp');


let engineId: string

export let functionMap: { [key: string]: Function } = {}
export let functionIdMap: { [key: string]: Function } = {}
export let functionNameIdMap: { [key: string]: string } = {}
export let moduleIdMap: { [key: string]: Module } = {}
export let lambdaIdMap: { [key: string]: Lambda } = {}

export function clearFnCache(id: string) {
    for (const cacheKey of Object.keys(require.cache)) {
        if (cacheKey.startsWith(`${FN_DIR}/${id}`)) {
            delete require.cache[cacheKey]
        }
    }
}

export async function registerFunction(fn: Function) {
    functionMap[fn.package + '/' + fn.name] = fn
    functionIdMap[fn.id] = fn
    functionNameIdMap[fn.package + '/' + fn.name] = fn.id

    await storeFunction(fn)

    clearFnCache(fn.id)
}

export async function registerModule(module: Module) {
    console.log('Store Module: ', module.id)
    await storeModule(module)
    console.log('Init Module: ', module.id)

    try {
        await initModule(module)
        console.log('Done Module: ', module.id)
        moduleIdMap[module.id] = module
    } catch (e) {
        console.log('Error Module: ', module.id)
        console.log(e)
    }

    clearFnCache(module.id)
}

export async function registerLambda(lambda: Lambda) {
    const extension = prepareExtensionFromLambda(lambda)

    await registerExtensions([extension])
    
    lambdaIdMap[lambda.id] = lambda
}

export async function registerFunctionTrigger(trigger: FunctionTrigger) {

}

export async function registerResourceRule(rule: ResourceRule) {

}

export async function loadAll() {
    console.log('begin load all')

    const modules$ = await moduleRepository.list()
    const modules = modules$.content.filter(module => module.engine.id === engineId)

    for (const module of modules) {
        await registerModule(module)
    }

    console.log('load functions')
    const functions$ = await functionRepository.list()
    const functions = functions$.content.filter(fn => fn.engine.id === engineId)

    console.log('registering functions')
    for (const fn of functions) {
        await registerFunction(fn)
    }

    const lambdas$ = await lambdaRepository.list()

    const lambdas = lambdas$.content.filter(lambda => functionIdMap[lambda.function.id])

    for (const lambda of lambdas) {
        await registerLambda(lambda)
    }

    // load functions
    // load modules
    // load lambdas
    // load triggers
    // load rules
    // register extensions

    // reloadFunction()
    // reloadModules()
    // reloadLambdas()

    // const functions = read<Function>('logic', FunctionName)

    // filter('logic', FunctionTriggerName, (record: FunctionTrigger) => functions.some(fn => fn.id === record.function.id))
    // const triggers = read<FunctionTrigger>('logic', FunctionTriggerName)

    // filter('logic', ResourceRuleName, (record: ResourceRule) => functions.some(fn => fn.id === record.conditionFunction.id))
    // const rules = read<ResourceRule>('logic', ResourceRuleName)

    // for (const cacheKey of Object.keys(require.cache)) {
    //     if (cacheKey.startsWith(FN_DIR)) {
    //         delete require.cache[cacheKey]
    //     }
    // }

    // let extensions: Extension[] = []

    // for (const trigger of triggers) {
    //     extensions.push(prepareExtensionFromTrigger(trigger))
    // }

    // for (const rule of rules) {
    //     extensions.push(prepareExtensionFromRule(rule))
    // }

    // extensions = extensions.filter((item, index) => extensions.findIndex(item2 => JSON.stringify(item) === JSON.stringify(item2)) === index)

    // await registerExtensions(extensions)

    // console.log('Configuring extensions: ', extensions.map(item => item.name))

    console.log('end load all')
}

function prepareExtensionFromTrigger(trigger: FunctionTrigger): Extension {
    const extension = {} as Extension
    extension.name = `${EXTENSION_NAME}_trigger_${trigger.namespace}_${trigger.resource}`
    extension.sync = !trigger.async
    const call = {} as components['schemas']['ExternalCall']
    const hCall = {} as components['schemas']['HttpCall']
    hCall.method = 'POST'
    hCall.uri = `${ENGINE_REMOTE_ADDR}/call/trigger`
    call.httpCall = hCall
    extension.call = call
    extension.responds = true

    if (trigger.order) {
        switch (trigger.order) {
            case 'before':
                extension.order = 10

                break
            case 'after':
                extension.order = 200
                break
            case 'instead':
                extension.order = 80
                extension.finalizes = true
                break
        }
    }

    const eventSelector = {} as components['schemas']['EventSelector']
    eventSelector.namespaces = [trigger.namespace]
    eventSelector.resources = [trigger.resource]
    let action: "CREATE" | "UPDATE" | "DELETE" | "GET" | "LIST" | "OPERATE"
    switch (trigger.action) {
        case 'create':
            action = 'CREATE'
            break
        case 'update':
            action = 'UPDATE'
            break
        case 'delete':
            action = 'DELETE'
            break
        default:
            throw new Error('Unknown action: ' + trigger.action)
    }
    eventSelector.actions = [action]
    extension.selector = eventSelector

    return extension
}

function prepareExtensionFromLambda(lambda: Lambda): Extension {
    const parts = lambda.eventSelectorPattern.split(':')

    const resourceFullName = parts[0]
    // const action = parts[1]
    const subParts = resourceFullName.split('/')
    let resourceName;
    let resourceNamespace;

    if (subParts.length === 1) {
        resourceName = subParts[0]
        resourceNamespace = 'default'
    } else if (subParts.length === 2) {
        resourceNamespace = subParts[0]
        resourceName = subParts[1]
    } else {
        throw new Error('Invalid resource name: ' + resourceFullName)
    }

    const extension = {} as Extension
    extension.name = `${EXTENSION_NAME}_lambda_${lambda.id}`
    extension.sync = true
    const call = {} as components['schemas']['ExternalCall']
    const hCall = {} as components['schemas']['HttpCall']
    hCall.method = 'POST'
    hCall.uri = `${ENGINE_REMOTE_ADDR}/call/lambda/${lambda.id}`
    call.httpCall = hCall
    extension.call = call
    extension.order = 85
    extension.finalizes = true
    extension.responds = true

    const eventSelector = {} as components['schemas']['EventSelector']
    eventSelector.namespaces = [resourceNamespace]
    eventSelector.resources = [resourceName]
    eventSelector.actions = ['CREATE']
    extension.selector = eventSelector

    return extension
}

function prepareExtensionFromRule(rule: ResourceRule): Extension {
    const extension = {} as Extension
    extension.name = `${EXTENSION_NAME}_rule_${rule.namespace}_${rule.resource}`
    extension.sync = true
    const call = {} as components['schemas']['ExternalCall']
    const hCall = {} as components['schemas']['HttpCall']
    hCall.method = 'POST'
    hCall.uri = `${ENGINE_REMOTE_ADDR}/call/rule`
    call.httpCall = hCall
    extension.call = call
    extension.order = 85

    const eventSelector = {} as components['schemas']['EventSelector']
    eventSelector.namespaces = [rule.namespace]
    eventSelector.resources = [rule.resource]
    eventSelector.actions = ['CREATE', 'UPDATE']
    extension.selector = eventSelector

    return extension
}

export async function initFunctionRegistry(_engineId: string) {
    engineId = _engineId
}

async function storeFunction(record: Function) {
    if (record.script) {
        console.log(record)
        const functionContent = scriptFunctionTemplate(record)

        fs.writeFileSync(path.join(FN_DIR + '/', record.id + '.js'), functionContent)
    } else if (record.module) {
        record.module = moduleIdMap[record.module.id]
        record.module.content = undefined
        const functionContent = moduleFunctionTemplate(record)

        fs.writeFileSync(path.join(FN_DIR + '/', record.id + '.js'), functionContent)
    }
}

async function storeModule(record: Module) {
    const content = Buffer.from(record.content, 'base64')

    const stream = new PassThrough()

    stream.end(content)

    const destination = path.join(FN_DIR + '/', record.id)

    fs.rmSync(destination, { recursive: true, force: true });

    mkdirp.sync(destination)

    await new Promise((resolve: any, reject) => {
        stream.pipe(extract(destination)).on('error', reject).on('close', resolve)
    })
}
