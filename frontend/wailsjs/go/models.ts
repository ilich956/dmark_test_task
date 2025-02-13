export namespace models {
	
	export class Task {
	    ID: number;
	    Title: string;
	    Description: string;
	    Deadline: string;
	    Priority: string;
	    Status: string;
	
	    static createFrom(source: any = {}) {
	        return new Task(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Title = source["Title"];
	        this.Description = source["Description"];
	        this.Deadline = source["Deadline"];
	        this.Priority = source["Priority"];
	        this.Status = source["Status"];
	    }
	}

}

