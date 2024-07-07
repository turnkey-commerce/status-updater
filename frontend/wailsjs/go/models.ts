export namespace main {
	
	export class Credentials {
	    smtpUser: string;
	    smtpPassword: string;
	
	    static createFrom(source: any = {}) {
	        return new Credentials(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.smtpUser = source["smtpUser"];
	        this.smtpPassword = source["smtpPassword"];
	    }
	}

}

