import { GetConfig } from '../../../bindings/window-service-watcher/internal/app/app';

export type ServiceStatus = 'running' | 'stopped' | 'error' | 'starting' | 'stopping';

export interface ServiceData {
	id: string;
	name: string;
	description: string;
	status: ServiceStatus;
	logsUrl: string | null;
}

export class Service {
	id: string;
	name: string;
	description: string;
	logsUrl: string | null;

	status = $state<ServiceStatus>('stopped');

	constructor(data: ServiceData) {
		this.id = data.id;
		this.name = data.name;
		this.description = data.description;
		this.status = data.status;
		this.logsUrl = data.logsUrl;
	}

	async get() {
		const config = await GetConfig();
		console.log('Config', config);
	}

	async start() {
		if (['running', 'starting'].includes(this.status)) return;
		this.status = 'starting';
		await new Promise((r) => setTimeout(r, 1500));
		this.status = 'running';
	}

	async stop() {
		if (['stopped', 'stopping'].includes(this.status)) return;

		this.status = 'stopping';
		await new Promise((r) => setTimeout(r, 1500));
		this.status = 'stopped';
	}

	async restart() {
		await this.stop();
		await this.start();
	}
}

export class ServiceStore {
	services = $state<Service[]>([]);

	healthyCount = $derived(() => this.services.filter((s) => s.status === 'running').length);
	totalCount = $derived(() => this.services.length);

	constructor() {
		// this.services = this.#getMockData().map((data) => new Service(data));
	}

	async init() {
		await GetConfig().then((configs) => {
			configs.Services.forEach((config) => {
				this.services.push(new Service({
					id: config.id,
					name: config.name || 'Unknown',
					description:  config.description ||'No description provided.',
					status: 'stopped',
					logsUrl: config.log_path ? `file://${config.log_path}` : null
				}));
			})
		}).catch((err) => {
			console.error('Error loading config in ServiceStore:', err);
		});
	}

	#getMockData(): ServiceData[] {
		return [
			{
				id: '1',
				name: 'Auth Service',
				description: 'Auth provider',
				status: 'running',
				logsUrl: 'http://localhost:8080/logs'
			},
			{
				id: '2',
				name: 'Payment Gateway',
				description: 'Payment processing',
				status: 'running',
				logsUrl: null
			},
			{
				id: '3',
				name: 'Email Worker',
				description: 'Email jobs',
				status: 'stopped',
				logsUrl: 'http://localhost:8082/logs'
			},
			{
				id: '4',
				name: 'Data Aggregator',
				description: 'Data stats',
				status: 'error',
				logsUrl: 'http://localhost:8083/logs'
			},
			{
				id: '5',
				name: 'Notification Svc',
				description: 'Push notifications',
				status: 'running',
				logsUrl: null
			},
			{
				id: '6',
				name: 'Legacy API',
				description: 'Old PHP backend',
				status: 'stopped',
				logsUrl: null
			}
		];
	}
}

export const serviceStore = new ServiceStore();
