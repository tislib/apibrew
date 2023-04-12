import {components} from "./schema";
import axios from "axios";


/////// #### abs #### //////

export interface Entity<T> {
    id?: string
}

interface Repository<T extends Entity<T>> {
    create(entity: T): Promise<T>;

    update(entity: T): Promise<T>;

    apply(entity: T): Promise<T>;

    get(id: string): Promise<T>;

    find(params: FindParams): Promise<{ total: number, content: { properties: T }[] }>;

    extend(extensionService: ExtensionService): RepositoryExtension<T>;
}

interface BooleanExpression {

}

interface FindParams {
    limit?: number;
    offset?: number;
    useHistory?: boolean;
    annotations?: { [key: string]: string };
    resolveReferences?: string[]; // default ["*"]
    query?: BooleanExpression | null;
}

export interface RepositoryExtension<T extends Entity<T>> {
    onCreate(handler: (elem: T) => Promise<T>): void;

    onUpdate(handler: (elem: T) => Promise<T>): void;

    onDelete(handler: (elem: T) => Promise<T>): void;

    onGet(handler: (id: string) => Promise<T>): void;

    onList(handler: () => Promise<{ properties: T }[]>): void;
}


/////// #### apply #### //////

export interface DhClientParams {
    Addr: string;
    Insecure?: boolean;
    token?: string;
}

export class DhClient {
    params: DhClientParams;

    constructor(params: DhClientParams) {
        this.params = params
    }

    public async authenticateWithUsernameAndPassword(username: string, password: string): Promise<void> {
        const authRequest: components["schemas"]["AuthenticationRequest"] = {
            username: username,
            password: password,
            term: "LONG"
        }

        const result = await axios.post<components["schemas"]["AuthenticationResponse"]>(`http://${this.params.Addr}/authentication/token`, authRequest);

        this.params.token = result.data.token!.content;
    }

    public newRepository<T extends Entity<T>>(namespace: string, resource: string): Repository<T> {
        return new RepositoryImpl<T>(this, {
            namespace: namespace,
            resource: resource,
            updateCheckVersion: false,
        });
    }

    public NewExtensionService(host: string, port: number): ExtensionService {
        return new ExtensionServiceImpl(host, port, host + ':' + port, this);
    }
}

type ExternalFunctionData = { [key: string]: any }

type ExternalFunction = (req: ExternalFunctionData) => Promise<ExternalFunctionData>;

interface FunctionCallRequest {
    name: string;
    request: ExternalFunctionData
}

interface FunctionCallResponse {
    response: ExternalFunctionData
}

interface ExtensionService {
    run(): Promise<void>;

    registerFunction(name: string, handler: ExternalFunction): void;

    getRemoteHost(): string;
}

class ExtensionServiceImpl implements ExtensionService {
    private host: string;
    private port: number;
    private remoteHost: string;
    private client: DhClient;
    private functions: { [key: string]: ExternalFunction };

    constructor(host: string, port: number, remoteHost: string, client: DhClient) {
        this.host = host;
        this.port = port
        this.remoteHost = remoteHost;
        this.client = client;
        this.functions = {};
    }

    getRemoteHost(): string {
        return this.remoteHost;
    }

    registerFunction(name: string, handler: ExternalFunction): void {
        this.functions[name] = handler;
    }

    async run(): Promise<void> {
        const express = require('express')
        const app = express()

        app.use(express.json())

        app.get('/', (req: any, res: any) => {
            res.send('ok')
        })

        app.post('/:name', async (req: any, res: any) => {
            const name = req.params.name

            const request: FunctionCallRequest = {
                name: name,
                request: req.body.content
            }

            const response = await this.functions[name](request.request)

            res.send({
                content: response
            })
        })

        console.log('starting extension service')
        app.listen(this.port, this.host, () => {
            console.log(`External service is listening on ${this.host}`)
        })
    }
}

interface RepositoryParams<T extends Entity<T>> {
    namespace: string,
    resource: string,
    updateCheckVersion: boolean;
}

export class RepositoryImpl<T extends Entity<T>> implements Repository<T> {
    private readonly client: DhClient;
    private readonly params: RepositoryParams<T>;

    constructor(client: DhClient, params: RepositoryParams<T>) {
        this.client = client;
        this.params = params;
    }

    async create(entity: T): Promise<T> {
        const result = await axios.post<T>(`http://${this.client.params.Addr}/records/${this.params.namespace}/${this.params.resource}`, entity, {
            headers: {
                Authorization: `Bearer ${this.client.params.token}`
            }
        });

        return result.data;
    }

    async update(entity: T): Promise<T> {
        const result = await axios.put<T>(`http://${this.client.params.Addr}/records/${this.params.namespace}/${this.params.resource}/${entity.id}`, entity, {
            headers: {
                Authorization: `Bearer ${this.client.params.token}`
            }
        });

        return result.data;
    }

    async get(id: string): Promise<T> {
        const result = await axios.get<T>(`http://${this.client.params.Addr}/records/${this.params.namespace}/${this.params.resource}/${id}`, {
            headers: {
                Authorization: `Bearer ${this.client.params.token}`
            }
        });

        return result.data;
    }

    public async apply<T>(entity: T): Promise<T> {
        const result = await axios.patch<T>(`http://${this.client.params.Addr}/${this.params.namespace}/${this.params.resource}`, entity, {
            headers: {
                Authorization: `Bearer ${this.client.params.token}`
            }
        });

        return result.data;
    }

    async find(params: FindParams): Promise<{ total: number, content: { properties: T }[] }> {
        if (!params.resolveReferences) {
            params.resolveReferences = ["*"];
        }

        const result = await axios.get<{ total: number, content: { properties: T }[] }>(`http://${this.client.params.Addr}/records/${this.params.namespace}/${this.params.resource}`, {
            headers: {
                Authorization: `Bearer ${this.client.params.token}`
            }
        });

        return result.data;
    }

    extend(extensionService: ExtensionService): RepositoryExtension<T> {
        return new RepositoryExtensionImpl<T>(this, extensionService, this.params.resource, this.params.namespace, this.client);
    }
}


// ## repository extension

export class RepositoryExtensionImpl<T extends Entity<T>> implements RepositoryExtension<T> {
    private repository: Repository<T>;
    private extension: ExtensionService;
    private resourceName: string;
    private namespace: string;
    private client: DhClient;
    private extensionRepository: RepositoryImpl<components["schemas"]["Extension"]>;

    constructor(
        repository: Repository<T>,
        extension: ExtensionService,
        resourceName: string,
        namespace: string,
        client: DhClient
    ) {
        this.repository = repository;
        this.extension = extension;
        this.resourceName = resourceName;
        this.namespace = namespace;
        this.client = client;
        this.extensionRepository = new RepositoryImpl<components["schemas"]["Extension"]>(client, {
            namespace: "system", resource: "extension", updateCheckVersion: false
        })
    }

    async onCreate(handler: (elem: T) => Promise<T>): Promise<void> {
        const extensionName = this.getExtensionName("OnCreate");

        this.extension.registerFunction(extensionName, async function (data: ExternalFunctionData) {
            const records = []

            console.log('data', data)

            for (const record of data.request.records) {
                const entity = await handler(record.properties)
                records.push({
                    properties: entity
                })
            }

            const response: ExternalFunctionData = {
                "response": {
                    '@type': 'type.googleapis.com/stub.CreateRecordResponse',
                    "records": records
                }
            }

            console.log("response", response)

            return response
        });

        const ext = {
            name: extensionName,
            namespace: this.namespace,
            resource: this.resourceName,
            instead: {
                create: {
                    kind: "httpCall",
                    uri: `http://${this.extension.getRemoteHost()}/${extensionName}`,
                    method: 'POST',
                },
            },
        };

        await this.extensionRepository.apply(ext)
    }

    async onUpdate(handler: (elem: T) => Promise<T>): Promise<void> {
        //TODO implement me
        throw new Error("Method not implemented.");
    }

    async onDelete(handler: (elem: T) => Promise<T>): Promise<void> {
        //TODO implement me
        throw new Error("Method not implemented.");
    }

    async onGet(handler: (id: string) => Promise<T>): Promise<void> {
        //TODO implement me
        throw new Error("Method not implemented.");
    }

    async onList(handler: () => Promise<{ properties: T }[]>): Promise<void> {
        //TODO implement me
        throw new Error("Method not implemented.");
    }

    private getExtensionName(action: string): string {
        return `${this.namespace}-${this.resourceName}-${action}`;
    }
}