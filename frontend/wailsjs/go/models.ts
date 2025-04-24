export namespace main {
	
	export class CommandResult {
	    output: string;
	    currentDirectory: string;
	    error: string;
	    isBuiltinCommand: boolean;
	
	    static createFrom(source: any = {}) {
	        return new CommandResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.output = source["output"];
	        this.currentDirectory = source["currentDirectory"];
	        this.error = source["error"];
	        this.isBuiltinCommand = source["isBuiltinCommand"];
	    }
	}

}

