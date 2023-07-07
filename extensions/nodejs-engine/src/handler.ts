import { functionIdMap, functionMap, lambdaIdMap, registerFunction, registerFunctionTrigger, registerLambda, registerModule } from "./function-registry";
import { ResourceRule, ResourceRuleName } from "./model/resource-rule";
import { FunctionTrigger, FunctionTriggerName } from "./model/function-trigger";
import { components } from "./model/base-schema";
import { executeFunction, executeLambda, locateFunction } from "./function-execute";
import { Lambda, LambdaResource } from "./model/lambda";
import { functionRepository, functionTriggerRepository, lambdaRepository, moduleRepository } from "./client";

type Event = components['schemas']['Event']

const { VM } = require('vm2');

export async function handleFunctionExecutionCall(event: Event) {
    for (const record of event.records) {
        try {
            const packageName = record.properties.function.package
            const name = record.properties.function.name
            const input = record.properties.input
            const fn = locateFunction(packageName, name)

            if (!fn) {
                throw new Error('Function not found')
            }

            const output = (await executeFunction(fn, input) ?? { ok: true });

            record.properties.output = output

            record.properties.status = 'success'
        } catch (e) {
            console.error(e)

            record.properties.error = e.message
            record.properties.status = 'error'
        }
    }

    return event;
}

export async function handleFunctionCall(event: Event) {
    console.log('trigger function', (event.records ?? [{}])[0]['id'])

    return await handleFunctionExecutionCall(event)
}

export async function handleLambdaCall(event: Event, lambdaId: string) {
    console.log('trigger lambda', lambdaId)

    const lambda = lambdaIdMap[lambdaId]

    if (!lambda) {
        throw new Error('Lambda not found')
        return
    }

    for (const record of event.records) {
        try {
            await executeLambda(lambda, record.properties)
        } catch (e) {
            console.error(e)
        }
    }

}

export async function handleReload(event: Event) {
    console.log('trigger reload', event.resource.namespace, event.resource.name, (event.records ?? [{}])[0]['id'])

    for (const record of event.records ?? []) {
        switch (`${event.resource.namespace}/${event.resource.name}`) {
            case 'logic/Function':
                const fn = await functionRepository.get(record.id!)
                await registerFunction(fn)

                console.log('reloaded function', fn.id)
                break
            case 'logic/Module':
                const module = await moduleRepository.get(record.id!)
                await registerModule(module)

                console.log('reloaded module', module.id)
                break
            case 'logic/Lambda':
                const lambda = await lambdaRepository.get(record.id!)
                await registerLambda(lambda)

                console.log('reloaded lambda', lambda.id)
                break
            case 'logic/FunctionTrigger':
                const functionTrigger = await functionTriggerRepository.get(record.id!)
                await registerFunctionTrigger(functionTrigger)

                console.log('reloaded functionTrigger', functionTrigger.id)
                break
        }
    }
}