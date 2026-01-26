import { CheckUpdate, Update } from '../../../wailsjs/go/app/App';
import type { app } from '../../../wailsjs/go/models';

export interface Setting {
	runBackground: boolean;
	notifications: boolean;
	runOnStartup: boolean;
	autoCheckUpdate: boolean;
}

export class SystemStore {
	setting = $state<Setting>({
		runBackground: false,
		notifications: false,
		runOnStartup: false,
		autoCheckUpdate: false
	});
	initialized = $state<boolean>(false);
	version = $state<app.UpdateInfo>({
		available: false,
		currentVersion: 'v0.0.0',
		latestVersion: 'v0.0.0',
		downloadUrl: ''
	});

	async init() {
		if (this.initialized) return;
		try {
			const version = await this.checkUpdate();
			this.version = { ...this.version, ...version };
			this.initialized = true;
		} catch (err) {
			console.error('Error loading system info:', err);
		}
	}

	async checkUpdate() {
		const version = await CheckUpdate();
		return version;
	}

	async update() {
		const url = this.version?.downloadUrl ?? '';
		if (!url) return;

		await Update(url);
	}
}

export const systemStore = new SystemStore();
