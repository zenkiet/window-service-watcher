import {
	GetConfig,
	StartService,
	StopService
} from '../../../bindings/window-service-watcher/internal/app/app';
import { Events } from '@wailsio/runtime';

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

	listen(id: string) {
		Events.On('service-update-' + id, ({ data }) => {
			this.status = this.#mapBackendStatus(data);
		});
	}

	#mapBackendStatus(backendStatus: string): ServiceStatus {
		const statusMap: Record<string, ServiceStatus> = {
			Running: 'running',
			Stopped: 'stopped',
			Starting: 'starting',
			Stopping: 'stopping',
			Error: 'error',
			'Not Found': 'error',
			Unknown: 'error',
			Paused: 'stopped',
			Pausing: 'stopping',
			Resuming: 'starting'
		};
		return statusMap[backendStatus] || 'error';
	}

	async start() {
		if (['running', 'starting'].includes(this.status)) return;
		this.status = 'starting';

		try {
			await StartService(this.id);
		} catch (err) {
			console.error(`Error starting service ${this.id}:`, err);
			this.status = 'error';
		}
	}

	async stop() {
		if (['stopped', 'stopping'].includes(this.status)) return;
		this.status = 'stopping';

		try {
			await StopService(this.id);
		} catch (err) {
			console.error(`Error stopping service ${this.id}:`, err);
			this.status = 'error';
		}
	}

	async restart() {
		await this.stop();
		await new Promise((resolve) => setTimeout(resolve, 1000));
		await this.start();
	}
}

export class ServiceStore {
	services = $state<Service[]>([]);
	isLoading = $state<boolean>(true);

	healthyCount = $derived(() => this.services.filter((s) => s.status === 'running').length);
	totalCount = $derived(() => this.services.length);

	async init() {
		try {
			await new Promise((resolve) => setTimeout(resolve, 500));

			const configs = await GetConfig();

			configs.Services.forEach((config) => {
				const service = new Service({
					id: config.id,
					name: config.name || 'Unknown',
					description: config.description || 'No description provided.',
					status: 'stopped',
					logsUrl: config.log_path ? `file://${config.log_path}` : null
				});

				service.listen(config.id);
				this.services.push(service);
			});

			this.isLoading = false;
		} catch (err) {
			console.error('Error loading config in ServiceStore:', err);
		}
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
