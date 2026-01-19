export namespace domain {
	
	export class ServiceStatus {
	    name: string;
	    status: string;
	    is_healthy: boolean;
	
	    static createFrom(source: any = {}) {
	        return new ServiceStatus(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.status = source["status"];
	        this.is_healthy = source["is_healthy"];
	    }
	}

}

